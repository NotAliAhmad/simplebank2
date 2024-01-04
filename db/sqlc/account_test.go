package db

import (
	"context"
	"database/sql"
	"testing"
	"time"

	"github.com/notaliahmad/simplebank2/util"
	"github.com/stretchr/testify/require"
)

func createRandomAccount(t *testing.T) Account {
	args := CreateAccountParams{
		Owner:    util.RandomOwner(),
		Balance:  util.RandomMoney(),
		Currency: util.RandomCurrency(),
	}

	account, err := testQueries.CreateAccount(context.Background(), args)

	require.NoError(t, err, nil)
	require.NotEmpty(t, account)

	require.Equal(t, args.Owner, account.Owner)
	require.Equal(t, args.Balance, account.Balance)
	require.Equal(t, args.Currency, account.Currency)

	require.NotZero(t, account.ID)
	require.NotZero(t, account.CreatedAt)

	return account
}

// all tests must start with the test prefix with the uppercase letter 'T'
func TestCreateAccount(t *testing.T) {
	createRandomAccount(t)
}

func TestGetAccount(t *testing.T) {
	account := createRandomAccount(t)

	same_account, err := testQueries.GetAccount(context.Background(), account.ID)

	require.NoError(t, err)
	require.NotEmpty(t, account)

	require.Equal(t, account.Owner, same_account.Owner)
	require.Equal(t, account.Balance, same_account.Balance)
	require.Equal(t, account.Currency, same_account.Currency)
	require.Equal(t, account.ID, same_account.ID)
	require.WithinDuration(t, account.CreatedAt, same_account.CreatedAt, time.Second)

}

func TestUpdateAccount(t *testing.T) {
	account := createRandomAccount(t)
	args := UpdateAccountParams{
		ID:      account.ID,
		Balance: util.RandomMoney(),
	}
	same_account, err := testQueries.UpdateAccount(context.Background(), args)
	require.NoError(t, err)
	require.NotEmpty(t, same_account)

	require.Equal(t, account.ID, same_account.ID)
	require.Equal(t, account.Owner, same_account.Owner)
	require.Equal(t, args.Balance, same_account.Balance)
	require.Equal(t, account.Currency, same_account.Currency)
	require.WithinDuration(t, account.CreatedAt, same_account.CreatedAt, time.Second)
}

func TestDeleteAccount(t *testing.T) {
	account := createRandomAccount(t)
	testQueries.DeleteAccount(context.Background(), account.ID)

	deleted_account, err := testQueries.GetAccount(context.Background(), account.ID)

	require.Error(t, err)
	require.EqualError(t, err, sql.ErrNoRows.Error())

	require.Empty(t, deleted_account)
}

func TestListAccounts(t *testing.T) {
	for i := 0; i < 10; i++ {
		createRandomAccount(t)
	}
	args := ListAccountsParams{
		Limit:  5,
		Offset: 5,
	}

	accounts, err := testQueries.ListAccounts(context.Background(), args)
	require.NoError(t, err)
	require.Len(t, accounts, 5)

	for _, account := range accounts {
		require.NotEmpty(t, account)
	}
}
