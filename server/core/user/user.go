package main

import (
	"context"
	"fmt"
	"github.com/elastic/go-elasticsearch/v7"
	"github.com/elastic/go-elasticsearch/v7/esapi"
	"github.com/google/uuid"
	pb "github.com/nghiaump/SavingAssignmentZTF/protobuf"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"log"
	"math/rand"
	"net"
	"strings"
)

type UserServiceHandler struct {
	kycMap   map[string]*pb.KYC
	usersMap map[string]*pb.User
	esClient *elasticsearch.Client
}

func NewUserServiceHandler(client *elasticsearch.Client) *UserServiceHandler {
	handler := UserServiceHandler{}
	handler.esClient = client
	return &handler
}

func StartUserServer(handler *UserServiceHandler, port string) {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	log.Println("Creating UserServiceHandler")
	pb.RegisterUserServiceServer(s, handler)
	log.Printf("Starting gRPC User service listener on port " + port)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

func (handler *UserServiceHandler) RegisterUser(ctx context.Context, req *pb.RegisterUserRequest) (*pb.RegisterUserResponse, error) {
	out, err := uuid.NewUUID()
	log.Printf("Calling RegisterUser(): %v", out.String())
	if err != nil {
		return nil, status.Errorf(codes.Internal,
			"Error while generating UserID %v", err)
	}

	indexReq := esapi.IndexRequest{
		Index:   "user",
		Body:    strings.NewReader(fmt.Sprintf(`{"name" : "%s", "cccd": "%s"}`, req.UserName, req.IdCardNumber)),
		Refresh: "true",
	}

	indexRes, err2 := indexReq.Do(context.Background(), handler.esClient)
	if err2 != nil {
		log.Printf("Error indexing document: %v\n", err2)
	}
	defer indexRes.Body.Close()
	log.Printf("Indexed new User to ElasticSearch %v\n", indexRes)
	return &pb.RegisterUserResponse{Success: true, UserId: out.String()}, status.New(codes.OK, "").Err()
}

func (handler *UserServiceHandler) GetCurrentKYC(ctx context.Context, req *pb.GetCurrentKYCRequest) (*pb.GetCurrentKYCResponse, error) {
	log.Printf("Calling GetCurrentKYC for userID: %v", req.UserId)

	_, exists := handler.usersMap[req.UserId]
	if !exists {
		log.Printf("unregistered userID")
		return nil, status.New(codes.NotFound, "").Err()
	}

	var kycLevel int32
	if handler.kycMap == nil {
		handler.kycMap = make(map[string]*pb.KYC)
		kycLevel = GenKYC(req)
		handler.kycMap[req.UserId] = &pb.KYC{UserId: req.UserId, Level: kycLevel}
	} else {
		// Map existed
		value, exists := handler.kycMap[req.UserId]
		if !exists {
			kycLevel = GenKYC(req)
			handler.kycMap[req.UserId] = &pb.KYC{UserId: req.UserId, Level: kycLevel}
		} else {
			// KYC existed
			kycLevel = value.Level
		}
	}
	return &pb.GetCurrentKYCResponse{UserId: req.UserId, KycLevel: kycLevel}, status.New(codes.OK, "").Err()
}

func GenKYC(req *pb.GetCurrentKYCRequest) int32 {

	return int32(rand.Intn(3) + 1)
}
func GenKYCDefault() int32 {

	return int32(rand.Intn(3) + 1)
}
