package grpc

import (
	"context"

	authv1 "github.com/Bookil/Bookil-Proto/gen/golang/auth/v1"
)

func (a *Adapter) Register(ctx context.Context, request *authv1.RegisterRequest) (*authv1.RegisterResponse, error) {
	err := a.validator.ValidateRegisterInputs(request.FirstName, request.LastName, request.Email, request.Password)
	if err != nil {
		return nil, ErrInvalidInputs
	}

	userID, err := a.api.Register(ctx, request.FirstName, request.LastName, request.Email, request.Password)
	if err != nil {
		return nil, ErrFailedRegister
	}

	return &authv1.RegisterResponse{
		UserId: userID,
	}, nil
}

func (a *Adapter) SendVerificationCodeAgain(ctx context.Context, request *authv1.SendVerificationCodeAgainRequest) (*authv1.SendVerificationCodeAgainResponse, error) {
	err := a.validator.ValidateSendVerificationCode(request.Email)
	if err != nil {
		return nil, ErrInvalidInputs
	}

	err = a.api.SendVerificationCodeAgain(ctx, request.Email)
	if err != nil {
		return nil, ErrFailedRegister
	}

	return &authv1.SendVerificationCodeAgainResponse{}, nil
}

func (a *Adapter) VerifyEmail(ctx context.Context, request *authv1.VerifyEmailRequest) (*authv1.VerifyEmailResponse, error) {
	err := a.validator.ValidateVerifyEmailInputs(request.Email, request.VerificationCode)
	if err != nil {
		return nil, ErrInvalidInputs
	}

	err = a.api.VerifyEmail(ctx, request.Email, request.VerificationCode)
	if err != nil {
		return nil, ErrFailedVerifyEmil
	}

	return &authv1.VerifyEmailResponse{}, nil
}

func (a *Adapter) Login(ctx context.Context, request *authv1.LoginRequest) (*authv1.LoginResponse, error) {
	err := a.validator.ValidateLoginInputs(request.Email, request.Password)
	if err != nil {
		return nil, ErrInvalidInputs
	}

	accessToken, refreshToken, err := a.api.Login(ctx, request.Email, request.Password)
	if err != nil {
		return nil, ErrFailedLogin
	}

	return &authv1.LoginResponse{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}, nil
}

func (a *Adapter) Authentication(ctx context.Context, request *authv1.AuthenticationRequest) (*authv1.AuthenticationResponse, error) {
	err := a.validator.ValidateAuthenticateInputsAndAuthorization(request.AccessToken)
	if err != nil {
		return nil, ErrInvalidInputs
	}

	userID, err := a.api.Authenticate(ctx, request.AccessToken)
	if err != nil {
		return nil, ErrFailedAuthenticate
	}

	return &authv1.AuthenticationResponse{
		UserId: userID,
	}, nil
}

func (a *Adapter) RoleAuthorization(ctx context.Context, request *authv1.RoleAuthorizationRequest) (*authv1.RoleAuthorizationResponse, error) {
	err := a.validator.ValidateAuthenticateInputsAndAuthorization(request.AccessToken)
	if err != nil {
		return nil, ErrInvalidInputs
	}

	err = a.api.RoleAuthorization(ctx, request.AccessToken)
	if err != nil {
		return nil, ErrFailedAuthenticate
	}

	return &authv1.RoleAuthorizationResponse{}, nil
}

func (a *Adapter) RefreshToken(ctx context.Context, request *authv1.RefreshTokenRequest) (*authv1.RefreshTokenResponse, error) {
	err := a.validator.ValidateRefreshTokenInputs(request.UserId, request.RefreshToken)
	if err != nil {
		return nil, ErrInvalidInputs
	}

	accessToken, err := a.api.RefreshToken(ctx, request.UserId, request.RefreshToken)
	if err != nil {
		return nil, ErrFailedAuthenticate
	}

	return &authv1.RefreshTokenResponse{
		AccessToken: accessToken,
	}, nil
}

func (a *Adapter) ChangePassword(ctx context.Context, request *authv1.ChangePasswordRequest) (*authv1.ChangePasswordResponse, error) {
	err := a.validator.ValidateChangePasswordInputs(request.UserId, request.OldPassword, request.NewPassword)
	if err != nil {
		return nil, ErrInvalidInputs
	}

	err = a.api.ChangePassword(ctx, request.UserId, request.OldPassword, request.NewPassword)
	if err != nil {
		return nil, ErrFailedChargePassword
	}

	return &authv1.ChangePasswordResponse{}, nil
}

func (a *Adapter) ResetPassword(ctx context.Context, request *authv1.ResetPasswordRequest) (*authv1.ResetPasswordResponse, error) {
	err := a.validator.ValidateResetPasswordInputs(request.Email)
	if err != nil {
		return nil, ErrInvalidInputs
	}

	err = a.api.ResetPassword(ctx, request.Email)
	if err != nil {
		return nil, ErrFailedResetPassword
	}

	return &authv1.ResetPasswordResponse{}, nil
}

func (a *Adapter) SubmitResetPassword(ctx context.Context, request *authv1.SubmitResetPasswordRequest) (*authv1.SubmitResetPasswordResponse, error) {
	err := a.validator.ValidateSubmitResetPasswordInputs(request.ResetPasswordToken, request.NewPassword)
	if err != nil {
		return nil, ErrInvalidInputs
	}

	err = a.api.SubmitResetPassword(ctx, request.ResetPasswordToken, request.NewPassword)
	if err != nil {
		return nil, ErrFailedSubmitResetPassword
	}

	return &authv1.SubmitResetPasswordResponse{}, nil
}

func (a *Adapter) DeleteAccount(ctx context.Context, request *authv1.DeleteAccountRequest) (*authv1.DeleteAccountResponse, error) {
	err := a.validator.ValidateDeleteAccountInputs(request.UserId, request.Password)
	if err != nil {
		return nil, ErrInvalidInputs
	}

	err = a.api.DeleteAccount(ctx, request.UserId, request.Password)
	if err != nil {
		return nil, ErrFailedDeleteAccount
	}

	return &authv1.DeleteAccountResponse{}, nil
}
