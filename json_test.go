package goutils

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestToJson(t *testing.T) {
	t.Run("ok", func(t *testing.T) {
		structure := map[string]string{"value": "structure"}
		value := ToJson(structure)
		valueExpect := `{"value":"structure"}`
		assert.Equal(t, valueExpect, value)
	})
	t.Run("not ok", func(t *testing.T) {
		structure := make(chan bool)
		value := ToJson(structure)
		valueExpect := ""
		assert.Equal(t, valueExpect, value)
	})
}
