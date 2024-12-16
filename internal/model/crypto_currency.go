package model

import (
	"github.com/google/uuid"
)

type CryptoCurrency struct {
	ID          uuid.UUID
	APIID       string
	Symbol      string
	Name        string
	Description *string
	Supply      *float64
	MaxSupply   *float64
}