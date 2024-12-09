package task3

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSortNumbers(t *testing.T) {
	testsCases := []struct {
		Name   string
		Input  []int
		Expect []int
	}{
		{
			Name:   "Unsorted_Positive",
			Input:  []int{8, 23, 0, 3},
			Expect: []int{0, 3, 8, 23},
		},
		{
			Name:   "Unsorted_All",
			Input:  []int{4, -2, 0, 23, -4},
			Expect: []int{-4, -2, 0, 4, 23},
		},
		{
			Name:   "All_equals",
			Input:  []int{7, 7, 7, 7},
			Expect: []int{7, 7, 7, 7},
		},
	}

	for _, tt := range testsCases {
		t.Run(tt.Name, func(t *testing.T) {
			assert.Equal(t, tt.Expect, SortNumbers(tt.Input))
		})
	}
}
