package usecase

import (
	"github.com/danyukod/wallet-core-go/internal/entity"
	"github.com/danyukod/wallet-core-go/internal/usecase/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
)

func TestCreateAccountInteract_Execute(t *testing.T) {
	client, _ := entity.NewClient("John Doe", "jd@email.com")
	cm := &mocks.ClientGatewayMock{}
	cm.On("Get", client.ID).Return(client, nil)

	am := &mocks.AccountGatewayMock{}
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
