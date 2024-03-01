package main

import (
	"context"
	"fmt"
	pb "github.com/nghiaump/SavingAssignmentZTF/protobuf"
	"log"
)

func CreateWithdrawalRequest() pb.WithdrawalRequest {
	var withDrawalReq pb.WithdrawalRequest
	fmt.Print("userID: ")
	fmt.Scan(&withDrawalReq.UserId)
	fmt.Print("accountID: ")
	fmt.Scan(&withDrawalReq.AccountId)
	fmt.Print("withdrawal amount: ")
	fmt.Scan(&withDrawalReq.Amount)
	fmt.Print("withdrawal date(simulated): ")
	fmt.Scan(&withDrawalReq.Date)
	return withDrawalReq
}

func CreateAccountInquiryRequest() pb.AccountInquiryRequest {
	var accInquiryReq pb.AccountInquiryRequest
	fmt.Print("userID: ")
	fmt.Scan(&accInquiryReq.UserId)
	fmt.Print("accountID: ")
	fmt.Scan(&accInquiryReq.AccountId)
	return accInquiryReq
}

func CreateOpenSavingAccountRequest() pb.OpenSavingsAccountRequest {
	var accReq pb.OpenSavingsAccountRequest
	var termType string
	var term int32
	var termCombo int32

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
	}

	accReq.TermType = termType
	accReq.Term = term
	fmt.Print("Created Date: ")
	fmt.Scan(&accReq.CreatedDate)
	return accReq
}

func CreateKYCRequest() pb.GetCurrentKYCRequest {
	var kycReq pb.GetCurrentKYCRequest
	fmt.Print("UserID: ")
	fmt.Scan(&kycReq.UserId)
	fmt.Print("ID Card Number: ")
	fmt.Scan(&kycReq.IdCardNumber)
	return kycReq
}

func CreateRegisterUserRequest() pb.RegisterUserRequest {
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
	return userReg
}

func CreateAccountFilter() pb.Filter {
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
	fmt.Println("Page Size")
	fmt.Scan(&filter.PageSize)
	filter.PageIndex = 1
	return filter
}

func LoopForPaginate(filter pb.Filter, c pb.MidSavingServiceClient, ctx context.Context) {
	var nextPage int
	for {
		fmt.Println("See another page result?: ")
		fmt.Scan(&nextPage)
		if nextPage == -1 {
			break
		}
		filter.PageIndex = int32(nextPage)
		savingAccList, _ := c.SearchAccountsByFilter(ctx, &filter)

		log.Printf("Hits: %v", len(savingAccList.AccList))
		if len(savingAccList.AccList) > 0 {
			printSavingAccountsTable(savingAccList.AccList)
		}
	}
}

func PrintResult(savingAccList *pb.SavingAccountList) {
	fmt.Printf("Hits: %v\n", savingAccList.AggTotalHits)
	fmt.Printf("Total balance: %v\n", savingAccList.AggTotalBalance)
	fmt.Printf("Page %v:\n", savingAccList.PageIndex)
	if len(savingAccList.AccList) > 0 {
		printSavingAccountsTable(savingAccList.AccList)
	}
}

func printSavingAccountsTable(accList []*pb.SavingAccount) {
	headerFormat := "| %-40s | %-40s | %10s | %-5s | %-12s | %-25s |"
	rowFormat := "| %-40s | %-40s | %10d | %-5d | %-12d | %-25s |"
	divider := "+------------------------------------------+------------------------------------------+------------+-------+--------------+---------------------------+"

	fmt.Println(divider)
	fmt.Printf(headerFormat, "ID", "UserID", "Balance", "KYC", "Term in Days", "Due Date")
	fmt.Println("\n" + divider)

	for _, acc := range accList {
		fmt.Printf(rowFormat, acc.Id, acc.UserId, acc.Balance, acc.Kyc, acc.TermInDays, acc.DueDate)
		fmt.Println()
	}

	fmt.Println(divider)
}

func printUsersTable(users []*pb.User) {
	headerFormat := "| %-40s | %-40s | %-40s | %-12s | %-12s | %-5s | %-25s |"
	rowFormat := "| %-40s | %-40s | %-40s | %-12s | %-12s | %-5d | %-25s |"
	divider := "+------------------------------------------+------------------------------------------+------------------------------------------+--------------+--------------+-------+---------------------------+"

	fmt.Println(divider)
	fmt.Printf(headerFormat, "ID", "ID Card Number", "User Name", "DOB", "Address", "KYC Level", "Registered Date")
	fmt.Println()
	fmt.Println(divider)

	for _, user := range users {
		fmt.Printf(rowFormat, user.Id, user.IdCardNumber, user.UserName, user.Dob, user.Address, user.KycLevel, user.RegisteredDate)
		fmt.Println()
	}

	fmt.Println(divider)
}
