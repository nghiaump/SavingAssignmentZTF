package main

import (
	"flag"
	"github.com/golang/glog"
	pb "github.com/nghiaump/SavingAssignmentZTF/protobuf"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"os"
)

const MidPort = ":50050"
const UserPort = ":50051"
const SavingPort = ":50052"
const ContainerUserCoreEnv = "CONTAINER_USER_CORE_HOST"
const ContainerSavingCoreEnv = "CONTAINER_SAVING_CORE_HOST"

func main() {
	os.Args = append(os.Args, "-logtostderr=true")
	flag.Parse()
	defer glog.Flush()

	// Lấy giá trị của biến môi trường
	addressUserCore := os.Getenv(ContainerUserCoreEnv)
	if addressUserCore == "" {
		glog.Info("Biến môi trường CONTAINER_USER_CORE_HOST không được cung cấp.")
		return
	}
	addressSavingCore := os.Getenv(ContainerSavingCoreEnv)
	if addressSavingCore == "" {
		glog.Info("Biến môi trường CONTAINER_SAVING_CORE_HOST không được cung cấp.")
		return
	}

	connUserCore, errUserCore := grpc.Dial(addressUserCore+UserPort, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if errUserCore != nil {
		glog.Fatalf("Cannot connect to User Core: %v", errUserCore)
	}
	defer connUserCore.Close()

	connSavingCore, errSavingCore := grpc.Dial(addressSavingCore+SavingPort, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if errSavingCore != nil {
		glog.Fatalf("Cannot connect to User Core: %v", errSavingCore)
	}
	defer connSavingCore.Close()

	midServiceHandler := CreateMidServiceHandler(pb.NewUserServiceClient(connUserCore), pb.NewSavingsServiceClient(connSavingCore))

	StartMidServer(midServiceHandler, MidPort)
}
