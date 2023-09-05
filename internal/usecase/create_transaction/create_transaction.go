package create_transaction

import (
	"github.com/danyukod/wallet-core-go/internal/entity"
	"github.com/danyukod/wallet-core-go/internal/gateway"
)

type CreateTransactionInputDTO struct {
	AccountIdFrom string
	AccountIdTo   string
	Amount        float64
}

type CreateTransactionOutputDTO struct {
	ID string
}

type CreateTransactionUseCase interface {
	Execute(input *CreateTransactionInputDTO) (*CreateTransactionOutputDTO, error)
}

type CreateTransactionInteractor struct {
	gateway.TransactionGateway
	gateway.AccountGateway
}

func NewCreateTransactionInteractor(transactionGateway gateway.TransactionGateway, accountGateway gateway.AccountGateway) *CreateTransactionInteractor {
	return &CreateTransactionInteractor{
		transactionGateway,
		accountGateway,
	}
}

func (i *CreateTransactionInteractor) Execute(input *CreateTransactionInputDTO) (*CreateTransactionOutputDTO, error) {
	accountFrom, err := i.AccountGateway.FindById(input.AccountIdFrom)
	if err != nil {
		return nil, err
	}
	accountTo, err := i.AccountGateway.FindById(input.AccountIdTo)
	if err != nil {
		return nil, err
	}

	transaction, err := entity.NewTransaction(accountFrom, accountTo, input.Amount)
	if err != nil {
		return nil, err
	}

	err = i.TransactionGateway.Create(transaction)
	if err != nil {
		return nil, err
	}

	return &CreateTransactionOutputDTO{
		ID: transaction.ID,
	}, nil
}
