// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package rpc

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion7

// AuthServiceClient is the client API for AuthService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type AuthServiceClient interface {
	Register(ctx context.Context, in *RegisterRequest, opts ...grpc.CallOption) (*RegisterResponse, error)
	Login(ctx context.Context, in *LoginRequest, opts ...grpc.CallOption) (*LoginResponse, error)
	// ForgotPassword, then ResetPassword if succeed
	ForgotPassword(ctx context.Context, in *ForgotPasswordRequest, opts ...grpc.CallOption) (*ForgotPasswordResponse, error)
	ResetPassword(ctx context.Context, in *ResetPasswordRequest, opts ...grpc.CallOption) (*ResetPasswordResponse, error)
	// Refresh Access Token endpoint
	RefreshAccessToken(ctx context.Context, in *RefreshAccessTokenRequest, opts ...grpc.CallOption) (*RefreshAccessTokenResponse, error)
}

type authServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewAuthServiceClient(cc grpc.ClientConnInterface) AuthServiceClient {
	return &authServiceClient{cc}
}

var authServiceRegisterStreamDesc = &grpc.StreamDesc{
	StreamName: "Register",
}

func (c *authServiceClient) Register(ctx context.Context, in *RegisterRequest, opts ...grpc.CallOption) (*RegisterResponse, error) {
	out := new(RegisterResponse)
	err := c.cc.Invoke(ctx, "/proto.AuthService/Register", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

var authServiceLoginStreamDesc = &grpc.StreamDesc{
	StreamName: "Login",
}

func (c *authServiceClient) Login(ctx context.Context, in *LoginRequest, opts ...grpc.CallOption) (*LoginResponse, error) {
	out := new(LoginResponse)
	err := c.cc.Invoke(ctx, "/proto.AuthService/Login", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

var authServiceForgotPasswordStreamDesc = &grpc.StreamDesc{
	StreamName: "ForgotPassword",
}

func (c *authServiceClient) ForgotPassword(ctx context.Context, in *ForgotPasswordRequest, opts ...grpc.CallOption) (*ForgotPasswordResponse, error) {
	out := new(ForgotPasswordResponse)
	err := c.cc.Invoke(ctx, "/proto.AuthService/ForgotPassword", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

var authServiceResetPasswordStreamDesc = &grpc.StreamDesc{
	StreamName: "ResetPassword",
}

func (c *authServiceClient) ResetPassword(ctx context.Context, in *ResetPasswordRequest, opts ...grpc.CallOption) (*ResetPasswordResponse, error) {
	out := new(ResetPasswordResponse)
	err := c.cc.Invoke(ctx, "/proto.AuthService/ResetPassword", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

var authServiceRefreshAccessTokenStreamDesc = &grpc.StreamDesc{
	StreamName: "RefreshAccessToken",
}

func (c *authServiceClient) RefreshAccessToken(ctx context.Context, in *RefreshAccessTokenRequest, opts ...grpc.CallOption) (*RefreshAccessTokenResponse, error) {
	out := new(RefreshAccessTokenResponse)
	err := c.cc.Invoke(ctx, "/proto.AuthService/RefreshAccessToken", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AuthServiceService is the service API for AuthService service.
// Fields should be assigned to their respective handler implementations only before
// RegisterAuthServiceService is called.  Any unassigned fields will result in the
// handler for that method returning an Unimplemented error.
type AuthServiceService struct {
	Register func(context.Context, *RegisterRequest) (*RegisterResponse, error)
	Login    func(context.Context, *LoginRequest) (*LoginResponse, error)
	// ForgotPassword, then ResetPassword if succeed
	ForgotPassword func(context.Context, *ForgotPasswordRequest) (*ForgotPasswordResponse, error)
	ResetPassword  func(context.Context, *ResetPasswordRequest) (*ResetPasswordResponse, error)
	// Refresh Access Token endpoint
	RefreshAccessToken func(context.Context, *RefreshAccessTokenRequest) (*RefreshAccessTokenResponse, error)
}

func (s *AuthServiceService) register(_ interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RegisterRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return s.Register(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     s,
		FullMethod: "/proto.AuthService/Register",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return s.Register(ctx, req.(*RegisterRequest))
	}
	return interceptor(ctx, in, info, handler)
}
func (s *AuthServiceService) login(_ interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LoginRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return s.Login(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     s,
		FullMethod: "/proto.AuthService/Login",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return s.Login(ctx, req.(*LoginRequest))
	}
	return interceptor(ctx, in, info, handler)
}
func (s *AuthServiceService) forgotPassword(_ interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ForgotPasswordRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return s.ForgotPassword(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     s,
		FullMethod: "/proto.AuthService/ForgotPassword",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return s.ForgotPassword(ctx, req.(*ForgotPasswordRequest))
	}
	return interceptor(ctx, in, info, handler)
}
func (s *AuthServiceService) resetPassword(_ interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ResetPasswordRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return s.ResetPassword(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     s,
		FullMethod: "/proto.AuthService/ResetPassword",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return s.ResetPassword(ctx, req.(*ResetPasswordRequest))
	}
	return interceptor(ctx, in, info, handler)
}
func (s *AuthServiceService) refreshAccessToken(_ interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RefreshAccessTokenRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return s.RefreshAccessToken(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     s,
		FullMethod: "/proto.AuthService/RefreshAccessToken",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return s.RefreshAccessToken(ctx, req.(*RefreshAccessTokenRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// RegisterAuthServiceService registers a service implementation with a gRPC server.
func RegisterAuthServiceService(s grpc.ServiceRegistrar, srv *AuthServiceService) {
	srvCopy := *srv
	if srvCopy.Register == nil {
		srvCopy.Register = func(context.Context, *RegisterRequest) (*RegisterResponse, error) {
			return nil, status.Errorf(codes.Unimplemented, "method Register not implemented")
		}
	}
	if srvCopy.Login == nil {
		srvCopy.Login = func(context.Context, *LoginRequest) (*LoginResponse, error) {
			return nil, status.Errorf(codes.Unimplemented, "method Login not implemented")
		}
	}
	if srvCopy.ForgotPassword == nil {
		srvCopy.ForgotPassword = func(context.Context, *ForgotPasswordRequest) (*ForgotPasswordResponse, error) {
			return nil, status.Errorf(codes.Unimplemented, "method ForgotPassword not implemented")
		}
	}
	if srvCopy.ResetPassword == nil {
		srvCopy.ResetPassword = func(context.Context, *ResetPasswordRequest) (*ResetPasswordResponse, error) {
			return nil, status.Errorf(codes.Unimplemented, "method ResetPassword not implemented")
		}
	}
	if srvCopy.RefreshAccessToken == nil {
		srvCopy.RefreshAccessToken = func(context.Context, *RefreshAccessTokenRequest) (*RefreshAccessTokenResponse, error) {
			return nil, status.Errorf(codes.Unimplemented, "method RefreshAccessToken not implemented")
		}
	}
	sd := grpc.ServiceDesc{
		ServiceName: "proto.AuthService",
		Methods: []grpc.MethodDesc{
			{
				MethodName: "Register",
				Handler:    srvCopy.register,
			},
			{
				MethodName: "Login",
				Handler:    srvCopy.login,
			},
			{
				MethodName: "ForgotPassword",
				Handler:    srvCopy.forgotPassword,
			},
			{
				MethodName: "ResetPassword",
				Handler:    srvCopy.resetPassword,
			},
			{
				MethodName: "RefreshAccessToken",
				Handler:    srvCopy.refreshAccessToken,
			},
		},
		Streams:  []grpc.StreamDesc{},
		Metadata: "authn.proto",
	}

	s.RegisterService(&sd, nil)
}
