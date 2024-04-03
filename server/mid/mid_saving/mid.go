package main

import (
	"context"
	"github.com/confluentinc/confluent-kafka-go/v2/kafka"
	"github.com/golang/glog"
	pb "github.com/nghiaump/SavingAssignmentZTF/protobuf"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"net"
)

const (
	WithdrawalLimitKYC3 = 150000000
	WithdrawalLimitKYC2 = 100000000
)

type MidServiceHandler struct {
	userServiceClient   pb.UserServiceClient
	savingServiceClient pb.SavingServiceClient
	kafkaProducer       *kafka.Producer
}

func CreateMidServiceHandler(userServiceClient pb.UserServiceClient, savingServiceClient pb.SavingServiceClient, producer *kafka.Producer) *MidServiceHandler {
	midServiceHandler := MidServiceHandler{
		userServiceClient:   userServiceClient,
		savingServiceClient: savingServiceClient,
		kafkaProducer:       producer,
	}
	return &midServiceHandler
}

func StartMidServer(midHandler *MidServiceHandler, port string) {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		glog.Fatalf("StartMidServer: failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterMidSavingServiceServer(s, midHandler)
	glog.Info("StartMidServer: listening on port " + port)
	if err := s.Serve(lis); err != nil {
		glog.Fatalf("StartMidServer: failed to serve: %v", err)
	}
}

func (handler *MidServiceHandler) GetJWT(ctx context.Context, req *pb.LoginRequest) (*pb.JWT, error) {
	glog.Infof("Login: username:%v", req.Username)
	res, _ := handler.userServiceClient.GetJWT(ctx, req)

	if res.Token == "OK" {
		tokenString, err := CreateTokenString(req)
		if err == nil {
			glog.Infof("Login: Authenticated successfully, tokenstring: %v", tokenString)
			return &pb.JWT{
				Token: tokenString,
			}, nil
		}
	}

	glog.Info("Login: Failed to Authenticate!")
	return &pb.JWT{Token: ""}, nil
}

func (handler *MidServiceHandler) CreateUser(ctx context.Context, req *pb.CreateUserRequest) (*pb.CreateUserResponse, error) {
	glog.Info("RegisterUser")

	authErr := CheckJWTFromContext(ctx)
	if authErr != nil {
		return nil, authErr
	}

	glog.Infof("%v", req)
	res, err := handler.userServiceClient.CreateUser(ctx, req)

	if err != nil {
		glog.Infof("RegisterUser: Could not register new user: %v", err)
		return res, err
	}

	glog.Infof("RegisterUser:%v registered successfully", res.UserId)
	return res, status.New(codes.OK, "").Err()
}

func (handler *MidServiceHandler) GetCurrentKYC(ctx context.Context, req *pb.GetCurrentKYCRequest) (*pb.GetCurrentKYCResponse, error) {
	glog.Info("GetCurrentKYC:")
	res, err := handler.userServiceClient.GetCurrentKYC(ctx, req)

	if err != nil {
		glog.Infof("GetCurrentKYC: Get KYC level: %v", err)
		return nil, err
	}
	glog.Infof("GetCurrentKYC: User ID: %v, KYC level: %v", res.UserId, res.KycLevel)
	return res, status.New(codes.OK, "").Err()
}

func (handler *MidServiceHandler) CreateSavingsAccount(ctx context.Context, req *pb.OpenSavingsAccountRequest) (*pb.OpenSavingsAccountResponse, error) {
	glog.Info("OpenSavingAccount:")

	// authentication check
	authErr := CheckJWTFromContext(ctx)
	if authErr != nil {
		return nil, authErr
	}

	user, err := handler.userServiceClient.GetUserByIdCardNumber(ctx, &pb.IDCardNumber{
		IdCardNumber: req.IdCardNumber,
	})

	if user == nil || err != nil {
		glog.Infof("OpenSavingAccount: User not found: %v", err)
		return &pb.OpenSavingsAccountResponse{
			Success: false,
		}, status.New(codes.NotFound, "OpenSavingAccount: User not found").Err()
	}

	if user.KycLevel <= 1 {
		glog.Infof("OpenSavingAccount: No permission for KYC level %v", user.KycLevel)
		return &pb.OpenSavingsAccountResponse{
			Success: false,
			UserId:  req.UserId,
		}, status.Error(codes.PermissionDenied, "OpenSavingAccount: KYC level < 2")
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

	accRes, errOpenSaving := handler.savingServiceClient.CreateSavingsAccount(ctx, savingAcc)

	if errOpenSaving != nil {
		glog.Info("OpenSavingAccount: failed to call core saving")
		return &pb.OpenSavingsAccountResponse{Success: false}, status.Error(codes.Internal, "Failed to create new SavingAccount")
	}

	glog.Infof("OpenSavingAccount: success: \nUserID: %v, Account Detail: %v", req.UserId, accRes)

	// Kafka
	glog.Info("OpenSavingAccount: Producing Kafka message:")
	handler.ProduceNewMessage(accRes)

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

func (handler *MidServiceHandler) UpdateAccountByWithdrawal(ctx context.Context, req *pb.WithdrawalRequest) (*pb.WithdrawalResponse, error) {
	glog.Info("Withdrawal")

	// authentication check
	authErr := CheckJWTFromContext(ctx)
	if authErr != nil {
		return nil, authErr
	}

	accReq := &pb.AccountInquiryRequest{
		UserId:    req.UserId,
		AccountId: req.AccountId,
	}

	accRes, accErr := handler.GetAccount(ctx, accReq)
	if accErr != nil {
		glog.Infof("Withdrawal: Cannot inquire the account %v", accErr)
		return nil, status.Error(codes.Internal, "Withdrawal: Cannot inquire the account")
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
		}, status.Errorf(codes.Canceled, "Withdrawal: amount exceeded limit: %v", withdrawalLimit)
	}

	// Validate amount by balance
	if accRes.Balance >= req.Amount {
		// Use strategy for calculate the interest
		accPT := &SavingAccountPT{}
		accPT.ParseFrom(accRes)
		accPT.GetCalculator(req.Date) // Using Strategy Pattern: Early or OnTime

		// Check valid withdrawal date
		if accPT.Calculator == nil {
			glog.Info("Withdrawal: invalid date")
			return nil, status.Error(codes.PermissionDenied, "Withdrawal: invalid date")
		}
		totalWithdrawnAmount := int64(float64(req.Amount) * (accPT.CalculateRate(req.Date) + 1))

		// Update to core service
		updatedAcc, errUpdate := handler.savingServiceClient.UpdateAccountByWithdrawal(ctx, req)
		if errUpdate != nil {
			glog.Infof("Withdrawal: failed to update balance in core saving: %v", errUpdate)
			return nil, status.Error(codes.Internal, "Withdrawal: failed to update balance in core saving")
		}

		glog.Info("Withdrawal: success")
		return &pb.WithdrawalResponse{
			Success:         true,
			Acc:             updatedAcc,
			WithdrawnAmount: totalWithdrawnAmount,
		}, status.New(codes.OK, "").Err()

	} else {
		// No money enough
		return &pb.WithdrawalResponse{
			Success:         false,
			Acc:             accRes,
			WithdrawnAmount: 0,
		}, status.Errorf(codes.Aborted, "Withdrawal: Not enough balance. Remain balance: %v", accRes.Balance)
	}
}

func (handler *MidServiceHandler) GetAccount(ctx context.Context, req *pb.AccountInquiryRequest) (*pb.SavingAccount, error) {
	glog.Info("AccountInquiry")

	// authentication check
	authErr := CheckJWTFromContext(ctx)
	if authErr != nil {
		return nil, authErr
	}

	resAcc, _ := handler.savingServiceClient.GetAccountByID(ctx, &pb.AccID{
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

func (handler *MidServiceHandler) ListAccountsByUserID(ctx context.Context, req *pb.AccountInquiryRequest) (*pb.SavingAccountList, error) {
	glog.Infof("SearchAccountsByUserID: %v", req.UserId)

	// authentication check
	authErr := CheckJWTFromContext(ctx)
	if authErr != nil {
		return nil, authErr
	}

	savingAccList, _ := handler.savingServiceClient.ListAccountsByUserID(ctx, req)
	glog.Infof("SearchAccountsByUserID: Result received from core_saving: %v\n", len(savingAccList.AccList))
	for _, acc := range savingAccList.AccList {
		glog.Info(acc.Id)
	}
	return savingAccList, nil
}

func (handler *MidServiceHandler) ListAccountsByFilter(ctx context.Context, req *pb.Filter) (*pb.SavingAccountList, error) {
	glog.Infof("SearchAccountsByFilter: received request %v", req)

	// authentication check
	authErr := CheckJWTFromContext(ctx)
	if authErr != nil {
		return nil, authErr
	}

	// Validate the filter request
	valid := ValidateFilterRequest(req)
	if !valid {
		glog.Info("SearchAccountsByFilter: Invalid filters")
		return nil, status.Error(codes.InvalidArgument, "SearchAccountsByFilter: Invalid filters")
	}

	savingAccList, _ := handler.savingServiceClient.ListAccountsByFilter(ctx, req)

	glog.Infof("SearchAccountsByFilter: Result received from core_saving: %v\n", len(savingAccList.AccList))
	for _, acc := range savingAccList.AccList {
		glog.Info(acc.Id)
	}
	return savingAccList, nil
}

func (handler *MidServiceHandler) ListAccountsByIDCardNumber(ctx context.Context, req *pb.IDCardNumber) (*pb.SavingAccountList, error) {
	glog.Infof("SearchAccountByIDCardNumber: %v", req.IdCardNumber)

	// authentication check
	authErr := CheckJWTFromContext(ctx)
	if authErr != nil {
		return nil, authErr
	}

	user, _ := handler.GetUserByIdCardNumber(ctx, req)
	if user == nil {
		return nil, nil
	}
	accList, _ := handler.ListAccountsByUserID(ctx, &pb.AccountInquiryRequest{
		UserId:    user.Id,
		AccountId: "",
	})

	return accList, nil
}

func (handler *MidServiceHandler) GetAccountByID(ctx context.Context, req *pb.AccID) (*pb.SavingAccount, error) {
	glog.Infof("SearchAccountByID: %v", req.Id)

	// authentication check
	authErr := CheckJWTFromContext(ctx)
	if authErr != nil {
		return nil, authErr
	}

	resAcc, err := handler.savingServiceClient.GetAccountByID(ctx, req)
	return resAcc, err
}

// SEARCH USER ========

func (handler *MidServiceHandler) ListUserByNumberAccountRange(ctx context.Context, req *pb.NumberAccountRange) (*pb.ListUserWithAccounts, error) {
	glog.Info("SearchUserByNumberAccountRange")

	// authentication check
	authErr := CheckJWTFromContext(ctx)
	if authErr != nil {
		return nil, authErr
	}

	resList, err := handler.savingServiceClient.ListUsersByNumberAccountRange(ctx, req)
	return resList, err
}

func (handler *MidServiceHandler) GetUserByID(ctx context.Context, req *pb.UserID) (*pb.User, error) {
	glog.Infof("SearchUserByID: %v", req.Id)

	// authentication check
	authErr := CheckJWTFromContext(ctx)
	if authErr != nil {
		return nil, authErr
	}

	resUser, err := handler.userServiceClient.GetUserByID(ctx, req)
	return resUser, err
}

func (handler *MidServiceHandler) GetUserByIdCardNumber(ctx context.Context, req *pb.IDCardNumber) (*pb.User, error) {
	glog.Infof("SearchUserByIDCardNumber: %v", req.IdCardNumber)

	// authentication check
	authErr := CheckJWTFromContext(ctx)
	if authErr != nil {
		return nil, authErr
	}

	user, _ := handler.userServiceClient.GetUserByIdCardNumber(ctx, req)
	return user, nil
}

func (handler *MidServiceHandler) GetUserByAccountID(ctx context.Context, req *pb.AccountID) (*pb.User, error) {
	// authentication check
	authErr := CheckJWTFromContext(ctx)
	if authErr != nil {
		return nil, authErr
	}

	acc, err := handler.savingServiceClient.GetAccountByID(ctx, &pb.AccID{
		Id: req.AccountID,
	})
	if acc == nil || err != nil {
		return nil, err
	} else {
		user, err := handler.userServiceClient.GetUserByID(ctx, &pb.UserID{
			Id: acc.UserId,
		})

		if user == nil || err != nil {
			return nil, err
		} else {
			return user, nil
		}
	}
}

func (handler *MidServiceHandler) ListUsersByFilter(ctx context.Context, req *pb.UserFilter) (*pb.UserList, error) {
	// authentication check
	authErr := CheckJWTFromContext(ctx)
	if authErr != nil {
		return nil, authErr
	}

	users, _ := handler.userServiceClient.ListUsersByFilter(ctx, req)
	return users, nil
}

func (handler *MidServiceHandler) UpdateUser(ctx context.Context, req *pb.User) (*pb.User, error) {
	// TODO
	return nil, nil
}
func (handler *MidServiceHandler) DeleteUser(ctx context.Context, req *pb.User) (*pb.DeleteUserResponse, error) {
	// TODO
	return &pb.DeleteUserResponse{
		Success: true,
	}, nil
}
func (handler *MidServiceHandler) DeleteAccount(ctx context.Context, req *pb.SavingAccount) (*pb.DeleteAccountResponse, error) {
	// TODO
	return &pb.DeleteAccountResponse{
		Success: true,
	}, nil
}
