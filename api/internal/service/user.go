package service

import (
	"context"
	"errors"

	userSDK "github.com/clerk/clerk-sdk-go/v2/user"
	"github.com/neilchetna/mind-mapper/internal/models"
	"gorm.io/gorm"
)

type UserRepository interface {
	GetByClerkId(ctx context.Context, clerkId string) (*models.User, error)
	CreateUser(ctx context.Context, user *models.User) error
}

type UserService struct {
	userRepo UserRepository
}

func UserServiceBuilder(userRepo UserRepository) *UserService {
	return &UserService{userRepo}
}

func (s *UserService) SyncClerkUserToDatabase(ctx context.Context, clerkId string) (*models.User, error) {
	user, err := s.userRepo.GetByClerkId(ctx, clerkId)

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			clerkUser, err := userSDK.Get(ctx, clerkId)
			if err != nil {
				return nil, err
			}

			user = &models.User{
				Email:       clerkUser.EmailAddresses[0].EmailAddress,
				ClerkUserId: clerkId,
			}

			err = s.userRepo.CreateUser(ctx, user)
			if err != nil {
				return nil, err
			}
		} else {
			return nil, err
		}
	}

	return user, nil
}
