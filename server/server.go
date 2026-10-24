package server

import (
	"context"
	"errors"
	"grpc_mongo_crud/db"
	_ "grpc_mongo_crud/db"
	"grpc_mongo_crud/models"
	_ "grpc_mongo_crud/models"
	pb "grpc_mongo_crud/proto"

	"go.mongodb.org/mongo-driver/bson"
)

type UserServer struct {
	pb.UnimplementedUserServiceServer
}

func (s *UserServer) CreateUser(ctx context.Context, req *pb.UserRequest) (*pb.UserResponse, error) {
	user := models.User{
		UserId:   req.UserId,
		Name:     req.Name,
		Age:      req.Age,
		Email:    req.Email,
		Password: req.Password,
	}
	_, err := db.UserCollection.InsertOne(ctx, user)
	if err != nil {
		return nil, err
	}
	return &pb.UserResponse{
		UserId: req.UserId,
		Name:   req.Name,
		Age:    req.Age,
		Email:  req.Email,
	}, nil
}

func (s *UserServer) GetUser(ctx context.Context, req *pb.GetUserRequest) (*pb.UserResponse, error) {
	user := models.User{}
	err := db.UserCollection.FindOne(ctx, bson.M{"userId": req.UserId}).Decode(&user)
	if err != nil {
		return nil, errors.New("user not found")
	}
	return &pb.UserResponse{
		UserId: user.UserId,
		Name:   user.Name,
		Email:  user.Email,
		Age:    user.Age,
	}, nil
}

func (s *UserServer) UpdateUser(ctx context.Context, req *pb.UserRequest) (*pb.UserResponse, error) {
	filter := bson.M{"userId": req.UserId}
	update := bson.M{"$set": bson.M{
		"name":     req.Name,
		"age":      req.Age,
		"email":    req.Email,
		"password": req.Password,
	}}
	_, err := db.UserCollection.UpdateOne(ctx, filter, update)
	if err != nil {
		return nil, err
	}
	return &pb.UserResponse{
		UserId: req.UserId,
		Name:   req.Name,
		Age:    req.Age,
		Email:  req.Email,
	}, nil
}

func (s *UserServer) DeleteUser(ctx context.Context, req *pb.GetUserRequest) (*pb.DeleteResponse, error) {
	_, err := db.UserCollection.DeleteOne(ctx, bson.M{"userId": req.UserId})
	if err != nil {
		return nil, err
	}

	return &pb.DeleteResponse{Message: "User deleted successfully"}, nil
}
