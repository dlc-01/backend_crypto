package main

import (
	"context"
	"fmt"
	"github.com/dlc-01/BackendCrypto/internal/conf"
	"log"

	"github.com/dlc-01/BackendCrypto/internal/model"
	"github.com/dlc-01/BackendCrypto/pkg/coincap"
	"github.com/dlc-01/BackendCrypto/pkg/db/psql"
	"github.com/dlc-01/BackendCrypto/pkg/db/psql/repositories"
)

func main() {
	// Инициализация конфигурации и подключения к БД
	cfg, err := conf.InitConf()
	if err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}

	db, err := psql.NewSQLClient(*cfg)
	if err != nil {
		log.Fatalf("Failed to connect to the database: %v", err)
	}

	cryptoRepo := repositories.NewCryptoRepo(db)
	coinCapClient := coincap.NewCoinCapClient()

	// Получение данных первых 10 криптовалют из CoinCap API
	cryptos, err := coinCapClient.GetAllCryptos()
	if err != nil {
		log.Fatalf("Failed to fetch cryptos from CoinCap API: %v", err)
	}

	// Вставка первых 10 криптовалют в БД
	for i, crypto := range cryptos {
		if i >= 10 {
			break
		}

		// Создаем объект для сохранения в БД
		cryptoToSave := &model.CryptoCurrency{
			ID:                crypto.ID,
			APIID:             crypto.APIID,
			Rank:              crypto.Rank,
			Symbol:            crypto.Symbol,
			Name:              crypto.Name,
			Supply:            crypto.Supply,
			MaxSupply:         crypto.MaxSupply,
			MarketCapUsd:      crypto.MarketCapUsd,
			VolumeUsd24Hr:     crypto.VolumeUsd24Hr,
			PriceUsd:          crypto.PriceUsd,
			ChangePercent24Hr: crypto.ChangePercent24Hr,
			Vwap24Hr:          crypto.Vwap24Hr,
			ImageURL:          crypto.ImageURL,
			Description:       nil, // Описание можно добавить вручную
		}

		_, err := cryptoRepo.Create(context.Background(), cryptoToSave)
		if err != nil {
			log.Printf("Failed to insert crypto %s into database: %v", crypto.Symbol, err)
		} else {
			fmt.Printf("Successfully inserted crypto: %s\n", crypto.Symbol)
		}
	}
}
