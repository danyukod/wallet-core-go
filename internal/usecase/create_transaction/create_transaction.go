package usecase

import (
	"context"
	"github.com/danyukod/chave-pix-utils/pkg/events"
	"github.com/danyukod/chave-pix-utils/pkg/uow"
	"github.com/danyukod/wallet-core-go/internal/entity"
	"github.com/danyukod/wallet-core-go/internal/gateway"
)

type CreateTransactionInputDTO struct {
	AccountIdFrom string `json:"account_id_from"`
	AccountIdTo   string `json:"account_id_to"`
	Amount        int    `json:"amount"`
}

type CreateTransactionOutputDTO struct {
	ID            string `json:"id"`
	AccountIDFrom string `json:"account_id_from"`
	AccountIDTo   string `json:"account_id_to"`
	Amount        int    `json:"amount"`
}

type BalanceUpdatedOutputDTO struct {
	AccountIDFrom        string `json:"account_id_from"`
	AccountIDTo          string `json:"account_id_to"`
	BalanceAccountIDFrom int    `json:"balance_account_id_from"`
	BalanceAccountIDTo   int    `json:"balance_account_id_to"`
}

type CreateTransactionUseCase interface {
	Execute(ctx context.Context, input *CreateTransactionInputDTO) (*CreateTransactionOutputDTO, error)
}

type CreateTransactionInteractor struct {
	uow.UowInterface
	events.EventDispatcherInterface
	TransactionCreated events.EventInterface
	BalanceUpdated     events.EventInterface
}

func NewCreateTransactionInteract(
	uow uow.UowInterface,
	eventDispatcher events.EventDispatcherInterface,
	transactionCreated events.EventInterface,
	balanceUpdated events.EventInterface) *CreateTransactionInteractor {
	return &CreateTransactionInteractor{
		uow,
		eventDispatcher,
		transactionCreated,
		balanceUpdated,
	}
}

func (i *CreateTransactionInteractor) Execute(ctx context.Context, input *CreateTransactionInputDTO) (*CreateTransactionOutputDTO, error) {
	output := &CreateTransactionOutputDTO{}
	balanceUpdatedOutput := &BalanceUpdatedOutputDTO{}
	err := i.Do(ctx, func(_ *uow.Uow) error {
		accountRepository := i.getAccountRepository(ctx)
		transactionRepository := i.getTransactionRepository(ctx)

		accountFrom, err := accountRepository.FindById(input.AccountIdFrom)
		if err != nil {
			return err
		}
		accountTo, err := accountRepository.FindById(input.AccountIdTo)
		if err != nil {
			return err
		}

		transaction, err := entity.NewTransaction(accountFrom, accountTo, input.Amount)
		if err != nil {
			return err
		}

		err = accountRepository.UpdateBalance(accountFrom)
		if err != nil {
			return err
		}

		err = accountRepository.UpdateBalance(accountTo)
		if err != nil {
			return err
		}

		err = transactionRepository.Create(transaction)
		if err != nil {
			return err
		}

		output.ID = transaction.ID
		output.AccountIDFrom = input.AccountIdFrom
		output.AccountIDTo = input.AccountIdTo
		output.Amount = input.Amount

		balanceUpdatedOutput.AccountIDFrom = input.AccountIdFrom
		balanceUpdatedOutput.AccountIDTo = input.AccountIdTo
		balanceUpdatedOutput.BalanceAccountIDFrom = accountFrom.Balance
		balanceUpdatedOutput.BalanceAccountIDTo = accountTo.Balance
		return nil
	})
	if err != nil {
		return nil, err
	}

	i.TransactionCreated.SetPayload(output)
	i.EventDispatcherInterface.Dispatch(i.TransactionCreated)

	i.BalanceUpdated.SetPayload(balanceUpdatedOutput)
	i.EventDispatcherInterface.Dispatch(i.BalanceUpdated)

	return output, nil
}

func (i *CreateTransactionInteractor) getAccountRepository(ctx context.Context) gateway.AccountGateway {
	repo, err := i.GetRepository(ctx, "AccountDB")
	if err != nil {
		panic(err)
	}
	return repo.(gateway.AccountGateway)
}

func (i *CreateTransactionInteractor) getTransactionRepository(ctx context.Context) gateway.TransactionGateway {
	repo, err := i.GetRepository(ctx, "TransactionDB")
	if err != nil {
		panic(err)
	}
	return repo.(gateway.TransactionGateway)
}
