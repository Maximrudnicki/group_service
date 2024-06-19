package utils

import (
	"log"

	"group_service/cmd/config"
	pb "group_service/proto"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func GetUserIdFromToken(token string) (uint32, error) {
	loadConfig, err := config.LoadConfig(".")
	if err != nil {
		log.Fatal("🚀 Could not load environment variables", err)
	}
	// connect to auth_service as a client
	conn, err := grpc.Dial(loadConfig.AUTH_SERVICE, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Printf("Did not connect: %v", err)
		return 0, err
	}
	defer conn.Close()
	c := pb.NewAuthenticationServiceClient(conn)

	UserIdResponse, err := ValidateToken(c, token)

	if err != nil {
		return 0, err
	}

	return UserIdResponse.UserId, nil
}
