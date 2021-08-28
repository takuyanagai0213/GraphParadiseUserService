package main

import (
	"github.com/takuyanagai0213/GraphParadiseUserService/grpc"
)

func main() {
	grpc.NewUserGrpcServer()
	// dir, _ := os.Getwd()
	// 依存関係を定義
	// userPersistence := persistence.NewUserPersistence(database.DBConnect())
	// userUseCase := usecase.NewUserUseCase(userPersistence)
	// userHandler := handler.NewUserHandler(userUseCase)
	// // ルーティングの設定
	// router := httprouter.New()
	// router.GET("/api/users", userHandler.Index)
	// router.GET("/user/new", api.CreateUser)

	// サーバ起動
	// fmt.Println("========================")
	// fmt.Println("Server Start >> http://localhost:3000")
	// fmt.Println("========================")
	// log.Fatal(http.ListenAndServe(":80", router))

	// http.HandleFunc("/user/new", api.CreateUser)
	// http.Handle("/user/get", new(userHandler.Index))
	// http.HandleFunc("/user/update", api.UpdateUsers)
	// port
	// fmt.Println("========================")
	// fmt.Println("Server Start >> http://localhost:80")
	// fmt.Println("========================")
	// CORS対策
	// handler := cors.Default().Handler(router)
	// http.ListenAndServe(":80", handler)
}
