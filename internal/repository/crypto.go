package repository

import (
	"context"
	"github.com/dlc-01/BackendCrypto/internal/model"
)

type CryptoRepository interface {
	Create(ctx context.Context, crypto *model.CryptoCurrency) (*model.CryptoCurrency, error)
	GetAll(ctx context.Context) (*[]model.CryptoCurrency, error)
	GetByUUID(ctx context.Context, uuid string) (*model.CryptoCurrency, error)
	GetBySymbol(ctx context.Context, symbol string) (*model.CryptoCurrency, error)
	Delete(ctx context.Context, symbol string) error
	Update(ctx context.Context, crypto *model.CryptoCurrency) (*model.CryptoCurrency, error)
}
