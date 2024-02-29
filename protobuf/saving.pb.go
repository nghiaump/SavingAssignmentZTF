// Code generated by protoc-gen-go. DO NOT EDIT.
// source: saving.proto

package protobuf

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

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
	Kyc                  int32    `protobuf:"varint,10,opt,name=kyc,proto3" json:"kyc,omitempty" es:"kyc"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SavingAccount) Reset()         { *m = SavingAccount{} }
func (m *SavingAccount) String() string { return proto.CompactTextString(m) }
func (*SavingAccount) ProtoMessage()    {}
func (*SavingAccount) Descriptor() ([]byte, []int) {
	return fileDescriptor_7e8dd110a4bce988, []int{0}
}

func (m *SavingAccount) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SavingAccount.Unmarshal(m, b)
}
func (m *SavingAccount) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SavingAccount.Marshal(b, m, deterministic)
}
func (m *SavingAccount) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SavingAccount.Merge(m, src)
}
func (m *SavingAccount) XXX_Size() int {
	return xxx_messageInfo_SavingAccount.Size(m)
}
func (m *SavingAccount) XXX_DiscardUnknown() {
	xxx_messageInfo_SavingAccount.DiscardUnknown(m)
}

var xxx_messageInfo_SavingAccount proto.InternalMessageInfo

func (m *SavingAccount) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *SavingAccount) GetUserId() string {
	if m != nil {
		return m.UserId
	}
	return ""
}

func (m *SavingAccount) GetBalance() int64 {
	if m != nil {
		return m.Balance
	}
	return 0
}

func (m *SavingAccount) GetTermType() string {
	if m != nil {
		return m.TermType
	}
	return ""
}

func (m *SavingAccount) GetTerm() int32 {
	if m != nil {
		return m.Term
	}
	return 0
}

func (m *SavingAccount) GetTermInDays() int32 {
	if m != nil {
		return m.TermInDays
	}
	return 0
}

func (m *SavingAccount) GetCreatedDate() string {
	if m != nil {
		return m.CreatedDate
	}
	return ""
}

func (m *SavingAccount) GetDueDate() string {
	if m != nil {
		return m.DueDate
	}
	return ""
}

func (m *SavingAccount) GetRate() float32 {
	if m != nil {
		return m.Rate
	}
	return 0
}

func (m *SavingAccount) GetKyc() int32 {
	if m != nil {
		return m.Kyc
	}
	return 0
}

type AccID struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AccID) Reset()         { *m = AccID{} }
func (m *AccID) String() string { return proto.CompactTextString(m) }
func (*AccID) ProtoMessage()    {}
func (*AccID) Descriptor() ([]byte, []int) {
	return fileDescriptor_7e8dd110a4bce988, []int{1}
}

func (m *AccID) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AccID.Unmarshal(m, b)
}
func (m *AccID) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AccID.Marshal(b, m, deterministic)
}
func (m *AccID) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AccID.Merge(m, src)
}
func (m *AccID) XXX_Size() int {
	return xxx_messageInfo_AccID.Size(m)
}
func (m *AccID) XXX_DiscardUnknown() {
	xxx_messageInfo_AccID.DiscardUnknown(m)
}

var xxx_messageInfo_AccID proto.InternalMessageInfo

func (m *AccID) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

type Filter struct {
	Kyc                  int32    `protobuf:"varint,1,opt,name=kyc,proto3" json:"kyc,omitempty"`
	TermInDays           int32    `protobuf:"varint,2,opt,name=term_in_days,proto3" json:"term_in_days,omitempty"`
	DueDateEarliest      string   `protobuf:"bytes,3,opt,name=due_date_earliest,proto3" json:"due_date_earliest,omitempty"`
	DueDateLatest        string   `protobuf:"bytes,4,opt,name=due_date_latest,proto3" json:"due_date_latest,omitempty"`
	MinBalance           int64    `protobuf:"varint,5,opt,name=min_balance,proto3" json:"min_balance,omitempty"`
	PageSize             int32    `protobuf:"varint,6,opt,name=page_size,proto3" json:"page_size,omitempty"`
	PageIndex            int32    `protobuf:"varint,7,opt,name=page_index,proto3" json:"page_index,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Filter) Reset()         { *m = Filter{} }
func (m *Filter) String() string { return proto.CompactTextString(m) }
func (*Filter) ProtoMessage()    {}
func (*Filter) Descriptor() ([]byte, []int) {
	return fileDescriptor_7e8dd110a4bce988, []int{2}
}

func (m *Filter) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Filter.Unmarshal(m, b)
}
func (m *Filter) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Filter.Marshal(b, m, deterministic)
}
func (m *Filter) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Filter.Merge(m, src)
}
func (m *Filter) XXX_Size() int {
	return xxx_messageInfo_Filter.Size(m)
}
func (m *Filter) XXX_DiscardUnknown() {
	xxx_messageInfo_Filter.DiscardUnknown(m)
}

var xxx_messageInfo_Filter proto.InternalMessageInfo

func (m *Filter) GetKyc() int32 {
	if m != nil {
		return m.Kyc
	}
	return 0
}

func (m *Filter) GetTermInDays() int32 {
	if m != nil {
		return m.TermInDays
	}
	return 0
}

func (m *Filter) GetDueDateEarliest() string {
	if m != nil {
		return m.DueDateEarliest
	}
	return ""
}

func (m *Filter) GetDueDateLatest() string {
	if m != nil {
		return m.DueDateLatest
	}
	return ""
}

func (m *Filter) GetMinBalance() int64 {
	if m != nil {
		return m.MinBalance
	}
	return 0
}

func (m *Filter) GetPageSize() int32 {
	if m != nil {
		return m.PageSize
	}
	return 0
}

func (m *Filter) GetPageIndex() int32 {
	if m != nil {
		return m.PageIndex
	}
	return 0
}

type SavingAccountList struct {
	AccList              []*SavingAccount `protobuf:"bytes,1,rep,name=acc_list,proto3" json:"acc_list,omitempty"`
	XXX_NoUnkeyedLiteral struct{}         `json:"-"`
	XXX_unrecognized     []byte           `json:"-"`
	XXX_sizecache        int32            `json:"-"`
}

func (m *SavingAccountList) Reset()         { *m = SavingAccountList{} }
func (m *SavingAccountList) String() string { return proto.CompactTextString(m) }
func (*SavingAccountList) ProtoMessage()    {}
func (*SavingAccountList) Descriptor() ([]byte, []int) {
	return fileDescriptor_7e8dd110a4bce988, []int{3}
}

func (m *SavingAccountList) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SavingAccountList.Unmarshal(m, b)
}
func (m *SavingAccountList) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SavingAccountList.Marshal(b, m, deterministic)
}
func (m *SavingAccountList) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SavingAccountList.Merge(m, src)
}
func (m *SavingAccountList) XXX_Size() int {
	return xxx_messageInfo_SavingAccountList.Size(m)
}
func (m *SavingAccountList) XXX_DiscardUnknown() {
	xxx_messageInfo_SavingAccountList.DiscardUnknown(m)
}

var xxx_messageInfo_SavingAccountList proto.InternalMessageInfo

func (m *SavingAccountList) GetAccList() []*SavingAccount {
	if m != nil {
		return m.AccList
	}
	return nil
}

type OpenSavingsAccountRequest struct {
	UserId               string   `protobuf:"bytes,1,opt,name=user_id,proto3" json:"user_id,omitempty"`
	IdCardNumber         string   `protobuf:"bytes,2,opt,name=id_card_number,proto3" json:"id_card_number,omitempty"`
	Balance              int64    `protobuf:"varint,3,opt,name=balance,proto3" json:"balance,omitempty"`
	TermType             string   `protobuf:"bytes,4,opt,name=term_type,proto3" json:"term_type,omitempty"`
	Term                 int32    `protobuf:"varint,5,opt,name=term,proto3" json:"term,omitempty"`
	CreatedDate          string   `protobuf:"bytes,6,opt,name=created_date,proto3" json:"created_date,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *OpenSavingsAccountRequest) Reset()         { *m = OpenSavingsAccountRequest{} }
func (m *OpenSavingsAccountRequest) String() string { return proto.CompactTextString(m) }
func (*OpenSavingsAccountRequest) ProtoMessage()    {}
func (*OpenSavingsAccountRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_7e8dd110a4bce988, []int{4}
}

func (m *OpenSavingsAccountRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_OpenSavingsAccountRequest.Unmarshal(m, b)
}
func (m *OpenSavingsAccountRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_OpenSavingsAccountRequest.Marshal(b, m, deterministic)
}
func (m *OpenSavingsAccountRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_OpenSavingsAccountRequest.Merge(m, src)
}
func (m *OpenSavingsAccountRequest) XXX_Size() int {
	return xxx_messageInfo_OpenSavingsAccountRequest.Size(m)
}
func (m *OpenSavingsAccountRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_OpenSavingsAccountRequest.DiscardUnknown(m)
}

var xxx_messageInfo_OpenSavingsAccountRequest proto.InternalMessageInfo

func (m *OpenSavingsAccountRequest) GetUserId() string {
	if m != nil {
		return m.UserId
	}
	return ""
}

func (m *OpenSavingsAccountRequest) GetIdCardNumber() string {
	if m != nil {
		return m.IdCardNumber
	}
	return ""
}

func (m *OpenSavingsAccountRequest) GetBalance() int64 {
	if m != nil {
		return m.Balance
	}
	return 0
}

func (m *OpenSavingsAccountRequest) GetTermType() string {
	if m != nil {
		return m.TermType
	}
	return ""
}

func (m *OpenSavingsAccountRequest) GetTerm() int32 {
	if m != nil {
		return m.Term
	}
	return 0
}

func (m *OpenSavingsAccountRequest) GetCreatedDate() string {
	if m != nil {
		return m.CreatedDate
	}
	return ""
}

type OpenSavingsAccountResponse struct {
	Success              bool     `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
	UserId               string   `protobuf:"bytes,2,opt,name=user_id,proto3" json:"user_id,omitempty"`
	AccountId            string   `protobuf:"bytes,3,opt,name=account_id,proto3" json:"account_id,omitempty"`
	Balance              int64    `protobuf:"varint,4,opt,name=balance,proto3" json:"balance,omitempty"`
	Rate                 float32  `protobuf:"fixed32,5,opt,name=rate,proto3" json:"rate,omitempty"`
	CreatedDate          string   `protobuf:"bytes,6,opt,name=created_date,proto3" json:"created_date,omitempty"`
	DueDate              string   `protobuf:"bytes,7,opt,name=due_date,proto3" json:"due_date,omitempty"`
	ExpectedInterest     int64    `protobuf:"varint,8,opt,name=expected_interest,proto3" json:"expected_interest,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *OpenSavingsAccountResponse) Reset()         { *m = OpenSavingsAccountResponse{} }
func (m *OpenSavingsAccountResponse) String() string { return proto.CompactTextString(m) }
func (*OpenSavingsAccountResponse) ProtoMessage()    {}
func (*OpenSavingsAccountResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_7e8dd110a4bce988, []int{5}
}

func (m *OpenSavingsAccountResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_OpenSavingsAccountResponse.Unmarshal(m, b)
}
func (m *OpenSavingsAccountResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_OpenSavingsAccountResponse.Marshal(b, m, deterministic)
}
func (m *OpenSavingsAccountResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_OpenSavingsAccountResponse.Merge(m, src)
}
func (m *OpenSavingsAccountResponse) XXX_Size() int {
	return xxx_messageInfo_OpenSavingsAccountResponse.Size(m)
}
func (m *OpenSavingsAccountResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_OpenSavingsAccountResponse.DiscardUnknown(m)
}

var xxx_messageInfo_OpenSavingsAccountResponse proto.InternalMessageInfo

func (m *OpenSavingsAccountResponse) GetSuccess() bool {
	if m != nil {
		return m.Success
	}
	return false
}

func (m *OpenSavingsAccountResponse) GetUserId() string {
	if m != nil {
		return m.UserId
	}
	return ""
}

func (m *OpenSavingsAccountResponse) GetAccountId() string {
	if m != nil {
		return m.AccountId
	}
	return ""
}

func (m *OpenSavingsAccountResponse) GetBalance() int64 {
	if m != nil {
		return m.Balance
	}
	return 0
}

func (m *OpenSavingsAccountResponse) GetRate() float32 {
	if m != nil {
		return m.Rate
	}
	return 0
}

func (m *OpenSavingsAccountResponse) GetCreatedDate() string {
	if m != nil {
		return m.CreatedDate
	}
	return ""
}

func (m *OpenSavingsAccountResponse) GetDueDate() string {
	if m != nil {
		return m.DueDate
	}
	return ""
}

func (m *OpenSavingsAccountResponse) GetExpectedInterest() int64 {
	if m != nil {
		return m.ExpectedInterest
	}
	return 0
}

type WithdrawalRequest struct {
	UserId               string   `protobuf:"bytes,1,opt,name=user_id,proto3" json:"user_id,omitempty"`
	AccountId            string   `protobuf:"bytes,2,opt,name=account_id,proto3" json:"account_id,omitempty"`
	Amount               int64    `protobuf:"varint,3,opt,name=amount,proto3" json:"amount,omitempty"`
	Date                 string   `protobuf:"bytes,4,opt,name=date,proto3" json:"date,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *WithdrawalRequest) Reset()         { *m = WithdrawalRequest{} }
func (m *WithdrawalRequest) String() string { return proto.CompactTextString(m) }
func (*WithdrawalRequest) ProtoMessage()    {}
func (*WithdrawalRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_7e8dd110a4bce988, []int{6}
}

func (m *WithdrawalRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_WithdrawalRequest.Unmarshal(m, b)
}
func (m *WithdrawalRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_WithdrawalRequest.Marshal(b, m, deterministic)
}
func (m *WithdrawalRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_WithdrawalRequest.Merge(m, src)
}
func (m *WithdrawalRequest) XXX_Size() int {
	return xxx_messageInfo_WithdrawalRequest.Size(m)
}
func (m *WithdrawalRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_WithdrawalRequest.DiscardUnknown(m)
}

var xxx_messageInfo_WithdrawalRequest proto.InternalMessageInfo

func (m *WithdrawalRequest) GetUserId() string {
	if m != nil {
		return m.UserId
	}
	return ""
}

func (m *WithdrawalRequest) GetAccountId() string {
	if m != nil {
		return m.AccountId
	}
	return ""
}

func (m *WithdrawalRequest) GetAmount() int64 {
	if m != nil {
		return m.Amount
	}
	return 0
}

func (m *WithdrawalRequest) GetDate() string {
	if m != nil {
		return m.Date
	}
	return ""
}

type WithdrawalResponse struct {
	Success              bool           `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
	Acc                  *SavingAccount `protobuf:"bytes,2,opt,name=acc,proto3" json:"acc,omitempty"`
	WithdrawnAmount      int64          `protobuf:"varint,3,opt,name=withdrawn_amount,proto3" json:"withdrawn_amount,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *WithdrawalResponse) Reset()         { *m = WithdrawalResponse{} }
func (m *WithdrawalResponse) String() string { return proto.CompactTextString(m) }
func (*WithdrawalResponse) ProtoMessage()    {}
func (*WithdrawalResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_7e8dd110a4bce988, []int{7}
}

func (m *WithdrawalResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_WithdrawalResponse.Unmarshal(m, b)
}
func (m *WithdrawalResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_WithdrawalResponse.Marshal(b, m, deterministic)
}
func (m *WithdrawalResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_WithdrawalResponse.Merge(m, src)
}
func (m *WithdrawalResponse) XXX_Size() int {
	return xxx_messageInfo_WithdrawalResponse.Size(m)
}
func (m *WithdrawalResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_WithdrawalResponse.DiscardUnknown(m)
}

var xxx_messageInfo_WithdrawalResponse proto.InternalMessageInfo

func (m *WithdrawalResponse) GetSuccess() bool {
	if m != nil {
		return m.Success
	}
	return false
}

func (m *WithdrawalResponse) GetAcc() *SavingAccount {
	if m != nil {
		return m.Acc
	}
	return nil
}

func (m *WithdrawalResponse) GetWithdrawnAmount() int64 {
	if m != nil {
		return m.WithdrawnAmount
	}
	return 0
}

type AccountInquiryRequest struct {
	UserId               string   `protobuf:"bytes,1,opt,name=user_id,proto3" json:"user_id,omitempty"`
	AccountId            string   `protobuf:"bytes,2,opt,name=account_id,proto3" json:"account_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AccountInquiryRequest) Reset()         { *m = AccountInquiryRequest{} }
func (m *AccountInquiryRequest) String() string { return proto.CompactTextString(m) }
func (*AccountInquiryRequest) ProtoMessage()    {}
func (*AccountInquiryRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_7e8dd110a4bce988, []int{8}
}

func (m *AccountInquiryRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AccountInquiryRequest.Unmarshal(m, b)
}
func (m *AccountInquiryRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AccountInquiryRequest.Marshal(b, m, deterministic)
}
func (m *AccountInquiryRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AccountInquiryRequest.Merge(m, src)
}
func (m *AccountInquiryRequest) XXX_Size() int {
	return xxx_messageInfo_AccountInquiryRequest.Size(m)
}
func (m *AccountInquiryRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_AccountInquiryRequest.DiscardUnknown(m)
}

var xxx_messageInfo_AccountInquiryRequest proto.InternalMessageInfo

func (m *AccountInquiryRequest) GetUserId() string {
	if m != nil {
		return m.UserId
	}
	return ""
}

func (m *AccountInquiryRequest) GetAccountId() string {
	if m != nil {
		return m.AccountId
	}
	return ""
}

func init() {
	proto.RegisterType((*SavingAccount)(nil), "protobuf.SavingAccount")
	proto.RegisterType((*AccID)(nil), "protobuf.AccID")
	proto.RegisterType((*Filter)(nil), "protobuf.Filter")
	proto.RegisterType((*SavingAccountList)(nil), "protobuf.SavingAccountList")
	proto.RegisterType((*OpenSavingsAccountRequest)(nil), "protobuf.OpenSavingsAccountRequest")
	proto.RegisterType((*OpenSavingsAccountResponse)(nil), "protobuf.OpenSavingsAccountResponse")
	proto.RegisterType((*WithdrawalRequest)(nil), "protobuf.WithdrawalRequest")
	proto.RegisterType((*WithdrawalResponse)(nil), "protobuf.WithdrawalResponse")
	proto.RegisterType((*AccountInquiryRequest)(nil), "protobuf.AccountInquiryRequest")
}

func init() { proto.RegisterFile("saving.proto", fileDescriptor_7e8dd110a4bce988) }

var fileDescriptor_7e8dd110a4bce988 = []byte{
	// 656 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x54, 0xc1, 0x4e, 0xdb, 0x40,
	0x10, 0x95, 0x1d, 0x12, 0x92, 0x01, 0x52, 0xb2, 0x52, 0xc1, 0x85, 0x8a, 0x46, 0x3e, 0x54, 0x69,
	0x85, 0x38, 0xc0, 0x17, 0x80, 0x28, 0x2d, 0x55, 0xa5, 0xaa, 0x46, 0xb4, 0x47, 0x6b, 0xb3, 0x9e,
	0xc2, 0xaa, 0xc9, 0xc6, 0xec, 0xae, 0x01, 0xf7, 0xde, 0x1b, 0x52, 0xbf, 0xac, 0xbf, 0xd3, 0x73,
	0xb5, 0x6b, 0x3b, 0xb1, 0x63, 0x02, 0x95, 0xda, 0x53, 0x76, 0xde, 0x9b, 0x9d, 0x9d, 0x37, 0xf3,
	0x1c, 0x58, 0x55, 0xf4, 0x9a, 0x8b, 0x8b, 0xbd, 0x58, 0x4e, 0xf4, 0x84, 0xb4, 0xed, 0xcf, 0x30,
	0xf9, 0xea, 0xff, 0x74, 0x61, 0xed, 0xcc, 0x52, 0x87, 0x8c, 0x4d, 0x12, 0xa1, 0x49, 0x17, 0x5c,
	0x1e, 0x79, 0x4e, 0xdf, 0x19, 0x74, 0x02, 0x97, 0x47, 0xc4, 0x83, 0xe5, 0x44, 0xa1, 0x0c, 0x79,
	0xe4, 0xb9, 0x16, 0x2c, 0x42, 0xc3, 0x0c, 0xe9, 0x88, 0x0a, 0x86, 0x5e, 0xa3, 0xef, 0x0c, 0x1a,
	0x41, 0x11, 0x92, 0xe7, 0xd0, 0xd1, 0x28, 0xc7, 0xa1, 0x4e, 0x63, 0xf4, 0x96, 0xec, 0xad, 0x19,
	0x40, 0x08, 0x2c, 0x99, 0xc0, 0x6b, 0xf6, 0x9d, 0x41, 0x33, 0xb0, 0x67, 0xe2, 0xc3, 0xaa, 0x4d,
	0xe0, 0x22, 0x8c, 0x68, 0xaa, 0xbc, 0x96, 0xe5, 0x2a, 0x98, 0xc9, 0x61, 0x12, 0xa9, 0xc6, 0x28,
	0x8c, 0xa8, 0x46, 0x6f, 0xd9, 0x16, 0xae, 0x60, 0x64, 0x0b, 0xda, 0x51, 0x82, 0x19, 0xdf, 0xb6,
	0xfc, 0x34, 0x36, 0xef, 0x4a, 0x83, 0x77, 0xfa, 0xce, 0xc0, 0x0d, 0xec, 0x99, 0xac, 0x43, 0xe3,
	0x5b, 0xca, 0x3c, 0xb0, 0xcf, 0x99, 0xa3, 0xbf, 0x09, 0xcd, 0x43, 0xc6, 0x4e, 0x8f, 0xe7, 0x07,
	0xe1, 0xff, 0x76, 0xa0, 0x75, 0xc2, 0x47, 0x1a, 0x65, 0x71, 0xcb, 0x99, 0xde, 0xaa, 0xf5, 0xef,
	0xde, 0xd3, 0xff, 0x2e, 0xf4, 0x8a, 0x5e, 0x42, 0xa4, 0x72, 0xc4, 0x51, 0x69, 0x3b, 0xb9, 0x4e,
	0x50, 0x27, 0xc8, 0x00, 0x9e, 0x4c, 0xc1, 0x11, 0xd5, 0x26, 0x37, 0x9b, 0xe4, 0x3c, 0x4c, 0xfa,
	0xb0, 0x32, 0xe6, 0x22, 0x2c, 0x76, 0xd1, 0xb4, 0xbb, 0x28, 0x43, 0x66, 0x1f, 0x31, 0xbd, 0xc0,
	0x50, 0xf1, 0xef, 0x98, 0x8f, 0x76, 0x06, 0x90, 0x1d, 0x00, 0x1b, 0x70, 0x11, 0xe1, 0xad, 0x9d,
	0x6a, 0x33, 0x28, 0x21, 0xfe, 0x3b, 0xe8, 0x55, 0x2c, 0xf2, 0x81, 0x2b, 0x4d, 0x0e, 0xa0, 0x4d,
	0x19, 0x0b, 0x47, 0x5c, 0x69, 0xcf, 0xe9, 0x37, 0x06, 0x2b, 0xfb, 0x9b, 0x7b, 0x85, 0xab, 0xf6,
	0x2a, 0xe9, 0xc1, 0x34, 0xd1, 0xff, 0xe5, 0xc0, 0xb3, 0x8f, 0x31, 0x8a, 0x8c, 0x57, 0x45, 0x02,
	0x5e, 0x25, 0x46, 0x47, 0xc9, 0x69, 0x4e, 0xd5, 0x69, 0x2f, 0xa1, 0xcb, 0xa3, 0x90, 0x51, 0x19,
	0x85, 0x22, 0x19, 0x0f, 0x51, 0xe6, 0x56, 0x9c, 0x43, 0xff, 0xb7, 0x23, 0x2b, 0x6e, 0x6b, 0xd5,
	0xdd, 0xe6, 0xdf, 0xb9, 0xb0, 0x75, 0x9f, 0x1e, 0x15, 0x4f, 0x84, 0x42, 0xd3, 0x8e, 0x4a, 0x18,
	0x43, 0xa5, 0xac, 0xa0, 0x76, 0x50, 0x84, 0x0f, 0x7c, 0x54, 0x3b, 0x00, 0x34, 0x2b, 0x63, 0xc8,
	0xcc, 0x1d, 0x25, 0xa4, 0x2c, 0x71, 0xa9, 0x2a, 0xb1, 0xb0, 0x77, 0xb3, 0x64, 0xef, 0xbf, 0x10,
	0x51, 0xf9, 0x64, 0x96, 0xe7, 0x3e, 0x99, 0x5d, 0xe8, 0xe1, 0x6d, 0x8c, 0xcc, 0x24, 0x73, 0xa1,
	0x51, 0x1a, 0x1b, 0xb6, 0xed, 0xbb, 0x75, 0xc2, 0x4f, 0xa1, 0xf7, 0x85, 0xeb, 0xcb, 0x48, 0xd2,
	0x1b, 0x3a, 0x7a, 0x7c, 0xab, 0x55, 0xa9, 0x6e, 0x4d, 0xea, 0x06, 0xb4, 0xe8, 0xd8, 0x04, 0xf9,
	0x32, 0xf3, 0xc8, 0x08, 0xb5, 0xcd, 0x66, 0x6b, 0xb4, 0x67, 0xff, 0x87, 0x03, 0xa4, 0xfc, 0xf6,
	0xa3, 0x1b, 0x78, 0x05, 0x0d, 0xca, 0x98, 0x7d, 0xf5, 0x01, 0xeb, 0x9a, 0x1c, 0xf2, 0x1a, 0xd6,
	0x6f, 0xf2, 0xd2, 0x22, 0xac, 0x74, 0x54, 0xc3, 0xfd, 0x4f, 0xf0, 0x34, 0xbf, 0x7b, 0x2a, 0xae,
	0x12, 0x2e, 0xd3, 0x7f, 0x1e, 0xc3, 0xfe, 0x5d, 0x03, 0xba, 0xb9, 0xc1, 0xce, 0x50, 0x5e, 0x73,
	0x86, 0xe4, 0x04, 0x48, 0xdd, 0x76, 0x64, 0x91, 0x8a, 0xad, 0x45, 0x04, 0x79, 0x0f, 0xdd, 0x6a,
	0xb7, 0xe4, 0xc5, 0x2c, 0xf5, 0x5e, 0x1d, 0x8b, 0x6b, 0xbd, 0x81, 0xb5, 0xf3, 0xd8, 0xec, 0xe2,
	0x28, 0xf7, 0xe3, 0xf6, 0x2c, 0xb3, 0xe6, 0x8a, 0xc5, 0x65, 0x3e, 0xc3, 0xc6, 0x19, 0x52, 0xc9,
	0x2e, 0x73, 0x40, 0x1d, 0xa5, 0xe7, 0x0a, 0xe5, 0xe9, 0xf1, 0xe3, 0xad, 0x6d, 0x2f, 0xa8, 0x69,
	0xff, 0xaf, 0xde, 0xd6, 0xeb, 0x16, 0x7f, 0xe6, 0xb3, 0x6b, 0x19, 0xf2, 0x60, 0xa1, 0x61, 0xcb,
	0x72, 0x07, 0x7f, 0x02, 0x00, 0x00, 0xff, 0xff, 0x25, 0x5d, 0xdd, 0x36, 0x52, 0x07, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// SavingsServiceClient is the client API for SavingsService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type SavingsServiceClient interface {
	OpenSavingsAccount(ctx context.Context, in *SavingAccount, opts ...grpc.CallOption) (*SavingAccount, error)
	AccountInquiry(ctx context.Context, in *AccountInquiryRequest, opts ...grpc.CallOption) (*SavingAccount, error)
	UpdateBalance(ctx context.Context, in *WithdrawalRequest, opts ...grpc.CallOption) (*SavingAccount, error)
	SearchAccountsByUserID(ctx context.Context, in *AccountInquiryRequest, opts ...grpc.CallOption) (*SavingAccountList, error)
	SearchAccountsByFilter(ctx context.Context, in *Filter, opts ...grpc.CallOption) (*SavingAccountList, error)
}

type savingsServiceClient struct {
	cc *grpc.ClientConn
}

func NewSavingsServiceClient(cc *grpc.ClientConn) SavingsServiceClient {
	return &savingsServiceClient{cc}
}

func (c *savingsServiceClient) OpenSavingsAccount(ctx context.Context, in *SavingAccount, opts ...grpc.CallOption) (*SavingAccount, error) {
	out := new(SavingAccount)
	err := c.cc.Invoke(ctx, "/protobuf.SavingsService/OpenSavingsAccount", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *savingsServiceClient) AccountInquiry(ctx context.Context, in *AccountInquiryRequest, opts ...grpc.CallOption) (*SavingAccount, error) {
	out := new(SavingAccount)
	err := c.cc.Invoke(ctx, "/protobuf.SavingsService/AccountInquiry", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *savingsServiceClient) UpdateBalance(ctx context.Context, in *WithdrawalRequest, opts ...grpc.CallOption) (*SavingAccount, error) {
	out := new(SavingAccount)
	err := c.cc.Invoke(ctx, "/protobuf.SavingsService/UpdateBalance", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *savingsServiceClient) SearchAccountsByUserID(ctx context.Context, in *AccountInquiryRequest, opts ...grpc.CallOption) (*SavingAccountList, error) {
	out := new(SavingAccountList)
	err := c.cc.Invoke(ctx, "/protobuf.SavingsService/SearchAccountsByUserID", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *savingsServiceClient) SearchAccountsByFilter(ctx context.Context, in *Filter, opts ...grpc.CallOption) (*SavingAccountList, error) {
	out := new(SavingAccountList)
	err := c.cc.Invoke(ctx, "/protobuf.SavingsService/SearchAccountsByFilter", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// SavingsServiceServer is the server API for SavingsService service.
type SavingsServiceServer interface {
	OpenSavingsAccount(context.Context, *SavingAccount) (*SavingAccount, error)
	AccountInquiry(context.Context, *AccountInquiryRequest) (*SavingAccount, error)
	UpdateBalance(context.Context, *WithdrawalRequest) (*SavingAccount, error)
	SearchAccountsByUserID(context.Context, *AccountInquiryRequest) (*SavingAccountList, error)
	SearchAccountsByFilter(context.Context, *Filter) (*SavingAccountList, error)
}

// UnimplementedSavingsServiceServer can be embedded to have forward compatible implementations.
type UnimplementedSavingsServiceServer struct {
}

func (*UnimplementedSavingsServiceServer) OpenSavingsAccount(ctx context.Context, req *SavingAccount) (*SavingAccount, error) {
	return nil, status.Errorf(codes.Unimplemented, "method OpenSavingsAccount not implemented")
}
func (*UnimplementedSavingsServiceServer) AccountInquiry(ctx context.Context, req *AccountInquiryRequest) (*SavingAccount, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AccountInquiry not implemented")
}
func (*UnimplementedSavingsServiceServer) UpdateBalance(ctx context.Context, req *WithdrawalRequest) (*SavingAccount, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateBalance not implemented")
}
func (*UnimplementedSavingsServiceServer) SearchAccountsByUserID(ctx context.Context, req *AccountInquiryRequest) (*SavingAccountList, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SearchAccountsByUserID not implemented")
}
func (*UnimplementedSavingsServiceServer) SearchAccountsByFilter(ctx context.Context, req *Filter) (*SavingAccountList, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SearchAccountsByFilter not implemented")
}

func RegisterSavingsServiceServer(s *grpc.Server, srv SavingsServiceServer) {
	s.RegisterService(&_SavingsService_serviceDesc, srv)
}

func _SavingsService_OpenSavingsAccount_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SavingAccount)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SavingsServiceServer).OpenSavingsAccount(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protobuf.SavingsService/OpenSavingsAccount",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SavingsServiceServer).OpenSavingsAccount(ctx, req.(*SavingAccount))
	}
	return interceptor(ctx, in, info, handler)
}

func _SavingsService_AccountInquiry_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AccountInquiryRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SavingsServiceServer).AccountInquiry(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protobuf.SavingsService/AccountInquiry",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SavingsServiceServer).AccountInquiry(ctx, req.(*AccountInquiryRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SavingsService_UpdateBalance_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(WithdrawalRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SavingsServiceServer).UpdateBalance(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protobuf.SavingsService/UpdateBalance",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SavingsServiceServer).UpdateBalance(ctx, req.(*WithdrawalRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SavingsService_SearchAccountsByUserID_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AccountInquiryRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SavingsServiceServer).SearchAccountsByUserID(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protobuf.SavingsService/SearchAccountsByUserID",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SavingsServiceServer).SearchAccountsByUserID(ctx, req.(*AccountInquiryRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SavingsService_SearchAccountsByFilter_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Filter)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SavingsServiceServer).SearchAccountsByFilter(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protobuf.SavingsService/SearchAccountsByFilter",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SavingsServiceServer).SearchAccountsByFilter(ctx, req.(*Filter))
	}
	return interceptor(ctx, in, info, handler)
}

var _SavingsService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "protobuf.SavingsService",
	HandlerType: (*SavingsServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "OpenSavingsAccount",
			Handler:    _SavingsService_OpenSavingsAccount_Handler,
		},
		{
			MethodName: "AccountInquiry",
			Handler:    _SavingsService_AccountInquiry_Handler,
		},
		{
			MethodName: "UpdateBalance",
			Handler:    _SavingsService_UpdateBalance_Handler,
		},
		{
			MethodName: "SearchAccountsByUserID",
			Handler:    _SavingsService_SearchAccountsByUserID_Handler,
		},
		{
			MethodName: "SearchAccountsByFilter",
			Handler:    _SavingsService_SearchAccountsByFilter_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "saving.proto",
}
