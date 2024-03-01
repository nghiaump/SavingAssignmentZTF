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
	PageSize             int32            `protobuf:"varint,2,opt,name=page_size,json=pageSize,proto3" json:"page_size,omitempty"`
	PageIndex            int32            `protobuf:"varint,3,opt,name=page_index,json=pageIndex,proto3" json:"page_index,omitempty"`
	AggTotalHits         int64            `protobuf:"varint,4,opt,name=agg_total_hits,json=aggTotalHits,proto3" json:"agg_total_hits,omitempty"`
	AggTotalBalance      int64            `protobuf:"varint,5,opt,name=agg_total_balance,json=aggTotalBalance,proto3" json:"agg_total_balance,omitempty"`
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

func (m *SavingAccountList) GetPageSize() int32 {
	if m != nil {
		return m.PageSize
	}
	return 0
}

func (m *SavingAccountList) GetPageIndex() int32 {
	if m != nil {
		return m.PageIndex
	}
	return 0
}

func (m *SavingAccountList) GetAggTotalHits() int64 {
	if m != nil {
		return m.AggTotalHits
	}
	return 0
}

func (m *SavingAccountList) GetAggTotalBalance() int64 {
	if m != nil {
		return m.AggTotalBalance
	}
	return 0
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
	// 736 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x55, 0x4f, 0x4f, 0xdb, 0x48,
	0x14, 0x97, 0x6d, 0x12, 0x92, 0x07, 0x04, 0x32, 0xd2, 0x82, 0x17, 0x76, 0xd9, 0xc8, 0x5a, 0x55,
	0x29, 0x42, 0x1c, 0xe0, 0xdc, 0x03, 0x28, 0xa5, 0x4d, 0x55, 0xa9, 0xaa, 0x53, 0xda, 0xa3, 0x35,
	0x19, 0x4f, 0xc3, 0xa8, 0x89, 0x13, 0x66, 0xc6, 0x80, 0x7b, 0xef, 0xad, 0x52, 0x3f, 0x58, 0xd5,
	0x6b, 0x3f, 0x4a, 0xcf, 0xd5, 0x8c, 0xed, 0xd8, 0x13, 0x13, 0xa8, 0xd4, 0x9e, 0x32, 0xef, 0xf7,
	0xfe, 0xfd, 0xde, 0xbf, 0x18, 0xd6, 0x05, 0xbe, 0x66, 0xd1, 0xe8, 0x68, 0xc6, 0xa7, 0x72, 0x8a,
	0x1a, 0xfa, 0x67, 0x18, 0xbf, 0xf7, 0xbe, 0xd8, 0xb0, 0x31, 0xd0, 0xaa, 0x53, 0x42, 0xa6, 0x71,
	0x24, 0x51, 0x0b, 0x6c, 0x16, 0xba, 0x56, 0xc7, 0xea, 0x36, 0x7d, 0x9b, 0x85, 0xc8, 0x85, 0xd5,
	0x58, 0x50, 0x1e, 0xb0, 0xd0, 0xb5, 0x35, 0x98, 0x8b, 0x4a, 0x33, 0xc4, 0x63, 0x1c, 0x11, 0xea,
	0x3a, 0x1d, 0xab, 0xeb, 0xf8, 0xb9, 0x88, 0xfe, 0x81, 0xa6, 0xa4, 0x7c, 0x12, 0xc8, 0x64, 0x46,
	0xdd, 0x15, 0xed, 0x55, 0x00, 0x08, 0xc1, 0x8a, 0x12, 0xdc, 0x5a, 0xc7, 0xea, 0xd6, 0x7c, 0xfd,
	0x46, 0x1e, 0xac, 0x6b, 0x03, 0x16, 0x05, 0x21, 0x4e, 0x84, 0x5b, 0xd7, 0x3a, 0x03, 0x53, 0x36,
	0x84, 0x53, 0x2c, 0x69, 0x18, 0x84, 0x58, 0x52, 0x77, 0x55, 0x07, 0x36, 0x30, 0xb4, 0x0b, 0x8d,
	0x30, 0xa6, 0xa9, 0xbe, 0xa1, 0xf5, 0x73, 0x59, 0xe5, 0xe5, 0x0a, 0x6f, 0x76, 0xac, 0xae, 0xed,
	0xeb, 0x37, 0xda, 0x02, 0xe7, 0x43, 0x42, 0x5c, 0xd0, 0xe9, 0xd4, 0xd3, 0xdb, 0x81, 0xda, 0x29,
	0x21, 0xfd, 0xde, 0x62, 0x23, 0xbc, 0x1f, 0x16, 0xd4, 0xcf, 0xd9, 0x58, 0x52, 0x9e, 0x7b, 0x59,
	0x73, 0xaf, 0x0a, 0x7f, 0xfb, 0x0e, 0xfe, 0x87, 0xd0, 0xce, 0xb9, 0x04, 0x14, 0xf3, 0x31, 0xa3,
	0x42, 0xea, 0xce, 0x35, 0xfd, 0xaa, 0x02, 0x75, 0x61, 0x73, 0x0e, 0x8e, 0xb1, 0x54, 0xb6, 0x69,
	0x27, 0x17, 0x61, 0xd4, 0x81, 0xb5, 0x09, 0x8b, 0x82, 0x7c, 0x16, 0x35, 0x3d, 0x8b, 0x32, 0xa4,
	0xe6, 0x31, 0xc3, 0x23, 0x1a, 0x08, 0xf6, 0x91, 0x66, 0xad, 0x2d, 0x00, 0xb4, 0x0f, 0xa0, 0x05,
	0x16, 0x85, 0xf4, 0x56, 0x77, 0xb5, 0xe6, 0x97, 0x10, 0xef, 0xbb, 0x05, 0x6d, 0x63, 0x47, 0x5e,
	0x32, 0x21, 0xd1, 0x09, 0x34, 0x30, 0x21, 0xc1, 0x98, 0x09, 0xe9, 0x5a, 0x1d, 0xa7, 0xbb, 0x76,
	0xbc, 0x73, 0x94, 0xaf, 0xd5, 0x91, 0x61, 0xee, 0xcf, 0x0d, 0xd1, 0x5e, 0x99, 0x48, 0xda, 0xa3,
	0x86, 0x02, 0x06, 0x8a, 0xc7, 0xbf, 0x06, 0x0f, 0xa7, 0xa0, 0xd9, 0x57, 0x00, 0xfa, 0x1f, 0x5a,
	0x78, 0x34, 0x0a, 0xe4, 0x54, 0xe2, 0x71, 0x70, 0xc9, 0xa4, 0xd0, 0xfd, 0x70, 0xfc, 0x75, 0x3c,
	0x1a, 0xbd, 0x51, 0xe0, 0x73, 0x26, 0x05, 0x3a, 0x80, 0x76, 0x61, 0x65, 0xb6, 0x64, 0x33, 0x37,
	0x3c, 0x4b, 0x61, 0xef, 0x9b, 0x05, 0x7f, 0xbf, 0x9a, 0xd1, 0x28, 0x65, 0x2b, 0x72, 0xba, 0xf4,
	0x2a, 0x56, 0x6d, 0x2d, 0x2d, 0xbe, 0x65, 0x2e, 0xfe, 0x23, 0x68, 0xb1, 0x30, 0x20, 0x98, 0x87,
	0x41, 0x14, 0x4f, 0x86, 0x94, 0x67, 0x97, 0xb1, 0x80, 0xfe, 0xe9, 0x03, 0x31, 0x96, 0xbf, 0x5e,
	0x5d, 0x7e, 0xef, 0xb3, 0x0d, 0xbb, 0x77, 0xd5, 0x23, 0x66, 0xd3, 0x48, 0x50, 0x45, 0x47, 0xc4,
	0x84, 0x50, 0x21, 0x74, 0x41, 0x0d, 0x3f, 0x17, 0xef, 0xb9, 0xf1, 0x7d, 0x00, 0x9c, 0x86, 0x51,
	0xca, 0x74, 0x59, 0x4b, 0x48, 0xb9, 0xc4, 0x15, 0xb3, 0xc4, 0xfc, 0xda, 0x6a, 0xa5, 0x6b, 0xfb,
	0x85, 0x22, 0x8c, 0x0b, 0x5e, 0x5d, 0xb8, 0xe0, 0x43, 0x68, 0xd3, 0xdb, 0x19, 0x25, 0xca, 0x98,
	0x45, 0x92, 0x72, 0x75, 0x15, 0x0d, 0x9d, 0xb7, 0xaa, 0xf0, 0x12, 0x68, 0xbf, 0x63, 0xf2, 0x32,
	0xe4, 0xf8, 0x06, 0x8f, 0x1f, 0x9e, 0xaa, 0x59, 0xaa, 0x5d, 0x29, 0x75, 0x1b, 0xea, 0x78, 0xa2,
	0x84, 0x6c, 0x98, 0x99, 0xa4, 0x0a, 0xd5, 0x64, 0xd3, 0x31, 0xea, 0xb7, 0xf7, 0xc9, 0x02, 0x54,
	0xce, 0xfd, 0xe0, 0x04, 0x1e, 0x83, 0x83, 0x09, 0xd1, 0x59, 0xef, 0x39, 0x24, 0x65, 0x83, 0x0e,
	0x60, 0xeb, 0x26, 0x0b, 0x1d, 0x05, 0x06, 0xa3, 0x0a, 0xee, 0xbd, 0x86, 0xbf, 0x32, 0xdf, 0x7e,
	0x74, 0x15, 0x33, 0x9e, 0xfc, 0x76, 0x1b, 0x8e, 0xbf, 0x3a, 0xd0, 0xca, 0x16, 0x6c, 0x40, 0xf9,
	0x35, 0x23, 0x14, 0x9d, 0x03, 0xaa, 0xae, 0x1d, 0x5a, 0x56, 0xc5, 0xee, 0x32, 0x05, 0x7a, 0x01,
	0x2d, 0x93, 0x2d, 0xfa, 0xaf, 0x30, 0xbd, 0xb3, 0x8e, 0xe5, 0xb1, 0x9e, 0xc2, 0xc6, 0xc5, 0x4c,
	0xcd, 0x22, 0x3b, 0x76, 0xb4, 0x57, 0x58, 0x56, 0xb6, 0x62, 0x79, 0x98, 0x27, 0xd0, 0x1e, 0x50,
	0xcc, 0xc9, 0x65, 0x06, 0x9c, 0x25, 0xfd, 0x1e, 0xda, 0x34, 0x58, 0xf5, 0x7b, 0xcb, 0xdd, 0xdf,
	0xc2, 0xb6, 0xe1, 0x2e, 0xce, 0x92, 0x0b, 0x41, 0x79, 0xbf, 0xf7, 0x70, 0x65, 0x7b, 0x4b, 0x62,
	0xea, 0x3f, 0xdf, 0x67, 0xd5, 0xb8, 0xf9, 0xa7, 0xa9, 0x70, 0x4b, 0x91, 0x7b, 0x03, 0x0d, 0xeb,
	0x5a, 0x77, 0xf2, 0x33, 0x00, 0x00, 0xff, 0xff, 0xbe, 0x11, 0xbb, 0xda, 0x20, 0x08, 0x00, 0x00,
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
	SearchAccountByID(ctx context.Context, in *AccID, opts ...grpc.CallOption) (*SavingAccount, error)
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

func (c *savingsServiceClient) SearchAccountByID(ctx context.Context, in *AccID, opts ...grpc.CallOption) (*SavingAccount, error) {
	out := new(SavingAccount)
	err := c.cc.Invoke(ctx, "/protobuf.SavingsService/SearchAccountByID", in, out, opts...)
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
	SearchAccountByID(context.Context, *AccID) (*SavingAccount, error)
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
func (*UnimplementedSavingsServiceServer) SearchAccountByID(ctx context.Context, req *AccID) (*SavingAccount, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SearchAccountByID not implemented")
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

func _SavingsService_SearchAccountByID_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AccID)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SavingsServiceServer).SearchAccountByID(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protobuf.SavingsService/SearchAccountByID",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SavingsServiceServer).SearchAccountByID(ctx, req.(*AccID))
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
			MethodName: "SearchAccountByID",
			Handler:    _SavingsService_SearchAccountByID_Handler,
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
