package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/rs/cors"
	"github.com/takuyanagai0213/GraphParadiseUserService/database"
	"github.com/takuyanagai0213/GraphParadiseUserService/infrastructure/persistence"
	"github.com/takuyanagai0213/GraphParadiseUserService/interfaces/handler"
	"github.com/takuyanagai0213/GraphParadiseUserService/usecase"
)

func main() {
	// gRPCサーバーの起動
	// grpc.NewUserGrpcServer()

	// 依存関係を定義
	userPersistence := persistence.NewUserPersistence(database.DBConnect())
	userUseCase := usecase.NewUserUseCase(userPersistence)
	userHandler := handler.NewUserHandler(userUseCase)
	// // ルーティングの設定
	router := httprouter.New()
	router.GET("/api/users", userHandler.Index)
	router.GET("/user/new", userHandler.CreateUser)

	// サーバ起動
	fmt.Println("========================")
	fmt.Println("Server Start >> http://localhost:80")
	fmt.Println("========================")
	log.Fatal(http.ListenAndServe(":80", router))
	// CORS対策
	handler := cors.Default().Handler(router)
	http.ListenAndServe(":80", handler)
}
