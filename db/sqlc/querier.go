// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.23.0

package db

import (
	"context"
)

type Querier interface {
	CreateAccount(ctx context.Context, arg CreateAccountParams) (Account, error)
	CreateEntrie(ctx context.Context, arg CreateEntrieParams) (Entry, error)
	CreateTransfer(ctx context.Context, arg CreateTransferParams) (Transfer, error)
	CreateUser(ctx context.Context, arg CreateUserParams) (User, error)
	GetAccount(ctx context.Context, id int64) (Account, error)
	GetAccountByOwner(ctx context.Context, arg GetAccountByOwnerParams) (Account, error)
	GetAccountForUpdate(ctx context.Context, id int64) (Account, error)
	GetEntrie(ctx context.Context, id int64) (Entry, error)
	GetTransfer(ctx context.Context, id int64) (Transfer, error)
	GetUser(ctx context.Context, username string) (User, error)
	GetUserByEmail(ctx context.Context, email string) (User, error)
	ListAccount(ctx context.Context, arg ListAccountParams) ([]Account, error)
	UpdateAccountBalace(ctx context.Context, arg UpdateAccountBalaceParams) (Account, error)
}

var _ Querier = (*Queries)(nil)
