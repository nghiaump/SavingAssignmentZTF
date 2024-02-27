package main

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/elastic/go-elasticsearch/v7"
	"github.com/elastic/go-elasticsearch/v7/esapi"
	"github.com/google/uuid"
	pb "github.com/nghiaump/SavingAssignmentZTF/protobuf"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"log"
	"net"
	"reflect"
	"strings"
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
	out, err := uuid.NewUUID()
	if err != nil {
		return nil, status.Errorf(codes.Internal,
			"Error while generating AccountID", err)
	}
	log.Printf("Calling OpenSavingAccount(), userID: %v, accountID: %v", req.UserId, out.String())

	req.Id = out.String()
	val := reflect.ValueOf(req)
	if val.Kind() == reflect.Ptr {
		val = val.Elem()
	}

	if val.Kind() != reflect.Struct {
		fmt.Println("Error: val.Kind() != reflect.Struct")
	}

	doc := make(map[string]interface{})

	for i := 0; i < val.NumField(); i++ {
		field := val.Type().Field(i)
		fieldName := field.Tag.Get(ESDocumentTag)

		if fieldName != "" {
			doc[fieldName] = val.Field(i).Interface()
		}
	}

	// Chuyển đổi map thành chuỗi JSON
	jsonStr, err := json.Marshal(doc)
	if err != nil {
		// TODO
	}

	log.Printf("Test json marshal %v", string(jsonStr))
	indexReq := esapi.IndexRequest{
		Index:   ESSavingIndex,
		Body:    strings.NewReader(string(jsonStr)),
		Refresh: "true",
	}

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
	//_ := []pb.SavingAccount{}
	accList := GetAllAccountsByUserIDHelper(req.UserId, handler.esClient)
	// Print the ID and document source for each hit.

	return &pb.SavingAccountList{
		AccList: accList,
	}, nil
}
