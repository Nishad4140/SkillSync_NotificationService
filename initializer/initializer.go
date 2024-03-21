package initializer

import (
	"github.com/Nishad4140/SkillSync_NotificationService/internal/adapters"
	"github.com/Nishad4140/SkillSync_NotificationService/internal/service"
	"github.com/Nishad4140/SkillSync_NotificationService/internal/usecase"
	"go.mongodb.org/mongo-driver/mongo"
)

func Initializer(db *mongo.Database) *service.NotificationService {
	adapter := adapters.NewNotificationAdapter(db)
	usecase := usecase.NewNotificationUsecase(adapter)
	service := service.NewNotificationService(adapter, usecase)
	return service
}
