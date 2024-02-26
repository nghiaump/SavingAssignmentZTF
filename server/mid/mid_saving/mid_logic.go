package main

import (
	"context"
	pb "github.com/nghiaump/SavingAssignmentZTF/protobuf"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"log"
	"net"
)

const (
	WithdrawalLimitKYC3 = 150000000
	WithdrawalLimitKYC2 = 100000000
)

type MidServiceHandler struct {
	userServiceClient   pb.UserServiceClient
	savingServiceClient pb.SavingsServiceClient
}

func CreateMidServiceHandler(userServiceClient pb.UserServiceClient, savingServiceClient pb.SavingsServiceClient) *MidServiceHandler {
	midServiceHandler := MidServiceHandler{
		userServiceClient:   userServiceClient,
		savingServiceClient: savingServiceClient,
	}
	return &midServiceHandler
}

func StartMidServer(midHandler *MidServiceHandler, port string) {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterMidSavingServiceServer(s, midHandler)
	log.Printf("Starting gRPC mid_saving service listener on port " + port)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

func (handler *MidServiceHandler) RegisterUser(ctx context.Context, req *pb.RegisterUserRequest) (*pb.RegisterUserResponse, error) {
	log.Println("Calling user service for RegisterUser()")
	log.Printf("%v", req)
	res, err := handler.userServiceClient.RegisterUser(ctx, req)

	if err != nil {
		log.Printf("Could not register new user: %v", err)
		return nil, status.Error(codes.Internal, "User register Failed")
	}

	log.Printf("User ID: %v registered successfully", res.UserId)
	return res, status.New(codes.OK, "").Err()
}

func (handler *MidServiceHandler) GetCurrentKYC(ctx context.Context, req *pb.GetCurrentKYCRequest) (*pb.GetCurrentKYCResponse, error) {
	log.Println("Calling User service for CurrentKYC()")
	res, err := handler.userServiceClient.GetCurrentKYC(ctx, req)

	if err != nil {
		log.Printf("Could not Get KYC level: %v", err)
		return nil, status.Error(codes.Internal, "Get KYC level failed")
	}

	log.Printf("User ID: %v, KYC level: %v", res.UserId, res.KycLevel)
	return res, status.New(codes.OK, "").Err()
}

func (handler *MidServiceHandler) OpenSavingsAccount(ctx context.Context, req *pb.OpenSavingsAccountRequest) (*pb.OpenSavingsAccountResponse, error) {
	log.Println("Calling User service to get KYC level")
	kycRes, err := handler.GetCurrentKYC(ctx, &pb.GetCurrentKYCRequest{
		UserId: req.UserId,
	})
	if err != nil {
		log.Printf("Cannot verify the KYC level from User Core: %v", err)
		return nil, status.Error(codes.NotFound, "KYC level verify failed")
	}

	if kycRes.KycLevel <= 1 {
		log.Printf("No permission for KYC level %v", kycRes.KycLevel)
		return &pb.OpenSavingsAccountResponse{
			Success: false,
			UserId:  req.UserId,
		}, status.Error(codes.PermissionDenied, "KYC level < 2")
	}

	dueDate := CalculateDueDate(req.TermType, int(req.Term), req.CreatedDate)
	interestRate := FindFixedInterestRate(req.TermType, req.Term, kycRes.KycLevel)
	expectedInterest := int64(float64(req.Balance) * float64(CalculateOnTimeInterest(req.TermType, req.Term, interestRate)))

	savingAcc := &pb.SavingAccount{
		Id:          "",
		UserId:      req.UserId,
		Balance:     req.Balance,
		TermType:    req.TermType,
		Term:        req.Term,
		CreatedDate: req.CreatedDate,
		DueDate:     dueDate,
		Rate:        interestRate,
	}

	log.Println("Calling Saving service for OpenSavingAccount()")
	accRes, errOpenSaving := handler.savingServiceClient.OpenSavingsAccount(ctx, savingAcc)

	if errOpenSaving != nil {
		log.Println("Cannot create new SavingAccount")
		return nil, status.Error(codes.Internal, "Failed to create new SavingAccount")
	}

	log.Printf("Created new SavingAccount successfully\nUserID: %v, Account Detail: %v", req.UserId, accRes)
	return &pb.OpenSavingsAccountResponse{
		Success:          true,
		UserId:           accRes.UserId,
		AccountId:        accRes.Id,
		Balance:          accRes.Balance,
		Rate:             accRes.Rate,
		CreatedDate:      accRes.CreatedDate,
		DueDate:          accRes.DueDate,
		ExpectedInterest: expectedInterest,
	}, status.New(codes.OK, "").Err()

}

func (handler *MidServiceHandler) Withdrawal(ctx context.Context, req *pb.WithdrawalRequest) (*pb.WithdrawalResponse, error) {
	log.Println("Calling User service for validating AccountInquiry()")
	accReq := &pb.AccountInquiryRequest{
		UserId:    req.UserId,
		AccountId: req.AccountId,
	}

	log.Println("Calling Saving service for AccountInquiry()")
	accRes, accErr := handler.AccountInquiry(ctx, accReq)
	if accErr != nil {
		log.Printf("Cannot inquire the account for withdrawal, error %v", accErr)
		return nil, status.Error(codes.Internal, "Cannot inquire the account")
	}

	var withdrawalLimit int64
	kycRes, err := handler.GetCurrentKYC(ctx, &pb.GetCurrentKYCRequest{
		UserId: req.UserId,
	})
	if err != nil {
		log.Printf("Cannot verify the KYC level from User Core: %v", err)
		return nil, status.Error(codes.NotFound, "KYC level verify failed")
	}
	if kycRes.KycLevel >= 3 {
		withdrawalLimit = WithdrawalLimitKYC3
	} else {
		withdrawalLimit = WithdrawalLimitKYC2
	}

	// Validated amount
	if req.Amount > withdrawalLimit {
		return &pb.WithdrawalResponse{
			Success:         false,
			Acc:             nil,
			WithdrawnAmount: 0,
		}, status.Errorf(codes.Canceled, "withdrawal amount exceeded limit: %v", withdrawalLimit)
	}

	if accRes.Balance >= req.Amount {
		updatedAcc, errUpdate := handler.savingServiceClient.UpdateBalance(ctx, &pb.WithdrawalRequest{
			UserId:    req.UserId,
			AccountId: req.AccountId,
			Amount:    req.Amount,
			Date:      req.Date,
		})

		if errUpdate != nil {
			log.Printf("error occurs when update balance in Saving service: %v", errUpdate)
			return nil, status.Error(codes.Internal, "Failed to update new balance")
		}

		// Use strategy for calculate the interest
		accPT := &SavingAccountPT{}
		accPT.ParseFrom(accRes)
		accPT.GetCalculator(req.Date) // Using Strategy Pattern: Early or OnTime
		totalWithdrawnAmount := int64(float64(req.Amount) * (accPT.CalculateRate(req.Date) + 1))

		log.Println("Withdrawn successfully")
		return &pb.WithdrawalResponse{
			Success:         true,
			Acc:             updatedAcc,
			WithdrawnAmount: totalWithdrawnAmount,
		}, status.New(codes.OK, "").Err()

	} else {
		// No enough money
		return &pb.WithdrawalResponse{
			Success:         false,
			Acc:             accRes,
			WithdrawnAmount: 0,
		}, status.Errorf(codes.Aborted, "Not enough balance.\nRemaining balance: %v", accRes.Balance)
	}
}

func (handler *MidServiceHandler) AccountInquiry(ctx context.Context, req *pb.AccountInquiryRequest) (*pb.SavingAccount, error) {
	log.Println("Calling Saving service for AccountInquiry()")
	res, err := handler.savingServiceClient.AccountInquiry(ctx, req)
	if err != nil {
		log.Printf("Cannot inquire the SavingAccount: %v", err)
		return nil, status.Error(codes.Internal, "Cannot inquire the account")
	}
	log.Printf("Inquired account successfully\nAccountID: %v, Detail: %v", req.AccountId, res)
	return res, status.New(codes.OK, "").Err()
}

func (handler *MidServiceHandler) GetAllAccountsByUserID(ctx context.Context, req *pb.AccountInquiryRequest) (*pb.SavingAccountList, error) {
	log.Printf("Calling GetAllAcc for userID: %v", req.UserId)
	res, _ := handler.savingServiceClient.GetAllAccountsByUserID(ctx, req)
	log.Printf("Result received from core_saving: %v\n", res)
	//for _, hit := range r["hits"].(map[string]interface{})["hits"].([]interface{}) {
	//	log.Printf(" * ID=%s, %s", hit.(map[string]interface{})["_id"], hit.(map[string]interface{})["_source"])
	//}
	return nil, nil
}
