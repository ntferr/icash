// tag:unit
package entities

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var (
	bankSuccess = Bank{
		Name: "Itau",
		Code: "341",
	}

	bankCodeEmpty = Bank{
		Name: "Itau",
		Code: "",
	}

	bankCodeOver3 = Bank{
		Name: "Itau",
		Code: "1234",
	}
)

func TestBankEntities(t *testing.T) {
	t.Run("when you create the correct object, should works", func(t *testing.T) {
		actual := bankSuccess
		err := actual.Validate()
		assert.Nil(t, err)
		assert.Equal(t, bankSuccess, actual)
	})

	t.Run("when code empty, should fails", func(t *testing.T) {
		actual := bankCodeEmpty
		err := actual.Validate()
		assert.Equal(t, err.Error(), "bank code is required")
	})

	t.Run("when code is over 3 characters, should fails", func(t *testing.T) {
		actual := bankCodeOver3
		err := actual.Validate()
		assert.Equal(t, err.Error(), "insert a bank code valid")
	})
}
