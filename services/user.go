package services

import (
	"crypto/md5"
	"encoding/hex"
	"errors"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"strings"
	"tarobackend/models"
	pb "tarobackend/proto"
	"tarobackend/utils"
	"time"
)

type UserReq struct {
	PageIndex  int64  `json:"page_index"`
	PageSize   int64  `json:"page_size"`
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

func ListUser(req *UserReq) ([]models.TaroUser, int64, error) {
	engine := utils.Engine_mysql
	var (
		users []models.TaroUser
		err   error
		count int64
	)
	m := new(models.TaroUser)
	if len(req.SearchName) != 0 {
		err = engine.Table("taro_user").
			Where("user_name = ?", req.SearchName).
			Limit(int(req.PageSize), int((req.PageIndex-1)*req.PageSize)).
			Find(&users)
		count, _ = engine.Where("user_name = ?", req.SearchName).Count(m)
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
		err error
		count int64
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
	// Set up a connection to the server.
	conn, err := grpc.Dial(beego.AppConfig.String("fabric_service"), grpc.WithInsecure())
	if err != nil {
		logs.Error("RegisterUser: did not connect: %v", err)
		return -1, err
	}
	//defer conn.Close()
	c := pb.NewFabricServiceClient(conn)
	// Contact the server and print out its response.
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()
	r, err := c.Register(ctx, &pb.RegisterReq{Username: req.Username})
	if err != nil {
		logs.Error("RegisterUser: could not Register: %v", err)
		return -1, err
	}
	if r.GetCode() == 0 {
		user := models.TaroUser{UserStatus: 1}
		engine := utils.Engine_mysql
		_, err = engine.ID(req.Userid).Update(&user)
		if err != nil {
			logs.Error("RegisterUser: User Status Update Error")
			return -1, err
		}
	}
	return r.GetCode(), nil
}

func DownloadCard(req *pb.DownloadReq) (*pb.DownloadResp, error) {
	// Set up a connection to the server.
	conn, err := grpc.Dial(beego.AppConfig.String("fabric_service"), grpc.WithInsecure())
	if err != nil {
		logs.Error("DownloadCard: did not connect: %v", err)
		return nil, err
	}
	//defer conn.Close()
	c := pb.NewFabricServiceClient(conn)
	// Contact the server and print out its response.
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()
	r, err := c.Download(ctx, &pb.DownloadReq{Username: req.Username})
	if err != nil {
		logs.Error("DownloadCard: could not Download: %v", err)
		return nil, err
	}
	if len(r.Card) == 0 {
		logs.Error("DownloadCard: Card is Empty")
		return nil, errors.New("DownloadCard: Card is Empty")
	}

	md5Inst := md5.New()
	md5Inst.Write([]byte(r.Card))
	md5Sum := md5Inst.Sum([]byte(""))
	engine := utils.Engine_mysql
	user := new(models.TaroUser)
	user.UserHash = hex.EncodeToString(md5Sum)
	_, err = engine.ID(req.Userid).Update(user)
	if err != nil {
		logs.Error("DownloadCard: User Hash Update Error")
		return nil, err
	}
	return r, nil
}

func Login(req *pb.LoginReq) (int64, error) {
	if len(req.Userhash) == 0 {
		logs.Error("Login: User Hash Empty")
		return -1, nil
	}
	fromdb := new(models.TaroUser)
	engine := utils.Engine_mysql
	has, err := engine.Table("taro_user").
		Where("user_name = ?", req.Username).Get(fromdb)
	if err != nil {
		logs.Error("Login: User Hash Get Error")
		return -1, err
	}
	if has && len(fromdb.UserHash) != 0 && fromdb.UserHash == req.Userhash {
		// Set up a connection to the server.
		conn, err := grpc.Dial(beego.AppConfig.String("fabric_service"), grpc.WithInsecure())
		if err != nil {
			logs.Error("Login: did not connect: %v", err)
			return -1, err
		}
		//defer conn.Close()
		c := pb.NewFabricServiceClient(conn)
		// Contact the server and print out its response.
		ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
		defer cancel()
		r, err := c.Login(ctx, &pb.LoginReq{Username: req.Username})
		if err != nil {
			logs.Error("Login: could not Download: %v", err)
			return -1, err
		}
		return r.GetCode(), nil
	} else {
		return -1, nil
	}
}
