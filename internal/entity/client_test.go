package entity

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCreateNewClient(t *testing.T) {
	client, err := NewClient("John Doe", "j@j.com")
	assert.Nil(t, err)
	assert.NotNil(t, client)
	assert.Equal(t, "John Doe", client.Name)
	assert.Equal(t, "j@j.com", client.Email)
}

func TestCreateNewClientArgsAreInvalid(t *testing.T) {
	client, err := NewClient("", "")
	assert.NotNil(t, err)
	assert.Nil(t, client)
}

func TestUpdateClient(t *testing.T) {
	client, _ := NewClient("John Doe", "j@j.com")

	err := client.Update("John Doe Updated", "j@j.com")
	assert.Nil(t, err)
	assert.Equal(t, "John Doe Updated", client.Name)
	assert.Equal(t, "j@j.com", client.Email)
}

func TestUpdateClientArgsAreInvalid(t *testing.T) {
	client, _ := NewClient("John Doe", "j@j")
	err := client.Update("John Doe Updated", "")
	assert.NotNil(t, err)
	assert.Error(t, err, "invalid email")

	err = client.Update("", "j@j")
	assert.NotNil(t, err)
	assert.Error(t, err, "invalid name")
}
