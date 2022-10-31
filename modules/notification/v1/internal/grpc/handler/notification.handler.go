package handler

import (
	notificationv1 "grpc-starter/api/notification/v1"
	"grpc-starter/modules/notification/v1/service"
)

type NotificationHandler struct {
	notificationv1.UnimplementedNotificationServiceServer

	notificationCreatorService service.CreatorServiceUseCase
	notificationFinderService  service.EmailFinderUseCase
}

func NewNotificationHandler(notificationCreatorService service.CreatorServiceUseCase, notificationFinderService service.EmailFinderUseCase) *NotificationHandler {
	return &NotificationHandler{notificationCreatorService: notificationCreatorService, notificationFinderService: notificationFinderService}
}
