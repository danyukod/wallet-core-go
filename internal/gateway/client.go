package gateway

import (
	"github.com/danyukod/wallet-core-go/internal/database"
	"github.com/danyukod/wallet-core-go/internal/entity"
)

type ClientGateway interface {
	Get(id string) (*entity.Client, error)
	Save(client *entity.Client) error
}

type ClientRepository struct {
	database.ClientDB
}

func NewClientGateway(clientDB database.ClientDB) ClientGateway {
	return &ClientRepository{clientDB}
}

func (c ClientRepository) Get(id string) (*entity.Client, error) {
	client, err := c.GetClient(id)
	if err != nil {
		return nil, err
	}
	return client, nil
}

func (c ClientRepository) Save(client *entity.Client) error {
	err := c.SaveClient(client)
	if err != nil {
		return err
	}
	return nil
}
