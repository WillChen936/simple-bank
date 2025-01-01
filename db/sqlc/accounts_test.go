package db

import (
	"context"
	"testing"
	"time"

	"github.com/WillChen936/simple-bank/utils"
	"github.com/stretchr/testify/require"
)

func TestCreateAccount(t *testing.T) {
	account := createRandomAccount(t)
	testQueries.DeleteAccount(context.Background(), account.ID)
}

func TestGetAccount(t *testing.T) {
	account1 := createRandomAccount(t)

	account2, err := testQueries.GetAccount(context.Background(), account1.ID)
	require.NoError(t, err)
	require.NotEmpty(t, account2)
	require.Equal(t, account1.ID, account2.ID)
	require.Equal(t, account1.Owner, account2.Owner)
	require.Equal(t, account1.Balance, account2.Balance)
	require.Equal(t, account1.Currency, account2.Currency)
	require.WithinDuration(t, account1.CreatedAt, account2.CreatedAt, time.Second)
}

// func TestDeleteAccount(t *testing.T) {
// 	account := createRandomAccount(t)
// 	err := testQueries.DeleteAccount(context.Background(), account.ID)

// 	require.NoError(t, err)
// }

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
