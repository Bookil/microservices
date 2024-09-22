package grpc

import (
	"context"

	emailv1 "github.com/Bookil/Bookil-Proto/gen/golang/email/v1"
)

func (a *Adapter) SendVerificationCode(ctx context.Context, req *emailv1.SendVerificationCodeRequest) (*emailv1.SendVerificationCodeResponse, error) {
	err := a.api.SendVerificationCode(req.GetEmail(), req.GetName(), req.GetVerificationCode())
	if err != nil {
		return nil, err
	}

	return &emailv1.SendVerificationCodeResponse{}, nil
}

func (a *Adapter) SendResetPassword(ctx context.Context, req *emailv1.SendResetPasswordRequest) (*emailv1.SendResetPasswordResponse, error) {
	err := a.api.SendResetPassword(req.GetEmail(), req.GetName(), req.GetUrl(), req.GetExpiry())
	if err != nil {
		return nil, err
	}

	return &emailv1.SendResetPasswordResponse{}, nil
}

func (a *Adapter) SendWelcome(ctx context.Context, req *emailv1.SendWelcomeRequest) (*emailv1.SendWelcomeResponse, error) {
	err := a.api.SendWelcome(req.GetEmail(), req.GetName())
	if err != nil {
		return nil, err
	}

	return &emailv1.SendWelcomeResponse{}, nil
}
