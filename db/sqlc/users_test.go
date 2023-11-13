package db

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/techschool/simplebank/uti"
)

func createRandomUser(t *testing.T) User {
	arg := CreateUserParams{
		Username:     uti.RandomOwner(),
		HashPassword: "secret",
		FullName:     uti.RandomString(20),
		Email:        uti.RandomString(25),
	}

	user, err := testQueries.CreateUser(context.Background(), arg)

	require.NoError(t, err)
	require.NotEmpty(t, user)

	require.Equal(t, arg.Username, user.Username)
	require.Equal(t, arg.HashPassword, user.HashPassword)
	require.Equal(t, arg.FullName, user.FullName)
	require.Equal(t, arg.Email, user.Email)
	require.NotEmpty(t, user.PasswordChangedAt)
	require.NotEmpty(t, user.CreatedAt)

	return user
}

func TestCreateUser(t *testing.T) {
	user := createRandomUser(t)
	require.NotEmpty(t, user)
}
