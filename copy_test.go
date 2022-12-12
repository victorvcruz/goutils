package goutils

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

type origin struct {
	Test1 string
	test2 int
	test3 bool
	test4 float64
}

type toCopy struct {
	Test1 string
	test2 int
	test3 bool
}

func TestCopy(t *testing.T) {
	t.Run("copy", func(t *testing.T) {
		originTest := origin{}
		toCopyTest := toCopy{
			Test1: "copy",
			test2: 2,
			test3: true,
		}
		test := Copy(originTest, toCopyTest).(origin)
		originExpect := origin{
			Test1: "copy",
			test2: 2,
			test3: true,
			test4: 0,
		}
		assert.Equal(t, originExpect, test)
	})
}
