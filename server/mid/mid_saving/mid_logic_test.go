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

}
