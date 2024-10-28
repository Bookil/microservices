package validation

import (
	"strings"
	"sync"

	"github.com/Bookil/microservices/auth/internal/application/core/domain"
	"github.com/go-playground/locales/en"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	en_translations "github.com/go-playground/validator/v10/translations/en"
)

var (
	locale_EN         = en.New()
	uni               = ut.New(locale_EN, locale_EN)
	ValiTranslator, _ = uni.GetTranslator("en")
)

var (
	valiLock     = &sync.Mutex{}
	valiInstance *Validator
)

type Validator struct {
	vi *validator.Validate
}

func NewValidator() *Validator {
	if valiInstance == nil {
		valiLock.Lock()
		defer valiLock.Unlock()

		vi := validator.New()

		valiInstance = &Validator{vi}
		valiInstance.translateOverride()
	}
	return valiInstance
}

func (v *Validator) ValidateRegisterInputs(firstName, lastName, email, password string) error {
	registerInputs := &registerInputs{
		FirstName: strings.TrimSpace(firstName),
		LastName:  strings.TrimSpace(lastName),
		Email:     strings.TrimSpace(email),
		Password:  strings.TrimSpace(password),
	}

	err := v.validate(registerInputs)

	return err
}

func (v *Validator) ValidateLoginInputs(email, password string) error {
	loginInputs := &loginInputs{
		Email:    strings.TrimSpace(email),
		Password: strings.TrimSpace(password),
	}

	err := v.validate(loginInputs)

	return err
}

func (v *Validator) ValidateSendVerificationCode(email string) error {
	sendVerificationCodeAgainInputs := &sendVerificationCodeAgainInputs{
		Email: strings.TrimSpace(email),
	}

	err := v.validate(sendVerificationCodeAgainInputs)

	return err
}

func (v *Validator) ValidateVerifyEmailInputs(email, verificationCode string) error {
	verifyEmailInputs := &verifyEmailInputs{
		Email:          strings.TrimSpace(email),
		ValidationCode: strings.TrimSpace(verificationCode),
	}

	err := v.validate(verifyEmailInputs)

	return err
}

func (v *Validator) ValidateChangePasswordInputs(UserID domain.UserID, oldPassword string, newPassword string) error {
	changePasswordInputs := &changePasswordInputs{
		UserID:      strings.TrimSpace(UserID),
		OldPassword: strings.TrimSpace(oldPassword),
		NewPassword: strings.TrimSpace(newPassword),
	}

	err := v.validate(changePasswordInputs)

	return err
}

func (v *Validator) ValidateAuthenticateInputsAndAuthorization(accessToken string) error {
	authenticateInputs := &authenticateAndAuthorizationInputs{
		AccessToken: strings.TrimSpace(accessToken),
	}

	err := v.validate(authenticateInputs)

	return err
}

func (v *Validator) ValidateRefreshTokenInputs(UserID domain.UserID, refreshToken string) error {
	refreshTokenInputs := &refreshTokenInputs{
		UserID:       strings.TrimSpace(UserID),
		RefreshToken: strings.TrimSpace(refreshToken),
	}

	err := v.validate(refreshTokenInputs)

	return err
}

func (v *Validator) ValidateResetPasswordInputs(email string) error {
	resetPasswordInputs := &resetPasswordInputs{
		Email: strings.TrimSpace(email),
	}

	err := v.validate(resetPasswordInputs)

	return err
}

func (v *Validator) ValidateSubmitResetPasswordInputs(resetPasswordToken string, newPassword string) error {
	submitResetPasswordInputs := &submitResetPasswordInputs{
		ResetPasswordToken: strings.TrimSpace(resetPasswordToken),
		Password:           strings.TrimSpace(newPassword),
	}

	err := v.validate(submitResetPasswordInputs)

	return err
}

func (v *Validator) ValidateDeleteAccountInputs(UserID domain.UserID, password string) error {
	deleteAccountInputs := &deleteAccountInputs{
		UserID:   strings.TrimSpace(UserID),
		Password: strings.TrimSpace(password),
	}

	err := v.validate(deleteAccountInputs)

	return err
}

func (v *Validator) validate(s interface{}) error {
	err := v.vi.Struct(s)

	if err == nil {
		return nil
	}

	return err
}

func (v *Validator) translateOverride() {
	en_translations.RegisterDefaultTranslations(v.vi, ValiTranslator) // nolint
}
