package models

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

var testTable = []struct {
	input    *Review
	expected float32
}{
	{
		&Review{
			EntryDate: 1250953369000, //2009
		},
		0.5,
	},
	{
		&Review{
			EntryDate: 1550953369000, //2019
		},
		0.9,
	},
}

func TestReview_ReviewWeight(t *testing.T) {
	for _, tt := range testTable {
		t.Run("should calculate review weight value", func(t *testing.T) {
			got := tt.input.ReviewWeight()
			assert.Equal(t, tt.expected,  got)
		})
	}
}
