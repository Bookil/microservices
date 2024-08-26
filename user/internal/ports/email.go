package ports

type EmailPort interface {
	SendVerificationCode(email,code string)error
	SendResetPassword(email string)error
	SendWellCome(email string)error
}