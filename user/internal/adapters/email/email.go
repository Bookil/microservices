package email

import (
	"log"
	"time"

	"github.com/Bookil/microservices/user/config"
)

type Adapter struct{}

func NewEmailAdapter() *Adapter {
	return &Adapter{}
}

func (a *Adapter) SendResetPassword(url string, duration time.Duration, email string) error {
	if config.CurrentEnv == config.Development || config.CurrentEnv == config.Test {
		log.Println("email sent")
	}
	return nil
}

func (a *Adapter) SendVerificationCode(email, code string) error {
	if config.CurrentEnv == config.Development || config.CurrentEnv == config.Test {
		log.Println("email sent")
	}
	return nil
}

func (a *Adapter) SendWellCome(email string) error {
	if config.CurrentEnv == config.Development || config.CurrentEnv == config.Test {
		log.Println("email sent")
	}
	return nil
}
