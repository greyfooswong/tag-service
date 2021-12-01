package main

import (
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"log"
	"net"
	pb "tag-service/proto"
	"tag-service/server"
)

func main() {
	s := grpc.NewServer()
	pb.RegisterTagServiceServer(s, server.NewTagServer())
	reflection.Register(s)
	lis, err := net.Listen("tcp", ":"+"7999")
	if err != nil {
		log.Fatalf("net.Listen err : %v", err)
	}
	err = s.Serve(lis)
	if err != nil {
		log.Fatalf("server.Serve err : %v", err)
	}
}
