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

# curl 构造登录 POST 请求 请求体为 json 格式
curl -k --header "Content-Type: application/json" \
  --request POST \
  --data '{"username":"zhao","userhash":"317fd62b83164a6d0cd2b27808941da2"}' \
  https://localhost:4433/user/login

# 构造策略判定 POST 请求 请求体为 json 格式
curl -k --header "Content-Type: application/json charset=UTF-8" \
  --request POST \
  --data '{"username":"zhao","userhash":"317fd62b83164a6d0cd2b27808941da2","policysub":"zhao","policyobj":"data4","policyact":"exec"}' \
  https://localhost:4433/policy/check

# xorm 自动生成 model代码
xorm reverse mysql root:123456@/taro?charset=utf8 templates/goxorm