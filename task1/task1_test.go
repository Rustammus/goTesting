package task1

import "testing"

func TestSumInt(t *testing.T) {

	testsCases := []struct {
		Name   string
		A      int
		B      int
		Expect int
	}{
		{"1+2=3", 1, 2, 3},
		{"3+2=5", 3, 2, 5},
		{"10-2=8", 10, -2, 8},
	}

	for _, tt := range testsCases {
		t.Run(tt.Name, func(t *testing.T) {
			result := SumInt(tt.A, tt.B)
			if result != tt.Expect {
				t.Errorf("got %d, want %d", result, tt.Expect)
			}
		})
	}
}
