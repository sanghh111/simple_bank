package db

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/techschool/simplebank/uti"
)

func TestCreateEntries(t *testing.T) {
	account := RandomAccount(t)
	arg := CreateEntrieParams{
		AccountID: account.ID,
		Amnount:   uti.RandomInt(10, 100),
	}

	entries, err := testQueries.CreateEntrie(context.Background(), arg)

	require.NoError(t, err)
	require.Equal(t, account.ID, entries.AccountID)
	require.NotEmpty(t, entries.CreatedAt)
	require.NotEmpty(t, entries.ID)
}
