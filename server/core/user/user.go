package main

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"github.com/elastic/go-elasticsearch/v7"
	"github.com/elastic/go-elasticsearch/v7/esapi"
	"github.com/golang/glog"
	pb "github.com/nghiaump/SavingAssignmentZTF/protobuf"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
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

func NewUserServiceHandler(client *elasticsearch.Client, db *sql.DB) *UserServiceHandler {
	handler := UserServiceHandler{}
	handler.esClient = client
	handler.db = db
	return &handler
}

func StartUserServer(handler *UserServiceHandler, port string) {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		glog.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	glog.Info("Creating UserServiceHandler")
	pb.RegisterUserServiceServer(s, handler)
	glog.Infof("Starting gRPC User service listener on port %v", port)
	if err := s.Serve(lis); err != nil {
		glog.Fatalf("failed to serve: %v", err)
	}
}

func (handler *UserServiceHandler) Login(ctx context.Context, req *pb.LoginRequest) (*pb.JWTToken, error) {
	// TODO
	if req.Username == "admin" && req.Password == "nguyendainghia" {
		return &pb.JWTToken{
			Token: "OK",
		}, nil
	}
	return &pb.JWTToken{
		Token: "INVALID",
	}, nil
}

func (handler *UserServiceHandler) RegisterUser(ctx context.Context, req *pb.RegisterUserRequest) (*pb.RegisterUserResponse, error) {
	existed := handler.CheckExistingUser(ctx, req.IdCardNumber)
	if existed {
		return &pb.RegisterUserResponse{Success: false, UserId: ""}, status.Error(codes.AlreadyExists, "")
	}

	glog.Infof("RegisterUser: %v", req)
	newUser := FillUserFromRegisterRequest(req)

	// ElasticSearch
	doc := CreateDocument(newUser)
	indexReq := CreateIndexRequest(ESUserIndex, doc)
	glog.Infof("RegisterUser: Indexed new User to ElasticSearch\n")
	indexRes, err2 := indexReq.Do(context.Background(), handler.esClient)
	if err2 != nil {
		glog.Infof("RegisterUser: Failed to index document: %v\n", err2)
	}
	defer indexRes.Body.Close()

	// SQL
	glog.Info("RegisterUser: Write User to MySQL database")
	errSQL := handler.SQLCreateUser(newUser)
	if errSQL != nil {
		glog.Infof("RegisterUser: Failed to write to MySQL database: %v\n", errSQL)
	} else {
		glog.Infof("RegisterUser: Write new user to MySQL database successfully\n")
	}

	return &pb.RegisterUserResponse{Success: true, UserId: req.UserName}, status.New(codes.OK, "").Err()
}

func CreateIndexRequest(indexName string, doc map[string]interface{}) esapi.IndexRequest {
	jsonStr, err := json.Marshal(doc)
	if err != nil {
		// TODO
	}

	glog.Infof("Test json marshal %v", string(jsonStr))
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

	glog.Info("check existing user %v", resUser)
	if resUser != nil {
		return true
	}
	return false
}

func FillUserFromRegisterRequest(req *pb.RegisterUserRequest) *pb.User {
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
	glog.Infof("Calling GetCurrentKYC for userID: %v", req.UserId)
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
	glog.Infof("Search User by userID: %v", req.Id)
	resUser := SearchOneUserByUniqueTextField("user_name", req.Id, handler.esClient)
	if resUser == nil {
		return nil, status.Error(codes.NotFound, "")
	} else {
		return resUser, nil
	}
}

func (handler *UserServiceHandler) SearchUserByIdCardNumber(ctx context.Context, req *pb.IDCardNumber) (*pb.User, error) {
	glog.Infof("Search User by ID Card Number %v", req.IdCardNumber)
	resUser := SearchOneUserByUniqueTextField("id_card_number", req.IdCardNumber, handler.esClient)
	if resUser == nil {
		return nil, status.Error(codes.NotFound, "")
	}
	return resUser, nil
}

func (handler *UserServiceHandler) SearchUserByAccountID(ctx context.Context, req *pb.AccountID) (*pb.User, error) {
	glog.Infof("Search User by AccountID %v", req.AccountID)
	resUser := SearchOneUserByUniqueTextField("account_id", req.AccountID, handler.esClient)
	if resUser == nil {
		return nil, status.Error(codes.NotFound, "")
	}
	return resUser, nil
}

func (handler *UserServiceHandler) SearchUserByFilter(ctx context.Context, req *pb.UserFilter) (*pb.UserList, error) {
	glog.Infof("SearchUsersByFilters %v", req)
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
