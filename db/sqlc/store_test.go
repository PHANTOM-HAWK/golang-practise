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
	n := 100
	err := make(chan error)
	result := make(chan resultTransferTx)
	store := *NewStore(testDB)
	for i := 0; i < n; i++ {
		go func() {
			arg := TransferTxParams{Amount: amount}
			if i%2 == 0 {
				arg.FromAccountID = account1.ID
				arg.ToAccountID = account2.ID
			} else {
				arg.ToAccountID = account1.ID
				arg.FromAccountID = account2.ID
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
		// assert.Equal(t, res.Transfer.FromAccountID, account1.ID)
		// assert.Equal(t, res.Transfer.ToAccountID, account2.ID)
		// assert.Equal(t, res.Transfer.Amount, amount)
		//the diff in bw each trans should be equal
		//get balance

		diff1 := account1.Balance - res.FromAccount.Balance
		diff2 := res.ToAccount.Balance - account2.Balance
		assert.Equal(t, diff1, diff2)
	}
}
