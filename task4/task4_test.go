package task4

import (
	"math"
	"testing"
	"testing/quick"
)

func TestSquareRootQuick(t *testing.T) {
	f := func(x float64) bool {
		if x < 0 {
			return math.IsNaN(SquareRoot(x))
		}
		return math.Abs(math.Sqrt(x)-SquareRoot(x)) < 1e-9
	}
	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}
