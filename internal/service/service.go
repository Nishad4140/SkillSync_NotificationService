package service

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/Nishad4140/SkillSync_NotificationService/internal/adapters"
	"github.com/Nishad4140/SkillSync_NotificationService/internal/helper"
	"github.com/Nishad4140/SkillSync_NotificationService/internal/usecase"
	"github.com/Nishad4140/SkillSync_ProtoFiles/pb"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"google.golang.org/protobuf/types/known/emptypb"
)

type NotificationService struct {
	adapters adapters.AdapterInterface
	usecase  usecase.UsecaseInterface
	pb.UnimplementedNotificationServiceServer
}

func NewNotificationService(adapter adapters.AdapterInterface, usecase usecase.UsecaseInterface) *NotificationService {
	return &NotificationService{
		adapters: adapter,
		usecase:  usecase,
	}
}

func (email *NotificationService) SendOTP(ctx context.Context, req *pb.SendOTPRequest) (*emptypb.Empty, error) {
	helper.SendOTP(req.Email)
	return &emptypb.Empty{}, nil
}

func (email *NotificationService) VerifyOTP(ctx context.Context, req *pb.VerifyOTPRequest) (*pb.VerifyOTPResponse, error) {
	verified := helper.VerifyOTP(req.Email, req.Otp)
	res := &pb.VerifyOTPResponse{
		IsVerified: verified,
	}
	return res, nil
}

func (email *NotificationService) AddNotification(ctx context.Context, req *pb.AddNotificationRequest) (*emptypb.Empty, error) {
	if req.UserId == "" {
		return nil, fmt.Errorf("please provide valid user id")
	}
	var message primitive.M
	if err := json.Unmarshal([]byte(req.Notification), &message); err != nil {
		return nil, fmt.Errorf("failed to parse message JSON: %v", err)
	}
	if err := email.adapters.AddNotification(req.UserId, message); err != nil {
		return nil, err
	}
	return nil, nil
}

func (email *NotificationService) GetAllNotification(req *pb.GetNotificationsByUserId, srv pb.NotificationService_GetAllNotificationServer) error {
	notifications, err := email.adapters.GetAllNotifications(req.UserId)
	if err != nil {
		return err
	}

	for _, notification := range notifications {
		notification, ok := notification["message"].(string)
		if !ok {
			return fmt.Errorf("notification field is not a string in notification: %v", notification)
		}
		res := &pb.NotificationResponse{
			Notification: notification,
		}
		if err := srv.Send(res); err != nil {
			return err
		}
	}
	return nil
}
