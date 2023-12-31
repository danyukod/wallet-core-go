package entity

import (
	"errors"
	"github.com/google/uuid"
	"time"
)

type Transaction struct {
	ID          string
	AccountFrom *Account
	AccountTo   *Account
	Amount      int
	CreatedAt   time.Time
}

func NewTransaction(accountFrom, accountTo *Account, amount int) (*Transaction, error) {
	transaction := &Transaction{
		ID:          uuid.New().String(),
		AccountFrom: accountFrom,
		AccountTo:   accountTo,
		Amount:      amount,
		CreatedAt:   time.Now(),
	}
	err := transaction.Validate()
	if err != nil {
		return nil, err
	}
	transaction.Commit()
	return transaction, nil
}

func (t *Transaction) Commit() {
	t.AccountFrom.Debit(t.Amount)
	t.AccountTo.Credit(t.Amount)
}

func (t *Transaction) Validate() error {
	if t.AccountFrom == nil {
		return errors.New("invalid account from")
	}
	if t.AccountTo == nil {
		return errors.New("invalid account to")
	}
	if t.AccountFrom.ID == t.AccountTo.ID {
		return errors.New("account from and account to must be different")
	}
	if t.Amount <= 0 {
		return errors.New("invalid amount")
	}
	if t.AccountFrom.Balance < t.Amount {
		return errors.New("insufficient funds")
	}
	return nil
}
