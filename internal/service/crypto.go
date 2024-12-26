package service

import (
	"context"
	"github.com/dlc-01/BackendCrypto/internal/model"
)

type ICryptoService interface {
	GetAll(ctx context.Context) (*[]model.CryptoCurrency, error)
	GetBySymbol(ctx context.Context, symb string) (*model.CryptoCurrency, error)
	Create(ctx context.Context, symbol, description string) (*model.CryptoCurrency, error)
	Update(ctx context.Context, crypto *model.CryptoCurrency) error
	Delete(ctx context.Context, user *model.CryptoCurrency) error
}
