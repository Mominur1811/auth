package email

import (
	"auth-repo/authentication"
	"auth-repo/config"
	"fmt"
)

type EmailRepo interface {
	authentication.EmailRepo
}

type emailRepo struct {
	mailConf *config.Email
}

func NewEmailService(mailConf *config.Email) EmailRepo {
	fmt.Println(mailConf)
	return &emailRepo{
		mailConf: mailConf,
	}
}
