# 安全管理系统

## 项目概况

| 项目           | 描述                                                         | github地址                                                   | docker 容器名  |
| -------------- | ------------------------------------------------------------ | ------------------------------------------------------------ | -------------- |
| web 前端       | vue.js，提供用户管理、资源管理、权限管理等功能，在用户管理中包括用户证书的创建、下载、注销等 | [tarofrontend 分支： master](https://github.com/theChildinus/tarofrontend) | 部署在nginx中  |
| web 后端       | beego + casbin + grpc (client)，web的后端服务，将数据持久化到数据库中，作为 grpc 的客户端调用 fabric_service，请求 Fabric CA 服务 | [tarobackend 分支：master](https://github.com/theChildinus/tarobackend) | goweb          |
| fabric_service | fabric-java-sdk + grpc (server)，作为 grpc 的服务端 处理 web后台的 Fabric CA 请求，同时作为 Fabric CA Client 与 Fabric CA Server 通信 | [fabric-service-client 分支：master](https://github.com/theChildinus/fabric-service-client) | fabric_service |
| API 示例代码   | 提供 用户登录、权限校验、策略获取、证书验证等功能            | [LoginCheckApp 分支：master](https://github.com/theChildinus/LoginCheckApp) | 无             |
| 数据库         | mysql                                                        | 无                                                           | mysql5_7       |
| Fabric CA      | Fabric CA Server，提供基本的 Fabric CA 服务                  | 无                                                           | ca.example.com |

## 项目构建

| 项目（文件夹名）                       | 开发环境                                                     | 正式环境                                                     |
| -------------------------------------- | ------------------------------------------------------------ | ------------------------------------------------------------ |
| web 前端tarofrontend）                 | `npm run dev`                                                | `npm run build` 将生成的 `dist` 文件夹复制到 tarobackend 文件夹中 |
| web 后端(tarobackend)                  | 修改 `conf/app.conf` 中的 `httpaddr`, `mysqlurl`, `fabric_service` 为 localhost，端口不变，Goland IDE 直接启动 | 1. 注释 `compose.yaml`  中的 `command: sh -c 'cd /go/src && ./main'`。2. `docker-compose -f compose.yaml up -d` 重启 goweb服务。3. `docker exec -it goweb bash` 进入goweb容器，`cd src` , `go build main.go`  等待编译完成。4. 取消 `compose.yaml` 中的注释，再重启 goweb 服务 |
| fabric_service (fabric-service-client) | IDEA IDE 直接启动 FabricService                              | IDEA IDE Maven 工具中选择输入命令：`mvn assembly:assembly` , 将生成的后缀为 `with-dependencies`的jar包 拷贝到 tarobackend 文件夹中 |

## 部署

- 修改 `conf/app.conf`  本地主机信息部分
- 修改 `conf/nginx.conf` 第33行 `server_name` 为本机IP，保留 `https://`
- 一键部署 `docker-compose -f compose.yaml up -d` 

## 一些命令

#### 使用 grpc 结合 protobuf 自动生成代码 XXXXpb.go

```sh
protoc -I. \
  -I$GOPATH/src \
  --go_out=plugins=grpc:. \
  ./proto/fabric_service.proto
```

#### 使用 grpc-gateway 自动生成 XXXXpb.gw.go

其中 `googleapis` 的路径为 go mod 本地仓库中对应项目的路径

``` bash
protoc -I. \
  -I$GOPATH/pkg/mod/github.com/grpc-ecosystem/grpc-gateway@v1.11.2/third_party/googleapis \
  --grpc-gateway_out=logtostderr=true:. \
  ./proto/fabric_service.proto
```

#### curl 构造 POST 请求示例

`-k` 忽略 https 证书验证

```bash
curl -k --header "Content-Type: application/json" \
  --request POST \
  --data '{"username":"zhao","userhash":"317fd62b83164a6d0cd2b27808941da2"}' \
  https://localhost:4433/user/login
```

#### xorm 自动生成 model 代码

mysql 数据库表结构更改后，可使用该命令自动更新 goweb 中 model 层代码

```bash
xorm reverse mysql "root:zky_taro_mysql@(127.0.0.1:33069)/taro?charset=utf8" templates/goxorm
```

#### openssl 操作

```bash
# 从证书中提取公钥
openssl x509 -in zhao.crt -pubkey -nocert -out zhao_pub.pem

# 公钥格式转换
openssl pkcs8 -topk8 -nocrypt -in zhao.pem -out zhao2.pem
```

#### fabric-service 自动生成 grpc 代码

Idea 打开 maven工具，Plugins中选择 `protobuf`，运行 `protobuf:compile` 和 `protobuf:compile-custom`

## 注意事项

1. 外网访问本机端口时，需要在 windows 或 linux 防火墙中添加规则，允许访问。
2. 工程 fabric-service-client 与 tarobackend 中的 fabric_service.proto 文件的`message` `rpc` 内容要保持一致
3. 开发环境时，证书等文件产生在 fabric-service-client 的 card 文件夹下，正式环境时，产生在 tarobackend 的 card 文件夹下