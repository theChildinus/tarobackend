# 建议单条复制到 terminal 中使用
# grpc 结合 protobuf 自动生成 pb.go
protoc -I. \
  -I$GOPATH/src \
  -I$GOPATH/pkg/mod/github.com/grpc-ecosystem/grpc-gateway@v1.11.2/third_party/googleapis \
  --go_out=plugins=grpc:. \
  ./proto/fabric_service.proto

# grpc 结合 protobuf 和 grpc-gateway 自动生成 pb.gw.go
protoc -I. \
  -I$GOPATH/pkg/mod/github.com/grpc-ecosystem/grpc-gateway@v1.11.2/third_party/googleapis \
  --grpc-gateway_out=logtostderr=true:. \
  ./proto/fabric_service.proto

# curl 构造POST请求 请求体为 json
curl --header "Content-Type: application/json" \
  --request POST \
  --data '{"username":"xyz","password":"xyz"}' \
  http://localhost:3000/api/login

# xorm 自动生成 model代码
xorm reverse mysql root:123456@/taro?charset=utf8 templates/goxorm