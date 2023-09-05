package create_transaction

import (
	"github.com/danyukod/wallet-core-go/internal/entity"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
)

type TransactionGatewayMock struct {
	mock.Mock
}

type AccountGatewayMock struct {
	mock.Mock
}

func (m *AccountGatewayMock) FindById(id string) (*entity.Account, error) {
	args := m.Called(id)
	return args.Get(0).(*entity.Account), args.Error(1)
}

func (m *AccountGatewayMock) Save(account *entity.Account) error {
	args := m.Called(account)
	return args.Error(0)
}

func (m *TransactionGatewayMock) Create(transaction *entity.Transaction) error {
	args := m.Called(transaction)
	return args.Error(0)
}

func TestCreateTransactionInteractor_Execute(t *testing.T) {
	clientFrom, _ := entity.NewClient("client1", "client1@email.com")
	clientTo, _ := entity.NewClient("client2", "client2@email.com")
	accountFrom := entity.NewAccount(clientFrom)
	accountTo := entity.NewAccount(clientTo)

	accountFrom.Credit(1000)
	accountTo.Credit(100)

	ma := &AccountGatewayMock{}

	ma.On("FindById", accountFrom.ID).Return(accountFrom, nil)
	ma.On("FindById", accountTo.ID).Return(accountTo, nil)

	mt := &TransactionGatewayMock{}

	mt.On("Create", mock.Anything).Return(nil)

	inputDto := CreateTransactionInputDTO{
		AccountIdFrom: accountFrom.ID,
		AccountIdTo:   accountTo.ID,
		Amount:        100,
	}

	i := NewCreateTransactionInteractor(mt, ma)

	output, err := i.Execute(&inputDto)

	assert.NotNil(t, output)
	assert.Nil(t, err)
	ma.AssertExpectations(t)
	mt.AssertExpectations(t)
	ma.AssertNumberOfCalls(t, "FindById", 2)
	mt.AssertNumberOfCalls(t, "Create", 1)

}
