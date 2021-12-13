package ses

import (
	"envio-email/api/request"
	"envio-email/config"
	"fmt"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ses"
)

//	create a new aws session and returns the session var
func startNewSession() *session.Session {
	sess, err := session.NewSession()
	if err != nil {
		fmt.Println("failed to create session,", err)
	}
	return sess
}

//	aws configuration
func SetConfiguration() {
	os.Setenv("AWS_REGION", config.AWS_REGION)
	os.Setenv("AWS_ACCESS_KEY_ID", config.AWS_KEY_ID)
	os.Setenv("AWS_SECRET_ACCESS_KEY", config.AWS_SECRET_KEY)
}

//	create and send text or html email to single receipents.
func SendEmail(emailData request.Email) *ses.SendEmailOutput {
	// start a new aws session
	sess := startNewSession()
	// start a new ses session
	svc := ses.New(sess)

	params := &ses.SendEmailInput{
		Destination: &ses.Destination{ // Required
			ToAddresses: []*string{
				aws.String(emailData.To), // Required
				// More values...
			},
		},
		Message: &ses.Message{ // Required
			Body: &ses.Body{ // Required
				Text: &ses.Content{
					Data:    aws.String(emailData.Text), // Required
					Charset: aws.String("UTF-8"),
				},
			},
			Subject: &ses.Content{ // Required
				Data:    aws.String(emailData.Subject), // Required
				Charset: aws.String("UTF-8"),
			},
		},
		Source: aws.String(emailData.From), // Required

		ReplyToAddresses: []*string{
			aws.String(emailData.ReplyTo), // Required
			// More values...
		},
	}

	// send email
	resp, err := svc.SendEmail(params)

	if err != nil {
		fmt.Println(err.Error())
	}
	return resp
}
