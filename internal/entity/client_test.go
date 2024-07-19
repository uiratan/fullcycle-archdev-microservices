package entity

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateNewCliente(t *testing.T) {
	client, err := NewClient("Uiratan", "u@u.com")
	assert.Nil(t, err)
	assert.NotNil(t, client)
	assert.Equal(t, "Uiratan", client.Name)
	assert.Equal(t, "u@u.com", client.Email)
}

func TestCreateNewClientWhenArgsAreInvalid(t *testing.T) {
	client, err := NewClient("", "")
	assert.NotNil(t, err)
	assert.Nil(t, client)
}
