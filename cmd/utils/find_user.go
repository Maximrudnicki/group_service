package utils

import (
	"context"
	"log"

	pb "group_service/proto"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func FindUser(userId uint32) (*pb.UserResponse, error) {
	// connect to auth_service as a client
	conn, err := grpc.Dial("0.0.0.0:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Printf("Did not connect: %v", err)
		return nil, err
	}
	defer conn.Close()
	c := pb.NewAuthenticationServiceClient(conn)

	req := &pb.FindUserRequest{
		UserId: userId,
	}

	res, err := c.FindUser(context.Background(), req)
	if err != nil {
		log.Printf("Error happened while getting ID: %v\n", err)
		return nil, err
	}

	return res, nil
}
