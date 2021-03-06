protoc -I/usr/local/include -I. \
  --go_out=plugins=micro:$GOPATH/src/github.com/ewanvalentine/mgo-proto-test \
  -I$GOPATH/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis \
  proto/greeter/greeter.proto

protoc -I/usr/local/include -I. \
  -I$GOPATH/src \
  -I$GOPATH/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis \
  --grpc-gateway_out=logtostderr=true:$GOPATH/src/github.com/ewanvalentine/mgo-proto-test/api \
  proto/greeter/greeter.proto

protoc -I/usr/local/include -I. \
  -I$GOPATH/src \
  -I$GOPATH/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis \
  --go_out=plugins=grpc:$GOPATH/src/github.com/ewanvalentine/mgo-proto-test/api \
  proto/greeter/greeter.proto

protoc -I/usr/local/include -I. \
  -I$GOPATH/src \
  -I$GOPATH/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis \
  --swagger_out=logtostderr=true:$GOPATH/src/github.com/ewanvalentine/mgo-proto-test/api \
  proto/greeter/greeter.proto

protoc-go-inject-tag -input=proto/greeter/greeter.pb.go
protoc-go-inject-tag -input=api/proto/greeter/greeter.pb.go

GOOS=linux GOARCH=amd64 go build
