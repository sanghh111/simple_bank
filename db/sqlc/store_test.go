package db

import (
	"context"
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestTransfer(t *testing.T) {
	store := NewStore(testDB)

	from_account := RandomAccount(t)
	to_account := RandomAccount(t)

	// run n concurence tranfer tranctions
	n := 5
	ammount := int64(10)

	errs := make(chan error)
	results := make(chan TransferResult)
	fmt.Println(">> Before:", from_account.Balance, to_account.Balance)
	for i := 0; i < n; i++ {
		go func() {
			result, err := store.TransferTx(context.Background(), TransferTxParams{
				FromAccountID: from_account.ID,
				ToAccountID:   to_account.ID,
				Amount:        ammount,
			})

			errs <- err
			results <- result
		}()
	}

	// check results
	existed := make(map[int]bool)
	for i := 0; i < n; i++ {

		err := <-errs
		require.NoError(t, err)

		result := <-results
		require.NotEmpty(t, result)

		transfer := result.Transfer
		require.Equal(t, transfer.Amnount, ammount)
		require.NotEmpty(t, transfer.ID)
		require.Equal(t, transfer.FromAccountID, from_account.ID)
		require.NotEmpty(t, transfer.ToAccountID, to_account.ID)
		require.NotEmpty(t, transfer.CreatedAt)

		// check Enties
		from_entry := result.FromEntry
		require.Equal(t, from_entry.Amnount, ammount)
		require.NotEmpty(t, from_entry.ID)
		require.NotEmpty(t, from_entry.CreatedAt)
		require.Equal(t, from_entry.AccountID, from_account.ID)

		to_entry := result.ToEntry
		require.Equal(t, to_entry.Amnount, ammount)
		require.NotEmpty(t, to_entry.ID)
		require.NotEmpty(t, to_entry.CreatedAt)
		require.Equal(t, to_entry.AccountID, to_account.ID)

		// Check Account
		fromAccount := result.FromAccount
		require.NotEmpty(t, fromAccount)
		require.Equal(t, fromAccount.ID, from_account.ID)

		toAccount := result.ToAccount
		require.NotEmpty(t, toAccount)
		require.Equal(t, toAccount.ID, to_account.ID)

		fmt.Println(">> tx: ", fromAccount.Balance, toAccount.Balance)
		diff1 := from_account.Balance - fromAccount.Balance
		diff2 := toAccount.Balance - to_account.Balance
		require.Equal(t, diff1, diff2)
		require.True(t, diff1 > 0)
		require.True(t, diff1%ammount == 0)

		k := int(diff1 / ammount)
		require.True(t, k >= 1 && k <= n)
		require.NotContains(t, existed, k)
		existed[k] = true
	}

	// CHECK THE FINAL UPDATED BALANCES
	updateFromAccount, err := testQueries.GetAccount(context.Background(), from_account.ID)
	require.NoError(t, err)

	updateToAccount, err := testQueries.GetAccount(context.Background(), to_account.ID)
	require.NoError(t, err)

	fmt.Println(">> After:", updateFromAccount.Balance, updateToAccount.Balance)
	require.Equal(t, from_account.Balance-int64(n)*ammount, updateFromAccount.Balance)
	require.Equal(t, to_account.Balance+int64(n)*ammount, updateToAccount.Balance)
}
