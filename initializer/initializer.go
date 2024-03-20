package initializer

import (
	"github.com/Nishad4140/SkillSync_NotificationService/internal/adapters"
	"github.com/Nishad4140/SkillSync_NotificationService/internal/service"
	"github.com/Nishad4140/SkillSync_NotificationService/internal/usecase"
	"gorm.io/gorm"
)

func Initializer(db *gorm.DB) *service.NotificationService {
	adapter := adapters.NewNotificationAdapter(db)
	usecase := usecase.NewNotificationUsecase(adapter)
	service := service.NewNotificationService(adapter, usecase)
	return service
}
