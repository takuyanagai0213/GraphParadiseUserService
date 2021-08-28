package grpc

import (
	"context"
	"fmt"
	"log"
	"net"
	"os"
	"os/signal"

	"github.com/takuyanagai0213/GraphParadiseUserService/grpc/user"
	"google.golang.org/grpc"

	"github.com/takuyanagai0213/GraphParadiseUserService/database"
	"github.com/takuyanagai0213/GraphParadiseUserService/domain/repository"
)

type server struct {
}

func NewUserGrpcServer() {

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
