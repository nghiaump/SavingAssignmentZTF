package main

import (
	"context"
	"github.com/golang/mock/gomock"
	pb "github.com/nghiaump/SavingAssignmentZTF/protobuf"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"testing"
)

func TestRegUser(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()

	userServiceMock := pb.NewMockUserServiceClient(controller)
	mid := CreateMidServiceHandler(userServiceMock, nil)

	t.Run("Case success", func(t *testing.T) {
		userServiceMock.EXPECT().RegisterUser(gomock.Any(), gomock.Any()).Return(&pb.RegisterUserResponse{
			Success: true,
			UserId:  "abcd",
		}, nil)

		expected := pb.RegisterUserResponse{
			Success: true,
			UserId:  "abcd",
		}

		result, err := mid.RegisterUser(context.TODO(), &pb.RegisterUserRequest{
			IdCardNumber: "",
			UserName:     "",
		})

		if (result.UserId != expected.GetUserId()) || err != nil {
			t.Errorf("Expected %v but output is %v", expected, result)
		}
	})

	t.Run("Case Error", func(t *testing.T) {
		userServiceMock.EXPECT().RegisterUser(gomock.Any(), gomock.Any()).Return(&pb.RegisterUserResponse{
			Success: true,
			UserId:  "abcd",
		}, status.Error(codes.Unknown, ""))

		expectedErr := status.Error(codes.Internal, "User register Failed")

		_, err := mid.RegisterUser(context.TODO(), &pb.RegisterUserRequest{
			IdCardNumber: "",
			UserName:     "",
		})

		if err.Error() != expectedErr.Error() {
			t.Errorf("Expected %v but output is %v", expectedErr.Error(), err.Error())
		}
	})
}

func TestGetKYC(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()

	userServiceMock := pb.NewMockUserServiceClient(controller)
	mid := CreateMidServiceHandler(userServiceMock, nil)

	t.Run("Case success", func(t *testing.T) {
		userServiceMock.EXPECT().GetCurrentKYC(gomock.Any(), gomock.Any()).Return(&pb.GetCurrentKYCResponse{
			UserId:   "",
			KycLevel: 1,
		}, nil)

		expected := 1

		result, _ := mid.GetCurrentKYC(context.TODO(), &pb.GetCurrentKYCRequest{
			UserId: "",
		})

		if result.KycLevel != int32(expected) {
			t.Errorf("Expected %v but output is %v", expected, result)
		}
	})

	t.Run("Case Error", func(t *testing.T) {
		userServiceMock.EXPECT().GetCurrentKYC(gomock.Any(), gomock.Any()).Return(&pb.GetCurrentKYCResponse{
			UserId:   "",
			KycLevel: 1,
		}, status.Error(codes.Aborted, ""))

		expectedErr := status.Error(codes.Internal, "Get KYC level failed")

		_, err := mid.GetCurrentKYC(context.TODO(), &pb.GetCurrentKYCRequest{
			UserId: "123456",
		})

		if err.Error() != expectedErr.Error() {
			t.Errorf("Expected %v but output is %v", expectedErr, err)
		}
	})
}

func TestOpenSavingAccount(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()

	userServiceMock := pb.NewMockUserServiceClient(controller)
	savingServiceMock := pb.NewMockSavingsServiceClient(controller)
	mid := CreateMidServiceHandler(userServiceMock, savingServiceMock)

	t.Run("Case error1", func(t *testing.T) {
		userServiceMock.EXPECT().GetCurrentKYC(gomock.Any(), gomock.Any()).Return(&pb.GetCurrentKYCResponse{
			UserId:   "",
			KycLevel: 1,
		}, status.Error(codes.NotFound, ""))

		expectedErr := status.Error(codes.NotFound, "KYC level verify failed")

		_, err := mid.OpenSavingsAccount(context.TODO(), &pb.OpenSavingsAccountRequest{
			UserId:      "abv-der",
			Balance:     1000000,
			TermType:    "MONTHS",
			Term:        3,
			CreatedDate: "01022024",
		})

		if err.Error() != expectedErr.Error() {
			t.Errorf("Expected %v but output is %v", expectedErr, err)
		}
	})

	t.Run("Case KYC 1", func(t *testing.T) {
		userServiceMock.EXPECT().GetCurrentKYC(gomock.Any(), gomock.Any()).Return(&pb.GetCurrentKYCResponse{
			UserId:   "",
			KycLevel: 1,
		}, nil)

		expectedErr := status.Error(codes.PermissionDenied, "KYC level < 2")

		_, err := mid.OpenSavingsAccount(context.TODO(), &pb.OpenSavingsAccountRequest{
			UserId:      "abv-der",
			Balance:     1000000,
			TermType:    "MONTHS",
			Term:        3,
			CreatedDate: "01022024",
		})

		if err.Error() != expectedErr.Error() {
			t.Errorf("Expected %v but output is %v", expectedErr, err)
		}
	})

	t.Run("Case Open OK 1", func(t *testing.T) {
		userServiceMock.EXPECT().GetCurrentKYC(gomock.Any(), gomock.Any()).Return(&pb.GetCurrentKYCResponse{
			UserId:   "",
			KycLevel: 2,
		}, nil)

		savingServiceMock.EXPECT().OpenSavingsAccount(gomock.Any(), gomock.Any()).Return(&pb.SavingAccount{
			UserID:      "abc-123",
			Balance:     10000000,
			TermType:    "DAYS",
			Term:        21,
			CreatedDate: "01022024",
			DueDate:     "22022024",
			Rate:        17500,
		}, nil)

		res, _ := mid.OpenSavingsAccount(context.TODO(), &pb.OpenSavingsAccountRequest{
			UserId:      "abc-123",
			Balance:     10000000,
			TermType:    "DAYS",
			Term:        21,
			CreatedDate: "01022024",
		})

		expectedRes := &pb.SavingAccount{
			UserID:      "abc-123",
			Balance:     10000000,
			TermType:    "DAYS",
			Term:        21,
			CreatedDate: "01022024",
			DueDate:     "22022024",
			Rate:        17500,
		}

		if res.UserId != expectedRes.UserID {
			t.Errorf("Expected UserID %v but output is %v", expectedRes, res)
		}
		if res.Balance != expectedRes.Balance {
			t.Errorf("Expected balance %v but output is %v", expectedRes, res)
		}
		if res.Rate != expectedRes.Rate {
			t.Errorf("Expected rate %v but output is %v", expectedRes, res)
		}
		if res.CreatedDate != expectedRes.CreatedDate {
			t.Errorf("Expected createdDate %v but output is %v", expectedRes, res)
		}
		if res.DueDate != expectedRes.DueDate {
			t.Errorf("Expected dueDate %v but output is %v", expectedRes, res)
		}
	})

}
