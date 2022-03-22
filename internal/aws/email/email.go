package email

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ses"
)

type (
	Emailer interface {
		Build(to, subject, html string) *ses.SendEmailInput
		Send(input *ses.SendEmailInput) (*ses.SendEmailOutput, error)
	}

	emailService struct {
		client *ses.SES
		from   string
	}
)

func New(from string) Emailer {
	sess, err := session.NewSession(&aws.Config{
        Region: aws.String("eu-central-1")},
    )

	if err != nil {
		return &emailService{}
	}

	return &emailService{
		client: ses.New(sess),
		from:   from,
	}
}

func (s *emailService) Build(to, subject, html string) *ses.SendEmailInput {
	charSet := "UTF-8"

	return &ses.SendEmailInput{
		Destination: &ses.Destination{
			CcAddresses: []*string{},
			ToAddresses: []*string{
				aws.String(to),
			},
		},
		Message: &ses.Message{
			Body: &ses.Body{
				Html: &ses.Content{
					Charset: aws.String(charSet),
					Data:    aws.String(html),
				},
			},
			Subject: &ses.Content{
				Charset: aws.String(charSet),
				Data:    aws.String(subject),
			},
		},
		Source: aws.String(s.from),
	}
}

func (s *emailService) Send(input *ses.SendEmailInput) (*ses.SendEmailOutput, error) {
	return s.client.SendEmail(input)
}
