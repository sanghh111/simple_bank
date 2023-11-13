package uti

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func HashPassWord(t *testing.T) (string, string) {
	passWord := RandomString(20)
	hashPassword, err := HashPassword(passWord)
	require.NoError(t, err)
	require.NotEmpty(t, hashPassword)
	return passWord, hashPassword
}

func TestHashPassWord(t *testing.T) {
	HashPassWord(t)
}

func TestCheckPassWord(t *testing.T) {
	pass, hash := HashPassWord(t)
	is_check := CheckPassword(pass, hash)
	require.Equal(t, true, is_check)
}
