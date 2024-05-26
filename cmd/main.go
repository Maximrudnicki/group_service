package main

import (
	"context"
	"fmt"
	"log"
	"net"

	"group_service/cmd/config"
	"group_service/cmd/repository"
	pb "group_service/proto"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"google.golang.org/grpc"
)

type Server struct {
	pb.GroupServiceServer
	GroupRepository      repository.GroupRepository
	StatisticsRepository repository.StatisticsRepository
}

func main() {
	loadConfig, err := config.LoadConfig(".")
	if err != nil {
		log.Fatal("🚀 Could not load environment variables", err)
	}

	mc := fmt.Sprintf("mongodb://%s:%s", loadConfig.MONGODB_HOST, loadConfig.MONGODB_PORT) // "mongodb://localhost:27017"

	clientOptions := options.Client().ApplyURI(mc)

	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Println("Error connecting:", err)
	}

	groupCollection := client.Database(loadConfig.MONGODB_DB).Collection("groups")
	statisticsCollection := client.Database(loadConfig.MONGODB_DB).Collection("statistics")

	//Init Repository
	groupRepository := repository.NewGroupRepositoryImpl(groupCollection)
	statisticsRepository := repository.NewStatisticsRepositoryImpl(statisticsCollection)

	lis, err := net.Listen("tcp", loadConfig.GRPCPort)
	if err != nil {
		log.Fatalf("Failed to listen: %v\n", err)
	}

	log.Printf("Listening at %s\n", loadConfig.GRPCPort)

	opts := []grpc.ServerOption{}

	s := grpc.NewServer(opts...)

	pb.RegisterGroupServiceServer(s, &Server{
		GroupRepository: groupRepository, 
		StatisticsRepository: statisticsRepository,
	})

	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v\n", err)
	}
}
