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
	const ActionSearchAccountsByIDCardNumber = 8

	const ActionSearchUserByIDCardNumber = 10
	const ActionSearchUserByAccountID = 11
	const ActionSearchUsersByFilters = 12

	currentAction := ActionRegisterUser
	for {
		fmt.Println("Input action:")
		fmt.Println("1. Register User: using IDCardNumber, name")
		fmt.Println("2. Check KYC level: using UserID, IDCardNumber")
		fmt.Println("3. Open Saving Account: using UserID,...")
		fmt.Println("4. Checking Saving Account: using AccountID, UserID")
		fmt.Println("5. Withdraw")
		fmt.Println("6. Search all saving account by UserID")
		fmt.Printf("%v. Search accounts by Filters\n", ActionSearchAccountsByFilters)
		fmt.Printf("%v. Search User by ID card number\n", ActionSearchUserByIDCardNumber)
		fmt.Printf("%v. Search User by Account ID\n", ActionSearchUserByAccountID)
		fmt.Printf("%v. Search User by UserFilters\n", ActionSearchUsersByFilters)

		fmt.Scan(&currentAction)
		ctx := context.Background()

		switch currentAction {
		case ActionRegisterUser:
			// Register User
			{
				var userReg pb.RegisterUserRequest
				fmt.Print("IDCardnumber: ")
				fmt.Scan(&userReg.IdCardNumber)
				fmt.Print("Name: ")
				fmt.Scan(&userReg.UserName)
				fmt.Print("DOB: ")
				fmt.Scan(&userReg.Dob)
				fmt.Print("Address: ")
				fmt.Scan(&userReg.Address)
				fmt.Print("Phone number: ")
				fmt.Scan(&userReg.PhoneNumber)
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
				var kycReq pb.GetCurrentKYCRequest
				fmt.Print("UserID: ")
				fmt.Scan(&kycReq.UserId)
				fmt.Print("ID Card Number: ")
				fmt.Scan(&kycReq.IdCardNumber)
				kycRes, errKYC := c.GetCurrentKYC(ctx, &kycReq)
				if errKYC != nil {
					log.Printf("Could not get user KYC level: %v", errKYC.Error())
				} else {

					log.Printf("User ID: %v, KYC level: %v", kycRes.UserId, kycRes.KycLevel)
				}
			}
		case ActionCreateSavingAccount:
			{

				var termType string
				var term int32
				var termCombo int32
				var accReq pb.OpenSavingsAccountRequest
				fmt.Print("UserID: ")
				fmt.Scan(&accReq.UserId)
				fmt.Print("IDCardNumber: ")
				fmt.Scan(&accReq.IdCardNumber)
				fmt.Print("Balance: ")
				fmt.Scan(&accReq.Balance)
				fmt.Println("Term: ")
				fmt.Println("1: 21 DAYS")
				fmt.Println("2: 3 MONTHS")
				fmt.Println("3: 6 MONTHS")
				fmt.Println("4: 12 MONTHS")

				fmt.Scan(&termCombo)
				switch termCombo {
				case 1:
					{
						termType = "DAYS"
						term = 21
					}
				case 2:
					{
						termType = "MONTHS"
						term = 3
					}
				case 3:
					{
						termType = "MONTHS"
						term = 6
					}
				case 4:
					{
						termType = "MONTHS"
						term = 12
					}
					continue
				}

				accReq.TermType = termType
				accReq.Term = term
				fmt.Print("Created Date: ")
				fmt.Scan(&accReq.CreatedDate)
				accRes, errOpen := c.OpenSavingsAccount(ctx, &accReq)
				if errOpen != nil {
					log.Printf("Could not open account: %v", errOpen.Error())
				} else {
					log.Println(accRes)
				}

			}
		case ActionInquireSavingAccount:
			{
				var userID string
				var accountID string
				fmt.Print("userID: ")
				fmt.Scan(&userID)
				fmt.Print("accountID: ")
				fmt.Scan(&accountID)
				accRes, errInquire := c.AccountInquiry(ctx, &pb.AccountInquiryRequest{
					UserId:    userID,
					AccountId: accountID,
				})
				if errInquire != nil {
					log.Printf("Cannot inquire the account id %v", accountID)
					log.Printf("Error detail: %v", errInquire.Error())
				} else {
					log.Printf("accountID: %v \nDetail: %v", accountID, accRes)
				}

			}
		case ActionWithdrawal:
			{
				var userID string
				var accountID string
				var amount int64
				var withDrawDate string
				fmt.Print("userID: ")
				fmt.Scan(&userID)
				fmt.Print("accountID: ")
				fmt.Scan(&accountID)
				fmt.Print("withdraw amount: ")
				fmt.Scan(&amount)
				fmt.Print("withdraw date(simulated): ")
				fmt.Scan(&withDrawDate)

				withDrawRes, withDrawErr := c.Withdrawal(ctx, &pb.WithdrawalRequest{
					UserId:    userID,
					AccountId: accountID,
					Amount:    amount,
					Date:      withDrawDate,
				})

				if withDrawErr != nil {
					log.Printf("Withdrawn failed with err %v", withDrawErr.Error())
				} else {
					log.Println("Withdrawn successfully")
					log.Printf("Amount: %v, remain %v", withDrawRes.WithdrawnAmount, withDrawRes.Acc)
				}
			}

		case ActionSearchAllAccountsByUserID:
			{
				var userID string
				fmt.Print("userID: ")
				fmt.Scan(&userID)

				log.Printf("Calling Get All Acc for userID %v\n", userID)
				savingAccList, _ := c.SearchAccountsByUserID(ctx, &pb.AccountInquiryRequest{
					UserId:    userID,
					AccountId: "",
				})

				log.Printf("Hits: %v", len(savingAccList.AccList))
				for _, acc := range savingAccList.AccList {
					log.Printf(acc.Id)
				}

			}

		case ActionSearchAccountsByFilters:
			{
				var filter pb.Filter
				fmt.Println("Input KYC")
				fmt.Scan(&filter.Kyc)
				fmt.Println("Input TermInDays")
				fmt.Scan(&filter.TermInDays)
				fmt.Println("Input DueDateRange - earliest date")
				fmt.Scan(&filter.DueDateEarliest)
				fmt.Println("Input DueDateRange - latest date")
				fmt.Scan(&filter.DueDateLatest)
				fmt.Println("Input minimum balance")
				fmt.Scan(&filter.MinBalance)

				log.Println("Calling SearchAccountsByFilters ")
				savingAccList, _ := c.SearchAccountsByFilter(ctx, &filter)

				log.Printf("Hits: %v", len(savingAccList.AccList))
				for _, acc := range savingAccList.AccList {
					log.Printf(acc.Id)
				}
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
