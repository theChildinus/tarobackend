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
	"math/rand"
	"strconv"
	"strings"
	"tarobackend/models"
	pb "tarobackend/proto"
	"tarobackend/utils"
	"time"
)

var curRandom int64

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
	List  []string `json:"list"`
	Count int64    `json:"count"`
}

type LoginReq struct {
	UserName string `json:"name"`
	UserSecret string `json:"secret"`
}

type LoginResp struct {
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
			And("user_role = ?", req.SearchType).
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
	engine := utils.Engine_mysql
	res, err := engine.InsertOne(r)
	if err != nil {
		logs.Error("CreateUser: Table User InsertOne Error")
		return false, err
	}
	if res == 0 {
		logs.Error("CreateUser: User InsertOne Failed")
		return false, errors.New("CreateUser: User InsertOne Failed")
	}
	enf := utils.Enforcer
	success := enf.AddRoleForUser(r.UserName, r.UserRole)
	_ = enf.SavePolicy()
	if !success {
		logs.Error("CreateUser: User Add Role Failed")
		return false, errors.New("CreateUser: User Add Role Failed")
	}
	return true, nil
}

func DeleteUserById(r *models.TaroUser) (bool, error) {
	engine := utils.Engine_mysql
	m := new(models.TaroUser)
	_, err := engine.ID(r.UserId).Delete(m)
	if err != nil {
		logs.Error("DeleteUserById: Table User Delete Error")
		return false, err
	}
	enf := utils.Enforcer
	success := enf.DeleteRoleForUser(r.UserName, r.UserRole)
	_ = enf.SavePolicy()
	if !success {
		logs.Error("DeleteUserById: Delete User Role Error")
		return false, errors.New("DeleteUserById: Delete User Role Error")
	}
	return true, nil
}

func UpdateUser(r *models.TaroUser) (bool, error) {
	engine := utils.Engine_mysql
	old := new(models.TaroUser)
	var ret bool
	has, err := engine.Table("taro_user").
		Where("user_id = ?", r.UserId).Get(old)
	if err != nil {
		logs.Error("UpdateUser: Table User Get Error")
		return false, err
	}
	if has {
		enf := utils.Enforcer
		ret1 := enf.DeleteRoleForUser(old.UserName, old.UserRole)
		ret2 := enf.AddRoleForUser(r.UserName, r.UserRole)
		_ = enf.SavePolicy()
		ret = ret1 && ret2
	}
	_, err = engine.ID(r.UserId).Update(r)
	if err != nil {
		logs.Error("UpdatePolicy: Table Policy Update Error")
		return false, err
	}
	return ret, nil
}

func ListUserNameAndRole() ([]string, int64, error) {
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
	return append(names, roles...), count, nil
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
	r, err := c.Register(ctx, &pb.RegisterReq{Name: req.Name, Type: "user"})
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
	r, err := c.Download(ctx, &pb.DownloadReq{Name: req.Name, Type: "user"})
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
	fromdb := new(models.TaroUser)
	engine := utils.Engine_mysql
	has, err := engine.Table("taro_user").
		Where("user_name = ?", req.Name).Get(fromdb)
	if err != nil {
		logs.Error("VerifyIdentity: User Info Get Error")
		return -1, err
	}
	if !has {
		logs.Error("VerifyIdentity: User Doesn't Exist")
		return -1, err
	}
	if len(req.Sign) == 0 || req.Rand == 0 {
		randnum := rand.Int63()
		md5Inst := md5.New()
		nameWithNum := req.Name + strconv.FormatInt(randnum, 10)
		md5Inst.Write([]byte(nameWithNum))
		md5Sum := md5Inst.Sum([]byte(""))
		engine := utils.Engine_mysql
		user := new(models.TaroUser)
		user.UserHash = hex.EncodeToString(md5Sum)
		// fmt.Println("namewithNum: ", nameWithNum, "userhash: ", user.UserHash)
		_, err = engine.Where("user_name = ?", req.Name).Update(user)
		if err != nil {
			logs.Error("VerifyIdentity: User Hash Update Error")
			return -1, err
		}
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
	r, err := c.VerifyIdentity(ctx, &pb.VerifyIdentityReq{Name: req.Name, Rand: req.Rand, Sign: req.Sign, Type: "user"})
	if err != nil {
		logs.Error("VerifyIdentity: could not verify: %v", err)
		return -1, err
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
	r, err := c.Revoke(ctx, &pb.RevokeReq{Name: req.Name, Type: "user"})
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
	r, err := c.VerifyCert(ctx, &pb.VerifyCertReq{Name: req.Name, Certcontent: req.Certcontent, Type: "user"})
	if err != nil {
		logs.Error("VerifyCert: could not Verify: %v", err)
		return -1, err
	}
	return r.GetCode(), nil
}

func Login(req *LoginReq, Token string) (string, error) {

	// Create a new token object, specifying signing method and the claims
	// you would like it to contain.
	if len(req.UserSecret) != 0 && len(Token) == 0 {
		// verify UserSecret
		token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
			"foo": "bar",
			"nbf": time.Now().Unix(),
		})

		// Sign and get the complete encoded token as a string using the secret
		bytes, _ := base64.StdEncoding.DecodeString(req.UserSecret)
		if tokenString, err := token.SignedString(bytes); err == nil {
			tokenMap[req.UserName] = tokenString
			fmt.Println(tokenString, err)
			return tokenString, nil
		} else {
			return "", err
		}
	} else if len(Token) != 0 {
		if t, ok := tokenMap[req.UserName]; ok {
			if t == Token {
				return "0", nil
			} else {
				return "-1", nil
			}
		}
	}
	return "", nil
}

func Logout(req *LogoutReq, token string) (int64, error) {
	// verify token
	// delete in tokenMap
	if tokenMap[req.UserName] == token {
		delete(tokenMap, req.UserName)
	}
	return 0, nil
}
