protoc -I. -I$GOPATH/src \
       --go_out=plugins=grpc:. \
       --grpc-gateway_out=. \
       *.proto

protoc-go-inject-tag -input="*.pb.go"

