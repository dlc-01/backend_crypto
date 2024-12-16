package repositories

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/dlc-01/BackendCrypto/internal/model/projectError"
	"github.com/lib/pq"

	"github.com/dlc-01/BackendCrypto/internal/model"
	"github.com/dlc-01/BackendCrypto/internal/repository"
	"github.com/dlc-01/BackendCrypto/pkg/db/psql/query"
)

var _ repository.UserRepository = (*UserRepo)(nil)

type UserRepo struct {
	*sql.DB
}

func NewUserRepo(client *sql.DB) *UserRepo {
	return &UserRepo{client}
}

func (u *UserRepo) Create(ctx context.Context, user *model.User) (*model.User, error) {
	tx, err := u.Begin()
	if err != nil {
		return nil, fmt.Errorf("%w : %s", projectError.ErrorStartingTransaction, err)
	}

	err = tx.QueryRowContext(ctx, query.CreateUser,
		user.FirstName,
		user.LastName,
		user.Username,
		user.Email,
		user.PasswordHash,
	).Scan(&user.ID, &user.CreatedAt, &user.UpdatedAt)
	if err != nil {
		if pqErr, ok := err.(*pq.Error); ok && pqErr.Code == "23505" {
			return nil, projectError.ErrorUserExist
		}
		return nil, fmt.Errorf("%w : %s", projectError.ErrorCantCreateUser, err)
	}

	if err = tx.Commit(); err != nil {
		return nil, fmt.Errorf("%w : %s", projectError.ErrorCommitTransaction, err)
	}

	return user, nil
}

func (u *UserRepo) Get(ctx context.Context, uuid string) (*model.User, error) {
	var user model.User

	err := u.QueryRowContext(ctx, query.GetUserByUUID, uuid).
		Scan(&user.ID, &user.FirstName, &user.LastName, &user.Username, &user.Email, &user.PasswordHash, &user.CreatedAt, &user.UpdatedAt)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, projectError.ErrorUserNotFound
		}
		return nil, fmt.Errorf("%w : %s", projectError.ErrorCantGetUser, err)
	}

	return &user, nil
}

func (u *UserRepo) GetByUsername(ctx context.Context, username string) (*model.User, error) {
	var user model.User

	err := u.QueryRowContext(ctx, query.GetUserByUsername, username).
		Scan(&user.ID, &user.FirstName, &user.LastName, &user.Username, &user.Email, &user.PasswordHash, &user.CreatedAt, &user.UpdatedAt)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, projectError.ErrorUserNotFound
		}
		return nil, fmt.Errorf("%w : %s", projectError.ErrorCantGetUser, err)
	}

	return &user, nil
}

func (u *UserRepo) GetByEmail(ctx context.Context, email string) (*model.User, error) {
	var user model.User

	err := u.QueryRowContext(ctx, query.GetUserByEmail, email).
		Scan(&user.ID, &user.FirstName, &user.LastName, &user.Username, &user.Email, &user.PasswordHash, &user.CreatedAt, &user.UpdatedAt)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, projectError.ErrorUserNotFound
		}
		return nil, fmt.Errorf("%w : %s", projectError.ErrorCantGetUser, err)
	}

	return &user, nil
}

func (u *UserRepo) Update(ctx context.Context, user *model.User) (*model.User, error) {
	tx, err := u.Begin()
	if err != nil {
		return nil, fmt.Errorf("%w : %s", projectError.ErrorStartingTransaction, err)
	}

	_, err = tx.ExecContext(ctx, query.UpdateUser,
		user.ID,
		user.FirstName,
		user.LastName,
		user.Username,
		user.Email,
		user.PasswordHash,
	)
	if err != nil {
		return nil, fmt.Errorf("%w : %s", projectError.ErrorCantUpdateUser, err)
	}

	if err = tx.Commit(); err != nil {
		return nil, fmt.Errorf("%w : %s", projectError.ErrorCommitTransaction, err)
	}

	return user, nil
}

func (u *UserRepo) Delete(ctx context.Context, uuid string) error {
	tx, err := u.Begin()
	if err != nil {
		return fmt.Errorf("%w : %s", projectError.ErrorStartingTransaction, err)
	}

	_, err = tx.ExecContext(ctx, query.DeleteUser, uuid)
	if err != nil {
		return fmt.Errorf("%w : %s", projectError.ErrorCantDeleteUser, err)
	}

	if err = tx.Commit(); err != nil {
		return fmt.Errorf("%w : %s", projectError.ErrorCommitTransaction, err)
	}

	return nil
}
