package main

import (
	"grpc_mongo_crud/db"
	"grpc_mongo_crud/server"
	"log"
	"net"

	pb "grpc_mongo_crud/proto"

	"google.golang.org/grpc"
)

func main() {
	err := db.InitMango()

	if err != nil {
		log.Fatal("Mongo Connection Failed ", err)
	}

	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatal("Failed to listen: ", err)
	}
	grpcServer := grpc.NewServer()
	pb.RegisterUserServiceServer(grpcServer, &server.UserServer{})
	log.Println("Grpc server is running on port 50051")
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatal("Failed to serve the grpc server", err)
	}
}
