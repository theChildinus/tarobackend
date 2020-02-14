package services

import (
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"github.com/pkg/sftp"
	"golang.org/x/crypto/ssh"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"net"
	"os"
	"path"
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

type IdentityNamesResp struct {
	List  []NameOptions `json:"list"`
	Count int64    `json:"count"`
}

type NameOptions struct {
	Value string `json:"value"`
	Label string `json:"label"`
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

func InstallIdentity(req *pb.InstallReq) (int64, error) {
	sftpClient, err := connect(req.User, req.Pw, req.Ip, 22)
	if err != nil {
		logs.Error("InstallIdentity: Connect Error: %v", err)
		return -1, err
	}
	defer sftpClient.Close()
	certFilePath := "./card/" + req.Name + "/" + req.Name + ".crt"
	skFilePath :=  "./card/" + req.Name + "/" + req.Name + ".pem"

	fmt.Println("cert: " + certFilePath)
	fmt.Println("sk: " + skFilePath)

	certFile, err := os.Open(certFilePath)
	if err != nil {
		logs.Error("InstallIdentity: Open File" + req.Name + ".crt Failed")
		return -1, err
	}
	defer certFile.Close()
	skFile, err := os.Open(skFilePath)
	if err != nil {
		logs.Error("InstallIdentity: Open File" + req.Name + ".pem Failed")
	}
	defer skFile.Close()

	buf := make([]byte, 1024)
	dstFile, err := sftpClient.Create(path.Join(req.Path, path.Base(certFilePath)))
	if err != nil {
		logs.Error("InstallIdentity: Remote CreateFile" + req.Name + ".crt Failed")
	}
	for {
		n, _ := certFile.Read(buf)
		if n == 0 {
			break
		}
		dstFile.Write(buf)
	}
	dstFile, err = sftpClient.Create(path.Join(req.Path, path.Base(skFilePath)))
	if err != nil {
		logs.Error("InstallIdentity: Remote CreateFile" + req.Name + ".pem Failed")
	}
	for {
		n, _ := skFile.Read(buf)
		if n == 0 {
			break
		}
		dstFile.Write(buf)
	}

	defer dstFile.Close()
	fmt.Println("copy file to remote server finished!")
	return 0, nil
}

func connect(user, password, host string, port int) (*sftp.Client, error) {
	var (
		auth         []ssh.AuthMethod
		addr         string
		clientConfig *ssh.ClientConfig
		sshClient    *ssh.Client
		sftpClient   *sftp.Client
		err          error
	)
	// get auth method
	auth = make([]ssh.AuthMethod, 0)
	auth = append(auth, ssh.Password(password))

	clientConfig = &ssh.ClientConfig{
		User:    user,
		Auth:    auth,
		Timeout: 30 * time.Second,
		HostKeyCallback: func(hostname string, remote net.Addr, key ssh.PublicKey) error {
			return nil
		},
	}

	// connet to ssh
	addr = fmt.Sprintf("%s:%d", host, port)

	if sshClient, err = ssh.Dial("tcp", addr, clientConfig); err != nil {
		return nil, err
	}

	// create sftp client
	if sftpClient, err = sftp.NewClient(sshClient); err != nil {
		return nil, err
	}

	return sftpClient, nil
}

func ListIdentityNames() ([]NameOptions, int64, error) {
	engine := utils.Engine_mysql
	var (
		names []string
		err error
	)
	err = engine.Table("taro_identity").Select("identity_name").Find(&names)
	if err != nil {
		logs.Error("ListUserNameAndRole: Table User Find Names Error")
		return nil, 0, err
	}
	var ins []NameOptions
	for _, v := range names {
		m := &NameOptions{Value: v, Label: v}
		ins = append(ins, *m)
	}
	strs := []NameOptions{{"Admin","Admin"},
		{"Readers","Readers"},
		{"Writers","Writers"},
	}
	ins = append(ins, strs...)
	return ins, int64(len(ins)), nil
}
