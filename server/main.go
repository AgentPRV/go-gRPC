package main

import (
	"log"
	"net"

	grpc "google.golang.org/grpc"
)

func StartServer() {
	// Create a listener on TCP port 50051
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	// Create a new gRPC server
	s := grpc.NewServer()

	// Register the server implementation
	RegisterLogServiceServer(s, &LogServer{})

	log.Println("Starting server on port :50051")

	// Start the server
	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}

func main() {
	StartServer()
}
