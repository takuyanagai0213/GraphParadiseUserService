package main

import (
	"github.com/takuyanagai0213/GraphParadiseUserService/grpc"
)

func main() {
	// gRPCサーバーの起動
	grpc.NewUserGrpcServer()
}
