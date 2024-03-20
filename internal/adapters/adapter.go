package adapters

import "gorm.io/gorm"

type NotificationAdapter struct {
	DB *gorm.DB
}

func NewNotificationAdapter(db *gorm.DB) *NotificationAdapter {
	return &NotificationAdapter{
		DB: db,
	}
}
