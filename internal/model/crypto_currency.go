package model

import (
	"github.com/google/uuid"
)

type CryptoCurrency struct {
	ID                uuid.UUID `json:"id"`                // Уникальный идентификатор (генерируется локально)
	APIID             string    `json:"apiId"`             // Идентификатор из API CoinCap
	Rank              int       `json:"rank"`              // Рейтинг по рыночной капитализации
	Symbol            string    `json:"symbol"`            // Символ криптовалюты
	Name              string    `json:"name"`              // Название криптовалюты
	Supply            float64   `json:"supply"`            // Общее количество монет
	MaxSupply         *float64  `json:"maxSupply"`         // Максимальное количество монет (опционально)
	MarketCapUsd      float64   `json:"marketCapUsd"`      // Рыночная капитализация в USD
	VolumeUsd24Hr     float64   `json:"volumeUsd24Hr"`     // Торговый объем за последние 24 часа в USD
	PriceUsd          float64   `json:"priceUsd"`          // Цена в USD
	ChangePercent24Hr float64   `json:"changePercent24Hr"` // Изменение цены за последние 24 часа (в процентах)
	Vwap24Hr          float64   `json:"vwap24Hr"`          // Средневзвешенная цена за последние 24 часа
	ImageURL          string    `json:"imageUrl"`          // URL изображения криптовалюты
	Description       *string   `json:"description"`       // Описание (добавляется из локальной БД)
}
