package repository

import (
	"context"
	"github.com/dlc-01/BackendCrypto/internal/model"
)

type UserRepository interface {
	Create(ctx context.Context, info *model.User) (*model.User, error)
	Get(ctx context.Context, uuid string) (*model.User, error)
	GetByEmail(ctx context.Context, email string) (*model.User, error)
	GetByUsername(ctx context.Context, username string) (*model.User, error)
	Update(ctx context.Context, info *model.User) (*model.User, error)
	Delete(ctx context.Context, uuid string) error
}
