// tag:unit
package snowflake

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

const snowflakeLength = 19

func TestSnowflake(t *testing.T) {
	t.Run("should return an id when call generate function", func(t *testing.T) {
		id, err := GenerateNew()
		assert.Nil(t, err)

		assert.NotNil(t, id)

		actualLength := len([]rune(*id))

		assert.Equal(t, actualLength, snowflakeLength)
	})
}
