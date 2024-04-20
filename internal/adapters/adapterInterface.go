package adapters

import "go.mongodb.org/mongo-driver/bson"

type AdapterInterface interface {
	AddNotification(userId string, notificationData bson.M) error
	GetAllNotifications(userId string) ([]bson.M, error)
}
