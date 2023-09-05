package gateway

import "github.com/danyukod/wallet-core-go/internal/entity"

type TransactionGateway interface {
	Create(transaction *entity.Transaction) error
}
