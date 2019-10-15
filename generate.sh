# pb.go
protoc -I. \
  -I$GOPATH/src \
  -I$GOPATH/pkg/mod/github.com/grpc-ecosystem/grpc-gateway@v1.11.2/third_party/googleapis \
  --go_out=plugins=grpc:. \
  ./proto/register.proto

## pb.gw.go
#protoc -I. \
#  -I$GOPATH/pkg/mod/github.com/grpc-ecosystem/grpc-gateway@v1.11.2/third_party/googleapis \
#  --grpc-gateway_out=logtostderr=true:. \
#  ./proto/register.proto

#curl --header "Content-Type: application/json" \
#  --request POST \
#  --data '{"username":"xyz","password":"xyz"}' \
#  http://localhost:3000/api/login