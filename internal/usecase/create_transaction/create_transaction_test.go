package usecase

import (
	"context"
	"github.com/danyukod/chave-pix-utils/pkg/events"
	"github.com/danyukod/wallet-core-go/internal/entity"
	"github.com/danyukod/wallet-core-go/internal/event"
	"github.com/danyukod/wallet-core-go/internal/usecase/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
)

func TestCreateTransactionInteractor_Execute(t *testing.T) {
	clientFrom, _ := entity.NewClient("client1", "client1@email.com")
	clientTo, _ := entity.NewClient("client2", "client2@email.com")
	accountFrom := entity.NewAccount(clientFrom)
	accountTo := entity.NewAccount(clientTo)

	accountFrom.Credit(1000)
	accountTo.Credit(100)

	mockUow := &mocks.UowMock{}
	mockUow.On("Do", mock.Anything, mock.Anything).Return(nil)

	inputDto := &CreateTransactionInputDTO{
		AccountIdFrom: accountFrom.ID,
		AccountIdTo:   accountTo.ID,
		Amount:        100,
	}

	dispatcher := events.NewEventDispatcher()
	event := event.NewTransactionCreated()
	ctx := context.Background()

	i := NewCreateTransactionInteract(mockUow, dispatcher, event)

	output, err := i.Execute(ctx, inputDto)

	assert.NotNil(t, output)
	assert.Nil(t, err)
	mockUow.AssertExpectations(t)
	mockUow.AssertNumberOfCalls(t, "Do", 1)
}
