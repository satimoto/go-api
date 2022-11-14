package notification

import (
	"log"

	"github.com/appleboy/go-fcm"
	"github.com/satimoto/go-datastore/pkg/util"
)

type Notification interface {
	SendNotification(*fcm.Message) (*fcm.Response, error)
	SendNotificationWithRetry(message *fcm.Message, retries int) (*fcm.Response, error)
}

type NotificationService struct {
	client *fcm.Client
}

func NewService(apiKey string) Notification {
	client, err := fcm.NewClient(apiKey)
	util.PanicOnError("API047", "Invalid FCM API key", err)

	return &NotificationService{
		client: client,
	}
}

func (s *NotificationService) SendNotification(message *fcm.Message) (*fcm.Response, error) {
	log.Printf("Sending notification: %v", message.To)
	log.Printf("Data=%#v", message.Data)
	return s.client.Send(message)
}

func (s *NotificationService) SendNotificationWithRetry(message *fcm.Message, retries int) (*fcm.Response, error) {
	log.Printf("Sending notification with retry: %v", message.To)
	log.Printf("Data=%#v", message.Data)
	return s.client.SendWithRetry(message, retries)
}
