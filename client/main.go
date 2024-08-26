package main

import (
	context "context"
	"log"
	"time"

	grpc "google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	LogViaGRPC()
}

func LogViaGRPC() {
	conn, err := grpc.NewClient("localhost:50051", grpc.WithTransportCredentials(insecure.NewCredentials()), grpc.WithBlock())
	if err != nil {
		log.Fatalf("Failed to dial gRPC server: %v", err)
	}
	defer conn.Close()

	// Use the conn variable to make gRPC calls
	// lets create a new client
	c := NewLogServiceClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	_, err = c.WriteLog(ctx, &LogRequest{
		LogEntry: &Log{
			Name: "Prakalp",
			Data: "this data is a test data",
		}})

	if err != nil {
		log.Fatalf("Failed to write log: %v", err)
	}
	log.Println("Successfully logged via gRPC")
}
