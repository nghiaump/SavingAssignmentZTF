package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	pb "github.com/nghiaump/SavingAssignmentZTF/protobuf"
)

func (handler *UserServiceHandler) InitMySQLClient() {
	db, err := sql.Open("mysql", "user:password@tcp(mysql-container:3306)/database_name")
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	_, err = db.Exec(`
        CREATE TABLE IF NOT EXISTS user (
            id VARCHAR(255) PRIMARY KEY,
            id_card_number VARCHAR(255),
            user_name VARCHAR(255),
            dob VARCHAR(255),
            gender INT,
            address VARCHAR(255),
            phone_number VARCHAR(255),
            kyc_level INT,
            registered_date VARCHAR(255)
        );
    `)
	if err != nil {
		panic(err.Error())
	}

	fmt.Println("Table 'user' has been created (if it didn't exist already).")
	handler.db = db
}

func (handler *UserServiceHandler) CreateUser(user *pb.User) error {
	_, err := handler.db.Exec(`
		INSERT INTO user (id, id_card_number, user_name, dob, gender, address, phone_number, kyc_level, registered_date)
		VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?)
	`, user.Id, user.IdCardNumber, user.UserName, user.Dob, user.Gender, user.Address, user.PhoneNumber, user.KycLevel, user.RegisteredDate)
	return err
}

func (handler *UserServiceHandler) GetUserById(id string) (*pb.User, error) {
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

func (handler *UserServiceHandler) UpdateUser(user *pb.User) error {
	_, err := handler.db.Exec(`
		UPDATE user
		SET id_card_number=?, user_name=?, dob=?, gender=?, address=?, phone_number=?, kyc_level=?, registered_date=?
		WHERE id=?
	`, user.IdCardNumber, user.UserName, user.Dob, user.Gender, user.Address, user.PhoneNumber, user.KycLevel, user.RegisteredDate, user.Id)
	return err
}

func (handler *UserServiceHandler) DeleteUserById(id string) error {
	_, err := handler.db.Exec("DELETE FROM user WHERE id=?", id)
	return err
}

func (handler *UserServiceHandler) QueryUsers(params map[string]interface{}) ([]*pb.User, error) {
	query, args := handler.queryGenerator(params)

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

func (handler *UserServiceHandler) queryGenerator(params map[string]interface{}) (string, []interface{}) {
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
