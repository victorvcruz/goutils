package goutils

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestStringNullable(t *testing.T) {
	t.Run("str nil", func(t *testing.T) {
		var str interface{} = nil
		assert.Equal(t, "", StringNullable(str))
	})
	t.Run("str not nil", func(t *testing.T) {
		var str interface{} = "1"
		assert.Equal(t, "1", StringNullable(str))
	})

}
