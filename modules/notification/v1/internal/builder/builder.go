// Package builder is used to build the handler.
package builder

import (
	"google.golang.org/grpc"
	"gorm.io/gorm"
	userv1 "grpc-starter/api/user/v1"
	"grpc-starter/common/config"
	"grpc-starter/modules/notification/v1/internal/grpc/handler"
	"grpc-starter/modules/notification/v1/internal/repository"
	"grpc-starter/modules/notification/v1/service"
)

func BuildNotificationHandler(cfg config.Config, db *gorm.DB, grpcConn grpc.ClientConn, client userv1.UserServiceClient) *handler.NotificationHandler {
	emailFinderRepo := repository.NewEmailFinder(db)
	emailCreatorRepo := repository.NewEmailSent(db)

	emailFinderService := service.NewEmailFinderService(cfg, emailFinderRepo, client)
	emailCreatorService := service.NewCreatorService(cfg, emailCreatorRepo)

	return handler.NewNotificationHandler(emailCreatorService, emailFinderService)
}
