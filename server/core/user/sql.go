package main

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	pb "github.com/nghiaump/SavingAssignmentZTF/protobuf"
	"log"
	"time"
)

func CreateMySQLClient() *sql.DB {
	// Kết nối vào MySQL
	db, err := sql.Open("mysql", "root:@tcp(mysql:3306)/")
	for err != nil {
		time.Sleep(5 * time.Second)
		log.Println("Try to reconnect to mysql database")
		db, err = sql.Open("mysql", "root:@tcp(mysql:3306)/")
	}
	//defer db.Close() // khong dong ket noi
	log.Println("Connected to sql container")

	// Tạo database "dbo.user" nếu chưa tồn tại
	_, err = db.Exec("CREATE DATABASE IF NOT EXISTS userdb")
	if err != nil {
		panic(err.Error())
	}
	log.Println("Database UserDB is ready!")

	// Chọn database "dbo.user" để thực hiện các thao tác tiếp theo
	_, err = db.Exec("USE userdb")
	if err != nil {
		panic(err.Error())
	}

	// Tạo bảng "user" nếu chưa tồn tại
	_, err = db.Exec(`
		CREATE TABLE IF NOT EXISTS user (
			id VARCHAR(255) PRIMARY KEY,
			id_card_number VARCHAR(255),
			user_name VARCHAR(255),
			dob DATETIME,
			gender INT,
			address VARCHAR(255),
			phone_number VARCHAR(255),
			kyc_level INT,
			registered_date DATETIME
		);
	`)
	if err != nil {
		return nil
	}
	log.Println("Table 'user' has been created (if it didn't exist already).")
	return db
}

func (handler *UserServiceHandler) SQLCreateUser(user *pb.User) error {
	dob, _ := ConvertDateFormat(user.Dob)
	registeredDate, _ := ConvertDateFormat(user.RegisteredDate)
	// Convert dates before writing to MySQL database
	_, err := handler.db.Exec(`
		INSERT INTO user (id, id_card_number, user_name, dob, gender, address, phone_number, kyc_level, registered_date)
		VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?)
	`, user.Id, user.IdCardNumber, user.UserName, dob, user.Gender, user.Address, user.PhoneNumber, user.KycLevel, registeredDate)
	return err
}

func (handler *UserServiceHandler) SQLGetUserById(id string) (*pb.User, error) {
	var user pb.User
	err := handler.db.QueryRow(`
		SELECT id, id_card_number, user_name, dob, gender, address, phone_number, kyc_level, registered_date
		FROM user
		WHERE id=?
	`, id).Scan(&user.Id, &user.IdCardNumber, &user.UserName, &user.Dob, &user.Gender, &user.Address, &user.PhoneNumber, &user.KycLevel, &user.RegisteredDate)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (handler *UserServiceHandler) SQLUpdateUser(user *pb.User) error {
	_, err := handler.db.Exec(`
		UPDATE user
		SET id_card_number=?, user_name=?, dob=?, gender=?, address=?, phone_number=?, kyc_level=?, registered_date=?
		WHERE id=?
	`, user.IdCardNumber, user.UserName, user.Dob, user.Gender, user.Address, user.PhoneNumber, user.KycLevel, user.RegisteredDate, user.Id)
	return err
}

func (handler *UserServiceHandler) SQLDeleteUserById(id string) error {
	_, err := handler.db.Exec("DELETE FROM user WHERE id=?", id)
	return err
}

func (handler *UserServiceHandler) SQLQueryUsers(params map[string]interface{}) ([]*pb.User, error) {
	query, args := handler.SQLQueryGenerator(params)

	rows, err := handler.db.Query(query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []*pb.User
	for rows.Next() {
		var user pb.User
		err := rows.Scan(&user.Id, &user.IdCardNumber, &user.UserName, &user.Dob, &user.Gender, &user.Address, &user.PhoneNumber, &user.KycLevel, &user.RegisteredDate)
		if err != nil {
			return nil, err
		}
		users = append(users, &user)
	}

	return users, nil
}

func (handler *UserServiceHandler) SQLQueryGenerator(params map[string]interface{}) (string, []interface{}) {
	query := "SELECT * FROM user WHERE 1=1"
	var args []interface{}

	for key, value := range params {
		switch key {
		case "id_card_number":
			query += "AND id_card_number=?"
			args = append(args, value)
		case "user_name":
			query += "AND user_name=?"
			args = append(args, value)
		case "gender":
			query += " AND gender=?"
			args = append(args, value)
		case "addressContains":
			query += " AND address LIKE ?"
			args = append(args, "%"+value.(string)+"%")
		case "phone_number":
			query += "AND phone_number=?"
			args = append(args, value)
		case "kyc":
			query += "AND kyc=?"
			args = append(args, value)
		case "registered_date_range":
			if dateRange, ok := value.(map[string]string); ok {
				if dateRange["date1"] != "" && dateRange["date2"] != "" {
					query += " AND registered_date BETWEEN ? AND ?"
					args = append(args, dateRange["date1"], dateRange["date2"])
				}
			}
		}
	}

	return query, args
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
