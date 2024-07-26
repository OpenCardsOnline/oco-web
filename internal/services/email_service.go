package services

import (
	"fmt"
	"log"

	"github.com/sendgrid/sendgrid-go"
	"github.com/sendgrid/sendgrid-go/helpers/mail"
)

type EmailService struct {
	apiKey string
}

func (_s *EmailService) SendEmail(fromUsername string, fromEmail string, subject string, toUsername string, toEmail string, htmlContent string) {
	from := mail.NewEmail(fromUsername, fromEmail)
	to := mail.NewEmail(toUsername, toEmail)
	plainTextContent := "Please Enable HTML To View This Email!"
	message := mail.NewSingleEmail(from, subject, to, plainTextContent, htmlContent)
	client := sendgrid.NewSendClient(_s.apiKey)
	_, err := client.Send(message)
	if err != nil {
		log.Println(err)
	} else {
		fmt.Println("Email Sent!")
	}
}
