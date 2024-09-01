package grpc

import (
	"context"

	userv1 "github.com/Bookil/Bookil-Proto/gen/golang/user/v1"
	"github.com/Bookil/microservices/user/internal/adapters/grpc/interceptor"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (a *Adapter) Register(ctx context.Context, request *userv1.RegisterRequest) (*userv1.RegisterResponse, error) {
	userID, err := a.api.Register(ctx, request.FisrtName, request.LastName, request.Email)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &userv1.RegisterResponse{
		UserId: userID,
	}, nil
}

func (a *Adapter) ChangePassword(ctx context.Context, request *userv1.ChangePasswordRequest) (*userv1.ChangePasswordResponse, error) {
	userID, ok := ctx.Value(interceptor.UserID{}).(string)
	if !ok {
		return nil, status.Error(codes.Internal, "an error occurred")
	}

	err := a.api.ChangePassword(ctx, userID, request.NewPassword, request.OldPassword)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &userv1.ChangePasswordResponse{}, nil
}

func (a *Adapter) Update(ctx context.Context, request *userv1.UpdateRequest) (*userv1.UpdateResponse, error) {
	userID, ok := ctx.Value(interceptor.UserID{}).(string)
	if !ok {
		return nil, status.Error(codes.Internal, "an error occurred")
	}

	err := a.api.Update(ctx, userID, request.NewFirstName, request.NewLastName)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &userv1.UpdateResponse{}, nil
}

func (a *Adapter) DeleteAccount(ctx context.Context, request *userv1.DeleteAccountRequest) (*userv1.DeleteAccountResponse, error) {
	userID, ok := ctx.Value(interceptor.UserID{}).(string)
	if !ok {
		return nil, status.Error(codes.Internal, "an error occurred")
	}
	
	err := a.api.DeleteAccount(ctx, userID, request.Password)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &userv1.DeleteAccountResponse{}, nil
}
