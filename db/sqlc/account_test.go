package db

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/techschool/simplebank/uti"
)

func RandomAccount(t *testing.T) Account {

	user := createRandomUser(t)

	arg := CreateAccountParams{
		Owner:    user.Username,
		Balance:  uti.RandomMoney(),
		Currency: "VND",
	}

	account, err := testQueries.CreateAccount(context.Background(), arg)

	require.NoError(t, err)
	require.NotEmpty(t, account)

	require.Equal(t, arg.Balance, account.Balance)
	require.Equal(t, arg.Currency, account.Currency)
	require.Equal(t, arg.Owner, account.Owner)

	require.NotEmpty(t, account.ID)
	require.NotEmpty(t, account.CreatedAt)

	return account
}

func TestCreateAccount(t *testing.T) {
	RandomAccount(t)
}

func TestGetAccount(t *testing.T) {
	account1 := RandomAccount(t)

	account2, err := testQueries.GetAccount(context.Background(), account1.ID)

	require.NoError(t, err)
	require.NotEmpty(t, account2)

	require.Equal(t, account1.ID, account2.ID)
	require.Equal(t, account1.Owner, account2.Owner)
	require.Equal(t, account1.Balance, account2.Balance)
	require.Equal(t, account1.Currency, account2.Currency)
	require.Equal(t, account1.CreatedAt, account2.CreatedAt)
}

func TestListAccount(t *testing.T) {
	for i := 0; i < 10; i++ {
		RandomAccount(t)
	}

	arg := ListAccountParams{
		Offset: 5,
		Limit:  5,
	}

	listAccount, err := testQueries.ListAccount(context.Background(), arg)
	require.NoError(t, err)
	require.Equal(t, len(listAccount), 5)

}

func TestUpdatBalance(t *testing.T) {
	balance := uti.RandomMoney()
	account := RandomAccount(t)

	_, error := testQueries.UpdateAccountBalace(context.Background(), UpdateAccountBalaceParams{
		Balance: account.Balance + (balance),
		ID:      account.ID,
	})

	require.NoError(t, error)

	acccount_updater, err := testQueries.GetAccount(context.Background(), account.ID)

	require.NoError(t, err)
	require.NotEmpty(t, acccount_updater)
	require.Equal(t, account.Balance+balance, acccount_updater.Balance)

}
