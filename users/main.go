package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"os"
	"os/signal"

	"github.com/takuyanagai0213/GraphParadiseUserService/database"
	"github.com/takuyanagai0213/GraphParadiseUserService/domain/repository"
	"github.com/takuyanagai0213/GraphParadiseUserService/grpc/user"

	"google.golang.org/grpc"
)

type server struct {
}

func main() {
	// dir, _ := os.Getwd()
	// 依存関係を定義
	// userPersistence := persistence.NewUserPersistence(database.DBConnect())
	// userUseCase := usecase.NewUserUseCase(userPersistence)
	// userHandler := handler.NewUserHandler(userUseCase)
	// // ルーティングの設定
	// router := httprouter.New()
	// router.GET("/api/users", userHandler.Index)
	// router.GET("/user/new", api.CreateUser)

	fmt.Println("Blog Service Started")
	lis, err := net.Listen("tcp", "0.0.0.0:50051")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	// grpc
	opts := []grpc.ServerOption{}
	s := grpc.NewServer(opts...)

	user.RegisterUserServiceServer(s, &server{})
	go func() {
		fmt.Println("Starting server ....")
		if err := s.Serve(lis); err != nil {
			log.Fatalf("Failed to serve: %v", err)
		}
	}()
	//Wait for Control C to Exit
	ch := make(chan os.Signal, 1)
	signal.Notify(ch, os.Interrupt)

	<-ch
	fmt.Println("Stopping the server")
	s.Stop()
	fmt.Println("Close the listener")
	lis.Close()
	fmt.Println("end of program")
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

// 検索
func (*server) Search(context.Context, *user.ListUserRequest) (*user.ListUserResponse, error) {
	var user_list []*repository.User

	// DB接続確認
	db_conn := database.DBConnect()
	if err := db_conn.Take(&user_list).Error; err != nil {
		return nil, err
	}

	db_conn.Find(&user_list)

	// 名前検索
	// if name != "" {
	// 	db = db.Where("name = ?", name).Find(&user)
	// }

	return &user.ListUserResponse{
		Name: "名前",
		// Profile: user_list,
	}, nil
}
