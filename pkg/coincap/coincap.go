package coincap

import (
	"fmt"
	"github.com/dlc-01/BackendCrypto/internal/model"
	"strconv"
	"strings"
	"time"

	"github.com/go-resty/resty/v2"
	"github.com/google/uuid"
)

// Базовый URL API CoinCap
const baseURL = "https://api.coincap.io/v2"

// Базовый URL для изображений криптовалют
const imageBaseURL = "https://assets.coincap.io/assets/icons/"

// CoinCapClient представляет клиента для взаимодействия с CoinCap API
type CoinCapClient struct {
	httpClient *resty.Client
}

// NewCoinCapClient создает новый экземпляр CoinCapClient
func NewCoinCapClient() *CoinCapClient {
	client := resty.New().
		SetBaseURL(baseURL).
		SetTimeout(10 * time.Second).      // Установка таймаута
		SetRetryCount(3).                  // Количество повторных попыток
		SetRetryWaitTime(2 * time.Second). // Время ожидания между попытками
		SetRetryMaxWaitTime(10 * time.Second)

	return &CoinCapClient{
		httpClient: client,
	}
}

// CryptoCurrency представляет структуру данных о криптовалюте
type CryptoCurrency struct {
	ID                uuid.UUID // Уникальный идентификатор
	APIID             string    // Идентификатор из API CoinCap
	Rank              int       // Рейтинг по рыночной капитализации
	Symbol            string    // Символ криптовалюты в нижнем регистре
	Name              string    // Название криптовалюты
	Supply            float64   // Общее количество монет
	MaxSupply         *float64  // Максимальное количество монет (опционально)
	MarketCapUsd      float64   // Рыночная капитализация в USD
	VolumeUsd24Hr     float64   // Торговый объем за последние 24 часа в USD
	PriceUsd          float64   // Цена в USD
	ChangePercent24Hr float64   // Изменение цены за последние 24 часа в процентах
	Vwap24Hr          float64   // Средневзвешенная цена за последние 24 часа
	ImageURL          string    // URL изображения криптовалюты
	Description       *string   // Описание криптовалюты (опционально)
}

// responseAsset представляет структуру данных о криптовалюте, получаемую от API
type responseAsset struct {
	ID                string  `json:"id"`
	Rank              string  `json:"rank"`
	Symbol            string  `json:"symbol"`
	Name              string  `json:"name"`
	Supply            string  `json:"supply"`
	MaxSupply         *string `json:"maxSupply"` // Может быть null
	MarketCapUsd      string  `json:"marketCapUsd"`
	VolumeUsd24Hr     string  `json:"volumeUsd24Hr"`
	PriceUsd          string  `json:"priceUsd"`
	ChangePercent24Hr string  `json:"changePercent24Hr"`
	Vwap24Hr          string  `json:"vwap24Hr"`
	// Добавьте другие поля при необходимости
}

// ResponseData представляет структуру ответа от CoinCap API
type ResponseData struct {
	Data []responseAsset `json:"data"`
}

// GetAllCryptos получает список всех криптовалют и возвращает их в виде []CryptoCurrency
func (c *CoinCapClient) GetAllCryptos() ([]model.CryptoCurrency, error) {
	var response ResponseData

	// Выполнение GET-запроса с автоматическим разбором ответа в структуру ResponseData
	resp, err := c.httpClient.R().
		SetResult(&response).
		Get("/assets")

	if err != nil {
		return nil, fmt.Errorf("ошибка при выполнении запроса к CoinCap API: %w", err)
	}

	// Проверка статуса ответа
	if resp.StatusCode() != 200 {
		return nil, fmt.Errorf("получен статус %d от CoinCap API", resp.StatusCode())
	}

	// Преобразование response.Data в []CryptoCurrency
	cryptos := make([]model.CryptoCurrency, 0, len(response.Data))
	for _, asset := range response.Data {
		crypto, err := mapAssetToCryptoCurrency(asset)
		if err != nil {
			// Логирование ошибки и пропуск данного актива
			fmt.Printf("Ошибка при преобразовании актива %s: %v\n", asset.ID, err)
			continue
		}
		cryptos = append(cryptos, crypto)
	}

	return cryptos, nil
}

// GetCryptoBySymbol получает информацию о криптовалюте по её символу и возвращает её в виде *CryptoCurrency
func (c *CoinCapClient) GetCryptoBySymbol(symbol string) (*model.CryptoCurrency, error) {
	var response ResponseData

	// Выполнение GET-запроса с поиском по символу и автоматическим разбором ответа
	resp, err := c.httpClient.R().
		SetQueryParam("search", symbol).
		SetResult(&response).
		Get("/assets")

	if err != nil {
		return nil, fmt.Errorf("ошибка при выполнении запроса к CoinCap API: %w", err)
	}

	// Проверка статуса ответа
	if resp.StatusCode() != 200 {
		return nil, fmt.Errorf("получен статус %d от CoinCap API", resp.StatusCode())
	}

	// Проверка наличия данных
	if len(response.Data) == 0 {
		return nil, fmt.Errorf("криптовалюта с символом %s не найдена", symbol)
	}

	// Преобразование первого найденного актива в CryptoCurrency
	crypto, err := mapAssetToCryptoCurrency(response.Data[0])
	if err != nil {
		return nil, fmt.Errorf("ошибка при преобразовании актива: %w", err)
	}

	return &crypto, nil
}

func mapAssetToCryptoCurrency(asset responseAsset) (model.CryptoCurrency, error) {
	// Генерация UUID для идентификации
	id := uuid.New()

	// Преобразование данных из API в структуру
	crypto := model.CryptoCurrency{
		ID:                id,
		APIID:             asset.ID,
		Rank:              parseInt(asset.Rank),
		Symbol:            asset.Symbol,
		Name:              asset.Name,
		Supply:            parseFloat(asset.Supply),
		MaxSupply:         parseOptionalFloat(asset.MaxSupply),
		MarketCapUsd:      parseFloat(asset.MarketCapUsd),
		VolumeUsd24Hr:     parseFloat(asset.VolumeUsd24Hr),
		PriceUsd:          parseFloat(asset.PriceUsd),
		ChangePercent24Hr: parseFloat(asset.ChangePercent24Hr),
		Vwap24Hr:          parseFloat(asset.Vwap24Hr),
		ImageURL:          fmt.Sprintf("%s%s@2x.png", imageBaseURL, strings.ToLower(asset.Symbol)),
		Description:       nil, // Описание добавляется из БД
	}
	return crypto, nil
}

// Пример функций для преобразования данных
func parseFloat(value string) float64 {
	result, _ := strconv.ParseFloat(value, 64)
	return result
}

func parseInt(value string) int {
	result, _ := strconv.Atoi(value)
	return result
}

func parseOptionalFloat(value *string) *float64 {
	if value == nil || *value == "" {
		return nil
	}
	parsed, _ := strconv.ParseFloat(*value, 64)
	return &parsed
}
