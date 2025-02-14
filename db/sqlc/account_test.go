package db

import (
	"context"
	"database/sql"
	"testing"
	"time"

	"github.com/mitchxxx/simplebank/util"
	"github.com/stretchr/testify/require"
)

//Create a random accounts for Unit test
func createRandomAccount(t *testing.T) Account {
	user := createRandomUser(t)
	arg := CreateAccountParams {
		Owner: user.Username,
		Balance: util.RandomMoney(),
		Currency: util.RandomCurrency(),
	}
	account, err := testQueries.CreateAccount(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, account)

	require.Equal(t, arg.Owner, account.Owner)
	require.Equal(t, arg.Balance, account.Balance)
	require.Equal(t, arg.Currency, account.Currency)

	require.NotZero(t, account.ID)
	require.NotZero(t, account.CreatedAt)

	return account
}


// Unit Test for CreateAccount function
func TestCreateAccount(t *testing.T) {
	createRandomAccount(t)
}

// Unit Test for GetAccount function
func TestGetAccount(t *testing.T) {
	account1 := createRandomAccount(t)
	account2, err := testQueries.GetAccount(context.Background(), account1.ID)

	require.NotEmpty(t, account2)
	require.NoError(t, err)

	require.Equal(t, account1.ID, account2.ID)
	require.Equal(t, account1.Owner, account2.Owner)
	require.Equal(t, account1.Balance, account2.Balance)
	require.Equal(t, account1.CreatedAt, account2.CreatedAt)
	require.WithinDuration(t, account1.CreatedAt, account2.CreatedAt, time.Second)
}

// Unit test for UpdateAccount function
func TestUpdateAccount( t *testing.T) {
	account1 := createRandomAccount(t)
	arg := UpdateAccountParams {
		ID: account1.ID,
		Balance: util.RandomMoney(),
	}
	account2, err := testQueries.UpdateAccount(context.Background(), arg)
	require.NotEmpty(t, account2)
	require.NoError(t, err)

	require.Equal(t, account1.ID, account2.ID)
	require.Equal(t, account1.Owner, account2.Owner)
	require.Equal(t, arg.Balance, account2.Balance)
	require.Equal(t, account1.CreatedAt, account2.CreatedAt)
	require.WithinDuration(t, account1.CreatedAt, account2.CreatedAt, time.Second)
}

// Unit test for DeleteAccount function

func TestDeleteAccout(t *testing.T) {
	account1 := createRandomAccount(t)
	err := testQueries.DeleteAccount(context.Background(), account1.ID)
	require.NoError(t, err)

	account2, err := testQueries.GetAccount(context.Background(), account1.ID)
	require.Error(t, err)
	require.EqualError(t, err, sql.ErrNoRows.Error())
	require.Empty(t, account2)
}

// Unit test for ListAccount function
func TestListAccount(t *testing.T){
	for i := 0 ; i < 10; i++ {
		createRandomAccount(t)
	}
	arg := ListAccountsParams {
		Limit: 5,
		Offset: 5,
	}

	accounts, err := testQueries.ListAccounts(context.Background(), arg)
	require.NoError(t, err)
	require.Len(t, accounts, 5)

	for _, account := range listAccounts {
		require.NotEmpty(t, account)
	}
	
}