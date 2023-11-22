package usecase

import (
	"github.com/danyukod/chave-pix-utils/pkg/events"
	"github.com/danyukod/wallet-core-go/internal/entity"
	"github.com/danyukod/wallet-core-go/internal/gateway"
)

type CreateTransactionInputDTO struct {
	AccountIdFrom string `json:"account_id_from"`
	AccountIdTo   string `json:"account_id_to"`
	Amount        int    `json:"amount"`
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
	events.EventDispatcherInterface
	events.EventInterface
}

func NewCreateTransactionInteract(transactionGateway gateway.TransactionGateway,
	accountGateway gateway.AccountGateway,
	eventDispatcher events.EventDispatcherInterface,
	transactionCreated events.EventInterface) *CreateTransactionInteractor {
	return &CreateTransactionInteractor{
		transactionGateway,
		accountGateway,
		eventDispatcher,
		transactionCreated,
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

	output := &CreateTransactionOutputDTO{
		ID: transaction.ID,
	}

	i.EventInterface.SetPayload(output)
	i.EventDispatcherInterface.Dispatch(i.EventInterface)

	return output, nil
}
