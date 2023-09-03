package usecase

import (
	"github.com/danyukod/wallet-core-go/internal/entity"
	"github.com/danyukod/wallet-core-go/internal/gateway"
	"time"
)

type CreateClientInputDTO struct {
	Name  string
	Email string
}

type CreateClientOutputDTO struct {
	ID        string
	Name      string
	Email     string
	CreatedAt time.Time
	UpdateAt  time.Time
}

type CreateClientUseCase interface {
	Execute(input CreateClientInputDTO) (*CreateClientOutputDTO, error)
}

type CreateClientInteract struct {
	gateway.ClientGateway
}

func NewCreateClientInteract(clientGateway gateway.ClientGateway) *CreateClientInteract {
	return &CreateClientInteract{clientGateway}
}

func (i *CreateClientInteract) Execute(input CreateClientInputDTO) (*CreateClientOutputDTO, error) {
	client, err := entity.NewClient(input.Name, input.Email)
	if err != nil {
		return nil, err
	}

	err = i.Save(client)
	if err != nil {
		return nil, err
	}

	return &CreateClientOutputDTO{
		ID:        client.ID,
		Name:      client.Name,
		Email:     client.Email,
		CreatedAt: client.CreatedAt,
		UpdateAt:  client.UpdatedAt,
	}, nil
}
