package grpc

import (
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

var (
	ErrFailedRegister   = status.Errorf(codes.Internal, "failed to register")
	ErrFailedLogin      = status.Errorf(codes.Internal, "failed to login")
	ErrFailedAuthenticate = status.Errorf(codes.Internal, "failed to authenticate")
	ErrFailedVerifyEmil = status.Errorf(codes.Internal, "failed to verify email")
	ErrFailedResetPassword = status.Errorf(codes.Internal, "failed to reset password")
	ErrFailedSubmitResetPassword = status.Errorf(codes.Internal, "failed to submit reset password")
	ErrFailedChargePassword = status.Errorf(codes.Internal, "failed to change password")
	ErrFailedDeleteAccount = status.Errorf(codes.Internal, "failed to delete account")
)
