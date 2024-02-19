package main

import (
	"context"
	"github.com/google/uuid"
	pb "github.com/nghiaump/SavingAssignmentZTF/protobuf"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"log"
	"net"
)

type SavingServiceHandler struct {
	accountMap map[string]*pb.SavingAccount
}

func NewSavingServiceHandler() *SavingServiceHandler {
	handler := SavingServiceHandler{}
	return &handler
}

func StartSavingServer(handler *SavingServiceHandler, port string) {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	log.Println("Creating SavingServiceHandler")
	pb.RegisterSavingsServiceServer(s, handler)
	log.Printf("Starting gRPC Saving Service listener on SavingPort " + port)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

func (handler *SavingServiceHandler) OpenSavingsAccount(ctx context.Context, req *pb.SavingAccount) (*pb.SavingAccount, error) {
	out, err := uuid.NewUUID()
	if err != nil {
		return nil, status.Errorf(codes.Internal,
			"Error while generating AccountID", err)
	}
	log.Printf("Calling OpenSavingAccount(), userID: %v, accountID: %v", req.UserID, out.String())

	if handler.accountMap == nil {
		handler.accountMap = make(map[string]*pb.SavingAccount)
	}

	req.Id = out.String()
	handler.accountMap[req.Id] = req
	return req, status.New(codes.OK, "").Err()
}

func (handler *SavingServiceHandler) AccountInquiry(ctx context.Context, req *pb.AccountInquiryRequest) (*pb.SavingAccount, error) {
	log.Printf("Calling AccoutAquiry(), userID: %v, accountID: %v", req.UserId, req.AccountId)

	acc, exists := handler.accountMap[req.AccountId]
	if exists {
		if req.UserId == acc.UserID {
			log.Printf("Account %v exist", req.AccountId)
			return &pb.SavingAccount{
				Id:          req.AccountId,
				UserID:      req.UserId,
				Balance:     acc.Balance,
				TermType:    acc.TermType,
				Term:        acc.Term,
				CreatedDate: acc.CreatedDate,
				DueDate:     acc.DueDate,
				Rate:        acc.Rate,
			}, status.New(codes.OK, "").Err()
		} else {
			log.Printf("Account %v exist, but userID %v mismatch", req.AccountId, req.UserId)
			return nil, status.Errorf(codes.PermissionDenied, "")
		}

	}
	return nil, status.Errorf(codes.NotFound, "")
}

func (handler *SavingServiceHandler) UpdateBalance(ctx context.Context, req *pb.WithdrawalRequest) (*pb.SavingAccount, error) {
	log.Printf("Updating balance for accountID %v", req.AccountId)
	acc, exists := handler.accountMap[req.AccountId]
	if exists {
		// Check logic done by mid_saving
		// Only update balance
		updatedAcc := &pb.SavingAccount{
			Id:          acc.Id,
			UserID:      acc.UserID,
			Balance:     acc.Balance - req.Amount,
			TermType:    acc.TermType,
			Term:        acc.Term,
			CreatedDate: acc.CreatedDate,
			DueDate:     acc.DueDate,
			Rate:        acc.Rate,
		}
		handler.accountMap[req.AccountId] = updatedAcc
		return updatedAcc, status.New(codes.OK, "").Err()
	}
	return nil, status.Errorf(codes.NotFound, "Account %v does not exist.", req.AccountId)
}
