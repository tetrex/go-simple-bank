package db

import (
	"context"
	"database/sql"
	"fmt"
)

type Store struct {
	*Queries
	db *sql.DB
}

// NewStore creates a new store
func NewStore(db *sql.DB) *Store {
	return &Store{
		db:      db,
		Queries: New(db),
	}
}

// ExecTx executes a function within a database transaction
func (store *Store) execTx(ctx context.Context, fn func(*Queries) error) error {
	tx, err := store.db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}

	q := New(tx)
	err = fn(q)
	if err != nil {
		if rbErr := tx.Rollback(); rbErr != nil {
			return fmt.Errorf("tx err: %v, rb err: %v", err, rbErr)
		}
		return err
	}

	return tx.Commit()
}

type TransferTxParams struct {
	FromAccountId int64 `json:"from_account_id"`
	ToAccountId   int64 `json:"to_account_id"`
	Amount        int64 `json:"amount"`
}
type TransferTxResult struct {
	Transfer    Transfer `json:"transfer"`
	FromAccount Account  `json:"from_account"`
	ToAccount   Account  `json:"to_account"`
	FromEntry   Entry    `json:"from_entry"`
	ToEntry     Entry    `json:"to_entry"`
}

// Transfertx performs money transfer from acc1 -> acc2
// it creates a transfer record , add account entries , and updated account balances in single db transaction
func (store *Store) TransferTx(ctx context.Context, args TransferTxParams) (TransferTxResult, error) {

	var result TransferTxResult

	store.execTx(ctx, func(q *Queries) error {
		/*
			transfer Amount from account1 To account2
			1 : create entry (entry) of -amount for account1
			2 : update account1 for -amount
			3 : create entry of +amount for account2
			4 : update account2 for +amount
			5 : create transfer entry for account1 and account2 for amount
		*/

		fEntry, err := q.CreateEntry(ctx, CreateEntryParams{
			AccountID: args.FromAccountId,
			Amount:    args.Amount,
		})
		if err != nil {
			return err
		}
		result.FromEntry = fEntry

		// -------------------
		fromAccount, err := q.GetAccount(ctx, args.FromAccountId)
		if err != nil {
			return err
		}
		fromAccountUpdateArgs := UpdateAccountParams{
			ID:      args.FromAccountId,
			Balance: fromAccount.Balance - args.Amount,
		}

		updatedFromAccount, err := q.UpdateAccount(ctx, fromAccountUpdateArgs)
		if err != nil {
			return err
		}
		result.FromAccount = updatedFromAccount
		// -------------------
		toAccountEntry := CreateEntryParams{
			AccountID: args.ToAccountId,
			Amount:    args.Amount,
		}
		tEntry, err := q.CreateEntry(ctx, toAccountEntry)
		if err != nil {
			return err
		}
		result.ToEntry = tEntry
		// -------------------
		toAccount, err := q.GetAccount(ctx, args.ToAccountId)
		if err != nil {
			return err
		}
		toAccountUpdateArgs := UpdateAccountParams{
			ID:      args.ToAccountId,
			Balance: toAccount.Balance + args.Amount,
		}

		updatedToAccount, err := q.UpdateAccount(ctx, toAccountUpdateArgs)
		if err != nil {
			return err
		}
		result.ToAccount = updatedToAccount
		// -------------------
		transferEntryArgs := CreateTransferParams{
			FromAccountID: args.FromAccountId,
			ToAccountID:   args.ToAccountId,
			Amount:        args.Amount,
		}

		transfer, err := q.CreateTransfer(context.Background(), transferEntryArgs)
		if err != nil {
			return err
		}
		result.Transfer = transfer
		return nil
	})

	return result, nil
}
