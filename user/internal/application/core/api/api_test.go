package api_test

import (
	"context"
	"errors"
	"testing"

	"github.com/Bookil/microservices/user/internal/application/core/api"
	"github.com/Bookil/microservices/user/internal/application/core/domain"
	"github.com/Bookil/microservices/user/mocks"
	"github.com/stretchr/testify/suite"
	"go.uber.org/mock/gomock"
)

type ApplicationTestSuit struct {
	suite.Suite

	app *api.Application

	db   *mocks.MockDBPort
	auth *mocks.MockAuthPort
}

var ErrUnknownError = errors.New("random error")

func TestApplicationTestSuit(t *testing.T) {
	suite.Run(t, new(ApplicationTestSuit))
}

func (a *ApplicationTestSuit) SetupSuite() {
	ctrl := gomock.NewController(a.T())

	mockDB := mocks.NewMockDBPort(ctrl)
	mockAuth := mocks.NewMockAuthPort(ctrl)

	app := api.NewApplication(mockAuth, mockDB)

	a.app = app
	a.db = mockDB
	a.auth = mockAuth
}

func (a *ApplicationTestSuit) TestRegister_Success() {
	ctx := context.TODO()
	firstName := "John"
	lastName := "Doe"
	email := "john.doe@example.com"
	newUser := domain.NewUser(firstName, lastName, email)

	a.db.EXPECT().Create(ctx, gomock.Any()).Return(newUser, nil)

	userID, err := a.app.Register(ctx, firstName, lastName, email)
	a.NotEmpty(userID)
	a.NoError(err)
}

func (a *ApplicationTestSuit) TestRegister_EmailRegistered() {
	ctx := context.TODO()
	firstName := "John"
	lastName := "Doe"
	email := "john.doe@example.com"

	a.db.EXPECT().Create(ctx, gomock.Any()).Return(nil, errors.New("unique email"))

	userID, err := a.app.Register(ctx, firstName, lastName, email)
	a.Empty(userID)
	a.Error(err)
}

func (a *ApplicationTestSuit) TestRegister_OtherError() {
	ctx := context.TODO()
	firstName := "John"
	lastName := "Doe"
	email := "john.doe@example.com"

	a.db.EXPECT().Create(ctx, gomock.Any()).Return(nil, ErrUnknownError)

	userID, err := a.app.Register(ctx, firstName, lastName, email)
	a.Empty(userID)
	a.Error(err)
}

func (a *ApplicationTestSuit) TestGetUserIDAndNameByEmail_Success() {
	ctx := context.TODO()
	email := "john.doe@example.com"
	savedUser := &domain.User{UserID: "123",FirstName: "amir",Email: email}

	a.db.EXPECT().GetUserByEmail(ctx, email).Return(savedUser, nil)

	userID,name,err := a.app.GetUserIDAndNameByEmail(ctx, email)
	a.NotEmpty(userID)
	a.NotEmpty(name)
	a.NoError(err)
}

func (a *ApplicationTestSuit) TestGetUserIDAndNameByEmail_UserNotFound() {
	ctx := context.TODO()
	email := "john.doe@example.com"

	a.db.EXPECT().GetUserByEmail(ctx, email).Return(nil,errors.New("not found"))

	userID, name,err := a.app.GetUserIDAndNameByEmail(ctx, email)
	a.Empty(userID)
	a.Empty(name)
	a.Error(err)
}

func (a *ApplicationTestSuit) TestGetUserIDAndNameByEmail_OtherError() {
	ctx := context.TODO()
	email := "john.doe@example.com"

	a.db.EXPECT().GetUserByEmail(ctx, email).Return(nil,ErrUnknownError)

	userID,name,err := a.app.GetUserIDAndNameByEmail(ctx, email)
	a.Empty(userID)
	a.Empty(name)
	a.Error(err)
}

func (a *ApplicationTestSuit) TestChangePassword_Success() {
	ctx := context.TODO()
	userID := "123"
	oldPassword := "oldPass"
	newPassword := "newPass"

	a.auth.EXPECT().ChangePassword(ctx, userID, oldPassword, newPassword).Return(nil)

	err := a.app.ChangePassword(ctx, userID, oldPassword, newPassword)
	a.NoError(err)
}

func (a *ApplicationTestSuit) TestChangePassword_Error() {
	ctx := context.TODO()
	userID := "123"
	oldPassword := "oldPass"
	newPassword := "newPass"

	a.auth.EXPECT().ChangePassword(ctx, userID, oldPassword, newPassword).Return(ErrUnknownError)

	err := a.app.ChangePassword(ctx, userID, oldPassword, newPassword)
	a.Error(err)
}

func (a *ApplicationTestSuit) TestUpdate_Success() {
	ctx := context.TODO()
	userID := "123"
	firstName := "John"
	lastName := "Doe"

	a.db.EXPECT().Update(ctx, userID, firstName, lastName).Return(nil, nil)

	err := a.app.Update(ctx, userID, firstName, lastName)
	a.NoError(err)
}

func (a *ApplicationTestSuit) TestUpdate_Error() {
	ctx := context.TODO()
	userID := "123"
	firstName := "John"
	lastName := "Doe"

	a.db.EXPECT().Update(ctx, userID, firstName, lastName).Return(nil, errors.New("update failed"))

	err := a.app.Update(ctx, userID, firstName, lastName)
	a.Error(err)
}

func (a *ApplicationTestSuit) TestDeleteAccount_Success() {
	ctx := context.TODO()
	userID := "123"
	password := "password"

	a.db.EXPECT().DeleteByID(ctx, userID).Return(nil)
	a.auth.EXPECT().DeleteAccount(ctx, userID, password).Return(nil)

	err := a.app.DeleteAccount(ctx, userID, password)
	a.NoError(err)
}

func (a *ApplicationTestSuit) TestDeleteAccount_DBError() {
	ctx := context.TODO()
	userID := "123"
	password := "password"

	a.db.EXPECT().DeleteByID(ctx, userID).Return(errors.New("delete failed"))

	err := a.app.DeleteAccount(ctx, userID, password)
	a.Error(err)
}

func (a *ApplicationTestSuit) TestDeleteAccount_AuthError() {
	ctx := context.TODO()
	userID := "123"
	password := "password"

	a.db.EXPECT().DeleteByID(ctx, userID).Return(nil)
	a.auth.EXPECT().DeleteAccount(ctx, userID, password).Return(errors.New("delete failed"))

	err := a.app.DeleteAccount(ctx, userID, password)
	a.Error(err)
}
