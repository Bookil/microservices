package api_test

import (
	"context"
	"errors"
	"testing"

	"github.com/Bookil/microservices/user/internal/application/core/api"
	"github.com/Bookil/microservices/user/internal/application/core/domain"
	"github.com/Bookil/microservices/user/mocks"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
)

type ApplicationTestSuit struct {
	suite.Suite

	api *api.Application

	mockedAuth  *mocks.MockedAuth
	mockedDB    *mocks.MockedDB
	mockedEmail *mocks.MockedEmail
}

var ErrUnknownError = errors.New("random error")

func TestApplicationTestSuit(t *testing.T) {
	suite.Run(t, new(ApplicationTestSuit))
}

func (a *ApplicationTestSuit) SetupSuite() {
	mockedAuth := &mocks.MockedAuth{}

	mockedDB := &mocks.MockedDB{}

	mockedEmail := &mocks.MockedEmail{}

	api := api.NewApplication(mockedAuth, mockedEmail, mockedDB)

	a.api = api

	a.mockedAuth = mockedAuth
	a.mockedDB = mockedDB
	a.mockedEmail = mockedEmail
}

func (a *ApplicationTestSuit) TestRegister() {
	ctx := context.TODO()

	firstName := "John"
	lastName := "Doe"
	email := "johnDoe@gmail.com"
	password := "password"

	user := domain.NewUser(firstName, lastName, email)

	mockDBCall := a.mockedDB.On("Create", mock.Anything, mock.Anything).Return(user, nil)
	mockAuthCall := a.mockedAuth.On("Register", mock.Anything, mock.Anything, mock.Anything).Return("1234567", nil)
	mockEmailCall := a.mockedEmail.On("SendVerificationCode", mock.Anything, mock.Anything).Return(nil)

	code, err := a.api.Register(ctx, firstName, lastName, email, password)

	a.NoError(err)
	a.NotEmpty(code)

	mockDBCall.Unset()
	mockAuthCall.Unset()
	mockEmailCall.Unset()
}

func (a *ApplicationTestSuit) TestRegisterShouldFailWhenAuthFails() {
	ctx := context.TODO()

	firstName := "John"
	lastName := "Doe"
	email := "johnDoe@gmail.com"
	password := "password"

	user := domain.NewUser(firstName, lastName, email)

	mockDBCall := a.mockedDB.On("Create", mock.Anything, mock.Anything).Return(user, nil)
	mockAuthCall := a.mockedAuth.On("Register", mock.Anything, mock.Anything, mock.Anything).Return("", ErrUnknownError)
	mockEmailCall := a.mockedEmail.On("SendVerificationCode", mock.Anything, mock.Anything).Return(nil)

	code, err := a.api.Register(ctx, firstName, lastName, email, password)

	a.ErrorIs(err, api.ErrRegisterFailed)
	a.Empty(code)

	mockDBCall.Unset()
	mockAuthCall.Unset()
	mockEmailCall.Unset()
}

func (a *ApplicationTestSuit) TestRegisterShouldFailWhenDBFails() {
	ctx := context.TODO()

	firstName := "John"
	lastName := "Doe"
	email := "johnDoe@gmail.com"
	password := "password"

	mockDBCall := a.mockedDB.On("Create", mock.Anything, mock.Anything).Return(&domain.User{}, ErrUnknownError)
	mockAuthCall := a.mockedAuth.On("Register", mock.Anything, mock.Anything, mock.Anything).Return("123456", nil)
	mockEmailCall := a.mockedEmail.On("SendVerificationCode", mock.Anything, mock.Anything).Return(nil)

	code, err := a.api.Register(ctx, firstName, lastName, email, password)

	a.ErrorIs(err, api.ErrRegisterFailed)
	a.Empty(code)

	mockDBCall.Unset()
	mockAuthCall.Unset()
	mockEmailCall.Unset()
}

func (a *ApplicationTestSuit) TestRegisterShouldFailWhenDBFailEmails() {
	ctx := context.TODO()

	firstName := "John"
	lastName := "Doe"
	email := "johnDoe@gmail.com"
	password := "password"

	mockDBCall := a.mockedDB.On("Create", mock.Anything, mock.Anything).Return(&domain.User{}, errors.New("Email and unique"))
	mockAuthCall := a.mockedAuth.On("Register", mock.Anything, mock.Anything, mock.Anything).Return("123456", nil)
	mockEmailCall := a.mockedEmail.On("SendVerificationCode", mock.Anything, mock.Anything).Return(nil)

	code, err := a.api.Register(ctx, firstName, lastName, email, password)

	a.ErrorIs(err, api.ErrEmailRegistered)
	a.Empty(code)

	mockDBCall.Unset()
	mockAuthCall.Unset()
	mockEmailCall.Unset()
}

func (a *ApplicationTestSuit) TestRegisterShouldFailWhenEmailFails() {
	ctx := context.TODO()

	firstName := "John"
	lastName := "Doe"
	email := "johnDoe@gmail.com"
	password := "password"

	user := domain.NewUser(firstName, lastName, email)

	mockDBCall := a.mockedDB.On("Create", mock.Anything, mock.Anything, mock.Anything).Return(user, nil)
	mockAuthCall := a.mockedAuth.On("Register", mock.Anything, mock.Anything).Return("1234567", nil)
	mockEmailCall := a.mockedEmail.On("SendVerificationCode", mock.Anything, mock.Anything).Return(ErrUnknownError)

	code, err := a.api.Register(ctx, firstName, lastName, email, password)

	a.ErrorIs(err, api.ErrRegisterFailed)
	a.Empty(code)

	mockDBCall.Unset()
	mockAuthCall.Unset()
	mockEmailCall.Unset()
}

func (a *ApplicationTestSuit) TestLogin() {
	ctx := context.TODO()

	firstName := "John"
	lastName := "Doe"
	email := "johnDoe@gmail.com"
	password := "password"

	user := domain.NewUser(firstName, lastName, email)

	mockDBCall := a.mockedDB.On("GetUserByEmail", mock.Anything, mock.Anything).Return(user, nil)
	mockAuthCall := a.mockedAuth.On("Login", mock.Anything, mock.Anything, mock.Anything).Return("valid token", "valid token", nil)
	mockEmailCall := a.mockedEmail.On("SendWellCome", mock.Anything, mock.Anything).Return(nil)

	accessToken, refreshToken, err := a.api.Login(ctx, email, password)

	a.NoError(err)
	a.NotEmpty(accessToken)
	a.NotEmpty(refreshToken)

	mockDBCall.Unset()
	mockAuthCall.Unset()
	mockEmailCall.Unset()
}

func (a *ApplicationTestSuit) TestLoginShouldFailWhenDBFails() {
	ctx := context.TODO()

	email := "johnDoe@gmail.com"
	password := "password"

	mockDBCall := a.mockedDB.On("GetUserByEmail", mock.Anything, mock.Anything).Return(&domain.User{}, ErrUnknownError)
	mockAuthCall := a.mockedAuth.On("Login", mock.Anything, mock.Anything, mock.Anything).Return("valid token", "valid token", nil)
	mockEmailCall := a.mockedEmail.On("SendWellCome", mock.Anything, mock.Anything).Return(nil)

	accessToken, refreshToken, err := a.api.Login(ctx, email, password)

	a.ErrorIs(err, api.ErrLoggingFailed)
	a.Empty(accessToken)
	a.Empty(refreshToken)

	mockDBCall.Unset()
	mockAuthCall.Unset()
	mockEmailCall.Unset()
}

func (a *ApplicationTestSuit) TestLoginShouldFailWhenDBFailInvalidEmail() {
	ctx := context.TODO()

	email := "johnDoe@gmail.com"
	password := "password"

	mockDBCall := a.mockedDB.On("GetUserByEmail", mock.Anything, mock.Anything).Return(&domain.User{}, errors.New("not found"))
	mockAuthCall := a.mockedAuth.On("Login", mock.Anything, mock.Anything, mock.Anything).Return("valid token", "valid token", nil)
	mockEmailCall := a.mockedEmail.On("SendWellCome", mock.Anything, mock.Anything).Return(nil)

	accessToken, refreshToken, err := a.api.Login(ctx, email, password)

	a.ErrorIs(err, api.ErrUserNotFindWithThisEmail)
	a.Empty(accessToken)
	a.Empty(refreshToken)

	mockDBCall.Unset()
	mockAuthCall.Unset()
	mockEmailCall.Unset()
}

func (a *ApplicationTestSuit) TestLoginShouldFailWhenAuthFails() {
	ctx := context.TODO()

	firstName := "John"
	lastName := "Doe"
	email := "johnDoe@gmail.com"
	password := "password"

	user := domain.NewUser(firstName, lastName, email)

	mockDBCall := a.mockedDB.On("GetUserByEmail", mock.Anything, mock.Anything).Return(user, nil)
	mockAuthCall := a.mockedAuth.On("Login", mock.Anything, mock.Anything, mock.Anything).Return("valid token", "valid token", nil)
	mockEmailCall := a.mockedEmail.On("SendWellCome", mock.Anything, mock.Anything).Return(ErrUnknownError)

	accessToken, refreshToken, err := a.api.Login(ctx, email, password)

	a.ErrorIs(err, api.ErrLoggingFailed)
	a.Empty(accessToken)
	a.Empty(refreshToken)

	mockDBCall.Unset()
	mockAuthCall.Unset()
	mockEmailCall.Unset()
}

func (a *ApplicationTestSuit) TestLoginShouldFailWhenEmailFails() {
	ctx := context.TODO()

	firstName := "John"
	lastName := "Doe"
	email := "johnDoe@gmail.com"
	password := "password"

	user := domain.NewUser(firstName, lastName, email)

	mockDBCall := a.mockedDB.On("GetUserByEmail", mock.Anything, mock.Anything).Return(user, nil)
	mockAuthCall := a.mockedAuth.On("Login", mock.Anything, mock.Anything, mock.Anything).Return("", "", ErrUnknownError)
	mockEmailCall := a.mockedEmail.On("SendWellCome", mock.Anything, mock.Anything).Return(nil)

	accessToken, refreshToken, err := a.api.Login(ctx, email, password)

	a.Error(err)
	a.Empty(accessToken)
	a.Empty(refreshToken)

	mockDBCall.Unset()
	mockAuthCall.Unset()
	mockEmailCall.Unset()
}

func (a *ApplicationTestSuit) TestChangePassword() {
	ctx := context.TODO()

	firstName := "John"
	lastName := "Doe"
	email := "johnDoe@gmail.com"

	user := domain.NewUser(firstName, lastName, email)

	mockAuthCallAuthentication := a.mockedAuth.On("Authenticate", mock.Anything, mock.Anything).Return(user.UserID, nil)
	mockAuthCallChangePassword := a.mockedAuth.On("ChangePassword", mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(nil)

	err := a.api.ChangePassword(ctx, "valid", "oldPassword", "newPassword")

	a.NoError(err)

	mockAuthCallAuthentication.Unset()
	mockAuthCallChangePassword.Unset()
}

func (a *ApplicationTestSuit) TestChangePasswordShouldFailWhenAuthenticateFails() {
	ctx := context.TODO()

	mockAuthCallAuthentication := a.mockedAuth.On("Authenticate", mock.Anything, mock.Anything).Return("", ErrUnknownError)
	mockAuthCallChangePassword := a.mockedAuth.On("ChangePassword", mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(nil)

	err := a.api.ChangePassword(ctx, "valid", "oldPassword", "newPassword")

	a.ErrorIs(err, api.ErrAccessDenied)

	mockAuthCallAuthentication.Unset()
	mockAuthCallChangePassword.Unset()
}

func (a *ApplicationTestSuit) TestChangePasswordShouldFailWhenChangePasswordFails() {
	ctx := context.TODO()

	firstName := "John"
	lastName := "Doe"
	email := "johnDoe@gmail.com"

	user := domain.NewUser(firstName, lastName, email)

	mockAuthCallAuthentication := a.mockedAuth.On("Authenticate", mock.Anything, mock.Anything).Return(user.UserID, nil)
	mockAuthCallChangePassword := a.mockedAuth.On("ChangePassword", mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(ErrUnknownError)

	err := a.api.ChangePassword(ctx, "valid", "oldPassword", "newPassword")

	a.ErrorIs(err, api.ErrChangingPasswordFailed)

	mockAuthCallAuthentication.Unset()
	mockAuthCallChangePassword.Unset()
}

func (a *ApplicationTestSuit) TestResetPassword() {
	ctx := context.TODO()

	firstName := "John"
	lastName := "Doe"
	email := "johnDoe@gmail.com"

	user := domain.NewUser(firstName, lastName, email)

	mockDBCall := a.mockedDB.On("GetUserByEmail", mock.Anything, mock.Anything).Return(user, nil)
	mockAuthCall := a.mockedAuth.On("ResetPassword", mock.Anything, mock.Anything).Return("token", 2, nil)
	mockEmailCall := a.mockedEmail.On("SendResetPassword", mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(nil)

	err := a.api.ResetPassword(ctx, user.UserID)

	a.NoError(err)

	mockDBCall.Unset()
	mockAuthCall.Unset()
	mockEmailCall.Unset()
}

func (a *ApplicationTestSuit) TestResetPasswordShouldFailWhenAuthFails() {
	ctx := context.TODO()

	firstName := "John"
	lastName := "Doe"
	email := "johnDoe@gmail.com"

	user := domain.NewUser(firstName, lastName, email)

	mockDBCall := a.mockedDB.On("GetUserByEmail", mock.Anything, mock.Anything).Return(user, nil)
	mockAuthCall := a.mockedAuth.On("ResetPassword", mock.Anything, mock.Anything).Return("", 0, ErrUnknownError)
	mockEmailCall := a.mockedEmail.On("SendResetPassword", mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(nil)

	err := a.api.ResetPassword(ctx, user.UserID)

	a.ErrorIs(err,api.ErrResetPasswordFailed)

	mockDBCall.Unset()
	mockAuthCall.Unset()
	mockEmailCall.Unset()
}

func (a *ApplicationTestSuit) TestResetPasswordShouldFailWhenDBFails() {
	ctx := context.TODO()

	firstName := "John"
	lastName := "Doe"
	email := "johnDoe@gmail.com"

	user := domain.NewUser(firstName, lastName, email)

	mockDBCall := a.mockedDB.On("GetUserByEmail", mock.Anything, mock.Anything).Return(&domain.User{}, ErrUnknownError)
	mockAuthCall := a.mockedAuth.On("ResetPassword", mock.Anything, mock.Anything).Return("token", 2, nil)
	mockEmailCall := a.mockedEmail.On("SendResetPassword", mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(nil)

	err := a.api.ResetPassword(ctx, user.UserID)

	a.ErrorIs(err,api.ErrResetPasswordFailed)

	mockDBCall.Unset()
	mockAuthCall.Unset()
	mockEmailCall.Unset()
}

func (a *ApplicationTestSuit) TestResetPasswordShouldFailWhenDBFailsInvalidEmail() {
	ctx := context.TODO()

	firstName := "John"
	lastName := "Doe"
	email := "johnDoe@gmail.com"

	user := domain.NewUser(firstName, lastName, email)

	mockDBCall := a.mockedDB.On("GetUserByEmail", mock.Anything, mock.Anything).Return(&domain.User{}, errors.New("found"))
	mockAuthCall := a.mockedAuth.On("ResetPassword", mock.Anything, mock.Anything).Return("token", 2, nil)
	mockEmailCall := a.mockedEmail.On("SendResetPassword", mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(nil)

	err := a.api.ResetPassword(ctx, user.UserID)

	a.ErrorIs(err,api.ErrUserNotFindWithThisEmail)

	mockDBCall.Unset()
	mockAuthCall.Unset()
	mockEmailCall.Unset()
}

func (a *ApplicationTestSuit) TestResetPasswordShouldFailWhenEmailFails() {
	ctx := context.TODO()

	firstName := "John"
	lastName := "Doe"
	email := "johnDoe@gmail.com"

	user := domain.NewUser(firstName, lastName, email)

	mockDBCall := a.mockedDB.On("GetUserByEmail", mock.Anything, mock.Anything).Return(user,nil)
	mockAuthCall := a.mockedAuth.On("ResetPassword", mock.Anything, mock.Anything).Return("token", 2, nil)
	mockEmailCall := a.mockedEmail.On("SendResetPassword", mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(ErrUnknownError)

	err := a.api.ResetPassword(ctx, user.UserID)

	a.ErrorIs(err,api.ErrResetPasswordFailed)

	mockDBCall.Unset()
	mockAuthCall.Unset()
	mockEmailCall.Unset()
}

func (a *ApplicationTestSuit) TestUpdate() {
	ctx := context.TODO()

	firstName := "John"
	lastName := "Doe"
	email := "johnDoe@gmail.com"

	user := domain.NewUser(firstName, lastName, email)

	mockDBCall := a.mockedDB.On("Update", mock.Anything, mock.Anything,mock.Anything, mock.Anything).Return(user,nil)
	mockAuthCall := a.mockedAuth.On("Authenticate", mock.Anything, mock.Anything).Return(user.UserID, nil)

	err := a.api.Update(ctx,"accessToken",firstName,lastName)

	a.NoError(err)

	mockDBCall.Unset()
	mockAuthCall.Unset()
}

func (a *ApplicationTestSuit) TestUpdateShouldFailWhenDBFails() {
	ctx := context.TODO()

	firstName := "John"
	lastName := "Doe"
	email := "johnDoe@gmail.com"

	user := domain.NewUser(firstName, lastName, email)

	mockDBCall := a.mockedDB.On("Update", mock.Anything, mock.Anything,mock.Anything, mock.Anything).Return(&domain.User{},ErrUnknownError)
	mockAuthCall := a.mockedAuth.On("Authenticate", mock.Anything, mock.Anything).Return(user.UserID, nil)

	err := a.api.Update(ctx,"accessToken",firstName,lastName)

	a.ErrorIs(err,api.ErrUpdateFailed)

	mockDBCall.Unset()
	mockAuthCall.Unset()
}

func (a *ApplicationTestSuit) TestUpdateShouldFailWhenAuthFails() {
	ctx := context.TODO()

	firstName := "John"
	lastName := "Doe"
	email := "johnDoe@gmail.com"

	user := domain.NewUser(firstName, lastName, email)

	mockDBCall := a.mockedDB.On("Update", mock.Anything, mock.Anything,mock.Anything, mock.Anything).Return(user,nil)
	mockAuthCall := a.mockedAuth.On("Authenticate", mock.Anything, mock.Anything).Return("", ErrUnknownError)

	err := a.api.Update(ctx,"accessToken",firstName,lastName)

	a.ErrorIs(err,api.ErrUpdateFailed)

	mockDBCall.Unset()
	mockAuthCall.Unset()
}

func (a *ApplicationTestSuit) TestDeleteAccount() {
	ctx := context.TODO()

	password := "password"

	mockDBCall := a.mockedDB.On("Delete", mock.Anything,mock.Anything).Return(nil)
	mockAuthCallAuthentication := a.mockedAuth.On("Authenticate", mock.Anything, mock.Anything).Return("id",nil)
	mockAuthCallDeleteAccount := a.mockedAuth.On("DeleteAccount", mock.Anything, mock.Anything, mock.Anything).Return(nil)

	err := a.api.DeleteAccount(ctx,"accessToken",password)

	a.NoError(err)

	mockDBCall.Unset()
	mockAuthCallAuthentication.Unset()
	mockAuthCallDeleteAccount.Unset()
}

func (a *ApplicationTestSuit) TestDeleteAccountShouldFailWhenDBFails() {
	ctx := context.TODO()

	password := "password"

	mockDBCall := a.mockedDB.On("Delete", mock.Anything,mock.Anything).Return(ErrUnknownError)
	mockAuthCallAuthentication := a.mockedAuth.On("Authenticate", mock.Anything, mock.Anything).Return("id",nil)
	mockAuthCallDeleteAccount := a.mockedAuth.On("DeleteAccount", mock.Anything, mock.Anything, mock.Anything).Return(nil)

	err := a.api.DeleteAccount(ctx,"accessToken",password)

	a.ErrorIs(err,api.ErrDeleteAccountFailed)

	mockDBCall.Unset()
	mockAuthCallAuthentication.Unset()
	mockAuthCallDeleteAccount.Unset()
}

func (a *ApplicationTestSuit) TestDeleteAccountShouldFailWhenAuthFailsAuthenticate() {
	ctx := context.TODO()

	password := "password"

	mockDBCall := a.mockedDB.On("Delete", mock.Anything,mock.Anything).Return(nil)
	mockAuthCallAuthentication := a.mockedAuth.On("Authenticate", mock.Anything, mock.Anything).Return("",ErrUnknownError)
	mockAuthCallDeleteAccount := a.mockedAuth.On("DeleteAccount", mock.Anything, mock.Anything, mock.Anything).Return(nil)

	err := a.api.DeleteAccount(ctx,"accessToken",password)

	a.ErrorIs(err,api.ErrDeleteAccountFailed)

	mockDBCall.Unset()
	mockAuthCallAuthentication.Unset()
	mockAuthCallDeleteAccount.Unset()
}

func (a *ApplicationTestSuit) TestDeleteAccountShouldFailWhenAuthFails() {
	ctx := context.TODO()

	password := "password"

	mockDBCall := a.mockedDB.On("Delete", mock.Anything,mock.Anything).Return(nil)
	mockAuthCallAuthentication := a.mockedAuth.On("Authenticate", mock.Anything, mock.Anything).Return("id",nil)
	mockAuthCallDeleteAccount := a.mockedAuth.On("DeleteAccount", mock.Anything, mock.Anything, mock.Anything).Return(ErrUnknownError)

	err := a.api.DeleteAccount(ctx,"accessToken",password)

	a.ErrorIs(err,api.ErrDeleteAccountFailed)

	mockDBCall.Unset()
	mockAuthCallAuthentication.Unset()
	mockAuthCallDeleteAccount.Unset()
}