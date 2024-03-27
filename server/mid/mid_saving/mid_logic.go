package main

import (
	"context"
	"errors"
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

func (handler *MidServiceHandler) Login(ctx context.Context, req *pb.LoginRequest) (*pb.JWTToken, error) {
	log.Printf("Login() %v\n", req)
	res, _ := handler.userServiceClient.Login(ctx, req)

	log.Println(res)
	if res.Token == "OK" {
		tokenString, err := CreateTokenString(req)
		if err == nil {
			log.Println("Authenticated successfully!")
			log.Printf("Return tokenstring OK: %v", tokenString)
			return &pb.JWTToken{
				Token: tokenString,
			}, nil
		}
	}
	return &pb.JWTToken{Token: ""}, nil
}

func (handler *MidServiceHandler) RegisterUser(ctx context.Context, req *pb.RegisterUserRequest) (*pb.RegisterUserResponse, error) {
	log.Println("RegisterUser()")

	tokenString := GetTokenFromContext(ctx)
	_, err := ValidateJWTToken(tokenString)
	if err != nil {
		log.Println("Invalid JWT token")
		return nil, errors.New("Invalid JWT token")
	}

	log.Printf("%v", req)
	res, err := handler.userServiceClient.RegisterUser(ctx, req)

	if err != nil {
		log.Printf("Could not register new user: %v", err)
		return res, err
	}

	log.Printf("User ID: %v registered successfully", res.UserId)
	return res, status.New(codes.OK, "").Err()
}

func (handler *MidServiceHandler) GetCurrentKYC(ctx context.Context, req *pb.GetCurrentKYCRequest) (*pb.GetCurrentKYCResponse, error) {
	log.Println("GetCurrentKYC()")
	res, err := handler.userServiceClient.GetCurrentKYC(ctx, req)

	if err != nil {
		log.Printf("Could not Get KYC level: %v", err)
		return nil, err
	}

	log.Printf("User ID: %v, KYC level: %v", res.UserId, res.KycLevel)
	return res, status.New(codes.OK, "").Err()
}

func (handler *MidServiceHandler) OpenSavingsAccount(ctx context.Context, req *pb.OpenSavingsAccountRequest) (*pb.OpenSavingsAccountResponse, error) {
	log.Println("OpenSavingAccount()")
	user, err := handler.userServiceClient.SearchUserByIdCardNumber(ctx, &pb.IDCardNumber{
		IdCardNumber: req.IdCardNumber,
	})

	if user == nil || err != nil {
		log.Printf("User not found: %v", err)
		return &pb.OpenSavingsAccountResponse{
			Success: false,
		}, status.New(codes.NotFound, "User not found").Err()
	}

	if user.KycLevel <= 1 {
		log.Printf("No permission for KYC level %v", user.KycLevel)
		return &pb.OpenSavingsAccountResponse{
			Success: false,
			UserId:  req.UserId,
		}, status.Error(codes.PermissionDenied, "KYC level < 2")
	}

	dueDate := CalculateDueDate(req.TermType, int(req.Term), req.CreatedDate)
	interestRate := FindFixedInterestRate(req.TermType, req.Term, user.KycLevel)
	expectedInterest := int64(float64(req.Balance) * float64(CalculateOnTimeInterest(req.TermType, req.Term, interestRate)))

	// Convert to ISO 8601 for indexing in elasticsearch
	dueDate, _ = ConvertToISO8601(dueDate)
	createdDate, _ := ConvertToISO8601(req.CreatedDate)

	savingAcc := &pb.SavingAccount{
		Id:          "",
		UserId:      req.UserId,
		Balance:     req.Balance,
		TermType:    req.TermType,
		Term:        req.Term,
		CreatedDate: createdDate,
		DueDate:     dueDate,
		Rate:        interestRate,
		Kyc:         user.KycLevel,
	}

	switch savingAcc.TermType {
	case "DAYS":
		savingAcc.TermInDays = savingAcc.Term
	case "MONTHS":
		savingAcc.TermInDays = savingAcc.Term * 30
	case "YEARS":
		savingAcc.TermInDays = savingAcc.Term * 360
	}

	log.Println("Calling Saving service for OpenSavingAccount()")
	accRes, errOpenSaving := handler.savingServiceClient.OpenSavingsAccount(ctx, savingAcc)

	if errOpenSaving != nil {
		log.Println("Cannot create new SavingAccount")
		return &pb.OpenSavingsAccountResponse{Success: false}, status.Error(codes.Internal, "Failed to create new SavingAccount")
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
	log.Println("Withdrawal()")
	log.Println("Calling Saving service for AccountInquiry()")
	accReq := &pb.AccountInquiryRequest{
		UserId:    req.UserId,
		AccountId: req.AccountId,
	}

	accRes, accErr := handler.AccountInquiry(ctx, accReq)
	if accErr != nil {
		log.Printf("Cannot inquire the account for withdrawal, error %v", accErr)
		return nil, status.Error(codes.Internal, "Cannot inquire the account")
	}

	var withdrawalLimit int64
	kyc := accRes.Kyc
	if kyc >= 3 {
		withdrawalLimit = WithdrawalLimitKYC3
	} else {
		withdrawalLimit = WithdrawalLimitKYC2
	}

	// Validated amount by limit
	if req.Amount > withdrawalLimit {
		return &pb.WithdrawalResponse{
			Success:         false,
			Acc:             nil,
			WithdrawnAmount: 0,
		}, status.Errorf(codes.Canceled, "withdrawal amount exceeded limit: %v", withdrawalLimit)
	}

	// Validate amount by balance
	if accRes.Balance >= req.Amount {
		updatedAcc, errUpdate := handler.savingServiceClient.UpdateBalance(ctx, req)

		if errUpdate != nil {
			log.Printf("error occurs when update balance in Saving service: %v", errUpdate)
			return nil, status.Error(codes.Internal, "Failed to update new balance")
		}

		// Use strategy for calculate the interest
		accPT := &SavingAccountPT{}
		accPT.ParseFrom(accRes)
		accPT.GetCalculator(req.Date) // Using Strategy Pattern: Early or OnTime

		log.Printf("Debug accPT: %v\n", accPT)
		totalWithdrawnAmount := int64(float64(req.Amount) * (accPT.CalculateRate(req.Date) + 1))

		log.Println("Withdrawal successfully")
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
	log.Println("AccountInquiry()")
	resAcc, _ := handler.savingServiceClient.SearchAccountByID(ctx, &pb.AccID{
		Id: req.AccountId,
	})

	if resAcc == nil {
		return nil, status.Error(codes.NotFound, "")
	} else if resAcc.UserId != req.UserId {
		return nil, status.Error(codes.PermissionDenied, "")
	} else {
		return resAcc, nil
	}
}

// SEARCH ACCOUNT=====

func (handler *MidServiceHandler) SearchAccountsByUserID(ctx context.Context, req *pb.AccountInquiryRequest) (*pb.SavingAccountList, error) {
	log.Printf("SearchAccountsByUserID(): %v", req.UserId)
	savingAccList, _ := handler.savingServiceClient.SearchAccountsByUserID(ctx, req)

	log.Printf("Result received from core_saving: %v\n", len(savingAccList.AccList))
	for _, acc := range savingAccList.AccList {
		log.Println(acc.Id)
	}
	return savingAccList, nil
}

func (handler *MidServiceHandler) SearchAccountsByFilter(ctx context.Context, req *pb.Filter) (*pb.SavingAccountList, error) {
	log.Printf("Calling SearchAccountsByFilters with request %v", req)

	// Validate the filter request
	valid := ValidateFilterRequest(req)
	if !valid {
		log.Println("Invalid filters")
		return nil, status.Error(codes.InvalidArgument, "")
	}

	savingAccList, _ := handler.savingServiceClient.SearchAccountsByFilter(ctx, req)

	log.Printf("Result received from core_saving: %v\n", len(savingAccList.AccList))
	for _, acc := range savingAccList.AccList {
		log.Println(acc.Id)
	}
	return savingAccList, nil
}

func (handler *MidServiceHandler) SearchAccountsByIDCardNumber(ctx context.Context, req *pb.IDCardNumber) (*pb.SavingAccountList, error) {
	log.Printf("Calling SearchAccountByIDCardNumber %v", req.IdCardNumber)
	user, _ := handler.SearchUserByIdCardNumber(ctx, req)
	if user == nil {
		return nil, nil
	}

	accList, _ := handler.SearchAccountsByUserID(ctx, &pb.AccountInquiryRequest{
		UserId:    user.Id,
		AccountId: "",
	})

	return accList, nil
}

func (handler *MidServiceHandler) SearchAccountByID(ctx context.Context, req *pb.AccID) (*pb.SavingAccount, error) {
	log.Printf("Search Account by accountID: %v", req.Id)
	resAcc, err := handler.savingServiceClient.SearchAccountByID(ctx, req)
	return resAcc, err
}

// SEARCH USER ========

func (handler *MidServiceHandler) SearchUserByNumberAccountRange(ctx context.Context, req *pb.NumberAccountRange) (*pb.ListUserWithAccounts, error) {
	log.Printf("SearchUserByNumberAccountRange()")
	resList, err := handler.savingServiceClient.SearchUserByNumberAccountRange(ctx, req)
	return resList, err
}

func (handler *MidServiceHandler) SearchUserByID(ctx context.Context, req *pb.UserID) (*pb.User, error) {
	log.Printf("Search User by userID: %v", req.Id)
	resUser, err := handler.userServiceClient.SearchUserByID(ctx, req)
	return resUser, err
}

func (handler *MidServiceHandler) SearchUserByIdCardNumber(ctx context.Context, req *pb.IDCardNumber) (*pb.User, error) {
	user, _ := handler.userServiceClient.SearchUserByIdCardNumber(ctx, req)
	return user, nil
}

func (handler *MidServiceHandler) SearchUserByAccountID(ctx context.Context, req *pb.AccountID) (*pb.User, error) {
	acc, err := handler.savingServiceClient.SearchAccountByID(ctx, &pb.AccID{
		Id: req.AccountID,
	})
	if acc == nil || err != nil {
		return nil, err
	} else {
		user, err := handler.userServiceClient.SearchUserByID(ctx, &pb.UserID{
			Id: acc.UserId,
		})

		if user == nil || err != nil {
			return nil, err
		} else {
			return user, nil
		}
	}
}

func (handler *MidServiceHandler) SearchUsersByFilter(ctx context.Context, req *pb.UserFilter) (*pb.UserList, error) {
	users, _ := handler.userServiceClient.SearchUserByFilter(ctx, req)
	return users, nil
}
