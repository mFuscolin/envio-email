package service

import (
	"envio-email/api/request"
	"envio-email/infrastructure/aws/ses"
	"fmt"
)

func SendEmail(emailData request.Email) {
	ses.SetConfiguration()

	resp := ses.SendEmail(emailData)

	fmt.Println(resp)
}
