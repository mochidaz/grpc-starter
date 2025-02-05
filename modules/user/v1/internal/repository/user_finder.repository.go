package repository

import (
	"context"

	"github.com/google/uuid"
	"github.com/pkg/errors"
	"gorm.io/gorm"

	"grpc-starter/common/cache"
	"grpc-starter/modules/user/v1/entity"
)

// UserFinderRepository defines dependencies for UserFinder
type UserFinderRepository struct {
	db    *gorm.DB
	cache cache.Cacheable
}

// NewUserFinderRepository creates a new UserFinder repository
func NewUserFinderRepository(
	db *gorm.DB,
	cache cache.Cacheable,
) *UserFinderRepository {
	return &UserFinderRepository{
		db:    db,
		cache: cache,
	}
}

// UserFinderRepositoryUseCase is use case for finding in user table
type UserFinderRepositoryUseCase interface {
	// FindByID finds user
	FindByID(ctx context.Context, refID uuid.UUID) (*entity.User, error)
	// FindByEmailPassword finds user by email and password
	FindByEmailPassword(ctx context.Context, email string, password string) (*entity.User, error)
	// FindByEmail finds user by email
	FindByEmail(ctx context.Context, email string) (*entity.User, error)
	// FindAllUsers
	FindAllUsers(ctx context.Context) ([]*entity.User, error)
}

// FindByID finds user
func (r *UserFinderRepository) FindByID(ctx context.Context, refID uuid.UUID) (*entity.User, error) {
	var result *entity.User
	if err := r.db.WithContext(ctx).Model(&entity.User{}).Where("id = ?", refID).First(&result).Error; err != nil {
		return nil, errors.Wrap(err, "[UserFinderRepository - FindByID] Error while finding user data")
	}

	return result, nil
}

// FindByEmailPassword finds user by email and password
func (r *UserFinderRepository) FindByEmailPassword(ctx context.Context, email string, password string) (*entity.User, error) {
	var result *entity.User
	if err := r.db.WithContext(ctx).Model(&entity.User{}).Where("email = ? AND password = ?", email, password).First(&result).Error; err != nil {
		return nil, errors.Wrap(err, "[UserFinderRepository - FindByEmailPassword] Error while finding user data")
	}

	return result, nil
}

// FindByEmail finds user by email
func (r *UserFinderRepository) FindByEmail(ctx context.Context, email string) (*entity.User, error) {
	var result *entity.User
	if err := r.db.WithContext(ctx).Model(&entity.User{}).Where("email = ?", email).First(&result).Error; err != nil {
		return nil, errors.Wrap(err, "[UserFinderRepository - FindByEmail] Error while finding user data")
	}

	return result, nil
}

// FindAllUsers
func (r *UserFinderRepository) FindAllUsers(ctx context.Context) ([]*entity.User, error) {
	var result []*entity.User
	if err := r.db.WithContext(ctx).Model(&entity.User{}).Find(&result).Error; err != nil {
		return nil, errors.Wrap(err, "[UserFinderRepository - FindAllUsers] Error while finding user data")
	}

	return result, nil
}
