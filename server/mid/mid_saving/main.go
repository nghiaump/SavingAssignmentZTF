package main

import (
	pb "github.com/nghiaump/SavingAssignmentZTF/protobuf"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
)

const MidPort = ":50050"

func main() {
	connUserCore, errUserCore := grpc.Dial(AddressUserCore, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if errUserCore != nil {
		log.Fatalf("Cannot connect to User Core: %v", errUserCore)
	}
	defer connUserCore.Close()

	connSavingCore, errSavingCore := grpc.Dial(AddressSavingCore, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if errSavingCore != nil {
		log.Fatalf("Cannot connect to User Core: %v", errSavingCore)
	}
	defer connSavingCore.Close()

	midServiceHandler := CreateMidServiceHandler(pb.NewUserServiceClient(connUserCore), pb.NewSavingsServiceClient(connSavingCore))

	StartMidServer(midServiceHandler, MidPort)
}
