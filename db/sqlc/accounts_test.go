package db

import (
	"context"
	"testing"

	"github.com/WillChen936/simple-bank/utils"
	"github.com/stretchr/testify/require"
)

func createRandomAccount(t *testing.T) Account {
	arg := CreateAccountParams{
		Owner:    utils.CreateRandomOwner(),
		Balance:  utils.CreateRandomMoney(),
		Currency: utils.CreateRandomCurrency(),
	}

	account, err := testQueries.CreateAccount(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, account)
	require.Equal(t, account.Owner, arg.Owner)
	require.Equal(t, account.Balance, arg.Balance)
	require.Equal(t, account.Currency, arg.Currency)

	return account
}

func TestCreateAccount(t *testing.T) {
	createRandomAccount(t)
}
