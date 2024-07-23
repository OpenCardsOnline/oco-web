package services

import (
	"fmt"
	"log"

	"github.com/opencardsonline/oco-web/config"
	"github.com/sendgrid/sendgrid-go"
	"github.com/sendgrid/sendgrid-go/helpers/mail"
)

func SendEmail(fromUsername string, fromEmail string, subject string, toUsername string, toEmail string, htmlContent string) {
	from := mail.NewEmail(fromUsername, fromEmail)
	to := mail.NewEmail(toUsername, toEmail)
	plainTextContent := "Please Enable HTML To View This Email!"
	message := mail.NewSingleEmail(from, subject, to, plainTextContent, htmlContent)
	client := sendgrid.NewSendClient(config.AppConfiguration.EmailAPIKey)
	_, err := client.Send(message)
	if err != nil {
		log.Println(err)
	} else {
		fmt.Println("Email Sent!")
	}
}
