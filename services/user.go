package services

import (
	"crypto/md5"
	"encoding/base64"
	"encoding/hex"
	"errors"
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"github.com/dgrijalva/jwt-go"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"gopkg.in/gomail.v2"
	"math/rand"
	"mime"
	"strconv"
	"strings"
	"tarobackend/models"
	pb "tarobackend/proto"
	"tarobackend/utils"
	"time"
)

type UserReq struct {
	PageIndex  int64  `json:"page_index"`
	PageSize   int64  `json:"page_size"`
	SearchType string `json:"search_type"`
	SearchName string `json:"search_name"`
}

type UserResp struct {
	List  []models.TaroUser `json:"list"`
	Count int64             `json:"count"`
}

type UserNameAndRoleResp struct {
	List  []NameOptions `json:"list"`
	Count int64         `json:"count"`
}

type DeleteIds struct {
	Ids []int `json:"ids"`
}

type LoginReq struct {
	UserName   string `json:"name"`
	UserRole   string `json:"role"`
	UserSecret string `json:"secret"`
}

type LoginResp struct {
	Code  int    `json:"code"`
	Token string `json:"token"`
}

type LogoutReq struct {
	UserName string `json:"name"`
}

var tokenMap = make(map[string]string, 10)

func ListUser(req *UserReq) ([]models.TaroUser, int64, error) {
	engine := utils.Engine_mysql
	var (
		users []models.TaroUser
		err   error
		count int64
	)
	m := new(models.TaroUser)
	if len(req.SearchType) != 0 {
		err = engine.Table("taro_user").
			Where("user_name like ? ", "%"+req.SearchName+"%").
			And("user_role like ?", "%"+req.SearchType+"%").
			Limit(int(req.PageSize), int((req.PageIndex-1)*req.PageSize)).
			Find(&users)
		count, _ = engine.Where("user_name like ? ", "%"+req.SearchName+"%").
			And("user_role = ?", req.SearchType).Count(m)
	} else {
		err = engine.Table("taro_user").
			Limit(int(req.PageSize), int((req.PageIndex-1)*req.PageSize)).
			Find(&users)
		count, _ = engine.Count(m)
	}

	if err != nil {
		logs.Error("ListUser: Table User Find Error")
		return nil, 0, err
	}
	return users, count, nil
}

func CreateUser(r *models.TaroUser) (bool, error) {
	roles := strings.Split(r.UserRole, "#")
	if len(roles) == 2 {
		rar := &RoleAllotReq{Name: r.UserName, Roles: roles}
		if ret, _ := RoleAllot(rar); !ret {
			return false, nil
		}
	}
	engine := utils.Engine_mysql
	has, _ := engine.Exist(&models.TaroUser{UserName: r.UserName})
	if has {
		return false, errors.New("CreateUser: User " + r.UserName + " Existed!")
	}
	res, err := engine.InsertOne(r)
	if err != nil {
		logs.Error("CreateUser: Table User InsertOne Error")
		return false, err
	}
	if res == 0 {
		logs.Error("CreateUser: User InsertOne Failed")
		return false, errors.New("CreateUser: User InsertOne Failed")
	}
	return true, nil
}

func DeleteUserById(ids []int) (bool, error) {
	engine := utils.Engine_mysql
	m := new(models.TaroUser)
	_, err := engine.Table("taro_user").In("user_id", ids).Delete(m)
	if err != nil {
		logs.Error("DeleteUserById: Table User Delete Error")
		return false, err
	}
	return true, nil
}

func UpdateUser(r *models.TaroUser) (bool, error) {
	roles := strings.Split(r.UserRole, "#")
	if len(roles) == 2 {
		rar := &RoleAllotReq{Name: r.UserName, Roles: roles}
		if ret, _ := RoleAllot(rar); !ret {
			return false, nil
		}
	}
	engine := utils.Engine_mysql
	old := new(models.TaroUser)
	_, err := engine.Table("taro_user").
		Where("user_id = ?", r.UserId).Get(old)
	if err != nil {
		logs.Error("UpdateUser: Table User Get Error")
		return false, err
	}
	_, err = engine.ID(r.UserId).Update(r)
	if err != nil {
		logs.Error("UpdatePolicy: Table Policy Update Error")
		return false, err
	}
	return true, nil
}

func ListUserNameAndRole() ([]NameOptions, int64, error) {
	engine := utils.Engine_mysql
	var (
		names, roles []string
		err          error
		count        int64
	)
	err = engine.Table("taro_user").Select("user_name").Find(&names)
	if err != nil {
		logs.Error("ListUserNameAndRole: Table User Find Names Error")
		return nil, 0, err
	}

	ev, err := GetEnumValue(&models.TaroEnum{EnumKey: "user_role"})
	if err != nil {
		logs.Error("ListUserNameAndRole: Table Enum Find Roles Error")
		return nil, 0, err
	}
	roles = strings.Split(ev.EnumValue, "##")
	count = int64(len(names) + len(roles))
	var ins []NameOptions
	for _, v := range names {
		m := &NameOptions{Value: v, Label: v}
		ins = append(ins, *m)
	}
	for _, v := range roles {
		m := &NameOptions{Value: v, Label: v}
		ins = append(ins, *m)
	}
	return ins, count, nil
}

func RegisterUser(req *pb.RegisterReq) (int64, error) {
	// TODO Register User Id/Name in UserTable？
	// Set up a connection to the server.
	conn, err := grpc.Dial(beego.AppConfig.String("fabric_service"), grpc.WithInsecure())
	if err != nil {
		logs.Error("RegisterUser: did not connect: %v", err)
		return -1, err
	}
	//defer conn.Close()
	c := pb.NewFabricServiceClient(conn)
	// Contact the server and print out its response.
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()
	r, err := c.Register(ctx, &pb.RegisterReq{Name: req.Name, Type: "iotuser"})
	if err != nil {
		logs.Error("RegisterUser: could not Register: %v", err)
		return -1, err
	}
	if r.GetCode() == 0 {
		user := models.TaroUser{UserStatus: 1}
		engine := utils.Engine_mysql
		_, err = engine.ID(req.Id).Update(&user)
		if err != nil {
			logs.Error("RegisterUser: User Status Update Error")
			return -1, err
		}
	}
	return r.GetCode(), nil
}

func InstallUser(req *models.TaroUser) (int64, error) {
	if len(req.UserName) == 0 {
		logs.Error("UserName is empty")
		return -1, errors.New("UserName is empty")
	}

	if len(req.UserPath) != 0 {
		if _, err := InstallIdentity(&pb.InstallReq{
			Name: req.UserName,
			Ip:   beego.AppConfig.String("local_host"),
			User: beego.AppConfig.String("local_username"),
			Pw:   beego.AppConfig.String("local_password"),
			Path: req.UserPath,
		}); err != nil {
			logs.Error(err.Error())
			return -1, err
		}
		return 0, nil
	}
	if len(req.UserEmail) != 0 {
		if _, err := SendEmail(req.UserName, req.UserEmail); err != nil {
			logs.Error(err.Error())
			return -1, err
		}
		return 0, nil
	}
	return -1, errors.New("UserPath & UserEmail are empty")
}

func SendEmail(username, useremail string) (int64, error) {
	certFilePath := "./card/" + username + "/" + username + ".crt"
	skFilePath := "./card/" + username + "/" + username + ".pem"
	m := gomail.NewMessage()
	m.SetAddressHeader("From", beego.AppConfig.String("sender_email"),
		beego.AppConfig.String("sender_name"))
	m.SetHeader("To", m.FormatAddress(useremail, username))
	m.SetHeader("Subject", beego.AppConfig.String("email_subject"))
	m.SetBody("text/html", beego.AppConfig.String("email_body"))
	m.Attach(certFilePath, gomail.SetHeader(map[string][]string{
		"Content-Disposition": []string{
			fmt.Sprintf(`attachment; filename="%s"`, mime.QEncoding.Encode("UTF-8", username+".crt")),
		},
	}))
	m.Attach(skFilePath, gomail.SetHeader(map[string][]string{
		"Content-Disposition": []string{
			fmt.Sprintf(`attachment; filename="%s"`, mime.QEncoding.Encode("UTF-8", username+".pem")),
		},
	}))
	smtp_port, _ := strconv.Atoi(beego.AppConfig.String("smtp_port"))
	d := gomail.NewDialer(beego.AppConfig.String("smtp_server"), smtp_port,
		beego.AppConfig.String("sender_email"),
		beego.AppConfig.String("sender_auth_code"),
	)
	if err := d.DialAndSend(m); err != nil {
		return -1, err
	}
	return 0, nil
}

func DownloadCert(req *pb.DownloadReq) (*pb.DownloadResp, error) {
	// Set up a connection to the server.
	conn, err := grpc.Dial(beego.AppConfig.String("fabric_service"), grpc.WithInsecure())
	if err != nil {
		logs.Error("DownloadCert: did not connect: %v", err)
		return nil, err
	}
	//defer conn.Close()
	c := pb.NewFabricServiceClient(conn)
	// Contact the server and print out its response.
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()
	r, err := c.Download(ctx, &pb.DownloadReq{Name: req.Name, Type: "iotuser"})
	if err != nil {
		logs.Error("DownloadCert: could not Download: %v", err)
		return nil, err
	}
	if len(r.Cert) == 0 {
		logs.Error("DownloadCert: Cert is Empty")
		return nil, errors.New("DownloadCert: Cert is Empty")
	}

	return r, nil
}

func VerifyIdentity(req *pb.VerifyIdentityReq) (int64, error) {

	engine := utils.Engine_mysql
	isUser, _ := engine.Exist(&models.TaroUser{UserName: req.Name})
	isIdentity, _ := engine.Exist(&models.TaroIdentity{IdentityName: req.Name})
	if !isUser && !isIdentity {
		logs.Error("VerifyIdentity: User Doesn't Exist")
		return -1, nil
	}
	if len(req.Sign) == 0 || req.Rand == 0 {
		randnum := rand.Int63()
		return randnum, nil
	}

	// Set up a connection to the server.
	conn, err := grpc.Dial(beego.AppConfig.String("fabric_service"), grpc.WithInsecure())
	if err != nil {
		logs.Error("VerifyIdentity: did not connect: %v", err)
		return -1, err
	}
	//defer conn.Close()
	c := pb.NewFabricServiceClient(conn)
	// Contact the server and print out its response.
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()
	r, err := c.VerifyIdentity(ctx, &pb.VerifyIdentityReq{Name: req.Name, Rand: req.Rand, Sign: req.Sign, Type: "iotuser"})
	if err != nil {
		logs.Error("VerifyIdentity: could not verify: %v", err)
		return -1, err
	}
	if r.Code == 0 {
		md5Inst := md5.New()
		bytes, _ := base64.StdEncoding.DecodeString(req.Sign)
		md5Inst.Write(bytes)
		md5Sum := md5Inst.Sum([]byte(""))
		if isUser {
			user := new(models.TaroUser)
			user.UserHash = hex.EncodeToString(md5Sum)
			logs.Info("isUser")
			_, err = engine.Where("user_name = ?", req.Name).Update(user)
			if err != nil {
				logs.Error("VerifyIdentity: User Hash Update Error")
				return -1, err
			}
		} else if isIdentity {
			identity := new(models.TaroIdentity)
			identity.IdentityHash = hex.EncodeToString(md5Sum)
			logs.Info("isIdentity")
			_, err = engine.Where("identity_name = ?", req.Name).Update(identity)
			if err != nil {
				logs.Error("VerifyIdentity: Identity Hash Update Error")
				return -1, err
			}
		}
	}
	return r.Code, nil
}

func RevokeUser(req *pb.RevokeReq) (int64, error) {
	// TODO Register User Id/Name in UserTable？
	// Set up a connection to the server.
	conn, err := grpc.Dial(beego.AppConfig.String("fabric_service"), grpc.WithInsecure())
	if err != nil {
		logs.Error("RevokeUser: did not connect: %v", err)
		return -1, err
	}
	//defer conn.Close()
	c := pb.NewFabricServiceClient(conn)
	// Contact the server and print out its response.
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()
	r, err := c.Revoke(ctx, &pb.RevokeReq{Name: req.Name, Type: "iotuser"})
	if err != nil {
		logs.Error("RevokeUser: could not Revoke: %v", err)
		return -1, err
	}
	if r.GetCode() == 0 {
		user := models.TaroUser{UserStatus: 0, UserHash: ""}
		engine := utils.Engine_mysql
		_, err = engine.ID(req.Id).Cols("user_status", "user_hash").Update(&user)
		if err != nil {
			logs.Error("RevokeUser: User Status Update Error")
			return -1, err
		}
	}
	return r.GetCode(), nil
}

func VerifyCert(req *pb.VerifyCertReq) (int64, error) {
	conn, err := grpc.Dial(beego.AppConfig.String("fabric_service"), grpc.WithInsecure())
	if err != nil {
		logs.Error("VerifyCert: did not connect: %v", err)
		return -1, err
	}
	//defer conn.Close()
	c := pb.NewFabricServiceClient(conn)
	// Contact the server and print out its response.
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()
	r, err := c.VerifyCert(ctx, &pb.VerifyCertReq{Name: req.Name, Certcontent: req.Certcontent, Type: "iotuser"})
	if err != nil {
		logs.Error("VerifyCert: could not Verify: %v", err)
		return -1, err
	}
	return r.GetCode(), nil
}

func Login(req *LoginReq, tokenStr string) (string, error) {
	// determine user has allocated role
	if len(req.UserRole) == 0 {
		return "-1", errors.New("UserRole empty")
	}
	user := new(models.TaroUser)
	if _, err := utils.Engine_mysql.Table("taro_user").
		Where("user_name = ?", req.UserName).Get(user); err != nil {
		return "-1", err
	}
	hasRole := false
	roles := strings.Split(user.UserRole, "#")
	for _, v := range roles {
		if strings.Contains(v, req.UserRole) {
			hasRole = true
			break
		}
	}
	if !hasRole {
		return "-1", errors.New("User doesn't has this role")
	}
	// Create a new token object, specifying signing method and the claims
	// you would like it to contain.
	// first time login, tokenStr not generate
	if len(req.UserSecret) != 0 && len(tokenStr) == 0 {
		// verify UserSecret
		conn, err := grpc.Dial(beego.AppConfig.String("fabric_service"), grpc.WithInsecure())
		if err != nil {
			logs.Error("Login: did not connect: %v", err)
			return "-1", err
		}
		//defer conn.Close()
		c := pb.NewFabricServiceClient(conn)
		// Contact the server and print out its response.
		ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
		defer cancel()
		// userSecret should be One Time Password
		r, err := c.VerifyIdentity(ctx,
			&pb.VerifyIdentityReq{Name: req.UserName, Rand: 123456, Sign: req.UserSecret, Type: "iotuser"})
		if err != nil {
			logs.Error("Login: could not verify: %v", err)
			return "-1", err
		}
		if r.Code != 0 {
			return "-1", nil
		}
		token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
			"sub": req.UserName,
			"nbf": time.Now().Unix(),
		})

		// Sign and get the complete encoded token as a string using the secret
		bytes, _ := base64.StdEncoding.DecodeString(req.UserSecret)
		if tokenString, err := token.SignedString(bytes); err == nil {
			tokenMap[req.UserName] = req.UserSecret
			logs.Info("TokenStr: ", tokenString)
			return tokenString, nil
		} else {
			return "", err
		}
		// login with tokenStr
	} else if len(tokenStr) != 0 {
		if t, ok := tokenMap[req.UserName]; ok {
			claims, err := parseToken(tokenStr, t)
			if err != nil {
				return "-1", err
			}
			if claims.(jwt.MapClaims)["sub"] == req.UserName {
				logs.Info("claims: ", claims.(jwt.MapClaims)["sub"])
				return "0", nil
			} else {
				fmt.Println("")
				return "-1", nil
			}
		}
	}
	return "-1", nil
}

func Logout(req *LogoutReq, tokenStr string) (int64, error) {
	// verify token
	claims, err := parseToken(tokenStr, tokenMap[req.UserName])
	if err != nil {
		return -1, err
	}
	if claims.(jwt.MapClaims)["sub"] == req.UserName {
		delete(tokenMap, req.UserName)
		return 0, nil
	} else {
		return -1, nil
	}
}

func parseToken(tokenStr string, secretKey string) (interface{}, error) {
	bytes, _ := base64.StdEncoding.DecodeString(secretKey)
	token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
		// Don't forget to validate the alg is what you expect:
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}

		// hmacSampleSecret is a []byte containing your secret, e.g. []byte("my_secret_key")
		return bytes, nil
	})
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return claims, nil
	} else {
		return nil, err
	}
}
