package entity

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCreateAccount(t *testing.T) {
	t.Run("should create a new account", func(t *testing.T) {
		client, _ := NewClient("John Doe", "j@j.com")
		account := NewAccount(client)
		assert.NotNil(t, account)
		assert.Equal(t, client.ID, account.Client.ID)
	})

	t.Run("should create a new account with nil client", func(t *testing.T) {
		account := NewAccount(nil)
		assert.Nil(t, account)
	})
}

func TestCredit(t *testing.T) {
	t.Run("should credit amount", func(t *testing.T) {
		client, _ := NewClient("John Doe", "j@j.com")
		account := NewAccount(client)
		account.Credit(100)
		assert.Equal(t, 100.0, account.Balance)
	})
}

func TestDebit(t *testing.T) {
	t.Run("should debit amount", func(t *testing.T) {
		client, _ := NewClient("John Doe", "j@j.com")
		account := NewAccount(client)
		account.Credit(100)
		account.Debit(50)
		assert.Equal(t, 50.0, account.Balance)
	})
}
