package db

import (
	"context"
	"database/sql"
)

type store struct {
	*Queries
	db *sql.DB
}

func NewStore(db *sql.DB) *store {
	return &store{
		Queries: New(db),
		db:      db,
	}
}

func (s *store) execTx(ctx context.Context, fn func(*Queries) error) error {
	tx, err := s.db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}

	q := New(tx)

	if err := fn(q); err != nil {
		if rbErr := tx.Rollback(); rbErr != nil {
			return rbErr
		}
		return err
	}

	return tx.Commit()
}

// it should be used to transfer money from one account to another
// step 1 create transfer record
// step 2 create entries for both accounts
// step 2 update both accounts
type TransferTxParams struct {
	FromAccountID int64 `json:"from_account_id"`
	ToAccountID   int64 `json:"to_account_id"`
	Amount        int64 `json:"amount"`
}
type resultTransferTx struct {
	Transfer    Transfer `json:"transfer"`
	FromAccount Account  `json:"from_account"`
	ToAccount   Account  `json:"to_account"`
	FromEntry   Entry    `json:"from_entry"`
	ToEntry     Entry    `json:"to_entry"`
}

func (s *store) TransferTx(ctx context.Context, arg TransferTxParams) (resultTransferTx, error) {
	var result resultTransferTx
	err := s.execTx(ctx, func(q *Queries) error {
		var err error
		result.Transfer, err = q.CreateTransfer(ctx, CreateTransferParams{FromAccountID: arg.FromAccountID, ToAccountID: arg.ToAccountID, Amount: arg.Amount})
		if err != nil {
			return err
		}
		result.FromEntry, err = q.CreateEntry(ctx, CreateEntryParams{AccountID: arg.FromAccountID, Amount: -arg.Amount})
		if err != nil {
			return err
		}
		result.ToEntry, err = q.CreateEntry(ctx, CreateEntryParams{AccountID: arg.ToAccountID, Amount: arg.Amount})
		if err != nil {
			return err
		}
		if arg.FromAccountID < arg.ToAccountID {
			result.FromAccount, result.ToAccount, err = q.AddToAccount(ctx, arg.FromAccountID, -arg.Amount, arg.ToAccountID, arg.Amount)
		} else {
			result.ToAccount, result.FromAccount, err = q.AddToAccount(ctx, arg.ToAccountID, arg.Amount, arg.FromAccountID, -arg.Amount)
		}

		// //first get accounts for update then update account
		// fromAccount, err := q.GetAccountForUpdate(ctx, arg.FromAccountID)
		// if err != nil {
		// 	return err
		// }
		// result.FromAccount, err = q.UpdateAccount(ctx, UpdateAccountParams{ID: fromAccount.ID, Balance: fromAccount.Balance - arg.Amount})
		// if err != nil {
		// 	return err
		// }

		// //2nd get to account and update it
		// toAccount, err := q.GetAccountForUpdate(ctx, arg.ToAccountID)
		// if err != nil {
		// 	return err
		// }
		// result.ToAccount, err = q.UpdateAccount(ctx, UpdateAccountParams{ID: toAccount.ID, Balance: toAccount.Balance + arg.Amount})
		// if err != nil {
		// 	return err
		// }
		return nil
	})

	return result, err
}

func (q *Queries) AddToAccount(ctx context.Context, account1 int64, amount1 int64, account2 int64, amount2 int64) (Account, Account, error) {
	account1ForUpdate, err := q.GetAccountForUpdate(ctx, account1)
	account2ForUpdate, err := q.GetAccountForUpdate(ctx, account2)
	retAccount1, err := q.UpdateAccount(ctx, UpdateAccountParams{ID: account1ForUpdate.ID, Balance: account1ForUpdate.Balance + amount1})
	retAccount2, err := q.UpdateAccount(ctx, UpdateAccountParams{ID: account2ForUpdate.ID, Balance: account2ForUpdate.Balance + amount2})
	return retAccount1, retAccount2, err
}
