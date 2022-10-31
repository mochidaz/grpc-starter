package v1

import (
	"google.golang.org/grpc"
	notificationv1 "grpc-starter/api/notification/v1"
	userv1 "grpc-starter/api/user/v1"

	"gorm.io/gorm"

	"grpc-starter/common/config"
	"grpc-starter/modules/notification/v1/internal/builder"
)

func InitNotification(server *grpc.Server, cfg config.Config, db *gorm.DB, grpcConn *grpc.ClientConn, client userv1.UserServiceClient) {
	notification := builder.BuildNotificationHandler(cfg, db, *grpcConn, client)

	notificationv1.RegisterNotificationServiceServer(server, notification)
}
