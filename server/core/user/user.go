package main

import (
	"context"
	"github.com/google/uuid"
	pb "github.com/nghiaump/savingaccproto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"log"
	"math/rand"
	"net"
)

type UserServiceHandler struct {
	usersMap map[string]*pb.User
	kycMap   map[string]*pb.KYC
}

func NewUserServiceHandler() *UserServiceHandler {
	handler := UserServiceHandler{}
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

	if handler.usersMap == nil {
		handler.usersMap = make(map[string]*pb.User)
	}

	handler.usersMap[out.String()] = &pb.User{
		UserId:         out.String(),
		IdCardNumber:   req.IdCardNumber,
		UserName:       req.UserName,
		KycLevel:       GenKYCDefault(), // TODO
		RegisteredDate: "01012024",      // TODO
	}
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
