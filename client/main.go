package main

import (
	"context"
	"fmt"
	pb "github.com/nghiaump/SavingAssignmentZTF/protobuf"
	"google.golang.org/grpc"
	"log"
	"os"
)

const (
	ContainerMidSavingEnv = "CONTAINER_MID_SAVING_HOST"
	MidPort               = ":50050"
)

func main() {
	addressMidSavingCore := os.Getenv(ContainerMidSavingEnv)
	if addressMidSavingCore == "" {
		fmt.Println("Biến môi trường CONTAINER_MID_SAVING_HOST không được cung cấp.")
		return
	} else {
		fmt.Printf("address mid container: %v", addressMidSavingCore)
	}
	conn, err := grpc.Dial(addressMidSavingCore+MidPort, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewMidSavingServiceClient(conn)

	const ActionRegisterUser = 1
	const ActionCheckKYC = 2
	const ActionCreateSavingAccount = 3
	const ActionInquireSavingAccount = 4
	const ActionWithdrawal = 5
	const ActionSearchAllAccountsByUserID = 6
	const ActionSearchAccountsByFilters = 7
	const ActionSearchUserByIDCardNumber = 8
	const ActionSearchUserByAccountID = 9
	const ActionSearchUsersByFilters = 10

	currentAction := ActionRegisterUser
	for {
		fmt.Println("Input action:")
		fmt.Printf("%v. Register User: using IDCardNumber, name\n", ActionRegisterUser)
		fmt.Printf("%v. Check KYC level: using UserID, IDCardNumber\n", ActionCheckKYC)
		fmt.Printf("%v. Open Saving Account: using UserID, ...\n", ActionCreateSavingAccount)
		fmt.Printf("%v. Checking Saving Account: using AccountID, UserID\n", ActionInquireSavingAccount)
		fmt.Printf("%v. Withdrawal\n", ActionWithdrawal)
		fmt.Printf("%v. Search all saving accounts by UserID\n", ActionSearchAllAccountsByUserID)
		fmt.Printf("%v. Search accounts by Filters ***NEW***\n", ActionSearchAccountsByFilters)
		fmt.Printf("%v. Search User by ID card number\n", ActionSearchUserByIDCardNumber)
		fmt.Printf("%v. Search User by Account ID\n", ActionSearchUserByAccountID)
		fmt.Printf("%v. Search Users by Filters\n", ActionSearchUsersByFilters)

		fmt.Scan(&currentAction)
		ctx := context.Background()

		switch currentAction {
		case ActionRegisterUser:
			// Register User
			{
				userReg := CreateRegisterUserRequest()

				res, errRegUser := c.RegisterUser(ctx, &userReg)
				if errRegUser != nil {
					log.Printf("Could not register new user: %v", errRegUser.Error())
				} else {
					log.Printf("Registered successfully\nUser ID: %s", res.UserId)
				}
			}
		case ActionCheckKYC:
			// Check KYC level
			{
				kycReq := CreateKYCRequest()

				kycRes, errKYC := c.GetCurrentKYC(ctx, &kycReq)
				if errKYC != nil {
					log.Printf("Could not get user KYC level: %v", errKYC.Error())
				} else {

					log.Printf("User ID: %v, KYC level: %v", kycRes.UserId, kycRes.KycLevel)
				}
			}
		case ActionCreateSavingAccount:
			{
				accReq := CreateOpenSavingAccountRequest()

				accRes, errOpen := c.OpenSavingsAccount(ctx, &accReq)
				if errOpen != nil {
					log.Printf("Could not open account: %v", errOpen.Error())
				} else {
					log.Println(accRes)
				}

			}
		case ActionInquireSavingAccount:
			{
				accInquiryReq := CreateAccountInquiryRequest()

				accRes, errInquire := c.AccountInquiry(ctx, &accInquiryReq)
				if errInquire != nil {
					log.Printf("Cannot inquire the account id %v", accInquiryReq.AccountId)
					log.Printf("Error detail: %v", errInquire.Error())
				} else {
					log.Printf("accountID: %v \nDetail: %v", accInquiryReq.AccountId, accRes)
				}

			}
		case ActionWithdrawal:
			{
				withDrawalReq := CreateWithdrawalRequest()

				withDrawRes, withDrawErr := c.Withdrawal(ctx, &withDrawalReq)
				if withDrawErr != nil {
					log.Printf("Withdrawn failed with err %v", withDrawErr.Error())
				} else {
					log.Println("Withdrawn successfully")
					log.Printf("Amount: %v\nremain %v", withDrawRes.WithdrawnAmount, withDrawRes.Acc)
				}
			}

		case ActionSearchAllAccountsByUserID:
			{
				var userID string
				fmt.Print("userID: ")
				fmt.Scan(&userID)

				log.Printf("Calling Search All Acc for userID %v\n", userID)
				savingAccList, _ := c.SearchAccountsByUserID(ctx, &pb.AccountInquiryRequest{
					UserId:    userID,
					AccountId: "",
				})

				log.Printf("Hits: %v", len(savingAccList.AccList))
				if len(savingAccList.AccList) > 0 {
					printSavingAccountsTable(savingAccList.AccList)
				}
			}

		case ActionSearchAccountsByFilters:
			{
				filter := CreateAccountFilter()
				log.Println("Calling SearchAccountsByFilters ")
				savingAccList, _ := c.SearchAccountsByFilter(ctx, &filter)
				PrintResult(savingAccList)
				LoopForPaginate(filter, c, ctx)
			}

		case ActionSearchUserByAccountID:
			{
				var accID string
				fmt.Println("Input account id: ")
				fmt.Scan(&accID)
				user, _ := c.SearchUserByAccountID(ctx, &pb.AccountID{
					AccountID: accID,
				})

				fmt.Printf("Search result: %v\n", user)
			}

		case ActionSearchUserByIDCardNumber:
			{
				var IDCardNumber string
				fmt.Println("Input id card number: ")
				fmt.Scan(&IDCardNumber)
				user, _ := c.SearchUserByIdCardNumber(ctx, &pb.IDCardNumber{
					IdCardNumber: IDCardNumber,
				})

				fmt.Printf("Search result: %v\n", user)
			}

		}

	}

}
