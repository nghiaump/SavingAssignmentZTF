package main

type SavingAccount struct {
	// @gotags: es:"id"
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty" es:"id"`
	// @gotags: es:"user_id"
	UserId string `protobuf:"bytes,2,opt,name=user_id,proto3" json:"user_id,omitempty" es:"user_id"`
	// @gotags: es:"balance"
	Balance  int64  `protobuf:"varint,3,opt,name=balance,proto3" json:"balance,omitempty" es:"balance"`
	TermType string `protobuf:"bytes,4,opt,name=term_type,proto3" json:"term_type,omitempty"`
	Term     int32  `protobuf:"varint,5,opt,name=term,proto3" json:"term,omitempty"`
	// @gotags: es:"term_in_days"
	TermInDays int32 `protobuf:"varint,6,opt,name=term_in_days,proto3" json:"term_in_days,omitempty" es:"term_in_days"`
	// @gotags: es:"created_date"
	CreatedDate string `protobuf:"bytes,7,opt,name=created_date,proto3" json:"created_date,omitempty" es:"created_date"`
	// @gotags: es:"due_date"
	DueDate string `protobuf:"bytes,8,opt,name=due_date,proto3" json:"due_date,omitempty" es:"due_date"`
	// @gotags: es:"rate"
	Rate float32 `protobuf:"fixed32,9,opt,name=rate,proto3" json:"rate,omitempty" es:"rate"`
	// @gotags: es:"kyc"
	Kyc int32 `protobuf:"varint,10,opt,name=kyc,proto3" json:"kyc,omitempty" es:"kyc"`
}
