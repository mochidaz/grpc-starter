package handler

import (
	"context"
	"github.com/google/uuid"
	notificationv1 "grpc-starter/api/notification/v1"
	"strconv"
)

func (h *NotificationHandler) GetNotificationByID(ctx context.Context, req *notificationv1.GetNotificationByIDRequest) (*notificationv1.GetNotificationResponse, error) {

	toInt, err := strconv.Atoi(req.GetId())

	if err != nil {
		return nil, err
	}

	notification, err := h.notificationFinderService.FindByID(ctx, toInt)

	if err != nil {
		return nil, err
	}

	return &notificationv1.GetNotificationResponse{
		Notification: &notificationv1.Notification{
			Id:        strconv.Itoa(int(notification.ID)),
			MId:       notification.MId,
			From:      notification.From,
			To:        notification.To,
			Title:     notification.Subject,
			Body:      notification.Content,
			Status:    notification.Status,
			Category:  notification.Category,
			Username:  notification.Username,
			CreatedAt: notification.CreatedAt.String(),
			UpdatedAt: notification.UpdatedAt.String(),
		},
	}, nil
}

func (h *NotificationHandler) GetNotificationByEmail(ctx context.Context, req *notificationv1.GetNotificationByEmailRequest) (*notificationv1.ListNotificationResponse, error) {

	notifications, err := h.notificationFinderService.FindByEmail(ctx, req.GetEmail())

	var notificationList []*notificationv1.Notification

	for _, notification := range notifications {
		notificationList = append(notificationList, &notificationv1.Notification{
			Id:        strconv.Itoa(int(notification.ID)),
			MId:       notification.MId,
			From:      notification.From,
			To:        notification.To,
			Title:     notification.Subject,
			Body:      notification.Content,
			Status:    notification.Status,
			Category:  notification.Category,
			Username:  notification.Username,
			CreatedAt: notification.CreatedAt.String(),
			UpdatedAt: notification.UpdatedAt.String(),
		})
	}

	if err != nil {
		return nil, err
	}

	return &notificationv1.ListNotificationResponse{
		Notification: notificationList,
	}, nil
}

func (h *NotificationHandler) GetNotificationByMID(ctx context.Context, req *notificationv1.GetNotificationByMIDRequest) (*notificationv1.ListNotificationResponse, error) {

	parse, err := uuid.Parse(req.GetMId())

	if err != nil {
		return nil, err
	}

	notifications, err := h.notificationFinderService.FindByUserID(ctx, parse)

	var notificationList []*notificationv1.Notification

	for _, notification := range notifications {
		notificationList = append(notificationList, &notificationv1.Notification{
			Id:        strconv.Itoa(int(notification.ID)),
			MId:       notification.MId,
			From:      notification.From,
			To:        notification.To,
			Title:     notification.Subject,
			Body:      notification.Content,
			Status:    notification.Status,
			Category:  notification.Category,
			Username:  notification.Username,
			CreatedAt: notification.CreatedAt.String(),
			UpdatedAt: notification.UpdatedAt.String(),
		})
	}

	if err != nil {
		return nil, err
	}

	return &notificationv1.ListNotificationResponse{
		Notification: notificationList,
	}, nil
}

func (h *NotificationHandler) ListNotifications(ctx context.Context, req *notificationv1.ListNotificationRequest) (*notificationv1.ListNotificationResponse, error) {

	notifications, err := h.notificationFinderService.GetAllEmailSent(ctx)

	var notificationList []*notificationv1.Notification

	for _, notification := range notifications {
		notificationList = append(notificationList, &notificationv1.Notification{
			Id:        strconv.Itoa(int(notification.ID)),
			MId:       notification.MId,
			From:      notification.From,
			To:        notification.To,
			Title:     notification.Subject,
			Body:      notification.Content,
			Status:    notification.Status,
			Category:  notification.Category,
			Username:  notification.Username,
			CreatedAt: notification.CreatedAt.String(),
			UpdatedAt: notification.UpdatedAt.String(),
		})
	}

	if err != nil {
		return nil, err
	}

	return &notificationv1.ListNotificationResponse{
		Notification: notificationList,
	}, nil
}
