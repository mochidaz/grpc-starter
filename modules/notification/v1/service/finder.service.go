package service

import (
	"context"
	"github.com/google/uuid"
	userv1 "grpc-starter/api/user/v1"
	"grpc-starter/common/config"
	"grpc-starter/modules/notification/v1/entity"
	"grpc-starter/modules/notification/v1/internal/repository"
)

type EmailSentResponse struct {
	entity.EmailSent
	Username string
}

type EmailFinderService struct {
	cfg             config.Config
	emailRepository repository.EmailFinderUseCase
	userService     userv1.UserServiceClient
}

//type EmailFinderUseCase interface {
//	FindByID(ctx context.Context, id int) (*entity.EmailSent, error)
//
//	FindByUserID(ctx context.Context, userID uuid.UUID) ([]*entity.EmailSent, error)
//
//	FindByUsername(ctx context.Context, username string) ([]*entity.EmailSent, error)
//
//	FindByEmail(ctx context.Context, email string) ([]*entity.EmailSent, error)
//
//	GetAllEmailSent(ctx context.Context) ([]*EmailSentResponse, error)
//}

type EmailFinderUseCase interface {
	FindByID(ctx context.Context, id int) (*entity.EmailWithUsername, error)

	FindByUserID(ctx context.Context, userID uuid.UUID) ([]*entity.EmailWithUsername, error)

	FindByEmail(ctx context.Context, email string) ([]*entity.EmailWithUsername, error)

	GetAllEmailSent(ctx context.Context) ([]*EmailSentResponse, error)
}

func NewEmailFinderService(cfg config.Config, emailRepository repository.EmailFinderUseCase, userService userv1.UserServiceClient) *EmailFinderService {
	return &EmailFinderService{cfg: cfg, emailRepository: emailRepository, userService: userService}
}

func (efs *EmailFinderService) FindByID(ctx context.Context, id int) (*entity.EmailWithUsername, error) {

	emailSent, err := efs.emailRepository.FindByID(ctx, id)

	if err != nil {
		return nil, err
	}

	user, err := efs.userService.GetProfile(ctx, &userv1.GetProfileRequest{Id: emailSent.MId})

	if err != nil {
		return nil, err
	}

	newEmail := &entity.EmailWithUsername{
		EmailSent: *emailSent,
		Username:  user.Data.Username,
	}

	if err != nil {
		return nil, err
	}

	return newEmail, nil

}

func (efs *EmailFinderService) FindByUserID(ctx context.Context, userID uuid.UUID) ([]*entity.EmailWithUsername, error) {

	var response []*entity.EmailWithUsername

	emailSent, err := efs.emailRepository.FindByUserID(ctx, userID)

	user, err := efs.userService.GetProfile(ctx, &userv1.GetProfileRequest{Id: emailSent[0].MId})

	if err != nil {
		return nil, err
	}

	for _, resp := range emailSent {
		response = append(response, &entity.EmailWithUsername{
			EmailSent: *resp,
			Username:  user.Data.Username,
		})
	}

	return response, nil
}

func (efs *EmailFinderService) FindByEmail(ctx context.Context, email string) ([]*entity.EmailWithUsername, error) {

	var response []*entity.EmailWithUsername

	emailSent, err := efs.emailRepository.FindByEmail(ctx, email)

	if err != nil {
		return nil, err
	}

	user, err := efs.userService.GetProfile(ctx, &userv1.GetProfileRequest{Id: emailSent[0].MId})

	if err != nil {
		return nil, err
	}

	for _, resp := range emailSent {
		response = append(response, &entity.EmailWithUsername{
			EmailSent: *resp,
			Username:  user.Data.Username,
		})
	}

	return response, nil
}

func (efs *EmailFinderService) GetAllEmailSent(ctx context.Context) ([]*EmailSentResponse, error) {

	var response []*EmailSentResponse

	users, err := efs.userService.GetAllUsers(ctx, &userv1.GetAllUsersRequest{})

	hashmap := make(map[string]string)

	if err != nil {
		return nil, err
	}

	for _, user := range users.Data {
		hashmap[user.GetUserId()] = user.Username
	}

	emailSent, err := efs.emailRepository.GetAllEmailSent(ctx)

	if err != nil {
		return nil, err
	}

	for _, resp := range emailSent {
		response = append(response, &EmailSentResponse{
			EmailSent: *resp,
			Username:  hashmap[resp.MId],
		})
	}

	return response, nil
}
