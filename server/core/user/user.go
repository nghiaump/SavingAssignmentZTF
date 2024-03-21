package main

import (
	"context"
	"database/sql"
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
	"math/rand"
	"net"
	"reflect"
	"strings"
)

const ESDocumentTag = "es"
const ESUserIndex = "user"

type UserServiceHandler struct {
	usersMap map[string]*pb.User
	esClient *elasticsearch.Client
	db       *sql.DB
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
	existed := handler.CheckExistingUser(ctx, req.IdCardNumber)

	if existed {
		return &pb.RegisterUserResponse{Success: false, UserId: ""}, status.Error(codes.AlreadyExists, "")
	}

	out, _ := uuid.NewUUID()
	log.Printf("Calling RegisterUser(): %v", out.String())
	newUser := FillUserFromRegisterRequest(req, out.String())

	doc := CreateDocument(newUser)
	indexReq := CreateIndexRequest(ESUserIndex, doc)

	indexRes, err2 := indexReq.Do(context.Background(), handler.esClient)
	if err2 != nil {
		log.Printf("Error indexing document: %v\n", err2)
	}
	defer indexRes.Body.Close()
	log.Printf("Indexed new User to ElasticSearch %v\n", indexRes)
	//
	return &pb.RegisterUserResponse{Success: true, UserId: req.UserName}, status.New(codes.OK, "").Err()
}

func CreateIndexRequest(indexName string, doc map[string]interface{}) esapi.IndexRequest {
	jsonStr, err := json.Marshal(doc)
	if err != nil {
		// TODO
	}

	log.Printf("Test json marshal %v", string(jsonStr))
	indexReq := esapi.IndexRequest{
		Index:   indexName,
		Body:    strings.NewReader(string(jsonStr)),
		Refresh: "true",
	}
	return indexReq
}

func CreateDocument(newUser *pb.User) map[string]interface{} {
	val := reflect.ValueOf(newUser)
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
	return doc
}

func (handler *UserServiceHandler) CheckExistingUser(ctx context.Context, IDCardNumber string) bool {
	resUser, _ := handler.SearchUserByIdCardNumber(ctx, &pb.IDCardNumber{
		IdCardNumber: IDCardNumber,
	})

	log.Printf("check existing user %v", resUser)
	if resUser != nil {
		return true
	}
	return false
}

func FillUserFromRegisterRequest(req *pb.RegisterUserRequest, id string) *pb.User {
	registeredDate, _ := ConvertToISO8601("01012024")
	dob, _ := ConvertToISO8601(req.Dob)
	return &pb.User{
		//Id:             id,
		Id:             req.UserName,
		IdCardNumber:   req.IdCardNumber,
		UserName:       req.UserName,
		KycLevel:       GenKYCDefault(),
		RegisteredDate: registeredDate,
		Dob:            dob,
		Address:        req.Address,
		PhoneNumber:    req.PhoneNumber,
	}
}

func (handler *UserServiceHandler) GetCurrentKYC(ctx context.Context, req *pb.GetCurrentKYCRequest) (*pb.GetCurrentKYCResponse, error) {
	log.Printf("Calling GetCurrentKYC for userID: %v", req.UserId)
	user, _ := handler.SearchUserByIdCardNumber(ctx, &pb.IDCardNumber{
		IdCardNumber: req.IdCardNumber,
	})
	if user == nil {
		user, err := handler.SearchUserByID(ctx, &pb.UserID{Id: req.UserId})
		if user == nil {
			return nil, err
		} else {
			return &pb.GetCurrentKYCResponse{
				UserId:       req.UserId,
				IdCardNumber: req.IdCardNumber,
				KycLevel:     user.KycLevel,
			}, nil
		}

	} else {
		return &pb.GetCurrentKYCResponse{
			UserId:       req.UserId,
			IdCardNumber: req.IdCardNumber,
			KycLevel:     user.KycLevel,
		}, nil
	}

}

func (handler *UserServiceHandler) SearchUserByID(ctx context.Context, req *pb.UserID) (*pb.User, error) {
	log.Printf("Search User by userID: %v", req.Id)
	resUser := SearchOneUserByUniqueTextField("id", req.Id, handler.esClient)
	if resUser == nil {
		return nil, status.Error(codes.NotFound, "")
	} else {
		return resUser, nil
	}
}

func (handler *UserServiceHandler) SearchUserByIdCardNumber(ctx context.Context, req *pb.IDCardNumber) (*pb.User, error) {
	log.Printf("Search User by ID Card Number %v", req.IdCardNumber)
	resUser := SearchOneUserByUniqueTextField("id_card_number", req.IdCardNumber, handler.esClient)
	if resUser == nil {
		return nil, status.Error(codes.NotFound, "")
	}
	return resUser, nil
}

func (handler *UserServiceHandler) SearchUserByAccountID(ctx context.Context, req *pb.AccountID) (*pb.User, error) {
	log.Printf("Search User by AccountID %v", req.AccountID)
	resUser := SearchOneUserByUniqueTextField("account_id", req.AccountID, handler.esClient)
	if resUser == nil {
		return nil, status.Error(codes.NotFound, "")
	}
	return resUser, nil
}

func (handler *UserServiceHandler) SearchUserByFilter(ctx context.Context, req *pb.UserFilter) (*pb.UserList, error) {
	log.Printf("SearchUsersByFilters %v", req)
	users := SearchUsersByFiltersHelper(req, handler.esClient)
	if len(users) < 1 {
		return &pb.UserList{
			UserList: users,
		}, status.Error(codes.NotFound, "")
	}
	return &pb.UserList{
		UserList: users,
	}, nil

}

func GenKYCDefault() int32 {

	return int32(rand.Intn(3) + 1)
}
