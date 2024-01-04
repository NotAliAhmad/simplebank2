package db

import (
	"context"
	"testing"
	"time"

	"github.com/notaliahmad/simplebank2/util"
	"github.com/stretchr/testify/require"
)

func TestCreateTransfer(t *testing.T) {
	from_account_id := createRandomAccount(t)
	to_account_id := createRandomAccount(t)

	args := CreateTransferParams{
		FromAccountID: from_account_id.ID,
		ToAccountID:   to_account_id.ID,
		Amount:        util.RandomMoney(),
	}

	transfer, err := testQueries.CreateTransfer(context.Background(), args)

	require.NoError(t, err)
	require.NotEmpty(t, transfer)

	require.NotZero(t, transfer.ID)
	require.Equal(t, from_account_id.ID, args.FromAccountID)
	require.Equal(t, to_account_id.ID, args.ToAccountID)
	require.Equal(t, args.Amount, transfer.Amount)
	require.NotEmpty(t, transfer.CreatedAt)
}

func TestGetTransfer(t *testing.T) {
	from_account_id := createRandomAccount(t)
	to_account_id := createRandomAccount(t)

	args := CreateTransferParams{
		FromAccountID: from_account_id.ID,
		ToAccountID:   to_account_id.ID,
		Amount:        util.RandomMoney(),
	}

	transfer, _ := testQueries.CreateTransfer(context.Background(), args)

	getTransfer, err := testQueries.GetTransfer(context.Background(), transfer.ID)

	require.NoError(t, err)
	require.NotEmpty(t, getTransfer)

	require.Equal(t, transfer.ID,getTransfer.ID)
	require.Equal(t, transfer.FromAccountID, getTransfer.FromAccountID)
	require.Equal(t, transfer.ToAccountID, getTransfer.ToAccountID)
	require.Equal(t, transfer.Amount, getTransfer.Amount)
	require.WithinDuration(t, transfer.CreatedAt, getTransfer.CreatedAt, time.Second)
}

func TestListTransfer(t *testing.T) {
	from_account_id := createRandomAccount(t)
	to_account_id := createRandomAccount(t)

	for i := 0; i < 10; i++ {
		args := CreateTransferParams{
			FromAccountID: from_account_id.ID,
			ToAccountID:   to_account_id.ID,
			Amount:        util.RandomMoney(),
		}
		testQueries.CreateTransfer(context.Background(), args)
	}

	getTransfers, err := testQueries.ListTransfers(context.Background(), ListTransfersParams{
		FromAccountID: from_account_id.ID,
		ToAccountID:   to_account_id.ID,
		Limit:         10,
	})

	require.NoError(t, err)
	require.NotEmpty(t, getTransfers)

	require.Len(t, getTransfers, 10)

	for _, transfer := range getTransfers {
		require.NotEmpty(t, transfer)
		require.True(t, transfer.FromAccountID == from_account_id.ID || transfer.ToAccountID == to_account_id.ID)
	}
}
