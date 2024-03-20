package usecase

import "github.com/Nishad4140/SkillSync_NotificationService/internal/adapters"

type NotificationUsecase struct {
	notificationAdapter adapters.AdapterInterface
}

func NewNotificationUsecase(notificationAdapter adapters.AdapterInterface) *NotificationUsecase {
	return &NotificationUsecase{
		notificationAdapter: notificationAdapter,
	}
}
