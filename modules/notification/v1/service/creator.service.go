package service

import (
	"context"
	userv1 "grpc-starter/api/user/v1"
	"grpc-starter/common/config"
	"grpc-starter/modules/notification/v1/entity"
	"grpc-starter/modules/notification/v1/internal/repository"
)

type CreatorService struct {
	cfg            config.Config
	mailRepository repository.EmailSentUseCase
	userService    userv1.UserServiceClient
}

type CreatorServiceUseCase interface {
	InsertEmailSent(ctx context.Context, email *entity.EmailSent) error

	UpdateEmailSent(ctx context.Context, mId, from, to, subject, content, status, category string) error
}

func NewCreatorService(cfg config.Config, mailRepository repository.EmailSentUseCase) *CreatorService {
	return &CreatorService{
		cfg:            cfg,
		mailRepository: mailRepository,
	}
}

func (cs *CreatorService) InsertEmailSent(ctx context.Context, email *entity.EmailSent) error {

	email.Status = "sent"

	return cs.mailRepository.Insert(ctx, email)
}

func (cs *CreatorService) UpdateEmailSent(ctx context.Context, mId, from, to, subject, content, status, category string) error {
	emailSent := entity.NewEmailSent(mId, from, to, subject, content, status, category, "system")
	return cs.mailRepository.UpdateStatus(ctx, emailSent)
}
