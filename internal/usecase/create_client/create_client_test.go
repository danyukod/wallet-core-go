package usecase

import (
	"github.com/danyukod/wallet-core-go/internal/usecase/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
)

func TestCreateClientInteractor_Execute(t *testing.T) {
	m := &mocks.ClientGatewayMock{}
	m.On("Save", mock.Anything).Return(nil)

	i := NewCreateClientInteract(m)

	input := CreateClientInputDTO{
		Name:  "John Doe",
		Email: "jd@email.com",
	}

	output, err := i.Execute(input)

	assert.NotNil(t, output)
	assert.Nil(t, err)
	assert.Equal(t, input.Name, output.Name)
	assert.Equal(t, input.Email, output.Email)
	m.AssertExpectations(t)
	m.AssertNumberOfCalls(t, "Save", 1)
}
