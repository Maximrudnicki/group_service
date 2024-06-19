package utils

import (
	"context"
	"log"

	"group_service/cmd/config"
	pb "group_service/proto"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func AddWordToStudent(studentId uint32, word string, definition string) (*pb.AddWordToStudentResponse, error) {
	loadConfig, err := config.LoadConfig(".")
	if err != nil {
		log.Fatal("🚀 Could not load environment variables", err)
	}
	conn, err := grpc.Dial(loadConfig.VOCAB_SERVICE, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Printf("Did not connect: %v", err)
		return nil, err
	}
	defer conn.Close()
	c := pb.NewVocabServiceClient(conn)

	req := &pb.AddWordToStudentRequest{
		Word: word,
		Definition: definition,
		UserId: studentId,
	}

	res, err := c.AddWordToStudent(context.Background(), req)
	if err != nil {
		log.Printf("Error happened while getting ID: %v\n", err)
		return nil, err
	}

	return res, nil
}
