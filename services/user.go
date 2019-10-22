package services

import (
	"crypto/md5"
	"encoding/hex"
	"errors"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"tarobackend/models"
	pb "tarobackend/proto"
	"tarobackend/utils"
	"time"
)

type UserReq struct {
	PageIndex    int64  `json:"page_index"`
	PageSize     int64  `json:"page_size"`
	SearchName   string `json:"search_name"`
	RegisterId   int64  `json:"register_id"`
	RegisterName string `json:"register_name"`
}

type UserResp struct {
	List  []models.TaroUser `json:"list"`
	Count int64             `json:"count"`
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

func CreateUser(r *models.TaroUser) (int64, error) {
	engine := utils.Engine_mysql
	res, err := engine.InsertOne(r)
	if err != nil {
		logs.Error("CreateUser: Table User InsertOne Error")
		return 0, err
	}
	if res == 0 {
		logs.Debug("CreateUser: User InsertOne failed")
	}
	return res, nil
}

func DeleteUserById(id int) error {
	engine := utils.Engine_mysql
	r := new(models.TaroUser)
	_, err := engine.ID(id).Delete(r)
	if err != nil {
		logs.Error("DeleteUserById: Table User Delete Error")
		return err
	}
	return nil
}

func UpdateUser(r *models.TaroUser) error {
	engine := utils.Engine_mysql
	_, err := engine.ID(r.UserId).Update(r)
	if err != nil {
		logs.Error("UpdateUser: Table User Update Error")
		return err
	}
	return nil
}

func RegisterUser(req *UserReq) (int64, error) {
	// Set up a connection to the server.
	conn, err := grpc.Dial(beego.AppConfig.String("fabric_service"), grpc.WithInsecure())
	if err != nil {
		logs.Error("did not connect: %v", err)
		return -1, err
	}
	//defer conn.Close()
	c := pb.NewFabricServiceClient(conn)
	// Contact the server and print out its response.
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()
	r, err := c.Register(ctx, &pb.RegisterReq{Username: req.RegisterName})
	if err != nil {
		logs.Error("RegisterUser: could not Register: %v", err)
		return -1, err
	}
	if r.GetCode() == 0 {
		user := models.TaroUser{UserStatus: 1}
		engine := utils.Engine_mysql
		_, err = engine.ID(req.RegisterId).Update(&user)
		if err != nil {
			logs.Error("RegisterUser: User Status Update Error")
			return -1, err
		}
	}
	return r.GetCode(), nil
}

func DownloadCard(req *UserReq) (*pb.DownloadResp, error) {
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
	r, err := c.Download(ctx, &pb.DownloadReq{Username: req.RegisterName})
	if err != nil {
		logs.Error("DownloadCard: could not Download: %v", err)
		return nil, err
	}
	if len(r.Card) == 0 {
		logs.Error("Card is Empty")
		return nil, errors.New("Card is Empty")
	}

	md5Inst := md5.New()
	md5Inst.Write([]byte(r.Card))
	md5Sum := md5Inst.Sum([]byte(""))
	engine := utils.Engine_mysql
	user := new(models.TaroUser)
	user.UserHash = hex.EncodeToString(md5Sum)
	_, err = engine.ID(req.RegisterId).Update(user)
	if err != nil {
		logs.Error("DownloadCard: User Hash Update Error")
		return nil, err
	}
	return r, nil
}
