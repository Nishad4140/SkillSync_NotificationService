package adapters

import (
	"go.mongodb.org/mongo-driver/mongo"
)

type NotificationAdapter struct {
	DB *mongo.Database
}

func NewNotificationAdapter(db *mongo.Database) *NotificationAdapter {
	return &NotificationAdapter{
		DB: db,
	}
}
