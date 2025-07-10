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
		result.Transfer, err = s.CreateTransfer(ctx, CreateTransferParams{FromAccountID: arg.FromAccountID, ToAccountID: arg.ToAccountID, Amount: arg.Amount})
		if err != nil {
			return err
		}
		result.FromEntry, err = s.CreateEntries(ctx, CreateEntriesParams{AccountID: arg.FromAccountID, Amount: -arg.Amount})
		if err != nil {
			return err
		}
		result.ToEntry, err = s.CreateEntries(ctx, CreateEntriesParams{AccountID: arg.ToAccountID, Amount: arg.Amount})
		if err != nil {
			return err
		}

		//TODO//
		return nil
	})

	return result, err
}
