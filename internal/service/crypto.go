package service

import (
	"context"
	"github.com/dlc-01/BackendCrypto/internal/model"
)

type ICryptoService interface {
	GetByUserID(ctx context.Context, user *model.User) (*model.User, error)

	GetByUsername(ctx context.Context, user *model.User) (*model.User, error)

	Update(ctx context.Context, user *model.User) error

	Delete(ctx context.Context, user *model.User) error
}
