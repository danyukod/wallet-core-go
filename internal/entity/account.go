package entity

import (
	"github.com/google/uuid"
	"time"
)

type Account struct {
	ID        string
	Client    *Client
	Balance   int
	CreatedAt time.Time
	UpdatedAt time.Time
}

func NewAccount(client *Client) *Account {
	if client == nil {
		return nil
	}
	account := &Account{
		ID:        uuid.New().String(),
		Client:    client,
		Balance:   0,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	return account
}

func (a *Account) Credit(amount int) {
	a.Balance += amount
	a.UpdatedAt = time.Now()
}

func (a *Account) Debit(amount int) {
	a.Balance -= amount
	a.UpdatedAt = time.Now()
}
