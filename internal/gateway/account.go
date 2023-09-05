package gateway

import "github.com/danyukod/wallet-core-go/internal/entity"

type AccountGateway interface {
	Save(account *entity.Account) error
	FindById(id string) (*entity.Account, error)
}
