package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"grpc_demo/gym"
	"log"
	"net"
)

const (
	port = ":50051"
)

// server继承自动生成的服务类
type server struct {
	gym.UnimplementedGymServer
}

// 服务端必须实现了相应的接口BodyBuilding
func (s *server) BodyBuilding(ctx context.Context, in *gym.Person) (*gym.Reply, error) {
	fmt.Printf("%s正在健身, 动作: %s\n", in.Name, in.Actions)
	return &gym.Reply{Code: 0, Msg: "ok"}, nil
}
func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	gym.RegisterGymServer(s, &server{})

	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
