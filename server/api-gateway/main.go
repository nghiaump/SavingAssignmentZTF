package main

import (
	"context"
	"flag"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	pb "github.com/nghiaump/SavingAssignmentZTF/protobuf"
	"github.com/rs/cors"
	"google.golang.org/grpc"
	"log"
	"net/http"
)

const MidAddress = "mid-saving:50050"

func run() error {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	// Register gRPC server endpoint
	// Note: Make sure the gRPC server is running properly and accessible
	mux := runtime.NewServeMux()
	opts := []grpc.DialOption{grpc.WithInsecure()}
	err := pb.RegisterMidSavingServiceHandlerFromEndpoint(ctx, mux, MidAddress, opts)
	if err != nil {
		return err
	}
	log.Println("api-gateway start 8081")
	handler := cors.Default().Handler(mux)
	// Start HTTP server (and proxy calls to gRPC server endpoint)
	return http.ListenAndServe(":8081", handler)
}

func main() {
	flag.Parse()
	if err := run(); err != nil {
		log.Printf("Error: %v", err)
	}
}
