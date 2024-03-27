package main

import (
	"bytes"
	"context"
	"database/sql"
	"encoding/json"
	"github.com/elastic/go-elasticsearch/v7"
	"github.com/elastic/go-elasticsearch/v7/esapi"
	"strings"

	_ "github.com/go-sql-driver/mysql"
	pb "github.com/nghiaump/SavingAssignmentZTF/protobuf"
	//"github.com/pingcap/tidb/store/tikv"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"log"
	"net"
)

const ESDocumentTag = "es"
const ESSavingIndex = "saving"
const KafkaTopicSavingAccount = "NewSavingAccountCreated"

type SavingServiceHandler struct {
	accountMap map[string]*pb.SavingAccount
	esClient   *elasticsearch.Client
	db         *sql.DB
}

func NewSavingServiceHandler(esClient *elasticsearch.Client, db *sql.DB) *SavingServiceHandler {
	handler := SavingServiceHandler{}
	handler.esClient = esClient
	handler.db = db
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
	// ElasticSearch
	indexReq := CreateIndexingRequest(req, ESSavingIndex)
	indexRes, err := indexReq.Do(context.Background(), handler.esClient)
	if err != nil {
		log.Printf("OpenSavingAccount() - Error indexing document: %v\n", err)
	}
	defer indexRes.Body.Close()
	log.Printf("OpenSavingAccount() - Indexed new Saving Account to ElasticSearch %v\n", indexRes)

	// MySQL
	newAccUser := &AccountUser{
		accountID: req.Id,
		userID:    req.UserId,
	}
	log.Printf("OpenSavingAccount() - Write account-user to MySQL database:\n%v", newAccUser)
	errSQL := handler.SQLCreateAccountUser(newAccUser)
	if errSQL != nil {
		log.Printf("OpenSavingAccount() - Error writing to MySQL database: %v\n", errSQL)
	} else {
		log.Printf("OpenSavingAccount() - Write new account-user to MySQL database successfully\n")
	}

	// Produce message to Kafka
	//errKafka := ProduceNewSavingAccountMessage(req)
	//if errKafka != nil {
	//	log.Println("Error producing Kafka message")
	//} else {
	//	log.Println("Produced new message to Kafka")
	//}

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
	newBalance := acc.Balance - req.Amount

	// Rut het tai khoan
	if newBalance == 0 {
		deleteReq := esapi.DeleteRequest{
			Index:      ESSavingIndex,
			DocumentID: docID,
		}

		_, err := deleteReq.Do(context.Background(), handler.esClient)
		if err != nil {
			log.Fatalf("Error deleting account: %s", err)
		}
		log.Printf("Deleted Saving Account from ElasticSearch %v\n", docID)
		return acc, status.New(codes.OK, "").Err()
	}
	updateData := map[string]interface{}{
		"doc": map[string]interface{}{
			"balance": newBalance, // Giá trị mới của trường balance
		},
	}

	// Nguoc lai, cap nhat so du
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
	log.Println(strings.Repeat("=", 37))
	log.Printf("SearchAccountsByFilters %v", req)
	accList, totalHits, totalBalance := SearchAccountsByFiltersWithPaging(req, handler.esClient)

	return &pb.SavingAccountList{
		AccList:         accList,
		PageSize:        req.PageSize,
		PageIndex:       req.PageIndex,
		AggTotalHits:    totalHits,
		AggTotalBalance: totalBalance,
	}, nil
}

func (handler *SavingServiceHandler) SearchUserByNumberAccountRange(ctx context.Context, req *pb.NumberAccountRange) (*pb.ListUserWithAccounts, error) {
	result, err := handler.GetUserHavingAccountNumber(int(req.MinNumber), int(req.MaxNumber))
	if err != nil {
		return nil, nil
	}

	var userGroup []*pb.UserWithAccounts
	for userID, accountIDs := range result {
		obj := &pb.UserWithAccounts{
			UserId:     userID,
			AccountIds: accountIDs,
		}
		userGroup = append(userGroup, obj)
	}

	return &pb.ListUserWithAccounts{
		UserGroup: userGroup,
	}, nil
}
