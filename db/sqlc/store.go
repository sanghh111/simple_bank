package db

import (
	"context"
	"database/sql"
	"fmt"
)

// Store provide all functions to execute db queries and transactions
type Store struct {
	*Queries
	db *sql.DB
}

// NewStore creates a New store
func NewStore(db *sql.DB) *Store {
	return &Store{
		db:      db,
		Queries: New(db),
	}
}

// execTX execute a function within a database trancaction
func (store *Store) execTx(ctx context.Context, fn func(*Queries) error) error {
	tx, err := store.db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}

	q := New(tx)
	err = fn(q)
	if err != nil {
		rbErr := tx.Rollback()
		if rbErr != nil {
			return fmt.Errorf("tx.err: %v, rb.err %v", err, rbErr)
		}

		return err
	}

	return tx.Commit()
}

// TransferTxParams  contains input parameters of the transfer tranction
type TransferTxParams struct {
	FromAccountID int64 `json:"from_account_id"`
	ToAccountID   int64 `json:"to_account_id"`
	Amount        int64 `json:"ammount"`
}

// TransferResult is the result of the transer tranction
type TransferResult struct {
	Transfer    Transfer `json:"transfer"`
	FromAccount Account  `json:"from_account"`
	ToAccount   Account  `json:"to_account"`
	FromEntry   Entry    `json:"from_enty"`
	ToEntry     Entry    `json:"to_entry"`
}

// TransferTx performs a money transfer from one account to the other
// It create a transfer record, add account entries, and update account balance within single database transaction
func (store *Store) TransferTx(ctx context.Context, arg TransferTxParams) (TransferResult, error) {
	var result TransferResult

	err := store.execTx(ctx, func(q *Queries) error {
		var err error

		result.Transfer, err = q.CreateTransfer(ctx, CreateTransferParams{
			FromAccountID: arg.FromAccountID,
			ToAccountID:   arg.ToAccountID,
			Amnount:       arg.Amount,
		})

		if err != nil {
			return err
		}

		result.FromEntry, err = q.CreateEntrie(ctx, CreateEntrieParams{
			AccountID: arg.FromAccountID,
			Amnount:   arg.Amount,
		})

		if err != nil {
			return err
		}

		result.ToEntry, err = q.CreateEntrie(ctx, CreateEntrieParams{
			AccountID: arg.ToAccountID,
			Amnount:   arg.Amount,
		})

		if err != nil {
			return err
		}

		// get account -> update accounts balance
		fromAccount, err := q.GetAccountForUpdate(ctx, arg.FromAccountID)
		if err != nil {
			return nil
		}

		result.FromAccount, err = q.UpdateAccountBalace(ctx, UpdateAccountBalaceParams{
			ID:      fromAccount.ID,
			Balance: fromAccount.Balance - arg.Amount,
		})

		if err != nil {
			return err
		}

		ToAccount, err := q.GetAccountForUpdate(ctx, arg.ToAccountID)
		if err != nil {
			return nil
		}

		result.ToAccount, err = q.UpdateAccountBalace(ctx, UpdateAccountBalaceParams{
			ID:      ToAccount.ID,
			Balance: ToAccount.Balance + arg.Amount,
		})

		if err != nil {
			return nil
		}

		return nil
	})

	return result, err
}
