package main

import (
	context "context"
	"fmt"
)

type LogServer struct {
	UnimplementedLogServiceServer
}

func (l *LogServer) WriteLog(ctx context.Context, req *LogRequest) (*LogResponse, error) {
	input := req.GetLogEntry()

	fmt.Println("Received log entry: ", input)

	return &LogResponse{Result: "logged!!"}, nil
}
