package utils

import (
	"github.com/stretchr/testify/assert"
	"testing"
)


var tableTest = []struct{
	input string
	expected bool
}{
	{
		"test",
		true,
	},
	{
		"test 1",
		false,
	},
}

func TestContains(t *testing.T) {
	input := []string{"test"}

	for _, tt := range tableTest {
		t.Run("should check if slice contains string", func(t *testing.T) {
			got := Contains(input, tt.input)
			assert.Equal(t, tt.expected,  got)
		})
	}
}


