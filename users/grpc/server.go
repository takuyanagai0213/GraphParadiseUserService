package grpc

import (
	"context"
	"fmt"
	"log"
	"net"
	"os"
	"os/signal"

	"github.com/takuyanagai0213/GraphParadiseUserService/domain/model"
	"github.com/takuyanagai0213/GraphParadiseUserService/grpc/user"
	"github.com/takuyanagai0213/GraphParadiseUserService/usecase"
	"google.golang.org/grpc"
)

type server struct {
	Usecase usecase.UserUseCase
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
func (s server) ReadUser(ctx context.Context, req *user.ReadUserRequest) (*user.ReadUserResponse, error) {
	name := req.GetName()
	fmt.Println(name)
	user_data, err := s.Usecase.GetUserByName(name)

	if err != nil {
		return nil, err
	}
	user_data_comp := makeGrpcUserProfile(&user_data)
	return &user.ReadUserResponse{
		User: user_data_comp,
	}, nil
}
func makeGrpcUserProfile(user_data *model.User) *user.User {
	gUser := &user.User{
		Name:     user_data.Name,
		Password: user_data.Password,
		Area:     "a",
	}
	return gUser
}
