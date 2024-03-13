// Code generated by protoc-gen-go. DO NOT EDIT.
// source: mid_saving.proto

package protobuf

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	_ "google.golang.org/genproto/googleapis/api/annotations"
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

func init() { proto.RegisterFile("mid_saving.proto", fileDescriptor_262d7cf1cc6f51b8) }

var fileDescriptor_262d7cf1cc6f51b8 = []byte{
	// 509 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x92, 0xdf, 0x6a, 0x13, 0x41,
	0x14, 0xc6, 0x89, 0x17, 0x22, 0x87, 0x1a, 0xe3, 0xa4, 0xad, 0xe9, 0x36, 0x6d, 0x64, 0x5a, 0x6f,
	0x0a, 0x49, 0xa0, 0xde, 0x89, 0x20, 0x36, 0x41, 0x09, 0xfe, 0x83, 0x06, 0x11, 0xaf, 0x64, 0xb2,
	0x7b, 0x4c, 0x86, 0x26, 0xb3, 0xe9, 0xcc, 0x6c, 0x4b, 0xae, 0x84, 0xe2, 0x1b, 0x78, 0xef, 0x4b,
	0xf9, 0x0a, 0x3e, 0x88, 0xcc, 0xcc, 0xae, 0xbb, 0xd3, 0xed, 0x46, 0xc5, 0xab, 0x85, 0xef, 0xfc,
	0xf9, 0x7d, 0xe7, 0xdb, 0x81, 0xc6, 0x82, 0x47, 0x9f, 0x14, 0xbb, 0xe0, 0x62, 0xda, 0x5b, 0xca,
	0x58, 0xc7, 0xe4, 0x8e, 0xfd, 0x4c, 0x92, 0xcf, 0x01, 0x24, 0x0a, 0xa5, 0x53, 0x83, 0x8d, 0x62,
	0x4f, 0xd0, 0x9e, 0xc6, 0xf1, 0x74, 0x8e, 0x7d, 0xb6, 0xe4, 0x7d, 0x26, 0x44, 0xac, 0x99, 0xe6,
	0xb1, 0x50, 0xae, 0x7a, 0xfc, 0x1d, 0xa0, 0xf1, 0x86, 0x47, 0x63, 0x3b, 0x31, 0x46, 0x79, 0xc1,
	0x43, 0x24, 0x0b, 0xd8, 0x38, 0xc5, 0x29, 0x57, 0x1a, 0xe5, 0x7b, 0x85, 0x92, 0xec, 0xf5, 0x32,
	0x4e, 0xaf, 0xa8, 0x9f, 0xe2, 0x79, 0x82, 0x4a, 0x07, 0xfb, 0x55, 0x65, 0xb5, 0x8c, 0x85, 0x42,
	0xba, 0x7f, 0xf5, 0xe3, 0xe7, 0xb7, 0x5b, 0x2d, 0xda, 0xb4, 0x1e, 0x8c, 0x51, 0xd5, 0x97, 0x69,
	0xe3, 0x93, 0xda, 0x11, 0x39, 0x83, 0xbb, 0x2f, 0x51, 0x0f, 0x12, 0x29, 0x51, 0xe8, 0x57, 0x1f,
	0x07, 0xa4, 0xb0, 0xd0, 0x2b, 0x64, 0xc0, 0x4e, 0x65, 0x3d, 0x25, 0xee, 0x58, 0x62, 0x93, 0xd6,
	0x0b, 0xc4, 0xb3, 0x55, 0x68, 0x60, 0x5f, 0x6b, 0x40, 0xde, 0x2d, 0x51, 0xb8, 0x8b, 0xd5, 0xf3,
	0x30, 0x8c, 0x13, 0xa1, 0xc9, 0x41, 0xbe, 0xb2, 0x5c, 0xcd, 0xb8, 0x87, 0xeb, 0x9b, 0x52, 0xf8,
	0x81, 0x85, 0xef, 0xd1, 0x96, 0x8b, 0xdc, 0x55, 0x55, 0xff, 0x92, 0xeb, 0x59, 0x24, 0xd9, 0x25,
	0x9b, 0x1b, 0x1b, 0x73, 0x80, 0x0f, 0xbf, 0x05, 0xb2, 0x9b, 0x2f, 0xce, 0xd5, 0x8c, 0xda, 0xbe,
	0xb9, 0xf8, 0x6f, 0xb4, 0x7a, 0xea, 0x72, 0x24, 0xce, 0x13, 0x2e, 0x57, 0xa4, 0x10, 0xa1, 0x5f,
	0xc9, 0xa8, 0x0f, 0xf2, 0x06, 0x77, 0x67, 0xda, 0x46, 0x1f, 0x5a, 0x60, 0x40, 0xb7, 0x7c, 0x20,
	0x77, 0xe3, 0x86, 0x16, 0xc1, 0xfd, 0x31, 0x32, 0x19, 0xce, 0xd2, 0x91, 0x93, 0xd5, 0x68, 0x48,
	0xee, 0x79, 0xc0, 0xd1, 0xb0, 0x1a, 0xf0, 0xc8, 0x02, 0x3a, 0x34, 0xf0, 0x01, 0xca, 0xae, 0xec,
	0x4e, 0x56, 0x5d, 0x1e, 0x19, 0xca, 0x17, 0x68, 0x7b, 0x14, 0x65, 0x30, 0x03, 0x26, 0xa3, 0xb7,
	0xc9, 0x62, 0x82, 0x92, 0x6c, 0xe7, 0xfb, 0x8b, 0x7a, 0xb0, 0x5b, 0xc1, 0x7d, 0xcd, 0x95, 0xa6,
	0x47, 0x96, 0x7d, 0x48, 0x3b, 0xd5, 0xec, 0x6e, 0xc8, 0xa4, 0x35, 0x70, 0x55, 0x83, 0xed, 0xeb,
	0x0e, 0xcc, 0xbb, 0x1f, 0x0d, 0xff, 0x9c, 0xee, 0xff, 0x99, 0x30, 0x0f, 0x3a, 0x4d, 0x61, 0x56,
	0xf6, 0xf0, 0x82, 0xcf, 0x35, 0x4a, 0xd2, 0xc8, 0x11, 0x4e, 0x59, 0x0f, 0xed, 0x58, 0xe8, 0x0e,
	0xdd, 0xbc, 0x09, 0x6a, 0x48, 0xc7, 0x50, 0x77, 0x24, 0x73, 0xa3, 0xfd, 0xa5, 0x05, 0x82, 0xbb,
	0x3b, 0xa8, 0xfb, 0x0a, 0x39, 0x81, 0x96, 0x37, 0x13, 0xfd, 0xc5, 0xff, 0xb9, 0xbe, 0xe3, 0x29,
	0x6c, 0x15, 0x77, 0x64, 0x79, 0x0e, 0x49, 0xb3, 0x1c, 0x72, 0xd9, 0xc1, 0x33, 0x68, 0xe6, 0xd3,
	0x79, 0x38, 0x9b, 0x7e, 0x5b, 0x1a, 0x10, 0xf1, 0x55, 0x93, 0xcb, 0xe4, 0xb6, 0x95, 0x1e, 0xff,
	0x0a, 0x00, 0x00, 0xff, 0xff, 0x40, 0x09, 0x35, 0x9b, 0x7d, 0x05, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// MidSavingServiceClient is the client API for MidSavingService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type MidSavingServiceClient interface {
	RegisterUser(ctx context.Context, in *RegisterUserRequest, opts ...grpc.CallOption) (*RegisterUserResponse, error)
	GetCurrentKYC(ctx context.Context, in *GetCurrentKYCRequest, opts ...grpc.CallOption) (*GetCurrentKYCResponse, error)
	OpenSavingsAccount(ctx context.Context, in *OpenSavingsAccountRequest, opts ...grpc.CallOption) (*OpenSavingsAccountResponse, error)
	Withdrawal(ctx context.Context, in *WithdrawalRequest, opts ...grpc.CallOption) (*WithdrawalResponse, error)
	AccountInquiry(ctx context.Context, in *AccountInquiryRequest, opts ...grpc.CallOption) (*SavingAccount, error)
	SearchAccountByID(ctx context.Context, in *AccID, opts ...grpc.CallOption) (*SavingAccount, error)
	SearchAccountsByIDCardNumber(ctx context.Context, in *IDCardNumber, opts ...grpc.CallOption) (*SavingAccountList, error)
	SearchAccountsByUserID(ctx context.Context, in *AccountInquiryRequest, opts ...grpc.CallOption) (*SavingAccountList, error)
	SearchAccountsByFilter(ctx context.Context, in *Filter, opts ...grpc.CallOption) (*SavingAccountList, error)
	SearchUserByID(ctx context.Context, in *UserID, opts ...grpc.CallOption) (*User, error)
	SearchUserByIdCardNumber(ctx context.Context, in *IDCardNumber, opts ...grpc.CallOption) (*User, error)
	SearchUserByAccountID(ctx context.Context, in *AccountID, opts ...grpc.CallOption) (*User, error)
	SearchUsersByFilter(ctx context.Context, in *UserFilter, opts ...grpc.CallOption) (*UserList, error)
}

type midSavingServiceClient struct {
	cc *grpc.ClientConn
}

func NewMidSavingServiceClient(cc *grpc.ClientConn) MidSavingServiceClient {
	return &midSavingServiceClient{cc}
}

func (c *midSavingServiceClient) RegisterUser(ctx context.Context, in *RegisterUserRequest, opts ...grpc.CallOption) (*RegisterUserResponse, error) {
	out := new(RegisterUserResponse)
	err := c.cc.Invoke(ctx, "/protobuf.MidSavingService/RegisterUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *midSavingServiceClient) GetCurrentKYC(ctx context.Context, in *GetCurrentKYCRequest, opts ...grpc.CallOption) (*GetCurrentKYCResponse, error) {
	out := new(GetCurrentKYCResponse)
	err := c.cc.Invoke(ctx, "/protobuf.MidSavingService/GetCurrentKYC", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *midSavingServiceClient) OpenSavingsAccount(ctx context.Context, in *OpenSavingsAccountRequest, opts ...grpc.CallOption) (*OpenSavingsAccountResponse, error) {
	out := new(OpenSavingsAccountResponse)
	err := c.cc.Invoke(ctx, "/protobuf.MidSavingService/OpenSavingsAccount", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *midSavingServiceClient) Withdrawal(ctx context.Context, in *WithdrawalRequest, opts ...grpc.CallOption) (*WithdrawalResponse, error) {
	out := new(WithdrawalResponse)
	err := c.cc.Invoke(ctx, "/protobuf.MidSavingService/Withdrawal", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *midSavingServiceClient) AccountInquiry(ctx context.Context, in *AccountInquiryRequest, opts ...grpc.CallOption) (*SavingAccount, error) {
	out := new(SavingAccount)
	err := c.cc.Invoke(ctx, "/protobuf.MidSavingService/AccountInquiry", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *midSavingServiceClient) SearchAccountByID(ctx context.Context, in *AccID, opts ...grpc.CallOption) (*SavingAccount, error) {
	out := new(SavingAccount)
	err := c.cc.Invoke(ctx, "/protobuf.MidSavingService/SearchAccountByID", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *midSavingServiceClient) SearchAccountsByIDCardNumber(ctx context.Context, in *IDCardNumber, opts ...grpc.CallOption) (*SavingAccountList, error) {
	out := new(SavingAccountList)
	err := c.cc.Invoke(ctx, "/protobuf.MidSavingService/SearchAccountsByIDCardNumber", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *midSavingServiceClient) SearchAccountsByUserID(ctx context.Context, in *AccountInquiryRequest, opts ...grpc.CallOption) (*SavingAccountList, error) {
	out := new(SavingAccountList)
	err := c.cc.Invoke(ctx, "/protobuf.MidSavingService/SearchAccountsByUserID", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *midSavingServiceClient) SearchAccountsByFilter(ctx context.Context, in *Filter, opts ...grpc.CallOption) (*SavingAccountList, error) {
	out := new(SavingAccountList)
	err := c.cc.Invoke(ctx, "/protobuf.MidSavingService/SearchAccountsByFilter", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *midSavingServiceClient) SearchUserByID(ctx context.Context, in *UserID, opts ...grpc.CallOption) (*User, error) {
	out := new(User)
	err := c.cc.Invoke(ctx, "/protobuf.MidSavingService/SearchUserByID", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *midSavingServiceClient) SearchUserByIdCardNumber(ctx context.Context, in *IDCardNumber, opts ...grpc.CallOption) (*User, error) {
	out := new(User)
	err := c.cc.Invoke(ctx, "/protobuf.MidSavingService/SearchUserByIdCardNumber", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *midSavingServiceClient) SearchUserByAccountID(ctx context.Context, in *AccountID, opts ...grpc.CallOption) (*User, error) {
	out := new(User)
	err := c.cc.Invoke(ctx, "/protobuf.MidSavingService/SearchUserByAccountID", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *midSavingServiceClient) SearchUsersByFilter(ctx context.Context, in *UserFilter, opts ...grpc.CallOption) (*UserList, error) {
	out := new(UserList)
	err := c.cc.Invoke(ctx, "/protobuf.MidSavingService/SearchUsersByFilter", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MidSavingServiceServer is the server API for MidSavingService service.
type MidSavingServiceServer interface {
	RegisterUser(context.Context, *RegisterUserRequest) (*RegisterUserResponse, error)
	GetCurrentKYC(context.Context, *GetCurrentKYCRequest) (*GetCurrentKYCResponse, error)
	OpenSavingsAccount(context.Context, *OpenSavingsAccountRequest) (*OpenSavingsAccountResponse, error)
	Withdrawal(context.Context, *WithdrawalRequest) (*WithdrawalResponse, error)
	AccountInquiry(context.Context, *AccountInquiryRequest) (*SavingAccount, error)
	SearchAccountByID(context.Context, *AccID) (*SavingAccount, error)
	SearchAccountsByIDCardNumber(context.Context, *IDCardNumber) (*SavingAccountList, error)
	SearchAccountsByUserID(context.Context, *AccountInquiryRequest) (*SavingAccountList, error)
	SearchAccountsByFilter(context.Context, *Filter) (*SavingAccountList, error)
	SearchUserByID(context.Context, *UserID) (*User, error)
	SearchUserByIdCardNumber(context.Context, *IDCardNumber) (*User, error)
	SearchUserByAccountID(context.Context, *AccountID) (*User, error)
	SearchUsersByFilter(context.Context, *UserFilter) (*UserList, error)
}

// UnimplementedMidSavingServiceServer can be embedded to have forward compatible implementations.
type UnimplementedMidSavingServiceServer struct {
}

func (*UnimplementedMidSavingServiceServer) RegisterUser(ctx context.Context, req *RegisterUserRequest) (*RegisterUserResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RegisterUser not implemented")
}
func (*UnimplementedMidSavingServiceServer) GetCurrentKYC(ctx context.Context, req *GetCurrentKYCRequest) (*GetCurrentKYCResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetCurrentKYC not implemented")
}
func (*UnimplementedMidSavingServiceServer) OpenSavingsAccount(ctx context.Context, req *OpenSavingsAccountRequest) (*OpenSavingsAccountResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method OpenSavingsAccount not implemented")
}
func (*UnimplementedMidSavingServiceServer) Withdrawal(ctx context.Context, req *WithdrawalRequest) (*WithdrawalResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Withdrawal not implemented")
}
func (*UnimplementedMidSavingServiceServer) AccountInquiry(ctx context.Context, req *AccountInquiryRequest) (*SavingAccount, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AccountInquiry not implemented")
}
func (*UnimplementedMidSavingServiceServer) SearchAccountByID(ctx context.Context, req *AccID) (*SavingAccount, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SearchAccountByID not implemented")
}
func (*UnimplementedMidSavingServiceServer) SearchAccountsByIDCardNumber(ctx context.Context, req *IDCardNumber) (*SavingAccountList, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SearchAccountsByIDCardNumber not implemented")
}
func (*UnimplementedMidSavingServiceServer) SearchAccountsByUserID(ctx context.Context, req *AccountInquiryRequest) (*SavingAccountList, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SearchAccountsByUserID not implemented")
}
func (*UnimplementedMidSavingServiceServer) SearchAccountsByFilter(ctx context.Context, req *Filter) (*SavingAccountList, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SearchAccountsByFilter not implemented")
}
func (*UnimplementedMidSavingServiceServer) SearchUserByID(ctx context.Context, req *UserID) (*User, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SearchUserByID not implemented")
}
func (*UnimplementedMidSavingServiceServer) SearchUserByIdCardNumber(ctx context.Context, req *IDCardNumber) (*User, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SearchUserByIdCardNumber not implemented")
}
func (*UnimplementedMidSavingServiceServer) SearchUserByAccountID(ctx context.Context, req *AccountID) (*User, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SearchUserByAccountID not implemented")
}
func (*UnimplementedMidSavingServiceServer) SearchUsersByFilter(ctx context.Context, req *UserFilter) (*UserList, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SearchUsersByFilter not implemented")
}

func RegisterMidSavingServiceServer(s *grpc.Server, srv MidSavingServiceServer) {
	s.RegisterService(&_MidSavingService_serviceDesc, srv)
}

func _MidSavingService_RegisterUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RegisterUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MidSavingServiceServer).RegisterUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protobuf.MidSavingService/RegisterUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MidSavingServiceServer).RegisterUser(ctx, req.(*RegisterUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MidSavingService_GetCurrentKYC_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetCurrentKYCRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MidSavingServiceServer).GetCurrentKYC(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protobuf.MidSavingService/GetCurrentKYC",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MidSavingServiceServer).GetCurrentKYC(ctx, req.(*GetCurrentKYCRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MidSavingService_OpenSavingsAccount_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(OpenSavingsAccountRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MidSavingServiceServer).OpenSavingsAccount(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protobuf.MidSavingService/OpenSavingsAccount",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MidSavingServiceServer).OpenSavingsAccount(ctx, req.(*OpenSavingsAccountRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MidSavingService_Withdrawal_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(WithdrawalRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MidSavingServiceServer).Withdrawal(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protobuf.MidSavingService/Withdrawal",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MidSavingServiceServer).Withdrawal(ctx, req.(*WithdrawalRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MidSavingService_AccountInquiry_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AccountInquiryRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MidSavingServiceServer).AccountInquiry(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protobuf.MidSavingService/AccountInquiry",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MidSavingServiceServer).AccountInquiry(ctx, req.(*AccountInquiryRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MidSavingService_SearchAccountByID_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AccID)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MidSavingServiceServer).SearchAccountByID(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protobuf.MidSavingService/SearchAccountByID",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MidSavingServiceServer).SearchAccountByID(ctx, req.(*AccID))
	}
	return interceptor(ctx, in, info, handler)
}

func _MidSavingService_SearchAccountsByIDCardNumber_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(IDCardNumber)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MidSavingServiceServer).SearchAccountsByIDCardNumber(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protobuf.MidSavingService/SearchAccountsByIDCardNumber",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MidSavingServiceServer).SearchAccountsByIDCardNumber(ctx, req.(*IDCardNumber))
	}
	return interceptor(ctx, in, info, handler)
}

func _MidSavingService_SearchAccountsByUserID_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AccountInquiryRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MidSavingServiceServer).SearchAccountsByUserID(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protobuf.MidSavingService/SearchAccountsByUserID",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MidSavingServiceServer).SearchAccountsByUserID(ctx, req.(*AccountInquiryRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MidSavingService_SearchAccountsByFilter_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Filter)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MidSavingServiceServer).SearchAccountsByFilter(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protobuf.MidSavingService/SearchAccountsByFilter",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MidSavingServiceServer).SearchAccountsByFilter(ctx, req.(*Filter))
	}
	return interceptor(ctx, in, info, handler)
}

func _MidSavingService_SearchUserByID_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserID)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MidSavingServiceServer).SearchUserByID(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protobuf.MidSavingService/SearchUserByID",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MidSavingServiceServer).SearchUserByID(ctx, req.(*UserID))
	}
	return interceptor(ctx, in, info, handler)
}

func _MidSavingService_SearchUserByIdCardNumber_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(IDCardNumber)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MidSavingServiceServer).SearchUserByIdCardNumber(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protobuf.MidSavingService/SearchUserByIdCardNumber",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MidSavingServiceServer).SearchUserByIdCardNumber(ctx, req.(*IDCardNumber))
	}
	return interceptor(ctx, in, info, handler)
}

func _MidSavingService_SearchUserByAccountID_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AccountID)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MidSavingServiceServer).SearchUserByAccountID(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protobuf.MidSavingService/SearchUserByAccountID",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MidSavingServiceServer).SearchUserByAccountID(ctx, req.(*AccountID))
	}
	return interceptor(ctx, in, info, handler)
}

func _MidSavingService_SearchUsersByFilter_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserFilter)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MidSavingServiceServer).SearchUsersByFilter(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protobuf.MidSavingService/SearchUsersByFilter",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MidSavingServiceServer).SearchUsersByFilter(ctx, req.(*UserFilter))
	}
	return interceptor(ctx, in, info, handler)
}

var _MidSavingService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "protobuf.MidSavingService",
	HandlerType: (*MidSavingServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "RegisterUser",
			Handler:    _MidSavingService_RegisterUser_Handler,
		},
		{
			MethodName: "GetCurrentKYC",
			Handler:    _MidSavingService_GetCurrentKYC_Handler,
		},
		{
			MethodName: "OpenSavingsAccount",
			Handler:    _MidSavingService_OpenSavingsAccount_Handler,
		},
		{
			MethodName: "Withdrawal",
			Handler:    _MidSavingService_Withdrawal_Handler,
		},
		{
			MethodName: "AccountInquiry",
			Handler:    _MidSavingService_AccountInquiry_Handler,
		},
		{
			MethodName: "SearchAccountByID",
			Handler:    _MidSavingService_SearchAccountByID_Handler,
		},
		{
			MethodName: "SearchAccountsByIDCardNumber",
			Handler:    _MidSavingService_SearchAccountsByIDCardNumber_Handler,
		},
		{
			MethodName: "SearchAccountsByUserID",
			Handler:    _MidSavingService_SearchAccountsByUserID_Handler,
		},
		{
			MethodName: "SearchAccountsByFilter",
			Handler:    _MidSavingService_SearchAccountsByFilter_Handler,
		},
		{
			MethodName: "SearchUserByID",
			Handler:    _MidSavingService_SearchUserByID_Handler,
		},
		{
			MethodName: "SearchUserByIdCardNumber",
			Handler:    _MidSavingService_SearchUserByIdCardNumber_Handler,
		},
		{
			MethodName: "SearchUserByAccountID",
			Handler:    _MidSavingService_SearchUserByAccountID_Handler,
		},
		{
			MethodName: "SearchUsersByFilter",
			Handler:    _MidSavingService_SearchUsersByFilter_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "mid_saving.proto",
}
