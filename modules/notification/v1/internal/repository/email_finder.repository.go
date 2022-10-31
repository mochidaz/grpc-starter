package repository

import (
	"context"
	"github.com/google/uuid"
	"gorm.io/gorm"
	"grpc-starter/modules/notification/v1/entity"
)

type EmailFinder struct {
	db *gorm.DB
}

type EmailFinderUseCase interface {
	FindByID(ctx context.Context, id int) (*entity.EmailSent, error)

	FindByUserID(ctx context.Context, userID uuid.UUID) ([]*entity.EmailSent, error)

	FindByEmail(ctx context.Context, email string) ([]*entity.EmailSent, error)

	GetAllEmailSent(ctx context.Context) ([]*entity.EmailSent, error)
}

func NewEmailFinder(db *gorm.DB) *EmailFinder {
	return &EmailFinder{db: db}
}

func (ef *EmailFinder) FindByID(ctx context.Context, id int) (*entity.EmailSent, error) {
	var emailSent entity.EmailSent

	if err := ef.db.WithContext(ctx).First(&emailSent, id).Error; err != nil {
		return nil, err
	}

	return &emailSent, nil
}

func (ef *EmailFinder) FindByUserID(ctx context.Context, userID uuid.UUID) ([]*entity.EmailSent, error) {
	var emailSent []*entity.EmailSent

	if err := ef.db.WithContext(ctx).Where("m_id = ?", userID).Find(&emailSent).Error; err != nil {
		return nil, err
	}

	return emailSent, nil
}

func (ef *EmailFinder) FindByEmail(ctx context.Context, email string) ([]*entity.EmailSent, error) {
	var emailSent []*entity.EmailSent

	if err := ef.db.WithContext(ctx).Where("\"to\" = ?", email).Find(&emailSent).Error; err != nil {
		return nil, err
	}

	return emailSent, nil
}

func (ef *EmailFinder) GetAllEmailSent(ctx context.Context) ([]*entity.EmailSent, error) {
	var emailSent []*entity.EmailSent

	if err := ef.db.WithContext(ctx).Find(&emailSent).Error; err != nil {
		return nil, err
	}

	return emailSent, nil
}
