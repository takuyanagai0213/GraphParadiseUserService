package main

import (
	"github.com/takuyanagai0213/GraphParadiseUserService/db"
	"github.com/takuyanagai0213/GraphParadiseUserService/grpc"
)

func main() {
	db.Init()
	// gRPCサーバーの起動
	grpc.NewUserGrpcServer()
	defer db.Close()
}
