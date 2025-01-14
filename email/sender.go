package email

import (
	"fmt"
	"net/smtp"
)

func (svc *emailRepo) Send(subject, body, recipient string) error {
	fmt.Println(svc.mailConf)
	message := []byte("Subject: " + subject + "\r\n" +
		"Content-Type: text/plain; charset=\"utf-8\"\r\n" +
		"\r\n" +
		body)

	fmt.Println(message)
	auth := smtp.PlainAuth("", svc.mailConf.SourceMail, svc.mailConf.AppPass, svc.mailConf.SmtpHost)
	fmt.Println(auth)

	to := []string{recipient}
	fmt.Println(to)
	err := smtp.SendMail(svc.mailConf.SmtpHost+":"+svc.mailConf.SmtpPort, auth, svc.mailConf.SourceMail, to, message)
	if err != nil {
		fmt.Println(err)
		return err
	}
	fmt.Println("Email Sent Successfully!")
	return nil
}
