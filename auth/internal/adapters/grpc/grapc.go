package grpc

import (
	"context"

	authv1 "github.com/Bookil/Bookil-Proto/gen/golang/auth/v1"
	"google.golang.org/protobuf/types/known/durationpb"
)

func (a *Adapter) Register(ctx context.Context, request *authv1.RegisterRequest) (*authv1.RegisterResponse, error) {
	verificationCode, err := a.api.Register(ctx, request.UserId, request.Password)
	if err != nil {
		return nil, ErrFailedRegister
	}

	return &authv1.RegisterResponse{
		VerificationCode: verificationCode,
	}, nil
}

func (a *Adapter) VerifyEmail(ctx context.Context, request *authv1.VerifyEmailRequest) (*authv1.VerifyEmailResponse, error) {
	err := a.api.VerifyEmail(ctx, request.UserId, request.VerificationCode)
	if err != nil {
		return nil, ErrFailedVerifyEmil
	}

	return &authv1.VerifyEmailResponse{}, nil
}

func (a *Adapter) Login(ctx context.Context, request *authv1.LoginRequest) (*authv1.LoginResponse, error) {
	accessToken, refreshToken, err := a.api.Login(ctx, request.UserId, request.Password)
	if err != nil {
		return nil, ErrFailedLogin
	}

	return &authv1.LoginResponse{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}, nil
}

func (a *Adapter) Authenticate(ctx context.Context, request *authv1.AuthenticationRequest) (*authv1.AuthenticationResponse, error) {
	userID, err := a.api.Authenticate(ctx, request.AccessToken)
	if err != nil {
		return nil, ErrFailedAuthenticate
	}

	return &authv1.AuthenticationResponse{
		UserId: userID,
	}, nil
}

func (a *Adapter) RefreshToken(ctx context.Context, request *authv1.RefreshTokenRequest) (*authv1.RefreshTokenResponse, error) {
	accessToken, err := a.api.RefreshToken(ctx, request.UserId, request.RefreshToken)
	if err != nil {
		return nil, ErrFailedAuthenticate
	}

	return &authv1.RefreshTokenResponse{
		AccessToken: accessToken,
	}, nil
}

func (a *Adapter) ChangePassword(ctx context.Context, request *authv1.ChangePasswordRequest) (*authv1.ChangePasswordResponse, error) {
	err := a.api.ChangePassword(ctx, request.UserId, request.NewPassword, request.OldPassword)
	if err != nil {
		return nil, ErrFailedChargePassword
	}

	return &authv1.ChangePasswordResponse{}, nil
}

func (a *Adapter) ResetPassword(ctx context.Context, request *authv1.ResetPasswordRequest) (*authv1.ResetPasswordResponse, error) {
	token, timeout, err := a.api.ResetPassword(ctx, request.UserId)
	if err != nil {
		return nil, ErrFailedResetPassword
	}

	timeoutProto := durationpb.New(timeout)
	return &authv1.ResetPasswordResponse{
		ResetPasswordToken: token,
		Timeout:            timeoutProto,
	}, nil
}

func (a *Adapter) SubmitResetPassword(ctx context.Context, request *authv1.SubmitResetPasswordRequest) (*authv1.SubmitResetPasswordResponse, error) {
	err := a.api.SubmitResetPassword(ctx, request.ResetPasswordToken, request.NewPassword)
	if err != nil {
		return nil, ErrFailedSubmitResetPassword
	}

	return &authv1.SubmitResetPasswordResponse{}, nil
}

func (a *Adapter) DeleteAccount(ctx context.Context, request *authv1.DeleteAccountRequest) (*authv1.DeleteAccountResponse, error) {
	err := a.api.DeleteAccount(ctx, request.UserId, request.Password)
	if err != nil {
		return nil, ErrFailedDeleteAccount
	}

	return &authv1.DeleteAccountResponse{}, nil
}
