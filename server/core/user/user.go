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
	"math/rand"
	"net"
	"reflect"
	"strings"
)

const ESDocumentTag = "es"
const ESUserIndex = "user"

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

	newUser := FillUserFromRegisterRequest(req, out.String())

	if handler.kycMap == nil {
		handler.kycMap = make(map[string]*pb.KYC)
	}
	handler.kycMap[out.String()] = &pb.KYC{
		UserId: out.String(),
		Level:  2,
	}

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

	// Chuyển đổi map thành chuỗi JSON
	jsonStr, err := json.Marshal(doc)
	if err != nil {
		// TODO
	}

	log.Printf("Test json marshal %v", string(jsonStr))
	indexReq := esapi.IndexRequest{
		Index:   ESUserIndex,
		Body:    strings.NewReader(string(jsonStr)),
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

func FillUserFromRegisterRequest(req *pb.RegisterUserRequest, id string) *pb.User {
	registeredDate, _ := ConvertToISO8601("01012024")
	return &pb.User{
		Id:             id,
		IdCardNumber:   req.IdCardNumber,
		UserName:       req.UserName,
		KycLevel:       2,              //TODO
		RegisteredDate: registeredDate, //TODO
	}
}

func (handler *UserServiceHandler) GetCurrentKYC(ctx context.Context, req *pb.GetCurrentKYCRequest) (*pb.GetCurrentKYCResponse, error) {
	log.Printf("Calling GetCurrentKYC for userID: %v", req.UserId)

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

func (handler *UserServiceHandler) SearchUserByFilter(ctx context.Context, req *pb.UserFilter) (*pb.UserList, error) {
	// TODO
	return nil, nil
}

func (handler *UserServiceHandler) SearchUserByIdCardNumber(ctx context.Context, req *pb.IDCardNumber) (*pb.User, error) {
	// TODO
	return nil, nil
}

func GenKYC(req *pb.GetCurrentKYCRequest) int32 {

	return int32(rand.Intn(3) + 1)
}
func GenKYCDefault() int32 {

	return int32(rand.Intn(3) + 1)
}
