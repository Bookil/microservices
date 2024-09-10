package api_test

import (
	"context"
	"errors"
	"testing"
	"time"

	"github.com/Bookil/microservices/auth/internal/adapters/auth_manager"
	"github.com/Bookil/microservices/auth/internal/application/core/api"
	"github.com/Bookil/microservices/auth/internal/application/core/domain"
	"github.com/Bookil/microservices/auth/mocks"
	"github.com/stretchr/testify/suite"
	"go.uber.org/mock/gomock"
)

type ApplicationTestSuit struct {
	suite.Suite

	api *api.Application

	user        *mocks.MockUserPorts
	authManger  *mocks.MockAuthManager
	hashManager *mocks.MockHashManager
	DB          *mocks.MockDBPort
	email       *mocks.MockEmailPort
}

var ErrUnknownError = errors.New("random error")

func TestApplicationTestSuit(t *testing.T) {
	suite.Run(t, new(ApplicationTestSuit))
}

func (a *ApplicationTestSuit) SetupSuite() {
	ctrl := gomock.NewController(a.T())

	mockedUser := mocks.NewMockUserPorts(ctrl)
	mockedAuthManger := mocks.NewMockAuthManager(ctrl)
	mockedHashManger := mocks.NewMockHashManager(ctrl)
	mockedDB := mocks.NewMockDBPort(ctrl)
	mockedEmail := mocks.NewMockEmailPort(ctrl)

	api := api.NewApplication(mockedDB, mockedUser, mockedEmail, mockedAuthManger, mockedHashManger)

	a.api = api

	a.authManger = mockedAuthManger
	a.hashManager = mockedHashManger
	a.user = mockedUser
	a.DB = mockedDB
	a.email = mockedEmail
}

func (a *ApplicationTestSuit) TestRegister_Success() {
	ctx := context.TODO()
	password := "password"
	email := "john.doe@example.com"
	hashedPassword := "hashedPassword"
	verificationCode := "verificationCode"

	a.user.EXPECT().Register(ctx, "John", "Doe", email).Return(email, nil)
	a.hashManager.EXPECT().HashPassword(password).Return(hashedPassword, nil)
	a.DB.EXPECT().Create(ctx, gomock.Any()).Return(nil, nil)
	a.authManger.EXPECT().GenerateVerificationCode(ctx, email).Return(verificationCode, nil)
	a.email.EXPECT().SendVerificationCode("john.doe@example.com", verificationCode).Return(nil)

	userID, err := a.api.Register(ctx, "John", "Doe", email, password)
	a.NotEmpty(userID)
	a.NoError(err)
}

func (a *ApplicationTestSuit) TestRegister_UserRegisterError() {
	ctx := context.TODO()

	a.user.EXPECT().Register(ctx, "John", "Doe", "john.doe@example.com").Return("", ErrUnknownError)

	userID, err := a.api.Register(ctx, "John", "Doe", "john.doe@example.com", "password")
	a.Empty(userID)
	a.Error(err)
}

func (a *ApplicationTestSuit) TestRegister_HashPasswordError() {
	ctx := context.TODO()
	email := "john.doe@example.com"

	a.user.EXPECT().Register(ctx, "John", "Doe", email).Return(email, nil)
	a.hashManager.EXPECT().HashPassword("password").Return("", ErrUnknownError)

	userID, err := a.api.Register(ctx, "John", "Doe", email, "password")
	a.Empty(userID)
	a.Error(err)
}

func (a *ApplicationTestSuit) TestRegister_CreateAuthStoreError() {
	ctx := context.TODO()
	email := "john.doe@example.com"
	password := "password"
	hashedPassword := "hashedPassword"

	a.user.EXPECT().Register(ctx, "John", "Doe", email).Return(email, nil)
	a.hashManager.EXPECT().HashPassword(password).Return(hashedPassword, nil)
	a.DB.EXPECT().Create(ctx, gomock.Any()).Return(nil, ErrUnknownError)

	userID, err := a.api.Register(ctx, "John", "Doe", email, password)
	a.Empty(userID)
	a.Error(err)
}

func (a *ApplicationTestSuit) TestRegister_GenerateVerificationCodeError() {
	ctx := context.TODO()
	email := "john.doe@example.com"
	password := "password"
	hashedPassword := "hashedPassword"

	a.user.EXPECT().Register(ctx, "John", "Doe", email).Return(email, nil)
	a.hashManager.EXPECT().HashPassword(password).Return(hashedPassword, nil)
	a.DB.EXPECT().Create(ctx, gomock.Any()).Return(nil, nil)
	a.authManger.EXPECT().GenerateVerificationCode(ctx, email).Return("", ErrUnknownError)

	email, err := a.api.Register(ctx, "John", "Doe", email, password)
	a.Empty(email)
	a.Error(err)
}

func (a *ApplicationTestSuit) TestRegister_SendVerificationCodeError() {
	ctx := context.TODO()
	email := "john.doe@example.com"
	password := "password"
	hashedPassword := "hashedPassword"
	verificationCode := "verificationCode"

	a.user.EXPECT().Register(ctx, "John", "Doe", email).Return(email, nil)
	a.hashManager.EXPECT().HashPassword(password).Return(hashedPassword, nil)
	a.DB.EXPECT().Create(ctx, gomock.Any()).Return(nil, nil)
	a.authManger.EXPECT().GenerateVerificationCode(ctx, email).Return(verificationCode, nil)
	a.email.EXPECT().SendVerificationCode(email, verificationCode).Return(ErrUnknownError)

	userID, err := a.api.Register(ctx, "John", "Doe", email, password)
	a.Empty(userID)
	a.Error(err)
}

func (a *ApplicationTestSuit) TestVerifyEmail_Success() {
	ctx := context.TODO()
	expectedUserID := "123456"
	email := "john.doe@example.com"
	verificationCode := "validCode"

	a.authManger.EXPECT().CompareVerificationCode(ctx, email, verificationCode).Return(true, nil)
	a.user.EXPECT().GetUserIDByEmail(ctx, email).Return(expectedUserID, nil)
	a.DB.EXPECT().VerifyEmail(ctx, expectedUserID).Return(nil, nil)

	err := a.api.VerifyEmail(ctx, email, verificationCode)
	a.NoError(err)
}

func (a *ApplicationTestSuit) TestVerifyEmail_CompareVerificationCodeError() {
	ctx := context.TODO()
	email := "john.doe@example.com"
	verificationCode := "invalidCode"

	a.authManger.EXPECT().CompareVerificationCode(ctx, email, verificationCode).Return(false, ErrUnknownError)

	err := a.api.VerifyEmail(ctx, email, verificationCode)
	a.Error(err)
}

func (a *ApplicationTestSuit) TestVerifyEmail_VerifyEmailUserError() {
	ctx := context.TODO()
	email := "john.doe@example.com"
	verificationCode := "validCode"

	a.authManger.EXPECT().CompareVerificationCode(ctx, email, verificationCode).Return(true, nil)
	a.user.EXPECT().GetUserIDByEmail(ctx, email).Return("", ErrUnknownError)

	err := a.api.VerifyEmail(ctx, email, verificationCode)
	a.Error(err)
}

func (a *ApplicationTestSuit) TestVerifyEmail_VerifyEmailDBError() {
	ctx := context.TODO()
	userID := "123456"
	email := "john.doe@example.com"
	verificationCode := "validCode"

	a.authManger.EXPECT().CompareVerificationCode(ctx, email, verificationCode).Return(true, nil)
	a.user.EXPECT().GetUserIDByEmail(ctx, email).Return(userID, nil)
	a.DB.EXPECT().VerifyEmail(ctx, userID).Return(nil, ErrUnknownError)


	err := a.api.VerifyEmail(ctx, email, verificationCode)
	a.Error(err)
}

func (a *ApplicationTestSuit) TestLogin_Success() {
	ctx := context.TODO()
	email := "john.doe@example.com"
	password := "password"
	userID := "123456"
	hashedPassword := "hashedPassword"
	accessToken := "accessToken"
	refreshToken := "refreshToken"

	auth := &domain.Auth{
		UserID:              userID,
		HashedPassword:      hashedPassword,
		IsEmailVerified:     true,
		AccountLockedUntil:  0,
		FailedLoginAttempts: 0,
	}

	a.user.EXPECT().GetUserIDByEmail(ctx, email).Return(userID, nil)
	a.DB.EXPECT().GetByID(ctx, userID).Return(auth, nil)
	a.hashManager.EXPECT().CheckPasswordHash(password, hashedPassword).Return(true)
	a.authManger.EXPECT().GenerateAccessToken(ctx, userID).Return(accessToken, nil)
	a.authManger.EXPECT().GenerateRefreshToken(ctx, userID).Return(refreshToken, nil)
	a.DB.EXPECT().ClearFailedLoginAttempts(ctx, userID).Return(nil, nil)
	a.email.EXPECT().SendWelcome(email).Return(nil)

	at, rt, err := a.api.Login(ctx, email, password)
	a.NotEmpty(at)
	a.NotEmpty(rt)
	a.NoError(err)
}

func (a *ApplicationTestSuit) TestLogin_GetUserIDByEmailError() {
	ctx := context.TODO()
	email := "john.doe@example.com"
	password := "password"

	a.user.EXPECT().GetUserIDByEmail(ctx, email).Return("", ErrUnknownError)

	at, rt, err := a.api.Login(ctx, email, password)
	a.Empty(at)
	a.Empty(rt)
	a.Error(err)
}

func (a *ApplicationTestSuit) TestLogin_GetByIDError() {
	ctx := context.TODO()
	email := "john.doe@example.com"
	password := "password"
	userID := "123456"

	a.user.EXPECT().GetUserIDByEmail(ctx, email).Return(userID, nil)
	a.DB.EXPECT().GetByID(ctx, userID).Return(nil, ErrUnknownError)

	at, rt, err := a.api.Login(ctx, email, password)
	a.Empty(at)
	a.Empty(rt)
	a.Error(err)
}

func (a *ApplicationTestSuit) TestLogin_EmailNotVerified() {
	ctx := context.TODO()
	email := "john.doe@example.com"
	password := "password"
	userID := "123456"
	hashedPassword := "hashedPassword"

	auth := &domain.Auth{
		UserID:              userID,
		HashedPassword:      hashedPassword,
		IsEmailVerified:     false,
		AccountLockedUntil:  0,
		FailedLoginAttempts: 0,
	}

	a.user.EXPECT().GetUserIDByEmail(ctx, email).Return(userID, nil)
	a.DB.EXPECT().GetByID(ctx, userID).Return(auth, nil)

	at, rt, err := a.api.Login(ctx, email, password)
	a.Empty(at)
	a.Empty(rt)
	a.Error(err)
}

func (a *ApplicationTestSuit) TestLogin_AccountLocked() {
	ctx := context.TODO()
	email := "john.doe@example.com"
	password := "password"
	userID := "123456"
	hashedPassword := "hashedPassword"
	futureTime := time.Now().Add(time.Hour).Unix()

	auth := &domain.Auth{
		UserID:              userID,
		HashedPassword:      hashedPassword,
		IsEmailVerified:     true,
		AccountLockedUntil:  futureTime,
		FailedLoginAttempts: 0,
	}

	a.user.EXPECT().GetUserIDByEmail(ctx, email).Return(userID, nil)
	a.DB.EXPECT().GetByID(ctx, userID).Return(auth, nil)

	at, rt, err := a.api.Login(ctx, email, password)
	a.Empty(at)
	a.Empty(rt)
	a.Error(err)
}

func (a *ApplicationTestSuit) TestLogin_TooManyFailedAttempts() {
	ctx := context.TODO()
	email := "john.doe@example.com"
	password := "password"
	userID := "123456"
	hashedPassword := "hashedPassword"

	auth := &domain.Auth{
		UserID:              userID,
		HashedPassword:      hashedPassword,
		IsEmailVerified:     true,
		AccountLockedUntil:  0,
		FailedLoginAttempts: api.MaximumFailedLoginAttempts,
	}

	a.user.EXPECT().GetUserIDByEmail(ctx, email).Return(userID, nil)
	a.DB.EXPECT().GetByID(ctx, userID).Return(auth, nil)
	a.DB.EXPECT().LockAccount(ctx, userID, api.LockAccountDuration).Return(nil, nil)

	at, rt, err := a.api.Login(ctx, email, password)
	a.Empty(at)
	a.Empty(rt)
	a.Error(err)
}

func (a *ApplicationTestSuit) TestLogin_TooManyFailedAttemptsAndFailedLockAccount() {
	ctx := context.TODO()
	email := "john.doe@example.com"
	password := "password"
	userID := "123456"
	hashedPassword := "hashedPassword"

	auth := &domain.Auth{
		UserID:              userID,
		HashedPassword:      hashedPassword,
		IsEmailVerified:     true,
		AccountLockedUntil:  0,
		FailedLoginAttempts: api.MaximumFailedLoginAttempts,
	}

	a.user.EXPECT().GetUserIDByEmail(ctx, email).Return(userID, nil)
	a.DB.EXPECT().GetByID(ctx, userID).Return(auth, nil)
	a.DB.EXPECT().LockAccount(ctx, userID, api.LockAccountDuration).Return(nil, ErrUnknownError)

	at, rt, err := a.api.Login(ctx, email, password)
	a.Empty(at)
	a.Empty(rt)
	a.Error(err)
}

func (a *ApplicationTestSuit) TestLogin_InvalidPassword() {
	ctx := context.TODO()
	email := "john.doe@example.com"
	password := "password"
	userID := "123456"
	hashedPassword := "hashedPassword"

	auth := &domain.Auth{
		UserID:              userID,
		HashedPassword:      hashedPassword,
		IsEmailVerified:     true,
		AccountLockedUntil:  0,
		FailedLoginAttempts: 0,
	}

	a.user.EXPECT().GetUserIDByEmail(ctx, email).Return(userID, nil)
	a.DB.EXPECT().GetByID(ctx, userID).Return(auth, nil)
	a.hashManager.EXPECT().CheckPasswordHash(password, hashedPassword).Return(false)
	a.DB.EXPECT().IncrementFailedLoginAttempts(ctx, userID).Return(nil, nil)

	at, rt, err := a.api.Login(ctx, email, password)
	a.Empty(at)
	a.Empty(rt)
	a.Error(err)
}

func (a *ApplicationTestSuit) TestLogin_InvalidPasswordAndIncrementFailedLoginAttempts() {
	ctx := context.TODO()
	email := "john.doe@example.com"
	password := "password"
	userID := "123456"
	hashedPassword := "hashedPassword"

	auth := &domain.Auth{
		UserID:              userID,
		HashedPassword:      hashedPassword,
		IsEmailVerified:     true,
		AccountLockedUntil:  0,
		FailedLoginAttempts: 0,
	}

	a.user.EXPECT().GetUserIDByEmail(ctx, email).Return(userID, nil)
	a.DB.EXPECT().GetByID(ctx, userID).Return(auth, nil)
	a.hashManager.EXPECT().CheckPasswordHash(password, hashedPassword).Return(false)
	a.DB.EXPECT().IncrementFailedLoginAttempts(ctx, userID).Return(nil, ErrUnknownError)

	at, rt, err := a.api.Login(ctx, email, password)
	a.Empty(at)
	a.Empty(rt)
	a.Error(err)
}

func (a *ApplicationTestSuit) TestLogin_GenerateAccessTokenError() {
	ctx := context.TODO()
	email := "john.doe@example.com"
	password := "password"
	userID := "123456"
	hashedPassword := "hashedPassword"

	auth := &domain.Auth{
		UserID:              userID,
		HashedPassword:      hashedPassword,
		IsEmailVerified:     true,
		AccountLockedUntil:  0,
		FailedLoginAttempts: 0,
	}

	a.user.EXPECT().GetUserIDByEmail(ctx, email).Return(userID, nil)
	a.DB.EXPECT().GetByID(ctx, userID).Return(auth, nil)
	a.hashManager.EXPECT().CheckPasswordHash(password, hashedPassword).Return(true)
	a.authManger.EXPECT().GenerateAccessToken(ctx, userID).Return("", ErrUnknownError)

	at, rt, err := a.api.Login(ctx, email, password)
	a.Empty(at)
	a.Empty(rt)
	a.Error(err)
}

func (a *ApplicationTestSuit) TestLogin_GenerateRefreshTokenError() {
	ctx := context.TODO()
	email := "john.doe@example.com"
	password := "password"
	userID := "123456"
	hashedPassword := "hashedPassword"
	accessToken := "accessToken"

	auth := &domain.Auth{
		UserID:              userID,
		HashedPassword:      hashedPassword,
		IsEmailVerified:     true,
		AccountLockedUntil:  0,
		FailedLoginAttempts: 0,
	}

	a.user.EXPECT().GetUserIDByEmail(ctx, email).Return(userID, nil)
	a.DB.EXPECT().GetByID(ctx, userID).Return(auth, nil)
	a.hashManager.EXPECT().CheckPasswordHash(password, hashedPassword).Return(true)
	a.authManger.EXPECT().GenerateAccessToken(ctx, userID).Return(accessToken, nil)
	a.authManger.EXPECT().GenerateRefreshToken(ctx, userID).Return("", ErrUnknownError)

	at, rt, err := a.api.Login(ctx, email, password)
	a.Empty(at)
	a.Empty(rt)
	a.Error(err)
}

func (a *ApplicationTestSuit) TestLogin_ClearFailedLoginAttemptsError() {
	ctx := context.TODO()
	email := "john.doe@example.com"
	password := "password"
	userID := "123456"
	hashedPassword := "hashedPassword"
	accessToken := "accessToken"
	refreshToken := "refreshToken"

	auth := &domain.Auth{
		UserID:              userID,
		HashedPassword:      hashedPassword,
		IsEmailVerified:     true,
		AccountLockedUntil:  0,
		FailedLoginAttempts: 0,
	}

	a.user.EXPECT().GetUserIDByEmail(ctx, email).Return(userID, nil)
	a.DB.EXPECT().GetByID(ctx, userID).Return(auth, nil)
	a.hashManager.EXPECT().CheckPasswordHash(password, hashedPassword).Return(true)
	a.authManger.EXPECT().GenerateAccessToken(ctx, userID).Return(accessToken, nil)
	a.authManger.EXPECT().GenerateRefreshToken(ctx, userID).Return(refreshToken, nil)
	a.DB.EXPECT().ClearFailedLoginAttempts(ctx, userID).Return(nil, ErrUnknownError)

	at, rt, err := a.api.Login(ctx, email, password)
	a.Empty(at)
	a.Empty(rt)
	a.Error(err)
}

func (a *ApplicationTestSuit) TestLogin_SendWelcomeEmailError() {
	ctx := context.TODO()
	email := "john.doe@example.com"
	password := "password"
	userID := "123456"
	hashedPassword := "hashedPassword"
	accessToken := "accessToken"
	refreshToken := "refreshToken"

	auth := &domain.Auth{
		UserID:              userID,
		HashedPassword:      hashedPassword,
		IsEmailVerified:     true,
		AccountLockedUntil:  0,
		FailedLoginAttempts: 0,
	}

	a.user.EXPECT().GetUserIDByEmail(ctx, email).Return(userID, nil)
	a.DB.EXPECT().GetByID(ctx, userID).Return(auth, nil)
	a.hashManager.EXPECT().CheckPasswordHash(password, hashedPassword).Return(true)
	a.authManger.EXPECT().GenerateAccessToken(ctx, userID).Return(accessToken, nil)
	a.authManger.EXPECT().GenerateRefreshToken(ctx, userID).Return(refreshToken, nil)
	a.DB.EXPECT().ClearFailedLoginAttempts(ctx, userID).Return(nil, nil)
	a.email.EXPECT().SendWelcome(email).Return(ErrUnknownError)

	at, rt, err := a.api.Login(ctx, email, password)
	a.Empty(at)
	a.Empty(rt)
	a.Error(err)
}

func (a *ApplicationTestSuit) TestAuthenticate_Success() {
	ctx := context.TODO()
	accessToken := "validAccessToken"
	expectedUserID := "123456"

	accessTokenClaims := &domain.AccessTokenClaims{
		UserID: expectedUserID,
	}

	a.authManger.EXPECT().DecodeAccessToken(ctx, accessToken).Return(accessTokenClaims, nil)

	userID, err := a.api.Authenticate(ctx, accessToken)
	a.Equal(expectedUserID, userID)
	a.NoError(err)
}

func (a *ApplicationTestSuit) TestAuthenticate_DecodeAccessTokenError() {
	ctx := context.TODO()
	accessToken := "invalidAccessToken"

	a.authManger.EXPECT().DecodeAccessToken(ctx, accessToken).Return(nil, ErrUnknownError)

	userID, err := a.api.Authenticate(ctx, accessToken)
	a.Empty(userID)
	a.Error(err)
}

func (a *ApplicationTestSuit) TestAuthenticate_EmptyUserID() {
	ctx := context.TODO()
	accessToken := "validAccessTokenWithEmptyUserID"

	accessTokenClaims := &domain.AccessTokenClaims{
		UserID: "",
	}

	a.authManger.EXPECT().DecodeAccessToken(ctx, accessToken).Return(accessTokenClaims, nil)

	userID, err := a.api.Authenticate(ctx, accessToken)
	a.Empty(userID)
	a.Error(err)
}

func (a *ApplicationTestSuit) TestChangePassword_Success() {
	ctx := context.TODO()
	userID := "123456"
	oldPassword := "oldPassword"
	newPassword := "newPassword"
	hashedPassword := "hashedPassword"
	newHashedPassword := "newHashedPassword"

	auth := &domain.Auth{
		UserID:         userID,
		HashedPassword: hashedPassword,
	}

	a.DB.EXPECT().GetByID(ctx, userID).Return(auth, nil)
	a.hashManager.EXPECT().CheckPasswordHash(oldPassword, hashedPassword).Return(true)
	a.hashManager.EXPECT().HashPassword(newPassword).Return(newHashedPassword, nil)
	a.DB.EXPECT().ChangePassword(ctx, userID, newHashedPassword).Return(nil, nil)

	err := a.api.ChangePassword(ctx, userID, oldPassword, newPassword)
	a.NoError(err)
}

func (a *ApplicationTestSuit) TestChangePassword_GetByIDError() {
	ctx := context.TODO()
	userID := "123456"
	oldPassword := "oldPassword"
	newPassword := "newPassword"

	a.DB.EXPECT().GetByID(ctx, userID).Return(nil, ErrUnknownError)

	err := a.api.ChangePassword(ctx, userID, oldPassword, newPassword)
	a.Error(err)
}

func (a *ApplicationTestSuit) TestChangePassword_InvalidOldPassword() {
	ctx := context.TODO()
	userID := "123456"
	oldPassword := "oldPassword"
	newPassword := "newPassword"
	hashedPassword := "hashedPassword"

	auth := &domain.Auth{
		UserID:         userID,
		HashedPassword: hashedPassword,
	}

	a.DB.EXPECT().GetByID(ctx, userID).Return(auth, nil)
	a.hashManager.EXPECT().CheckPasswordHash(oldPassword, hashedPassword).Return(false)

	err := a.api.ChangePassword(ctx, userID, oldPassword, newPassword)
	a.Error(err)
}

func (a *ApplicationTestSuit) TestChangePassword_HashPasswordError() {
	ctx := context.TODO()
	userID := "123456"
	oldPassword := "oldPassword"
	newPassword := "newPassword"
	hashedPassword := "hashedPassword"

	auth := &domain.Auth{
		UserID:         userID,
		HashedPassword: hashedPassword,
	}

	a.DB.EXPECT().GetByID(ctx, userID).Return(auth, nil)
	a.hashManager.EXPECT().CheckPasswordHash(oldPassword, hashedPassword).Return(true)
	a.hashManager.EXPECT().HashPassword(newPassword).Return("", ErrUnknownError)

	err := a.api.ChangePassword(ctx, userID, oldPassword, newPassword)
	a.Error(err)
}

func (a *ApplicationTestSuit) TestChangePassword_ChangePasswordError() {
	ctx := context.TODO()
	userID := "123456"
	oldPassword := "oldPassword"
	newPassword := "newPassword"
	hashedPassword := "hashedPassword"
	newHashedPassword := "newHashedPassword"

	auth := &domain.Auth{
		UserID:         userID,
		HashedPassword: hashedPassword,
	}

	a.DB.EXPECT().GetByID(ctx, userID).Return(auth, nil)
	a.hashManager.EXPECT().CheckPasswordHash(oldPassword, hashedPassword).Return(true)
	a.hashManager.EXPECT().HashPassword(newPassword).Return(newHashedPassword, nil)
	a.DB.EXPECT().ChangePassword(ctx, userID, newHashedPassword).Return(nil, ErrUnknownError)

	err := a.api.ChangePassword(ctx, userID, oldPassword, newPassword)
	a.Error(err)
}

func (a *ApplicationTestSuit) TestRefreshToken_Success() {
	ctx := context.TODO()
	userID := "123456"
	refreshToken := "validRefreshToken"
	newAccessToken := "newAccessToken"

	auth := &domain.Auth{
		UserID: userID,
	}

	a.authManger.EXPECT().DecodeRefreshToken(ctx, userID, refreshToken).Return(nil, nil)
	a.DB.EXPECT().GetByID(ctx, userID).Return(auth, nil)
	a.authManger.EXPECT().GenerateAccessToken(ctx, userID).Return(newAccessToken, nil)

	accessToken, err := a.api.RefreshToken(ctx, userID, refreshToken)
	a.NotEmpty(accessToken)
	a.NoError(err)
}

func (a *ApplicationTestSuit) TestRefreshToken_DecodeRefreshTokenError() {
	ctx := context.TODO()
	userID := "123456"
	refreshToken := "invalidRefreshToken"

	a.authManger.EXPECT().DecodeRefreshToken(ctx, userID, refreshToken).Return(nil, ErrUnknownError)

	accessToken, err := a.api.RefreshToken(ctx, userID, refreshToken)
	a.Empty(accessToken)
	a.Error(err)
}

func (a *ApplicationTestSuit) TestRefreshToken_GetByIDError() {
	ctx := context.TODO()
	userID := "123456"
	refreshToken := "validRefreshToken"

	a.authManger.EXPECT().DecodeRefreshToken(ctx, userID, refreshToken).Return(nil, nil)
	a.DB.EXPECT().GetByID(ctx, userID).Return(nil, ErrUnknownError)

	accessToken, err := a.api.RefreshToken(ctx, userID, refreshToken)
	a.Empty(accessToken)
	a.Error(err)
}

func (a *ApplicationTestSuit) TestRefreshToken_GenerateAccessTokenError() {
	ctx := context.TODO()
	userID := "123456"
	refreshToken := "validRefreshToken"

	auth := &domain.Auth{
		UserID: userID,
	}

	a.authManger.EXPECT().DecodeRefreshToken(ctx, userID, refreshToken).Return(nil, nil)
	a.DB.EXPECT().GetByID(ctx, userID).Return(auth, nil)
	a.authManger.EXPECT().GenerateAccessToken(ctx, userID).Return("", ErrUnknownError)

	accessToken, err := a.api.RefreshToken(ctx, userID, refreshToken)
	a.Empty(accessToken)
	a.Error(err)
}

func (a *ApplicationTestSuit) TestResetPassword_Success() {
	ctx := context.TODO()
	email := "john.doe@example.com"
	userID := "123456"
	resetPasswordToken := "resetPasswordToken"

	auth := &domain.Auth{
		UserID:              userID,
		IsEmailVerified:     true,
		FailedLoginAttempts: 0,
	}

	a.user.EXPECT().GetUserIDByEmail(ctx, email).Return(userID, nil)
	a.DB.EXPECT().GetByID(ctx, userID).Return(auth, nil)
	a.authManger.EXPECT().GenerateResetPasswordToken(ctx, userID).Return(resetPasswordToken, nil)
	a.email.EXPECT().SendResetPassword("example.com", resetPasswordToken, email, auth_manager.ResetPasswordTokenExpr).Return(nil)

	err := a.api.ResetPassword(ctx, email)
	a.NoError(err)
}

func (a *ApplicationTestSuit) TestResetPassword_GetUserIDByEmailError() {
	ctx := context.TODO()
	email := "john.doe@example.com"
	a.user.EXPECT().GetUserIDByEmail(ctx, email).Return("", ErrUnknownError)

	err := a.api.ResetPassword(ctx, email)
	a.Error(err)
}

func (a *ApplicationTestSuit) TestResetPassword_GetByIDError() {
	ctx := context.TODO()
	email := "john.doe@example.com"
	userID := "123456"

	a.user.EXPECT().GetUserIDByEmail(ctx, email).Return(userID, nil)
	a.DB.EXPECT().GetByID(ctx, userID).Return(nil, ErrUnknownError)

	err := a.api.ResetPassword(ctx, email)
	a.Error(err)
}

func (a *ApplicationTestSuit) TestResetPassword_EmailNotVerified() {
	ctx := context.TODO()
	email := "john.doe@example.com"
	userID := "123456"

	auth := &domain.Auth{
		UserID:              userID,
		IsEmailVerified:     false,
		FailedLoginAttempts: 0,
	}

	a.user.EXPECT().GetUserIDByEmail(ctx, email).Return(userID, nil)
	a.DB.EXPECT().GetByID(ctx, userID).Return(auth, nil)

	err := a.api.ResetPassword(ctx, email)
	a.Error(err)
}

func (a *ApplicationTestSuit) TestResetPassword_AccountLocked() {
	ctx := context.TODO()
	email := "john.doe@example.com"
	userID := "123456"

	auth := &domain.Auth{
		UserID:              userID,
		IsEmailVerified:     true,
		FailedLoginAttempts: api.MaximumFailedLoginAttempts,
	}

	a.user.EXPECT().GetUserIDByEmail(ctx, email).Return(userID, nil)
	a.DB.EXPECT().GetByID(ctx, userID).Return(auth, nil)

	err := a.api.ResetPassword(ctx, email)
	a.Error(err)
}

func (a *ApplicationTestSuit) TestResetPassword_GenerateResetPasswordTokenError() {
	ctx := context.TODO()
	email := "john.doe@example.com"
	userID := "123456"

	auth := &domain.Auth{
		UserID:              userID,
		IsEmailVerified:     true,
		FailedLoginAttempts: 0,
	}

	a.user.EXPECT().GetUserIDByEmail(ctx, email).Return(userID, nil)
	a.DB.EXPECT().GetByID(ctx, userID).Return(auth, nil)
	a.authManger.EXPECT().GenerateResetPasswordToken(ctx, userID).Return("", ErrUnknownError)

	err := a.api.ResetPassword(ctx, email)
	a.Error(err)
}

func (a *ApplicationTestSuit) TestResetPassword_SendResetPasswordEmailError() {
	ctx := context.TODO()
	email := "john.doe@example.com"
	userID := "123456"
	resetPasswordToken := "resetPasswordToken"

	auth := &domain.Auth{
		UserID:              userID,
		IsEmailVerified:     true,
		FailedLoginAttempts: 0,
	}

	a.user.EXPECT().GetUserIDByEmail(ctx, email).Return(userID, nil)
	a.DB.EXPECT().GetByID(ctx, userID).Return(auth, nil)
	a.authManger.EXPECT().GenerateResetPasswordToken(ctx, userID).Return(resetPasswordToken, nil)
	a.email.EXPECT().SendResetPassword("example.com", resetPasswordToken, email, auth_manager.ResetPasswordTokenExpr).Return(ErrUnknownError)

	err := a.api.ResetPassword(ctx, email)
	a.Error(err)
}

func (a *ApplicationTestSuit) TestSubmitResetPassword_Success() {
	ctx := context.TODO()
	resetPasswordToken := "validResetPasswordToken"
	newPassword := "newPassword"
	hashedPassword := "hashedPassword"
	userID := "123456"

	resetPasswordTokenClaims := &domain.ResetPasswordTokenClaims{
		UserID: userID,
	}

	a.authManger.EXPECT().DecodeResetPasswordToken(ctx, resetPasswordToken).Return(resetPasswordTokenClaims, nil)
	a.hashManager.EXPECT().HashPassword(newPassword).Return(hashedPassword, nil)
	a.DB.EXPECT().ChangePassword(ctx, userID, hashedPassword).Return(nil, nil)

	err := a.api.SubmitResetPassword(ctx, resetPasswordToken, newPassword)
	a.NoError(err)
}

func (a *ApplicationTestSuit) TestSubmitResetPassword_DecodeResetPasswordTokenError() {
	ctx := context.TODO()
	resetPasswordToken := "invalidResetPasswordToken"
	newPassword := "newPassword"

	a.authManger.EXPECT().DecodeResetPasswordToken(ctx, resetPasswordToken).Return(nil, ErrUnknownError)

	err := a.api.SubmitResetPassword(ctx, resetPasswordToken, newPassword)
	a.Error(err)
}

func (a *ApplicationTestSuit) TestSubmitResetPassword_HashPasswordError() {
	ctx := context.TODO()
	resetPasswordToken := "validResetPasswordToken"
	newPassword := "newPassword"
	userID := "123456"

	resetPasswordTokenClaims := &domain.ResetPasswordTokenClaims{
		UserID: userID,
	}

	a.authManger.EXPECT().DecodeResetPasswordToken(ctx, resetPasswordToken).Return(resetPasswordTokenClaims, nil)
	a.hashManager.EXPECT().HashPassword(newPassword).Return("", ErrUnknownError)

	err := a.api.SubmitResetPassword(ctx, resetPasswordToken, newPassword)
	a.Error(err)
}

func (a *ApplicationTestSuit) TestSubmitResetPassword_ChangePasswordError() {
	ctx := context.TODO()
	resetPasswordToken := "validResetPasswordToken"
	newPassword := "newPassword"
	hashedPassword := "hashedPassword"
	userID := "123456"

	resetPasswordTokenClaims := &domain.ResetPasswordTokenClaims{
		UserID: userID,
	}

	a.authManger.EXPECT().DecodeResetPasswordToken(ctx, resetPasswordToken).Return(resetPasswordTokenClaims, nil)
	a.hashManager.EXPECT().HashPassword(newPassword).Return(hashedPassword, nil)
	a.DB.EXPECT().ChangePassword(ctx, userID, hashedPassword).Return(nil, ErrUnknownError)

	err := a.api.SubmitResetPassword(ctx, resetPasswordToken, newPassword)
	a.Error(err)
}

func (a *ApplicationTestSuit) TestDeleteAccount_Success() {
	ctx := context.TODO()
	userID := "123456"
	password := "password"
	hashedPassword := "hashedPassword"

	auth := &domain.Auth{
		UserID:         userID,
		HashedPassword: hashedPassword,
	}

	a.DB.EXPECT().GetByID(ctx, userID).Return(auth, nil)
	a.hashManager.EXPECT().CheckPasswordHash(password, hashedPassword).Return(true)
	a.DB.EXPECT().DeleteByID(ctx, userID).Return(nil)

	err := a.api.DeleteAccount(ctx, userID, password)
	a.NoError(err)
}

func (a *ApplicationTestSuit) TestDeleteAccount_GetByIDError() {
	ctx := context.TODO()
	userID := "123456"
	password := "password"

	a.DB.EXPECT().GetByID(ctx, userID).Return(nil, ErrUnknownError)

	err := a.api.DeleteAccount(ctx, userID, password)
	a.Error(err)
}

func (a *ApplicationTestSuit) TestDeleteAccount_InvalidPassword() {
	ctx := context.TODO()
	userID := "123456"
	password := "password"
	hashedPassword := "hashedPassword"

	auth := &domain.Auth{
		UserID:         userID,
		HashedPassword: hashedPassword,
	}

	a.DB.EXPECT().GetByID(ctx, userID).Return(auth, nil)
	a.hashManager.EXPECT().CheckPasswordHash(password, hashedPassword).Return(false)

	err := a.api.DeleteAccount(ctx, userID, password)
	a.Error(err)
}

func (a *ApplicationTestSuit) TestDeleteAccount_DeleteByIDError() {
	ctx := context.TODO()
	userID := "123456"
	password := "password"
	hashedPassword := "hashedPassword"

	auth := &domain.Auth{
		UserID:         userID,
		HashedPassword: hashedPassword,
	}

	a.DB.EXPECT().GetByID(ctx, userID).Return(auth, nil)
	a.hashManager.EXPECT().CheckPasswordHash(password, hashedPassword).Return(true)
	a.DB.EXPECT().DeleteByID(ctx, userID).Return(ErrUnknownError)

	err := a.api.DeleteAccount(ctx, userID, password)
	a.Error(err)
}
