package service

import (
	"context"

	"github.com/Nishad4140/SkillSync_NotificationService/internal/adapters"
	"github.com/Nishad4140/SkillSync_NotificationService/internal/helper"
	"github.com/Nishad4140/SkillSync_NotificationService/internal/usecase"
	"github.com/Nishad4140/SkillSync_ProtoFiles/pb"
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
