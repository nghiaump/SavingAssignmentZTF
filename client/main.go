package main

import (
	"context"
	"fmt"
	pb "github.com/nghiaump/savingaccproto"
	"google.golang.org/grpc"
	"log"
)

const (
	addressMidSaving = "localhost:50050"
)

func main() {
	conn, err := grpc.Dial(addressMidSaving, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewMidSavingServiceClient(conn)

	const ActionRegisterUser = 1
	const ActionCheckKYC = 2
	const ActionCreateSavingAccout = 3
	const ActionInquireSavingAccount = 4
	const ActionWithdrawal = 5

	currentAction := ActionRegisterUser
	for {
		fmt.Println("Input action:")
		fmt.Println("1. Register User: using IDCardNumber, name")
		fmt.Println("2. Check KYC level: using UserID")
		fmt.Println("3. Open Saving Account: using UserID (complex),...")
		fmt.Println("4. Checking Saving Account: using AccountID (complex)")
		fmt.Println("5. Withdraw")
		fmt.Scan(&currentAction)
		ctx := context.Background()

		switch currentAction {
		case ActionRegisterUser:
			// Register User
			{
				var IDCardNumber string
				var name string
				fmt.Print("IDCardnumber: ")
				fmt.Scan(&IDCardNumber)
				fmt.Print("Name: ")
				fmt.Scan(&name)
				res, errRegUser := c.RegisterUser(ctx,
					&pb.RegisterUserRequest{
						IdCardNumber: IDCardNumber,
						UserName:     name,
					})
				if errRegUser != nil {
					log.Printf("Could not register new user: %v", errRegUser.Error())
				}
				log.Printf("Registered successfully\nUser ID: %s", res.UserId)
			}
		case ActionCheckKYC:
			// Check KYC level
			{
				var userID string
				fmt.Print("UserID: ")
				fmt.Scan(&userID)
				kycRes, errKYC := c.GetCurrentKYC(ctx, &pb.GetCurrentKYCRequest{
					UserId: userID,
				})
				if errKYC != nil {
					log.Printf("Could not get user KYC level: %v", errKYC.Error())
				}
				log.Printf("User ID: %v, KYC level: %v", kycRes.UserId, kycRes.KycLevel)
			}
		case ActionCreateSavingAccout:
			{
				var userID string
				var balance int64
				var termType string
				var term int32
				var termCombo int32
				var createdDate string
				fmt.Print("UserID: ")
				fmt.Scan(&userID)
				fmt.Print("Balance: ")
				fmt.Scan(&balance)
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
				fmt.Print("Created Date: ")
				fmt.Scan(&createdDate)
				accRes, errOpen := c.OpenSavingsAccount(ctx, &pb.OpenSavingsAccountRequest{
					UserId:      userID,
					Balance:     balance,
					TermType:    termType,
					Term:        term,
					CreatedDate: createdDate,
				})
				if errOpen != nil {
					log.Printf("Could not open account: %v", errOpen.Error())
				}
				log.Println(accRes)
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
				}
				log.Printf("accountID: %v \nDetail: %v", accountID, accRes)
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

		}

	}

}
