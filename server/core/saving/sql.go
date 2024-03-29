package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/golang/glog"
	"strings"
	"time"
)

type AccountUser struct {
	accountID string
	userID    string
}

func CreateMySQLClient() *sql.DB {
	db, err := sql.Open("mysql", "root:@tcp(mysql:3306)/")
	for err != nil {
		time.Sleep(5 * time.Second)
		glog.Info("Try to reconnect to mysql database")
		db, err = sql.Open("mysql", "root:@tcp(mysql:3306)/")
	}
	//defer db.Close() -> Phai xoa, khong duoc dong ket noi
	glog.Info("Connected to sql container")

	_, err = db.Exec("CREATE DATABASE IF NOT EXISTS account_db")
	if err != nil {
		panic(err.Error())
	}
	glog.Info("Database account_db is ready!")

	_, err = db.Exec("USE account_db")
	if err != nil {
		panic(err.Error())
	}

	_, err = db.Exec(`
		CREATE TABLE IF NOT EXISTS account_user (
			account_id VARCHAR(255) PRIMARY KEY,
			user_id VARCHAR(255)
		);
	`)

	if err != nil {
		glog.Info("Table 'account_user' cannot be created.")
		return nil
	}
	glog.Info("Table 'account_user' has been created (if it didn't exist already).")
	return db
}

func (handler *SavingServiceHandler) SQLCreateAccountUser(accountUser *AccountUser) error {
	_, err := handler.db.Exec(`
		INSERT INTO account_user (account_id, user_id)
		VALUES (?, ?)
	`, accountUser.accountID, accountUser.userID)
	if err != nil {
		glog.Info(err)
	}
	return err
}

func (handler *SavingServiceHandler) SQLGetAccountUserByAccountId(accountID string) (*AccountUser, error) {
	var accountUser AccountUser
	err := handler.db.QueryRow(`
		SELECT account_id, user_id
		FROM account_user
		WHERE account_id=?
	`, accountID).Scan(&accountUser.accountID, &accountUser.userID)
	if err != nil {
		return nil, err
	}
	return &accountUser, nil
}

func (handler *SavingServiceHandler) SQLUpdateAccountUser(accountUser *AccountUser) error {
	_, err := handler.db.Exec(`
		UPDATE account_user
		SET user_id=?
		WHERE account_id=?
	`, accountUser.userID, accountUser.accountID)
	return err
}

func (handler *SavingServiceHandler) SQLDeleteAccountUserByAccountId(accountID string) error {
	_, err := handler.db.Exec("DELETE FROM account_user WHERE account_id=?", accountID)
	return err
}

func (handler *SavingServiceHandler) GetUserHavingAccountNumber(minNum int, maxNum int) (map[string][]string, error) {
	glog.Info("min-num-acc: %v, max-num-acc: %v", minNum, maxNum)
	userAccounts := make(map[string][]string)

	query := fmt.Sprintf(`
        SELECT user_id, GROUP_CONCAT(account_id) AS account_ids
        FROM account_user
        GROUP BY user_id
        HAVING COUNT(account_id) BETWEEN %d AND %d
    `, minNum, maxNum)

	rows, err := handler.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var userID, accountIDs string
		if err := rows.Scan(&userID, &accountIDs); err != nil {
			return nil, err
		}

		accountIDArr := strings.Split(accountIDs, ",")
		userAccounts[userID] = accountIDArr
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}

	return userAccounts, nil
}

func ConvertDateFormat(dateISO string) (string, error) {
	// Chuyển đổi từ ISO8601 sang đối tượng time.Time
	t, err := time.Parse(time.RFC3339, dateISO)
	if err != nil {
		return "", err
	}

	// Chuyển đổi đối tượng time.Time sang chuỗi định dạng "YYYY-MM-DD"
	return t.Format("2006-01-02"), nil
}
