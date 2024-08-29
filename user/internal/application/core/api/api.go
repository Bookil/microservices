package api

import (
	"context"

	"github.com/Bookil/microservices/user/helper"
	"github.com/Bookil/microservices/user/internal/application/core/domain"
	"github.com/Bookil/microservices/user/internal/ports"
)

type Application struct {
	auth  ports.AuthPort
	email ports.EmailPort
	db    ports.DBPort
}

func NewApplication(auth ports.AuthPort, email ports.EmailPort, db ports.DBPort) *Application {
	return &Application{
		auth:  auth,
		email: email,
		db:    db,
	}
}

func (a *Application) Register(ctx context.Context, firstName, lastName, email, password string) (domain.UserID, error) {
	newUser := domain.NewUser(firstName, lastName, email)

	savedUser, err := a.db.Create(ctx, newUser)
	if err != nil {
		if helper.IsContains("email", err) && helper.IsContains("unique", err) {
			return "", ErrEmailRegistered
		}
		return "", ErrRegisterFailed
	}

	code, err := a.auth.Register(ctx, newUser.UserID, password)
	if err != nil {
		return "", ErrRegisterFailed
	}

	err = a.email.SendVerificationCode(savedUser.Email, code)
	if err != nil {
		return "", ErrRegisterFailed
	}

	return savedUser.UserID, nil
}

func (a *Application) Login(ctx context.Context, email, password string) (string, string, error) {
	gotUser, err := a.db.GetUserByEmail(ctx, email)
	if err != nil {
		if helper.IsContains("found", err) {
			return "", "", ErrUserNotFindWithThisEmail
		}
		return "", "", ErrLoggingFailed
	}

	accessToken, refreshToken, err := a.auth.Login(ctx, gotUser.UserID, password)
	if err != nil {
		return "", "", err
	}

	err = a.email.SendWellCome(gotUser.Email)
	if err != nil {
		return "", "", ErrLoggingFailed
	}

	return accessToken, refreshToken, nil
}

func (a *Application) ChangePassword(ctx context.Context, accessToken, oldPassword, newPassword string) error {
	userID, err := a.auth.Authenticate(ctx, accessToken)
	if err != nil {
		return ErrAccessDenied
	}

	err = a.auth.ChangePassword(ctx, userID, oldPassword, newPassword)
	if err != nil {
		return ErrChangingPasswordFailed
	}

	return nil
}

func (a *Application) ResetPassword(ctx context.Context, email string) error {
	gotUser, err := a.db.GetUserByEmail(ctx, email)
	if err != nil {
		if helper.IsContains("found", err) {
			return ErrUserNotFindWithThisEmail
		}
		return ErrResetPasswordFailed
	}

	token, duration, err := a.auth.ResetPassword(ctx, gotUser.UserID)
	if err != nil {
		return ErrResetPasswordFailed
	}

	err = a.email.SendResetPassword("example.com", token, gotUser.Email, duration)
	if err != nil {
		return ErrResetPasswordFailed
	}

	return nil
}

func (a *Application) Update(ctx context.Context, accessToken, firstName, lastName string) error {
	UserID, err := a.auth.Authenticate(ctx, accessToken)
	if err != nil {
		return ErrUpdateFailed
	}
	_, err = a.db.Update(ctx, UserID, firstName, lastName)
	if err != nil {
		return ErrUpdateFailed
	}

	return nil
}

func (a *Application) DeleteAccount(ctx context.Context, accessToken, password string) error {
	userID, err := a.auth.Authenticate(ctx, accessToken)
	if err != nil {
		return ErrDeleteAccountFailed
	}

	err = a.db.Delete(ctx, userID)
	if err != nil {
		return ErrDeleteAccountFailed
	}

	err = a.auth.DeleteAccount(ctx, userID, password)
	if err != nil {
		return ErrDeleteAccountFailed
	}

	return nil
}
