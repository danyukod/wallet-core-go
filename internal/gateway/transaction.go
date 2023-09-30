package gateway

import (
	"github.com/danyukod/wallet-core-go/internal/database"
	"github.com/danyukod/wallet-core-go/internal/entity"
)

type TransactionGateway interface {
	Create(transaction *entity.Transaction) error
}

type TransactionRepository struct {
	database.TransactionDB
}

func NewTransactionGateway(transactionDB database.TransactionDB) TransactionGateway {
	return &TransactionRepository{transactionDB}
}

func (t TransactionRepository) Create(transaction *entity.Transaction) error {
	err := t.Create(transaction)
	if err != nil {
		return err
	}
	return nil
}
