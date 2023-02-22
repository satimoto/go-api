package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"
	"log"
	"net/url"
	"os"
	"strings"
	"time"

	"github.com/google/uuid"
	"github.com/satimoto/go-api/graph"
	"github.com/satimoto/go-api/internal/template"
	"github.com/satimoto/go-api/internal/util"
	"github.com/satimoto/go-datastore/pkg/db"
	"github.com/vektah/gqlparser/v2/gqlerror"
)

// CreateEmailSubscription is the resolver for the createEmailSubscription field.
func (r *mutationResolver) CreateEmailSubscription(ctx context.Context, input graph.CreateEmailSubscriptionInput) (*db.EmailSubscription, error) {
	backgroundCtx := context.Background()
	emailSubscription, err := r.EmailSubscriptionRepository.CreateEmailSubscription(backgroundCtx, db.CreateEmailSubscriptionParams{
		Email:            strings.ToLower(input.Email),
		Locale:           util.DefaultString(input.Locale, "en"),
		VerificationCode: util.RandomVerificationCode(),
		UnsubscribeCode:  uuid.NewString(),
		IsVerified:       false,
		CreatedDate:      time.Now(),
	})

	if err != nil {
		return nil, gqlerror.Errorf("Email subscription already exists")
	}

	params := url.Values{}
	params.Add("email", emailSubscription.Email)
	params.Add("code", emailSubscription.VerificationCode)

	html, subject, err := template.ParseEmailTemplateWithLocale("verify-email", emailSubscription.Locale, template.VerifyEmailTemplateData{
		Url: fmt.Sprintf("%s%s?%s", os.Getenv("WEB_DOMAIN"), util.URLLocale("/verify", emailSubscription.Locale, "en"), params.Encode()),
	})

	if err != nil {
		log.Print(err.Error())
		return nil, gqlerror.Errorf("Error creating verification email")
	}

	sendEmailInput := r.Emailer.Build(emailSubscription.Email, subject, html)
	_, err = r.Emailer.Send(sendEmailInput)

	if err != nil {
		log.Print(err.Error())
		return nil, gqlerror.Errorf("Error sending verification email")
	}

	return &emailSubscription, nil
}

// VerifyEmailSubscription is the resolver for the verifyEmailSubscription field.
func (r *mutationResolver) VerifyEmailSubscription(ctx context.Context, input graph.VerifyEmailSubscriptionInput) (*db.EmailSubscription, error) {
	backgroundCtx := context.Background()
	emailSubscription, err := r.EmailSubscriptionRepository.GetEmailSubscriptionByEmail(backgroundCtx, strings.ToLower(input.Email))

	if err != nil {
		return nil, gqlerror.Errorf("Email subscription not found")
	}

	if emailSubscription.VerificationCode == input.VerificationCode {
		emailSubscription, err = r.EmailSubscriptionRepository.UpdateEmailSubscription(backgroundCtx, db.UpdateEmailSubscriptionParams{
			ID:               emailSubscription.ID,
			Email:            emailSubscription.Email,
			Locale:           emailSubscription.Locale,
			VerificationCode: emailSubscription.VerificationCode,
			IsVerified:       true,
		})

		if err != nil {
			return nil, gqlerror.Errorf("Error updating email subscription")
		}

		return &emailSubscription, nil
	}

	return nil, gqlerror.Errorf("Invalid verification code")
}
