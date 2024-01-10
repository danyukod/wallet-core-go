package gateway

import (
	"github.com/danyukod/wallet-core-go/internal/database"
	"github.com/danyukod/wallet-core-go/internal/entity"
)

type AccountGateway interface {
	Save(account *entity.Account) error
	FindById(id string) (*entity.Account, error)
	UpdateBalance(account *entity.Account) error
}

type AccountRepository struct {
	database.AccountDB
}

func NewAccountGateway(accountDB database.AccountDB) AccountGateway {
	return &AccountRepository{accountDB}
}

func (a AccountRepository) Save(account *entity.Account) error {
	err := a.SaveAccount(account)
	if err != nil {
		return err
	}
	return nil
}

func (a AccountRepository) FindById(id string) (*entity.Account, error) {
	account, err := a.FindByID(id)
	if err != nil {
		return nil, err
	}
	return account, nil
}

func (a AccountRepository) UpdateBalance(account *entity.Account) error {
	err := a.UpdateBalance(account)
	if err != nil {
		return err
	}
	return nil
}
