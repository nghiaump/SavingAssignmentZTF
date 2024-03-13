package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	pb "github.com/nghiaump/SavingAssignmentZTF/protobuf"
	"log"
)

func GetMySQLDB() *sql.DB {
	db, err := sql.Open("mysql", "user:password@tcp(mysql:3306)/savingaccount")
	if err != nil {
		log.Fatal(err)
		return nil
	}
	defer db.Close()

	// Kiểm tra kết nối
	err = db.Ping()
	if err != nil {
		log.Fatal(err)
		return nil
	}
	fmt.Println("Connected to MySQL database!")
	return db
}

func createTable(db *sql.DB) {
	createTableSQL := `
	CREATE TABLE IF NOT EXISTS savingaccount (
		id VARCHAR(255) PRIMARY KEY,
		user_id VARCHAR(255),
		balance BIGINT,
		term_type VARCHAR(255),
		term INT,
		term_in_days INT,
		created_date VARCHAR(255),
		due_date VARCHAR(255),
		rate FLOAT,
		kyc INT
	);`

	_, err := db.Exec(createTableSQL)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Table 'savingaccount' created successfully or already exists!")
}

func insertData(db *sql.DB) {
	// Example of inserting data
	savingAcc := &pb.SavingAccount{
		Id:          "123",
		UserId:      "user123",
		Balance:     1000,
		TermType:    "Type A",
		Term:        12,
		TermInDays:  365,
		CreatedDate: "2024-03-07",
		DueDate:     "2025-03-07",
		Rate:        5.0,
		Kyc:         1,
	}

	insertQuery := "INSERT INTO savingaccount (id, user_id, balance, term_type, term, term_in_days, created_date, due_date, rate, kyc) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?)"

	_, err := db.Exec(insertQuery, savingAcc.Id, savingAcc.UserId, savingAcc.Balance, savingAcc.TermType, savingAcc.Term, savingAcc.TermInDays, savingAcc.CreatedDate, savingAcc.DueDate, savingAcc.Rate, savingAcc.Kyc)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Data inserted successfully!")
}

func readData(db *sql.DB) {
	// Example of reading data
	rows, err := db.Query("SELECT id, user_id, balance, term_type, term, term_in_days, created_date, due_date, rate, kyc FROM savingaccount")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	var savingAcc pb.SavingAccount
	for rows.Next() {
		err := rows.Scan(&savingAcc.Id, &savingAcc.UserId, &savingAcc.Balance, &savingAcc.TermType, &savingAcc.Term, &savingAcc.TermInDays, &savingAcc.CreatedDate, &savingAcc.DueDate, &savingAcc.Rate, &savingAcc.Kyc)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("ID: %s, User ID: %s, Balance: %d\n", savingAcc.Id, savingAcc.UserId, savingAcc.Balance)
	}
}

func updateData(db *sql.DB) {
	// Example of updating data
	updateQuery := "UPDATE savingaccount SET balance = ? WHERE id = ?"

	_, err := db.Exec(updateQuery, 1500, "123")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Data updated successfully!")
}

func deleteData(db *sql.DB) {
	// Example of deleting data
	deleteQuery := "DELETE FROM savingaccount WHERE id = ?"

	_, err := db.Exec(deleteQuery, "123")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Data deleted successfully!")
}
