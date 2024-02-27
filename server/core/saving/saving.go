package main

import (
	"context"
	"github.com/elastic/go-elasticsearch/v7"
	pb "github.com/nghiaump/SavingAssignmentZTF/protobuf"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"log"
	"net"
)

const ESDocumentTag = "es"
const ESSavingIndex = "saving"

type SavingServiceHandler struct {
	accountMap map[string]*pb.SavingAccount
	esClient   *elasticsearch.Client
}

func NewSavingServiceHandler(client *elasticsearch.Client) *SavingServiceHandler {
	handler := SavingServiceHandler{}
	handler.esClient = client
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
	indexReq := CreateIndexingRequest(req)

	indexRes, err2 := indexReq.Do(context.Background(), handler.esClient)
	if err2 != nil {
		log.Printf("Error indexing document: %v\n", err2)
	}
	defer indexRes.Body.Close()
	log.Printf("Indexed new Saving Account to ElasticSearch %v\n", indexRes)
	return req, status.New(codes.OK, "").Err()
}

func (handler *SavingServiceHandler) AccountInquiry(ctx context.Context, req *pb.AccountInquiryRequest) (*pb.SavingAccount, error) {
	log.Printf("Calling AccoutAquiry(), userID: %v, accountID: %v", req.UserId, req.AccountId)

	acc, exists := handler.accountMap[req.AccountId]
	if exists {
		if req.UserId == acc.UserId {
			log.Printf("Account %v exist", req.AccountId)
			return &pb.SavingAccount{
				Id:          req.AccountId,
				UserId:      req.UserId,
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
			UserId:      acc.UserId,
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

func (handler *SavingServiceHandler) GetAllAccountsByUserID(ctx context.Context, req *pb.AccountInquiryRequest) (*pb.SavingAccountList, error) {
	log.Printf("Get all accounts by UserID %v", req.UserId)
	accList := GetAllAccountsByUserIDHelper(req.UserId, handler.esClient)

	return &pb.SavingAccountList{
		AccList: accList,
	}, nil
}

func (handler *SavingServiceHandler) SearchAccountsByFilter(ctx context.Context, req *pb.Filter) (*pb.SavingAccountList, error) {
	log.Printf("SearchAccountsByFilters %v", req)
	accList := SearchAccountsByFiltersHelper(req, handler.esClient)

	return &pb.SavingAccountList{
		AccList: accList,
	}, nil
}
