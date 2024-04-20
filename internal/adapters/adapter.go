package adapters

import (
	"context"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type NotificationAdapter struct {
	DB *mongo.Database
}

func NewNotificationAdapter(db *mongo.Database) *NotificationAdapter {
	return &NotificationAdapter{
		DB: db,
	}
}

func (email *NotificationAdapter) AddNotification(userId string, notificationData bson.M) error {
	collection := email.DB.Collection("notifications")
	if collection == nil {
		err := email.DB.CreateCollection(context.Background(), "notifications")
		if err != nil {
			return err
		}
	}

	notificationData["userId"] = userId
	notificationData["timestamp"] = time.Now()

	_, err := collection.InsertOne(context.Background(), notificationData)
	return err
}

func (email *NotificationAdapter) GetAllNotifications(userId string) ([]bson.M, error) {
	collection := email.DB.Collection("notifications")
	if collection == nil {
		return nil, fmt.Errorf("collection not found")
	}
	options := options.Find().SetSort(bson.D{{"timestamp", -1}})
	filter := bson.M{"userId": userId}
	cursor, err := collection.Find(context.Background(), filter, options)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(context.Background())
	var notifications []bson.M
	for cursor.Next(context.Background()) {
		var notification bson.M
		err = cursor.Decode(&notification)
		if err != nil {
			return nil, err
		}
		notifications = append(notifications, notification)
	}

	return notifications, nil
}
