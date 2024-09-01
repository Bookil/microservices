package validation_test

import (
	"log"
	"testing"

	"github.com/Bookil/microservices/user/internal/adapters/validation"
	"github.com/Bookil/microservices/user/internal/application/core/domain"
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
		lastName  string
		email     string
		password  string
		valid     bool
	}{
		{
			firstName: "",
			lastName:  "valid",
			email:     "valid@gmail.com",
			valid:     false,
		},
		{
			firstName: "valid",
			lastName:  "",
			email:     "valid@gmail.com",
			valid:     false,
		},
		{
			firstName: "valid",
			lastName:  "valid",
			email:     "",
			valid:     false,
		},
		{
			firstName: "valid",
			lastName:  "valid",
			email:     "valid@valid.com",
			valid:     true,
		},
	}

	validator := validation.NewValidator()

	for n, tc := range testCases {
		n++
		log.Printf("number %d,case:%v\n", n, tc)
		err := validator.ValidateRegisterInputs(tc.firstName, tc.lastName, tc.email)

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

	for n, tc := range testCases {
		n++
		log.Printf("number %d,case:%v\n", n, tc)

		err := validator.ValidateChangePasswordInputs(tc.userID, tc.oldPassword, tc.newPassword)

		if tc.valid {
			require.NoError(v.T(), err)
		} else {
			require.Error(v.T(), err)
		}

	}

}

func (v *VerificationTestSuite) TestValidateUpdate() {
	testCases := []struct {
		userID    domain.UserID
		firstName string
		lastName  string
		valid     bool
	}{
		{
			userID:    "",
			firstName: "valid",
			lastName:  "valid",
			valid:     false,
		},
		{
			userID:    "valid",
			firstName: "",
			lastName:  "valid",
			valid:     false,
		},
		{
			userID:    " ",
			firstName: "valid",
			lastName:  "valid",
			valid:     false,
		},
		{
			userID:    "valid",
			firstName: "valid",
			lastName:  "valid",
			valid:     true,
		},
	}

	validator := validation.NewValidator()

	for n, tc := range testCases {
		n++
		log.Printf("number %d,case:%v\n", n, tc)

		err := validator.ValidateUpdateInputs(tc.userID, tc.firstName, tc.lastName)

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

	for n, tc := range testCases {
		n++
		log.Printf("number %d,case:%v\n", n, tc)

		err := validator.ValidateDeleteAccountInputs(tc.userID, tc.password)

		if tc.valid {
			require.NoError(v.T(), err)
		} else {
			require.Error(v.T(), err)
		}

	}
}
