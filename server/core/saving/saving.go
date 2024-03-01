package main

import (
	"bytes"
	"context"
	"encoding/json"
	"github.com/elastic/go-elasticsearch/v7"
	"github.com/elastic/go-elasticsearch/v7/esapi"
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
	indexReq := CreateIndexingRequest(req, ESSavingIndex)
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
	acc, _ := handler.SearchAccountByID(ctx, &pb.AccID{
		Id: req.AccountId, // validated before
	})

	if acc == nil {
		return nil, status.Error(codes.NotFound, "")
	}

	docID := SearchDocIDByUniqueTextField("id", req.AccountId, handler.esClient)
	updateData := map[string]interface{}{
		"doc": map[string]interface{}{
			"balance": acc.Balance - req.Amount, // Giá trị mới của trường balance
		},
	}

	updateBody, err := json.Marshal(updateData)
	if err != nil {
		log.Fatalf("Error marshaling update data: %s", err)
	}

	updateReq := esapi.UpdateRequest{
		Index:      ESSavingIndex,
		DocumentID: docID,
		Body:       bytes.NewReader(updateBody),
	}

	res, err := updateReq.Do(context.Background(), handler.esClient)
	if err != nil {
		log.Fatalf("Error updating document: %s", err)
	}
	defer res.Body.Close()
	log.Printf("Updated Saving Account to ElasticSearch\n")
	return acc, status.New(codes.OK, "").Err()

}

func (handler *SavingServiceHandler) SearchAccountByID(ctx context.Context, req *pb.AccID) (*pb.SavingAccount, error) {
	log.Printf("Search Account by account ID: %v", req.Id)
	resAcc := SearchOneAccountByUniqueTextField("id", req.Id, handler.esClient)
	if resAcc == nil {
		return nil, status.Errorf(codes.NotFound, "")
	} else {
		return resAcc, nil
	}
}

func (handler *SavingServiceHandler) SearchAccountsByUserID(ctx context.Context, req *pb.AccountInquiryRequest) (*pb.SavingAccountList, error) {
	log.Printf("Search accounts by UserID %v", req.UserId)
	accList := GetAllAccountsByUserIDHelper(req.UserId, handler.esClient)

	return &pb.SavingAccountList{
		AccList: accList,
	}, nil
}

func (handler *SavingServiceHandler) SearchAccountsByFilter(ctx context.Context, req *pb.Filter) (*pb.SavingAccountList, error) {
	log.Printf("SearchAccountsByFilters %v", req)
	accList, totalHits, totalBalance := SearchAccountsByFiltersWithPaginate(req, handler.esClient)

	return &pb.SavingAccountList{
		AccList:         accList,
		PageSize:        req.PageSize,
		PageIndex:       req.PageIndex,
		AggTotalHits:    totalHits,
		AggTotalBalance: totalBalance,
	}, nil
}
