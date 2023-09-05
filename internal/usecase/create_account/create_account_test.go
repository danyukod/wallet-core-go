package usecase

import (
	"github.com/danyukod/wallet-core-go/internal/entity"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
)

type AccountGatewayMock struct {
	mock.Mock
}

func (m *AccountGatewayMock) Save(account *entity.Account) error {
	args := m.Called(account)
	return args.Error(0)
}

func (m *AccountGatewayMock) FindById(id string) (*entity.Account, error) {
	args := m.Called(id)
	return args.Get(0).(*entity.Account), args.Error(1)
}

type ClientGatewayMock struct {
	mock.Mock
}

func (m *ClientGatewayMock) Get(id string) (*entity.Client, error) {
	args := m.Called(id)
	return args.Get(0).(*entity.Client), args.Error(1)
}

func (m *ClientGatewayMock) Save(client *entity.Client) error {
	args := m.Called(client)
	return args.Error(0)
}

func TestCreateAccountInteract_Execute(t *testing.T) {
	client, _ := entity.NewClient("John Doe", "jd@email.com")
	cm := &ClientGatewayMock{}
	cm.On("Get", client.ID).Return(client, nil)

	am := &AccountGatewayMock{}
	am.On("Save", mock.Anything).Return(nil)

	i := NewCreateAccountInteract(am, cm)

	input := &CreateAccountInputDTO{ClientID: client.ID}

	output, err := i.Execute(input)
	assert.NotNil(t, output)
	assert.Nil(t, err)
	cm.AssertExpectations(t)
	am.AssertExpectations(t)
	cm.AssertNumberOfCalls(t, "Get", 1)
	am.AssertNumberOfCalls(t, "Save", 1)
}
