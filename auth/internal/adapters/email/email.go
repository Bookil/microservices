package email

import (
	"log"

	"github.com/Bookil/microservices/auth/config"
)

type Adapter struct{}

func NewAdapter()*Adapter{
	return &Adapter{}
}

func (a *Adapter)SendVerificationEmail(email, verifyEmailRedirectUrl, verifyEmailToken string) error{
	if config.CurrentEnv == config.Development || config.CurrentEnv == config.Test{
		sendEmailTestingAndDevelopment()
		return nil
	}

	panic("unimplemented")
}



func sendEmailTestingAndDevelopment(){
	log.Println("Email sent")
}