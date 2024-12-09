package task9

import (
	"errors"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestErrorProneFunction(t *testing.T) {

	testsCases := []struct {
		Name        string
		Input       int
		ExpectError error
	}{
		{
			Name:        "Positive_no_err",
			Input:       1,
			ExpectError: nil,
		},
		{
			Name:        "Zero_no_err",
			Input:       0,
			ExpectError: nil,
		},
		{
			Name:        "Negative_err",
			Input:       -1,
			ExpectError: errors.New("negative input not allowed"),
		},
	}

	for _, tt := range testsCases {
		t.Run(tt.Name, func(t *testing.T) {
			err := AcceptOnlyPositive(tt.Input)

			// assert.ErrorIs(t, tt.ExpectError, err)
			assert.Equal(t, tt.ExpectError, err)
		})
	}
}
