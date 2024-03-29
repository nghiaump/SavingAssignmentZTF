package main

import (
	"context"
	"flag"
	"github.com/golang/glog"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	pb "github.com/nghiaump/SavingAssignmentZTF/protobuf"
	"github.com/rs/cors"
	"google.golang.org/grpc"
	"net/http"
	"os"
)

const MidAddress = "mid-saving:50050"

func run() error {
	// Init glog
	os.Args = append(os.Args, "-logtostderr=true")
	flag.Parse()
	defer glog.Flush()

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
	glog.Info("api-gateway start 8081")

	// cors.Default() không đủ, đặc biệt pre-flight request
	//handler := cors.Default().Handler(mux)

	// tạo CORS handler với custom options
	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"}, // Change this to the appropriate origin or origins
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Authorization", "Content-Type"},
		AllowCredentials: true,
	})

	// Wrap the ServeMux with the CORS handler
	handler := c.Handler(mux)

	// Start HTTP server (and proxy calls to gRPC server endpoint)
	return http.ListenAndServe(":8081", handler)
}

func main() {
	flag.Parse()
	if err := run(); err != nil {
		glog.Errorf("Error: %v", err)
	}
}
