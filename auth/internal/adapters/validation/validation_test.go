package validation_test

import (
	"testing"

	"github.com/Bookil/microservices/auth/internal/adapters/validation"
	"github.com/Bookil/microservices/auth/internal/application/core/domain"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
)

type VerificationTestSuite struct {
	suite.Suite
	validator *validation.Validator
}

func TestVerificationTestSuite(t *testing.T) {
	suite.Run(t, new(VerificationTestSuite))
}

func (v *VerificationTestSuite) SetupSuite() {
	v.validator = validation.NewValidator()
}
func (v *VerificationTestSuite) TestValidateRegister() {
	testCases := []struct {
		firstName string
		lastName string
		email   string
		password string
		valid    bool
	}{
		{
			firstName: "",
			lastName: "valid",
			email:   "valid@gmail.com",
			password: "12345678",
			valid:    false,
		},
		{
			firstName: "valid",
			lastName: "",
			email:   "valid@gmail.com",
			password: "12345678",
			valid:    false,
		},
		{
			firstName: "valid",
			lastName: "valid",
			email:   "",
			password: "12345678",
			valid:    false,
		},
		{
			firstName: "valid",
			lastName: "valid",
			email:   "valid@gmail.com",
			password: "",
			valid:    false,
		},
		{
			firstName: "valid",
			lastName: "valid",
			email:   "valid@valid.com",
			password: "12345",
			valid:    false,
		},
		{
			firstName: "valid",
			lastName: "valid",
			email:   "valid@valid.com",
			password: "12345678",
			valid:    true,
		},
	}

	validator := validation.NewValidator()

	for _, tc := range testCases {
		err := validator.ValidateRegisterInputs(tc.firstName,tc.lastName,tc.email,tc.password)

		if tc.valid {
			require.NoError(v.T(), err)
		} else {
			require.Error(v.T(), err)
		}

	}

}

func (v *VerificationTestSuite) TestValidateLogin() {
	testCases := []struct {
		email   string
		password string
		valid    bool
	}{
		{
			email:   "",
			password: "12345678",
			valid:    false,
		},
		{
			email:   "valid@valid.com",
			password: "",
			valid:    false,
		},
		{
			email:   "invalid",
			password: "12345678",
			valid:    false,
		},
		{
			email:   "valid@valid.com",
			password: "12345678",
			valid:    true,
		},
	}

	validator := validation.NewValidator()

	for _, tc := range testCases {
		err := validator.ValidateLoginInputs(tc.email, tc.password)

		if tc.valid {
			require.NoError(v.T(), err)
		} else {
			require.Error(v.T(), err)
		}

	}

}

func (v *VerificationTestSuite) TestValidateVerifyEmail() {
	testCases := []struct {
		userID           domain.UserID
		verificationCode string
		valid            bool
	}{
		{
			userID:           "",
			verificationCode: "123456",
			valid:            false,
		},
		{
			userID:           "1234",
			verificationCode: "",
			valid:            false,
		},
		{
			userID:           "12344566",
			verificationCode: "1234",
			valid:            false,
		},
		{
			userID:           "123456789",
			verificationCode: "123456",
			valid:            true,
		},
	}

	validator := validation.NewValidator()

	for _, tc := range testCases {
		err := validator.ValidateVerifyEmailInputs(tc.userID, tc.verificationCode)

		if tc.valid {
			require.NoError(v.T(), err)
		} else {
			require.Error(v.T(), err)
		}

	}

}

func (v *VerificationTestSuite) TestValidateAuthenticate() {
	testCases := []struct {
		accessToken string
		valid       bool
	}{
		{
			accessToken: "",
			valid:       false,
		},
		{
			accessToken: "valid",
			valid:       true,
		},
	}

	validator := validation.NewValidator()

	for _, tc := range testCases {
		err := validator.ValidateAuthenticateInputs(tc.accessToken)

		if tc.valid {
			require.NoError(v.T(), err)
		} else {
			require.Error(v.T(), err)
		}

	}

}

func (v *VerificationTestSuite) TestValidateChangePassword() {
	testCases := []struct {
		userID      domain.UserID
		oldPassword string
		newPassword string
		valid       bool
	}{
		{
			userID:      "12345",
			oldPassword: "12345678",
			newPassword: "",
			valid:       false,
		},
		{
			userID:      "1234",
			oldPassword: "",
			newPassword: "12345678",
			valid:       false,
		},
		{
			userID:      "",
			oldPassword: "23456789",
			newPassword: "12345678",
			valid:       false,
		},
		{
			userID:      "12345",
			oldPassword: "23456789",
			newPassword: "1234567",
			valid:       false,
		},
		{
			userID:      "12345",
			oldPassword: "23456789",
			newPassword: "12345678",
			valid:       true,
		},
	}

	validator := validation.NewValidator()

	for _, tc := range testCases {
		err := validator.ValidateChangePasswordInputs(tc.userID, tc.oldPassword, tc.newPassword)

		if tc.valid {
			require.NoError(v.T(), err)
		} else {
			require.Error(v.T(), err)
		}

	}

}

func (v *VerificationTestSuite) TestValidateRefreshToken() {
	testCases := []struct {
		userID       domain.UserID
		refreshToken string
		valid        bool
	}{
		{
			userID:       "",
			refreshToken: "valid",
			valid:        false,
		},
		{
			userID:       "1234",
			refreshToken: "",
			valid:        false,
		},
		{
			userID:       "123456789",
			refreshToken: "valid",
			valid:        true,
		},
	}

	validator := validation.NewValidator()

	for _, tc := range testCases {
		err := validator.ValidateRefreshTokenInputs(tc.userID, tc.refreshToken)

		if tc.valid {
			require.NoError(v.T(), err)
		} else {
			require.Error(v.T(), err)
		}

	}

}

func (v *VerificationTestSuite) TestValidateResetPassword() {
	testCases := []struct {
		email string
		valid  bool
	}{
		{
			email: "",
			valid:  false,
		},
		{
			email: "invalid",
			valid: false,
		},
		{
			email: "valid@valid.com",
			valid:  true,
		},
	}

	validator := validation.NewValidator()

	for _, tc := range testCases {
		err := validator.ValidateResetPasswordInputs(tc.email)

		if tc.valid {
			require.NoError(v.T(), err)
		} else {
			require.Error(v.T(), err)
		}

	}

}

func (v *VerificationTestSuite) TestValidateSubmitResetPassword() {
	testCases := []struct {
		resetPasswordToken string
		newPassword        string
		valid              bool
	}{
		{
			resetPasswordToken: "valid",
			newPassword:        "",
			valid:              false,
		},
		{
			resetPasswordToken: "",
			newPassword:        "12345678",
			valid:              false,
		},
		{
			resetPasswordToken: "7412",
			newPassword:        "2345678",
			valid:              false,
		},
		{
			resetPasswordToken: "7412",
			newPassword:        "23456788",
			valid:              true,
		},
	}

	validator := validation.NewValidator()

	for _, tc := range testCases {
		err := validator.ValidateSubmitResetPasswordInputs(tc.resetPasswordToken, tc.newPassword)

		if tc.valid {
			require.NoError(v.T(), err)
		} else {
			require.Error(v.T(), err)
		}

	}

}

func (v *VerificationTestSuite) TestValidateDeleteAccount() {
	testCases := []struct {
		userID   string
		password string
		valid    bool
	}{
		{
			userID:   "",
			password: "12345678",
			valid:    false,
		},
		{
			userID:   "1234",
			password: "",
			valid:    false,
		},
		{
			userID:   "123456789",
			password: "12345678",
			valid:    true,
		},
	}

	validator := validation.NewValidator()

	for _, tc := range testCases {
		err := validator.ValidateDeleteAccountInputs(tc.userID, tc.password)

		if tc.valid {
			require.NoError(v.T(), err)
		} else {
			require.Error(v.T(), err)
		}

	}

}
