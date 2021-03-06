package grpc

import (
	"fmt"
	"log"
	"net"
	"os"
	"os/signal"

	"github.com/takuyanagai0213/GraphParadiseUserService/grpc/userservice"
	"github.com/takuyanagai0213/GraphParadiseUserService/usecase/interactor"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type server struct {
	Usecase interactor.UserInteractor
}

// NewUserGrpcServer gRPCサーバー起動
func NewUserGrpcServer() {
	lis, err := net.Listen("tcp", "0.0.0.0:50051")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	server := &server{}

	s := makeServer()

	userservice.RegisterUserServiceServer(s, server)

	// Register reflection service on gRPC server.
	reflection.Register(s)
	log.Println("main grpc server has started")

	go func() {
		if err := s.Serve(lis); err != nil {
			log.Fatalf("failed to serve: %v", err)
		}
	}()

	ch := make(chan os.Signal, 1)
	signal.Notify(ch, os.Interrupt)

	// Block until a sgnal is received
	<-ch
	fmt.Println("Stopping the server")
	s.Stop()
	fmt.Println("Closing the client")
	lis.Close()
	fmt.Println("End of Program")

}

func makeServer() *grpc.Server {
	s := grpc.NewServer(
		grpc.UnaryInterceptor(grpc.UnaryServerInterceptor(transmitStatusInterceptor)),
		// grpc_auth.UnaryServerInterceptor(authorization.AuthFunc),
		// ),
	)
	return s
}
