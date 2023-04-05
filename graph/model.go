// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package graph

import (
	"fmt"
	"io"
	"strconv"

	"github.com/satimoto/go-datastore/pkg/db"
	"github.com/satimoto/go-datastore/pkg/geom"
)

type AddtionalGeoLocation struct {
	Latitude  float64         `json:"latitude"`
	Longitude float64         `json:"longitude"`
	Name      *db.DisplayText `json:"name"`
}

type Channel struct {
	ChannelID string `json:"channelId"`
}

type CreateAuthentication struct {
	Code  string `json:"code"`
	LnURL string `json:"lnUrl"`
}

type CreateBusinessDetailInput struct {
	Name    string            `json:"name"`
	Website *string           `json:"website"`
	Logo    *CreateImageInput `json:"logo"`
}

type CreateChannelRequestInput struct {
	// This field must be encoded as base64.
	PaymentHash string `json:"paymentHash"`
	// This field must be encoded as base64.
	PaymentAddr string `json:"paymentAddr"`
	AmountMsat  string `json:"amountMsat"`
}

type CreateCredentialInput struct {
	ClientToken    *string                    `json:"clientToken"`
	URL            string                     `json:"url"`
	BusinessDetail *CreateBusinessDetailInput `json:"businessDetail"`
	CountryCode    string                     `json:"countryCode"`
	PartyID        string                     `json:"partyId"`
	IsHub          bool                       `json:"isHub"`
}

type CreateEmailSubscriptionInput struct {
	Email  string  `json:"email"`
	Locale *string `json:"locale"`
}

type CreateImageInput struct {
	URL       string  `json:"url"`
	Thumbnail *string `json:"thumbnail"`
	Category  string  `json:"category"`
	Type      string  `json:"type"`
	Width     *int    `json:"width"`
	Height    *int    `json:"height"`
}

type CreatePartyInput struct {
	CredentialID             int64  `json:"credentialId"`
	CountryCode              string `json:"countryCode"`
	PartyID                  string `json:"partyId"`
	IsIntermediateCdrCapable bool   `json:"isIntermediateCdrCapable"`
	PublishLocation          bool   `json:"publishLocation"`
	PublishNullTariff        bool   `json:"publishNullTariff"`
}

type CreateReferralInput struct {
	Code     string `json:"code"`
	Referrer string `json:"referrer"`
}

type CreateTokenInput struct {
	UID     string  `json:"uid"`
	Type    *string `json:"type"`
	Allowed *string `json:"allowed"`
}

type CreateUserInput struct {
	Code        string  `json:"code"`
	Pubkey      string  `json:"pubkey"`
	DeviceToken *string `json:"deviceToken"`
	Lsp         *bool   `json:"lsp"`
}

type ExchangeAuthentication struct {
	Token string `json:"token"`
}

type Geolocation struct {
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
	Name      *string `json:"name"`
}

type GetConnectorInput struct {
	ID         *int64  `json:"id"`
	Identifier *string `json:"identifier"`
}

type GetEvseInput struct {
	ID         *int64  `json:"id"`
	UID        *string `json:"uid"`
	EvseID     *string `json:"evseId"`
	Identifier *string `json:"identifier"`
}

type GetLocationInput struct {
	ID      *int64  `json:"id"`
	UID     *string `json:"uid"`
	Country *string `json:"country"`
}

type GetPoiInput struct {
	UID *string `json:"uid"`
}

type GetSessionInput struct {
	ID              *int64  `json:"id"`
	UID             *string `json:"uid"`
	AuthorizationID *string `json:"authorizationId"`
}

type GetTariffInput struct {
	ID  *int64  `json:"id"`
	UID *string `json:"uid"`
}

type ListLocation struct {
	UID                      string            `json:"uid"`
	Name                     string            `json:"name"`
	CountryCode              *string           `json:"countryCode"`
	PartyID                  *string           `json:"partyId"`
	Country                  string            `json:"country"`
	Geom                     geom.Geometry4326 `json:"geom"`
	AvailableEvses           int               `json:"availableEvses"`
	TotalEvses               int               `json:"totalEvses"`
	IsIntermediateCdrCapable bool              `json:"isIntermediateCdrCapable"`
	IsPublished              bool              `json:"isPublished"`
	IsRemoteCapable          bool              `json:"isRemoteCapable"`
	IsRfidCapable            bool              `json:"isRfidCapable"`
	AddedDate                string            `json:"addedDate"`
}

type ListLocationsInput struct {
	Country         *string  `json:"country"`
	Interval        *int     `json:"interval"`
	IsExperimental  *bool    `json:"isExperimental"`
	IsRemoteCapable *bool    `json:"isRemoteCapable"`
	IsRfidCapable   *bool    `json:"isRfidCapable"`
	Limit           *int     `json:"limit"`
	XMin            *float64 `json:"xMin"`
	XMax            *float64 `json:"xMax"`
	YMin            *float64 `json:"yMin"`
	YMax            *float64 `json:"yMax"`
}

type ListPoisInput struct {
	XMin *float64 `json:"xMin"`
	XMax *float64 `json:"xMax"`
	YMin *float64 `json:"yMin"`
	YMax *float64 `json:"yMax"`
}

type ListSessionInvoicesInput struct {
	IsSettled *bool `json:"isSettled"`
	IsExpired *bool `json:"isExpired"`
}

type PongUserInput struct {
	Pong string `json:"pong"`
}

type PublishLocationInput struct {
	ID           *int64  `json:"id"`
	CredentialID *int64  `json:"credentialId"`
	CountryCode  *string `json:"countryCode"`
	PartyID      *string `json:"partyId"`
	IsPublished  bool    `json:"isPublished"`
}

type Rate struct {
	Rate        string `json:"rate"`
	RateMsat    string `json:"rateMsat"`
	LastUpdated string `json:"lastUpdated"`
}

type RegisterCredentialInput struct {
	ID          int64   `json:"id"`
	ClientToken *string `json:"clientToken"`
}

type ResultID struct {
	ID int64 `json:"id"`
}

type ResultOk struct {
	Ok bool `json:"ok"`
}

type StartSession struct {
	ID              int64   `json:"id"`
	Status          string  `json:"status"`
	AuthorizationID string  `json:"authorizationId"`
	VerificationKey *string `json:"verificationKey"`
	LocationUID     string  `json:"locationUid"`
	EvseUID         *string `json:"evseUid"`
}

type StartSessionInput struct {
	LocationUID string  `json:"locationUid"`
	EvseUID     *string `json:"evseUid"`
}

type StopSession struct {
	ID         int64  `json:"id"`
	Status     string `json:"status"`
	SessionUID string `json:"sessionUid"`
}

type StopSessionInput struct {
	AuthorizationID string `json:"authorizationId"`
}

type SyncCredentialInput struct {
	ID          int64   `json:"id"`
	FromDate    *string `json:"fromDate"`
	CountryCode *string `json:"countryCode"`
	PartyID     *string `json:"partyId"`
	WithTariffs *bool   `json:"withTariffs"`
}

type TariffElement struct {
	PriceComponents []db.PriceComponent    `json:"priceComponents"`
	Restrictions    *db.ElementRestriction `json:"restrictions"`
}

type TextDescription struct {
	Text        string `json:"text"`
	Description string `json:"description"`
}

type UnregisterCredentialInput struct {
	ID int64 `json:"id"`
}

type UpdateInvoiceRequestInput struct {
	ID             int64  `json:"id"`
	PaymentRequest string `json:"paymentRequest"`
}

type UpdatePartyInput struct {
	CredentialID             int64  `json:"credentialId"`
	CountryCode              string `json:"countryCode"`
	PartyID                  string `json:"partyId"`
	IsIntermediateCdrCapable bool   `json:"isIntermediateCdrCapable"`
	PublishLocation          bool   `json:"publishLocation"`
	PublishNullTariff        bool   `json:"publishNullTariff"`
}

type UpdateSessionInput struct {
	UID         string `json:"uid"`
	IsConfirmed bool   `json:"isConfirmed"`
}

type UpdateTokenAuthorizationInput struct {
	AuthorizationID string `json:"authorizationId"`
	Authorized      bool   `json:"authorized"`
}

type UpdateTokensInput struct {
	UserID  int64   `json:"userId"`
	UID     *string `json:"uid"`
	Allowed string  `json:"allowed"`
}

type UpdateUserInput struct {
	DeviceToken     *string  `json:"deviceToken"`
	Name            *string  `json:"name"`
	Address         *string  `json:"address"`
	PostalCode      *string  `json:"postalCode"`
	City            *string  `json:"city"`
	BatteryCapacity *float64 `json:"batteryCapacity"`
	BatteryPowerAc  *float64 `json:"batteryPowerAc"`
	BatteryPowerDc  *float64 `json:"batteryPowerDc"`
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
