package api

import (
	"context"

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
		return "", err
	}

	code, err := a.auth.Register(ctx, newUser.UserID, password)
	if err != nil {
		return "", err
	}

	err = a.email.SendVerificationCode(savedUser.Email, code)
	if err != nil {
		return "", err
	}

	return savedUser.UserID, nil
}

func (a *Application) Login(ctx context.Context, email, password string) (string, string, error) {
	gotUser, err := a.db.GetUserByEmail(ctx, email)
	if err != nil {
		return "", "", err
	}

	accessToken, refreshToken, err := a.auth.Login(ctx, gotUser.UserID, password)
	if err != nil {
		return "", "", err
	}

	err = a.email.SendWellCome(gotUser.Email)
	if err != nil {
		return "", "", err
	}

	return accessToken, refreshToken, nil
}

func (a *Application) ChangePassword(ctx context.Context, accessToken, oldPassword, newPassword string) error {
	userID, err := a.auth.Authenticate(ctx, accessToken)
	if err != nil {
		return err
	}

	err = a.auth.ChangePassword(ctx, userID, oldPassword, newPassword)
	if err != nil {
		return err
	}

	return nil
}

func (a *Application) ResetPassword(ctx context.Context, email string) error {
	// gotUser, err := a.db.GetUserByEmail(ctx, email)
	// if err != nil {
	// 	return err
	// }

	// url, duration, err := a.auth.ResetPassword(ctx, gotUser.UserID)
	// if err != nil {
	// 	return err
	// }

	// err = a.email.SendResetPassword(url, duration, gotUser.Email)
	// if err != nil {
	// 	return err
	// }

	// return nil
	panic("not implemented")
}

func (a *Application) Update(ctx context.Context, accessToken, firstName, lastName string) error {
	UserID, err := a.auth.Authenticate(ctx, accessToken)
	if err != nil {
		return err
	}
	_, err = a.db.Update(ctx, UserID, firstName, lastName)
	if err != nil {
		return err
	}

	return nil
}

func (a *Application) DeleteAccount(ctx context.Context, accessToken, password string) error {
	userID, err := a.auth.Authenticate(ctx, accessToken)
	if err != nil {
		return err
	}

	err = a.db.Delete(ctx, userID)
	if err != nil {
		return err
	}

	err = a.auth.DeleteAccount(ctx, userID, password)
	if err != nil {
		return err
	}

	return nil
}
