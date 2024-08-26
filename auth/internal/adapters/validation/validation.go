package validation

import (
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

func (v *Validator) ValidateRegisterInputs(UserID domain.UserID, password string) error {
	registerInputs := &registerInputs{
		UserID:   UserID,
		Password: password,
	}

	err := v.validate(registerInputs)

	return err
}

func (v *Validator) ValidateLoginInputs(UserID domain.UserID, password string) error {
	loginInputs := &loginInputs{
		UserID:   UserID,
		Password: password,
	}

	err := v.validate(loginInputs)

	return err
}

func (v *Validator) ValidateVerifyEmailInputs(UserID domain.UserID, verificationCode string) error {
	verifyEmailInputs := &verifyEmailInputs{
		UserID:         UserID,
		ValidationCode: verificationCode,
	}

	err := v.validate(verifyEmailInputs)

	return err
}

func (v *Validator) ValidateChangePasswordInputs(UserID domain.UserID, oldPassword string, newPassword string) error {
	changePasswordInputs := &changePasswordInputs{
		UserID:      UserID,
		OldPassword: oldPassword,
		NewPassword: newPassword,
	}

	err := v.validate(changePasswordInputs)

	return err
}

func (v *Validator) ValidateAuthenticateInputs(accessToken string) error {
	authenticateInputs := &authenticateInputs{
		AccessToken: accessToken,
	}

	err := v.validate(authenticateInputs)

	return err
}

func (v *Validator) ValidateRefreshTokenInputs(UserID domain.UserID, refreshToken string) error {
refreshTokenInputs := &refreshTokenInputs{
	UserID:       UserID,
	RefreshToken: refreshToken,
}

err := v.validate(refreshTokenInputs)

return err
}

func (v *Validator) ValidateResetPasswordInputs(UserID domain.UserID) error {
	resetPasswordInputs := &resetPasswordInputs{
		UserID: UserID,
	}

	err := v.validate(resetPasswordInputs)

	return err
}

func (v *Validator) ValidateSubmitResetPasswordInputs(resetPasswordToken string, newPassword string) error {
	submitResetPasswordInputs := &submitResetPasswordInputs{
		ResetPasswordToken: resetPasswordToken,
		Password:           newPassword,
	}

	err := v.validate(submitResetPasswordInputs)

	return err
}

func (v *Validator) ValidateDeleteAccountInputs(UserID domain.UserID, password string) error {
	deleteAccountInputs := &deleteAccountInputs{
		UserID:   UserID,
		Password: password,
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