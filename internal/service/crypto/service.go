package service

import (
	"context"
	"fmt"
	"github.com/dlc-01/BackendCrypto/pkg/coincap"
	"github.com/go-resty/resty/v2"

	"github.com/dlc-01/BackendCrypto/internal/model"
	"github.com/dlc-01/BackendCrypto/internal/repository"
)

type CryptoService struct {
	cryptoRepo    repository.CryptoRepository
	coinCapServer coincap.CoinCapClient
	client        *resty.Client
}

func NewCryptoService(cryptoRepo repository.CryptoRepository, coinCapServer coincap.CoinCapClient) *CryptoService {
	return &CryptoService{
		cryptoRepo:    cryptoRepo,
		coinCapServer: coinCapServer,
	}
}

func (s *CryptoService) GetAll(ctx context.Context) (*[]model.CryptoCurrency, error) {
	cryptosFromDB, err := s.cryptoRepo.GetAll(ctx)
	if err != nil {
		return nil, fmt.Errorf("error while getting array of cryptocurrencies from DB: %w", err)
	}

	var updatedCryptos []model.CryptoCurrency
	for _, crypto := range *cryptosFromDB {
		// Получить данные из API для каждой криптовалюты
		cryptoFromAPI, err := s.coinCapServer.GetCryptoBySymbol(crypto.Symbol)
		if err != nil {
			// Логируем ошибку, но продолжаем
			fmt.Printf("Error fetching crypto %s from API: %v\n", crypto.Symbol, err)
			updatedCryptos = append(updatedCryptos, crypto)
			continue
		}

		// Обновляем описание из БД, если оно есть
		if crypto.Description != nil {
			cryptoFromAPI.Description = crypto.Description
		}

		// Сохраняем обновленные данные в БД
		_, err = s.cryptoRepo.Update(ctx, cryptoFromAPI)
		if err != nil {
			fmt.Printf("Error updating crypto %s in DB: %v\n", crypto.Symbol, err)
		}

		updatedCryptos = append(updatedCryptos, *cryptoFromAPI)
	}

	return &updatedCryptos, nil
}

func (s *CryptoService) Create(ctx context.Context, symbol, description string) (*model.CryptoCurrency, error) {
	// Получить данные о криптовалюте из CoinCap API
	cryptoFromAPI, err := s.coinCapServer.GetCryptoBySymbol(symbol)
	if err != nil {
		return nil, fmt.Errorf("error fetching cryptocurrency from CoinCap API: %w", err)
	}

	cryptoFromAPI.Description = &description

	// Сохранить данные в базу данных
	cryptoFromDB, err := s.cryptoRepo.Create(ctx, cryptoFromAPI)
	if err != nil {
		return nil, fmt.Errorf("error saving cryptocurrency to DB: %w", err)
	}

	return cryptoFromDB, nil
}

func (s *CryptoService) GetBySymbol(ctx context.Context, symbol string) (*model.CryptoCurrency, error) {
	// Проверить, есть ли криптовалюта в CoinCap API
	cryptoFromAPI, err := s.coinCapServer.GetCryptoBySymbol(symbol)
	if err == nil && cryptoFromAPI != nil {
		// Проверить наличие описания в БД
		cryptoFromDB, _ := s.cryptoRepo.GetBySymbol(ctx, symbol)
		if cryptoFromDB != nil && cryptoFromDB.Description != nil {
			cryptoFromAPI.Description = cryptoFromDB.Description
		}
		// Сохранить данные из CoinCap API в БД
		_, err := s.cryptoRepo.Update(ctx, cryptoFromAPI)
		if err != nil {
			return nil, fmt.Errorf("error updating cryptocurrency in DB: %w", err)
		}
		return cryptoFromAPI, nil
	}

	// Если не найдено в API, проверить в БД
	cryptoFromDB, err := s.cryptoRepo.GetBySymbol(ctx, symbol)
	if err != nil {
		return nil, fmt.Errorf("crypto not found in API and DB: %w", err)
	}

	return cryptoFromDB, nil
}

func (s *CryptoService) Update(ctx context.Context, crypto *model.CryptoCurrency) error {
	_, err := s.cryptoRepo.Update(ctx, crypto)
	if err != nil {
		return fmt.Errorf("error updating cryptocurrency: %w", err)
	}
	return nil
}

func (s *CryptoService) Delete(ctx context.Context, crypto *model.CryptoCurrency) error {
	err := s.cryptoRepo.Delete(ctx, crypto.Symbol)
	if err != nil {
		return fmt.Errorf("error while deliting cryptocurrency: %w", err)
	}
	return nil
}
