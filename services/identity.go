package services

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"tarobackend/models"
	pb "tarobackend/proto"
	"tarobackend/utils"
	"time"
)

type IdentityReq struct {
	PageIndex  int64  `json:"page_index"`
	PageSize   int64  `json:"page_size"`
	SearchType string `json:"search_type"`
	SearchName string `json:"search_name"`
}

type IdentityResp struct {
	List  []models.TaroIdentity `json:"list"`
	Count int64                 `json:"count"`
}

func ListIdentity(req *IdentityReq) ([]models.TaroIdentity, int64, error) {
	engine := utils.Engine_mysql
	var (
		Identitys []models.TaroIdentity
		err       error
		count     int64
	)
	m := new(models.TaroIdentity)
	if len(req.SearchType) != 0 {
		err = engine.Table("taro_identity").
			Where("identity_name like ? ", "%"+req.SearchName+"%").
			And("identity_type = ?", req.SearchType).
			Limit(int(req.PageSize), int((req.PageIndex-1)*req.PageSize)).
			Find(&Identitys)
		count, _ = engine.Where("identity_name like ? ", "%"+req.SearchName+"%").
			And("identity_type = ?", req.SearchType).Count(m)
	} else {
		err = engine.Table("taro_identity").
			Limit(int(req.PageSize), int((req.PageIndex-1)*req.PageSize)).
			Find(&Identitys)
		count, _ = engine.Count(m)
	}

	if err != nil {
		logs.Error("ListIdentity: Table Identity Find Error")
		return nil, 0, err
	}
	return Identitys, count, nil
}

func CreateIdentity(r *models.TaroIdentity) (int64, error) {
	engine := utils.Engine_mysql
	res, err := engine.InsertOne(r)
	if err != nil {
		logs.Error("CreateIdentity: Table Identity InsertOne Error")
		return 0, err
	}
	if res == 0 {
		logs.Debug("CreateIdentity: Identity InsertOne failed")
	}
	return res, nil
}

func DeleteIdentityById(id int) error {
	engine := utils.Engine_mysql
	r := new(models.TaroIdentity)
	_, err := engine.ID(id).Delete(r)
	if err != nil {
		logs.Error("DeleteIdentityById: Table Identity Delete Error")
		return err
	}
	return nil
}

func UpdateIdentity(r *models.TaroIdentity) error {
	engine := utils.Engine_mysql
	_, err := engine.ID(r.IdentityId).Update(r)
	if err != nil {
		logs.Error("UpdateIdentity: Table Identity Update Error")
		return err
	}
	return nil
}

func RegisterIdentity(req *pb.RegisterReq) (int64, error) {
	// TODO Register Identity Id/Name in IdentityTable？
	// Set up a connection to the server.
	conn, err := grpc.Dial(beego.AppConfig.String("fabric_service"), grpc.WithInsecure())
	if err != nil {
		logs.Error("RegisterIdentity: did not connect: %v", err)
		return -1, err
	}
	//defer conn.Close()
	c := pb.NewFabricServiceClient(conn)
	// Contact the server and print out its response.
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()
	r, err := c.Register(ctx, &pb.RegisterReq{
		Name:        req.Name,
		Secret:      req.Secret,
		Type:        req.Type,
		Affiliation: req.Affiliation,
		Attrs:       req.Attrs,
	})
	if err != nil {
		logs.Error("RegisterIdentity: could not Register: %v", err)
		return -1, err
	}
	if r.GetCode() == 0 {
		identity := models.TaroIdentity{IdentityStatus: 1}
		engine := utils.Engine_mysql
		_, err = engine.ID(req.Id).Update(&identity)
		if err != nil {
			logs.Error("RegisterIdentity: identity Status Update Error")
			return -1, err
		}
	}
	return r.GetCode(), nil
}

func EnrollIdentity(req *pb.EnrollReq) (int64, error) {
	// Set up a connection to the server.
	conn, err := grpc.Dial(beego.AppConfig.String("fabric_service"), grpc.WithInsecure())
	if err != nil {
		logs.Error("EnrollIdentity: did not connect: %v", err)
		return -1, err
	}
	//defer conn.Close()
	c := pb.NewFabricServiceClient(conn)
	// Contact the server and print out its response.
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()
	r, err := c.Enroll(ctx, &pb.EnrollReq{Name: req.Name, Secret: req.Secret, Attrs: req.Attrs, Type: req.Type})
	if err != nil {
		logs.Error("EnrollIdentity: could not Enroll: %v", err)
		return -1, err
	}
	if r.GetCode() == 0 {
		identity := models.TaroIdentity{IdentityStatus: 2}
		engine := utils.Engine_mysql
		_, err = engine.ID(req.Id).Update(&identity)
		if err != nil {
			logs.Error("EnrollIdentity: Identity Status Update Error")
			return -1, err
		}
	}
	return r.GetCode(), nil
}

func RevokeIdentity(req *pb.RevokeReq) (int64, error) {
	// TODO Register Identity Id/Name in IdentityTable？
	// Set up a connection to the server.
	conn, err := grpc.Dial(beego.AppConfig.String("fabric_service"), grpc.WithInsecure())
	if err != nil {
		logs.Error("RevokeIdentity: did not connect: %v", err)
		return -1, err
	}
	//defer conn.Close()
	c := pb.NewFabricServiceClient(conn)
	// Contact the server and print out its response.
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()
	r, err := c.Revoke(ctx, &pb.RevokeReq{Name: req.Name, Type: req.Type})
	if err != nil {
		logs.Error("RevokeIdentity: could not Revoke: %v", err)
		return -1, err
	}
	if r.GetCode() == 0 {
		identity := models.TaroIdentity{IdentityStatus: 1}
		engine := utils.Engine_mysql
		_, err = engine.ID(req.Id).Update(&identity)
		if err != nil {
			logs.Error("RevokeIdentity: Identity Status Update Error")
			return -1, err
		}
	}
	return r.GetCode(), nil
}
