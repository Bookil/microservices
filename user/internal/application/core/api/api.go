package api

import (
	"context"

	"github.com/Bookil/microservices/user/helper"
	"github.com/Bookil/microservices/user/internal/application/core/domain"
	"github.com/Bookil/microservices/user/internal/ports"
)

type Application struct {
	auth ports.AuthPort
	db   ports.DBPort
}

func NewApplication(auth ports.AuthPort, db ports.DBPort) *Application {
	return &Application{
		auth: auth,
		db:   db,
	}
}

func (a *Application) Register(ctx context.Context, firstName, lastName, email string) (domain.UserID, error) {
	newUser := domain.NewUser(firstName, lastName, email)

	savedUser, err := a.db.Create(ctx, newUser)
	if err != nil {
		if helper.IsContains("email", err) && helper.IsContains("unique", err) {
			return "", ErrEmailRegistered
		}
		return "", ErrRegisterFailed
	}

	return savedUser.UserID, nil
}

func (a *Application) GetUserIDByEmail(ctx context.Context, email string) (domain.UserID, error) {
	savedUser, err := a.db.GetUserByEmail(ctx, email)
	if err != nil {
		if helper.IsContains("found", err) {
			return "", ErrUserNotFindWithThisEmail
		}
		return "", ErrRegisterFailed
	}

	return savedUser.UserID, nil
}

func (a *Application) ChangePassword(ctx context.Context, userID, oldPassword, newPassword string) error {
	err := a.auth.ChangePassword(ctx, userID, oldPassword, newPassword)
	if err != nil {
		return ErrChangingPasswordFailed
	}

	return nil
}

func (a *Application) Update(ctx context.Context, userID, firstName, lastName string) error {
	_, err := a.db.Update(ctx, userID, firstName, lastName)
	if err != nil {
		return ErrUpdateFailed
	}

	return nil
}

func (a *Application) DeleteAccount(ctx context.Context, userID, password string) error {
	err := a.db.DeleteByID(ctx, userID)
	if err != nil {
		return ErrDeleteAccountFailed
	}

	err = a.auth.DeleteAccount(ctx, userID, password)
	if err != nil {
		return ErrDeleteAccountFailed
	}

	return nil
}
