package db

import (
	"context"
	"testing"

	"github.com/notaliahmad/simplebank2/util"
	"github.com/stretchr/testify/require"
)

func TestCreateEntry(t *testing.T) {
	account := createRandomAccount(t)

	args := CreateEntryParams{
		AccountID: account.ID,
		Amount:    util.RandomMoney(),
	}

	entry, err := testQueries.CreateEntry(context.Background(), args)

	require.NotEmpty(t, entry)
	require.NoError(t, err)

	require.Equal(t, args.AccountID, entry.AccountID)
	require.Equal(t, args.Amount, entry.Amount)

	require.NotZero(t, account.ID)
	require.NotZero(t, account.CreatedAt)
}

func TestGetEntry(t *testing.T) {
	account := createRandomAccount(t)

	args := CreateEntryParams{
		AccountID: account.ID,
		Amount:    util.RandomMoney(),
	}
	entry, _ := testQueries.CreateEntry(context.Background(), args)

	getEntry, err := testQueries.GetEntry(context.Background(), entry.ID)

	require.NotEmpty(t, getEntry)
	require.NoError(t, err)

	require.Equal(t, entry.AccountID, getEntry.AccountID)
	require.Equal(t, entry.Amount, getEntry.Amount)

	require.NotZero(t, getEntry.ID)
	require.NotZero(t, getEntry.CreatedAt)

}

func TestListEntry(t *testing.T) {
	account := createRandomAccount(t)

	for i := 0; i < 10; i++ {
		args := CreateEntryParams{
			AccountID: account.ID,
			Amount:    util.RandomMoney(),
		}
		testQueries.CreateEntry(context.Background(), args)
	}

	listEntries, err := testQueries.ListEntries(context.Background(), ListEntriesParams{
		AccountID: account.ID,
		Limit:     10,
	})

	require.NotEmpty(t, listEntries)
	require.NoError(t, err)

	require.Len(t, listEntries, 10)

}
