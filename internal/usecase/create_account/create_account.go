package usecase

import (
	"github.com/danyukod/wallet-core-go/internal/entity"
	"github.com/danyukod/wallet-core-go/internal/gateway"
)

type CreateAccountInputDTO struct {
	ClientID string
}

type CreateAccountOutputDTO struct {
	ID string
}

type CreateAccountUseCase interface {
	Execute(input *CreateAccountInputDTO) (*CreateAccountOutputDTO, error)
}

type CreateAccountInteract struct {
	gateway.AccountGateway
	gateway.ClientGateway
}

func NewCreateAccountInteract(accountGateway gateway.AccountGateway, clientGateway gateway.ClientGateway) *CreateAccountInteract {
	return &CreateAccountInteract{accountGateway, clientGateway}
}

func (i *CreateAccountInteract) Execute(input *CreateAccountInputDTO) (*CreateAccountOutputDTO, error) {
	client, err := i.ClientGateway.Get(input.ClientID)
	if err != nil {
		return nil, err
	}

	account := entity.NewAccount(client)
	err = i.AccountGateway.Save(account)
	if err != nil {
		return nil, err
	}

	return &CreateAccountOutputDTO{ID: account.ID}, nil

}
