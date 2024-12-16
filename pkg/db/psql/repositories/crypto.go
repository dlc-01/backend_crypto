package repositories

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/dlc-01/BackendCrypto/internal/model"
	"github.com/dlc-01/BackendCrypto/internal/model/projectError"
	"github.com/dlc-01/BackendCrypto/internal/repository"
	"github.com/dlc-01/BackendCrypto/pkg/db/psql/query"
	"github.com/jackc/pgx/v5"
	"github.com/lib/pq"
)

var _ repository.CryptoRepository = (*CryptoRepo)(nil)

type CryptoRepo struct {
	*sql.DB
}

func (c CryptoRepo) Create(ctx context.Context, crypto *model.CryptoCurrency) (*model.CryptoCurrency, error) {
	tx, err := c.Begin()
	if err != nil {
		return nil, fmt.Errorf("%w : %s", projectError.ErrorStartingTransaction, err)
	}

	err = tx.QueryRowContext(ctx, query.CreateCryptocurrency, &crypto.Symbol, &crypto.Name, &crypto.Description, &crypto.Supply, &crypto.MaxSupply).
		Scan(&crypto.ID)
	if err != nil {
		if errCode := pq.ErrorCode(err.Error()); errCode == "23505" {
			return nil, projectError.ErrorCryptoExist
		}
		return nil, fmt.Errorf("%w: %s", projectError.ErrorCantCreateCrypto, err)
	}

	err = tx.Commit()
	if err != nil {
		return nil, fmt.Errorf("%w : %s", projectError.ErrorCommitTransaction, err)
	}
	return crypto, nil
}

func (c CryptoRepo) GetAll(ctx context.Context) (*[]model.CryptoCurrency, error) {
	stored := make([]model.CryptoCurrency, 0)
	tx, err := c.Begin()
	if err != nil {
		return nil, fmt.Errorf("%w : %s", projectError.ErrorStartingTransaction, err)
	}

	row, err := tx.QueryContext(ctx, query.GetAllCryptocurrencies)
	if err != nil {
		if err == pgx.ErrNoRows {
			return nil, projectError.ErrorNoData
		}
		return nil, fmt.Errorf("%w, %s", projectError.ErrorCryptoNotFound, err)
	}

	for row.Next() {
		crypto := model.CryptoCurrency{}

		err = row.Scan(&crypto.ID,
			&crypto.Symbol,
			&crypto.Name,
			&crypto.Description,
			&crypto.Supply,
			&crypto.MaxSupply)
		if err != nil {
			if err == pgx.ErrNoRows {
				return nil, projectError.ErrorNoData
			}
			return nil, fmt.Errorf("%w, %s", projectError.ErrorCryptoNotFound, err)
		}
		stored = append(stored, crypto)
	}

	err = tx.Commit()
	if err != nil {
		return nil, fmt.Errorf("%w : %s", projectError.ErrorCommitTransaction, err)
	}
	return &stored, nil

}

func (c CryptoRepo) GetByUUID(ctx context.Context, uuid string) (*model.CryptoCurrency, error) {
	//TODO implement me
	panic("implement me")
}

func (c CryptoRepo) GetBySymbol(ctx context.Context, symbol string) (*model.CryptoCurrency, error) {
	var stored model.CryptoCurrency

	tx, err := c.Begin()
	if err != nil {
		return nil, fmt.Errorf("%w : %s", projectError.ErrorStartingTransaction, err)
	}

	err = tx.QueryRowContext(ctx, query.GetCryptocurrencyBySymbol, symbol).
		Scan(&stored.ID,
			&stored.Symbol,
			&stored.Name,
			&stored.Description,
			&stored.Supply,
			&stored.MaxSupply)
	if err != nil {
		if err == pgx.ErrNoRows {
			return nil, projectError.ErrorNoData
		}
		return nil, fmt.Errorf("%w: %s", projectError.ErrorCryptoNotFound, err)
	}

	err = tx.Commit()
	if err != nil {
		return nil, fmt.Errorf("%w : %s", projectError.ErrorCommitTransaction, err)
	}

	return &stored, nil
}

func (c CryptoRepo) Delete(ctx context.Context, symbol string) error {
	tx, err := c.Begin()
	if err != nil {
		return fmt.Errorf("%w : %s", projectError.ErrorStartingTransaction, err)
	}

	err = tx.QueryRowContext(ctx, query.DeleteCryptocurrency, symbol).Err()
	if err != nil {
		return fmt.Errorf("%w : %w", projectError.ErrorCantDeleteCrypto, err)
	}

	err = tx.Commit()
	if err != nil {
		return fmt.Errorf("%w : %s", projectError.ErrorCommitTransaction, err)
	}

	return nil
}

func (c CryptoRepo) Update(ctx context.Context, crypto *model.CryptoCurrency) (*model.CryptoCurrency, error) {
	tx, err := c.Begin()
	if err != nil {
		return nil, fmt.Errorf("%w : %s", projectError.ErrorStartingTransaction, err)
	}

	err = tx.QueryRowContext(ctx, query.UpdateCryptocurrency, crypto.ID,
		crypto.Symbol,
		crypto.Name,
		crypto.Description,
		crypto.Supply,
		crypto.MaxSupply).Err()
	if err != nil {
		return nil, fmt.Errorf("%w: %s", projectError.ErrorCantUpdateCrypto, err)
	}

	err = tx.Commit()
	if err != nil {
		return nil, fmt.Errorf("%w : %s", projectError.ErrorCommitTransaction, err)
	}

	return crypto, nil
}

func NewCryptoRepo(client *sql.DB) *CryptoRepo {
	return &CryptoRepo{client}
}
