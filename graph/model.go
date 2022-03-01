// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package graph

import (
	"fmt"
	"io"
	"strconv"
)

type CreateAuthentication struct {
	Code  string `json:"code"`
	LnURL string `json:"lnUrl"`
}

type CreateChannelRequestInput struct {
	// This field must be encoded as base64.
	PaymentHash string `json:"paymentHash"`
	// This field must be encoded as base64.
	PaymentAddr string `json:"paymentAddr"`
	AmountMsat  string `json:"amountMsat"`
}

type CreateEmailSubscriptionInput struct {
	Email  string  `json:"email"`
	Locale *string `json:"locale"`
}

type CreateUserInput struct {
	Code        string `json:"code"`
	Pubkey      string `json:"pubkey"`
	DeviceToken string `json:"deviceToken"`
}

type ExchangeAuthentication struct {
	Token string `json:"token"`
}

type UpdateUserInput struct {
	DeviceToken string `json:"deviceToken"`
}

type VerifyAuthentication struct {
	Verified bool `json:"verified"`
}

type VerifyEmailSubscriptionInput struct {
	Email            string `json:"email"`
	VerificationCode string `json:"verificationCode"`
}

type AuthenticationAction string

const (
	AuthenticationActionRegister AuthenticationAction = "register"
	AuthenticationActionLogin    AuthenticationAction = "login"
	AuthenticationActionLink     AuthenticationAction = "link"
	AuthenticationActionAuth     AuthenticationAction = "auth"
)

var AllAuthenticationAction = []AuthenticationAction{
	AuthenticationActionRegister,
	AuthenticationActionLogin,
	AuthenticationActionLink,
	AuthenticationActionAuth,
}

func (e AuthenticationAction) IsValid() bool {
	switch e {
	case AuthenticationActionRegister, AuthenticationActionLogin, AuthenticationActionLink, AuthenticationActionAuth:
		return true
	}
	return false
}

func (e AuthenticationAction) String() string {
	return string(e)
}

func (e *AuthenticationAction) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = AuthenticationAction(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid AuthenticationAction", str)
	}
	return nil
}

func (e AuthenticationAction) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}
