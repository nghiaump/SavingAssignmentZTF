// Code generated by protoc-gen-go. DO NOT EDIT.
// source: mid_saving.proto

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

func init() { proto.RegisterFile("mid_saving.proto", fileDescriptor_262d7cf1cc6f51b8) }

var fileDescriptor_262d7cf1cc6f51b8 = []byte{
	// 368 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x91, 0x4f, 0x4f, 0xea, 0x40,
	0x14, 0xc5, 0x77, 0xef, 0xbd, 0xdc, 0xf0, 0x10, 0x2f, 0x8a, 0xa6, 0xa0, 0x2e, 0x74, 0xcd, 0x02,
	0xb7, 0x1a, 0x23, 0x34, 0x92, 0xaa, 0xa8, 0xa1, 0x51, 0xe3, 0x8a, 0x94, 0xf6, 0x0a, 0x93, 0x40,
	0x0b, 0x33, 0x53, 0x0c, 0x1f, 0xd4, 0xef, 0x63, 0xfa, 0x2f, 0xd3, 0xa1, 0x16, 0x5d, 0x35, 0x39,
	0xe7, 0x9e, 0x5f, 0xef, 0xb9, 0x03, 0xb5, 0x39, 0xf3, 0x46, 0xc2, 0x59, 0x31, 0x7f, 0xd2, 0x5e,
	0xf0, 0x40, 0x06, 0xf8, 0x2f, 0xfe, 0x8c, 0xc3, 0x77, 0x03, 0x42, 0x41, 0x3c, 0x51, 0x8d, 0x4a,
	0x7e, 0xa6, 0xf3, 0xf9, 0x17, 0x6a, 0x03, 0xe6, 0xd9, 0xb1, 0x66, 0x13, 0x5f, 0x31, 0x97, 0x70,
	0x00, 0x95, 0x21, 0x4d, 0x98, 0x90, 0xc4, 0x9f, 0x05, 0x71, 0x3c, 0x6a, 0x67, 0xa4, 0x76, 0x5e,
	0x1f, 0xd2, 0x32, 0x24, 0x21, 0x8d, 0xe3, 0x32, 0x5b, 0x2c, 0x02, 0x5f, 0x10, 0x3e, 0xc1, 0xff,
	0x3e, 0xc9, 0x5e, 0xc8, 0x39, 0xf9, 0xf2, 0xee, 0xad, 0x87, 0xb9, 0x80, 0x66, 0x64, 0xc0, 0x93,
	0x52, 0x3f, 0x25, 0x8e, 0x00, 0x1f, 0x17, 0xe4, 0x27, 0x5b, 0x8b, 0x6b, 0xd7, 0x0d, 0x42, 0x5f,
	0xe2, 0xa9, 0x8a, 0x15, 0xdd, 0x8c, 0x7d, 0xb6, 0x7d, 0x28, 0xfd, 0x41, 0x1f, 0xe0, 0x95, 0xc9,
	0xa9, 0xc7, 0x9d, 0x0f, 0x67, 0x86, 0x4d, 0x95, 0x51, 0x6a, 0x06, 0x6c, 0x7d, 0x6f, 0xa6, 0xa0,
	0x5b, 0xa8, 0xa6, 0x6c, 0xcb, 0x5f, 0x86, 0x8c, 0xaf, 0x31, 0x57, 0x4e, 0x77, 0x32, 0xe0, 0x81,
	0x1a, 0x48, 0xb6, 0xcb, 0xfa, 0x5d, 0xc2, 0xae, 0x4d, 0x0e, 0x77, 0xa7, 0xa9, 0xd0, 0x5d, 0x5b,
	0x26, 0xee, 0x68, 0x38, 0xcb, 0x2c, 0x8f, 0xdb, 0xd0, 0xd2, 0xe2, 0x22, 0xca, 0xf7, 0x1c, 0xee,
	0x3d, 0x84, 0xf3, 0x31, 0x71, 0x6c, 0xa8, 0x60, 0x5e, 0x37, 0x9a, 0x25, 0xc0, 0x7b, 0x26, 0x24,
	0xbe, 0x40, 0x63, 0x13, 0x1a, 0xbd, 0xbd, 0x65, 0xfe, 0xdc, 0x73, 0x2b, 0xb7, 0x5f, 0xe4, 0xde,
	0xb0, 0x99, 0x24, 0x8e, 0x35, 0x15, 0x4b, 0x94, 0xed, 0xa0, 0x0e, 0x54, 0x13, 0x50, 0xb4, 0x56,
	0x7c, 0xb1, 0x1c, 0x20, 0x59, 0xd5, 0xa8, 0xea, 0x0a, 0x76, 0xe1, 0x50, 0xcb, 0x78, 0xbf, 0xb8,
	0xd2, 0x26, 0xe3, 0x02, 0xf6, 0xf3, 0x8c, 0xec, 0x04, 0x26, 0xd6, 0x8b, 0x77, 0x29, 0x6e, 0x70,
	0x05, 0x75, 0x95, 0x56, 0xdd, 0xf7, 0xf4, 0xb1, 0xb4, 0x3f, 0xea, 0x6a, 0x54, 0x7b, 0xfc, 0x27,
	0x96, 0xce, 0xbf, 0x02, 0x00, 0x00, 0xff, 0xff, 0xc1, 0x52, 0x14, 0xce, 0x16, 0x04, 0x00, 0x00,
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
