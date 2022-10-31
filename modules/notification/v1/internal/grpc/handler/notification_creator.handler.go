package handler

import (
	"context"
	notificationv1 "grpc-starter/api/notification/v1"
	"grpc-starter/modules/notification/v1/entity"
)

func (h *NotificationHandler) CreateNotification(ctx context.Context, req *notificationv1.CreateNotificationRequest) (*notificationv1.CreateNotificationResponse, error) {

	newEmail := entity.NewEmailSent(
		req.GetMId(),
		req.GetFrom(),
		req.GetTo(),
		req.GetTitle(),
		req.GetBody(),
		"not sent",
		req.GetCategory(),
		"system",
	)

	err := h.notificationCreatorService.InsertEmailSent(ctx, newEmail)

	if err != nil {
		return nil, err
	}

	return &notificationv1.CreateNotificationResponse{
		Title:     newEmail.Subject,
		Body:      newEmail.Content,
		To:        newEmail.To,
		From:      newEmail.From,
		Category:  newEmail.Category,
		CreatedAt: newEmail.CreatedAt.String(),
		UpdatedAt: newEmail.UpdatedAt.String(),
	}, nil
}
