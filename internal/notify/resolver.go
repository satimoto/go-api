package notify

import (
	"os"

	"github.com/satimoto/go-api/internal/notification"
	"github.com/satimoto/go-datastore/pkg/db"
)

type NotifyRepository interface{}

type NotifyResolver struct {
	Repository          NotifyRepository
	NotificationService notification.Notification
}

func NewResolver(repositoryService *db.RepositoryService) *NotifyResolver {
	notificationService := notification.NewService(os.Getenv("FCM_API_KEY"))

	return NewResolverWithServices(repositoryService, notificationService)
}

func NewResolverWithServices(repositoryService *db.RepositoryService, notificationService notification.Notification) *NotifyResolver {
	repo := NotifyRepository(repositoryService)

	return &NotifyResolver{
		Repository:          repo,
		NotificationService: notificationService,
	}
}
