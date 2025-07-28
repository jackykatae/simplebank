package db

import (
	"context"
	"testing"
	"time"

	"github.com/jackykatae/simplebank/util"
	"github.com/stretchr/testify/require"
)

func createRandomTransfer(t *testing.T, accountFrom, accountTo Account) Transfer {
	arg := CreateTranferParams{
		FromAccountID: accountFrom.ID,
		ToAccountID:   accountTo.ID,
		Amount:        util.RandomMoney(),
	}

	tranfer, err := testQueries.CreateTranfer(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, tranfer)

	require.Equal(t, arg.FromAccountID, tranfer.FromAccountID)
	require.Equal(t, arg.ToAccountID, tranfer.ToAccountID)
	require.Equal(t, arg.Amount, tranfer.Amount)

	require.NotZero(t, tranfer.ID)
	require.NotZero(t, tranfer.CreatedAt)

	return tranfer
}

func TestCreateTranfer(t *testing.T) {
	accout1 := createRandomAccount(t)
	accout2 := createRandomAccount(t)
	createRandomTransfer(t, accout1, accout2)
}

func TestGetTranfer(t *testing.T) {
	accout1 := createRandomAccount(t)
	accout2 := createRandomAccount(t)
	tranfer1 := createRandomTransfer(t, accout1, accout2)

	tranfer2, err := testQueries.GetTranfer(context.Background(), tranfer1.ID)
	require.NoError(t, err)
	require.NotEmpty(t, tranfer2)

	require.Equal(t, tranfer1.FromAccountID, tranfer2.FromAccountID)
	require.Equal(t, tranfer1.ToAccountID, tranfer2.ToAccountID)
	require.Equal(t, tranfer1.ID, tranfer2.ID)
	require.Equal(t, tranfer1.Amount, tranfer2.Amount)
	require.WithinDuration(t, tranfer1.CreatedAt.Time, tranfer2.CreatedAt.Time, time.Second)
}

func TestListTransfer(t *testing.T) {
	account1 := createRandomAccount(t)
	account2 := createRandomAccount(t)

	for i := 0; i < 5; i++ {
		createRandomTransfer(t, account1, account2)
		createRandomTransfer(t, account2, account1)
	}

	arg := ListTranfersParams{
		FromAccountID: account1.ID,
		ToAccountID:   account1.ID,
		Limit:         5,
		Offset:        5,
	}

	transfers, err := testQueries.ListTranfers(context.Background(), arg)
	require.NoError(t, err)
	require.Len(t, transfers, 5)

	for _, transfer := range transfers {
		require.NotEmpty(t, transfer)
		require.True(t, transfer.FromAccountID == account1.ID || transfer.ToAccountID == account1.ID)
	}
}
