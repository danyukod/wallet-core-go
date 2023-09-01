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
