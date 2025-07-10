package db

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_store_TransferTx(t *testing.T) {
	account1 := CreateAccount(t)
	account2 := CreateAccount(t)
	ctx := context.Background()
	amount := int64(10)
	n := 5
	err := make(chan error)
	result := make(chan resultTransferTx)
	store := *NewStore(testDB)
	for i := 0; i < n; i++ {
		go func() {
			arg := TransferTxParams{
				FromAccountID: account1.ID,
				ToAccountID:   account2.ID,
				Amount:        amount,
			}
			res, e := store.TransferTx(ctx, arg)
			result <- res
			err <- e
		}()
	}

	for i := 0; i < n; i++ {
		res := <-result
		e := <-err
		assert.NoError(t, e)
		assert.Equal(t, res.Transfer.FromAccountID, account1.ID)
		assert.Equal(t, res.Transfer.ToAccountID, account2.ID)
		assert.Equal(t, res.Transfer.Amount, amount)
	}
}
