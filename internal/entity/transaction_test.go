package entity

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCreateTransaction(t *testing.T) {
	client, _ := NewClient("John Doe", "j@j.com")
	accountFrom := NewAccount(client)
	client2, _ := NewClient("John Doe 2", "j@j.com")
	accountTo := NewAccount(client2)

	accountTo.Credit(100)
	accountFrom.Credit(100)

	transaction, err := NewTransaction(accountFrom, accountTo, 50)
	assert.NotNil(t, transaction)
	assert.Nil(t, err)
	assert.Equal(t, 50.0, accountFrom.Balance)
	assert.Equal(t, 150.0, accountTo.Balance)
}

func TestCreateTransactionWithInsufficientFunds(t *testing.T) {
	client, _ := NewClient("John Doe", "j@j.com")
	accountFrom := NewAccount(client)
	client2, _ := NewClient("John Doe 2", "j@j2.com")
	accountTo := NewAccount(client2)

	accountFrom.Credit(100)
	accountTo.Credit(100)

	transaction, err := NewTransaction(accountFrom, accountTo, 150)
	assert.Nil(t, transaction)
	assert.NotNil(t, err)
	assert.Equal(t, "insufficient funds", err.Error())
	assert.Equal(t, 100.0, accountFrom.Balance)
}
