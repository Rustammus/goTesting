package task7

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestAddDuration(t *testing.T) {

	parseOrFail := func(s string) time.Time {
		timeT, err := time.Parse(time.DateTime, s)
		assert.NoError(t, err, "parse time error")
		return timeT
	}

	testsCases := []struct {
		Name          string
		InputTime     time.Time
		InputDuration time.Duration
		ExpectTime    time.Time
	}{
		{
			Name:          "add_1_hour",
			InputTime:     parseOrFail("2020-02-20 12:03:10"),
			InputDuration: time.Hour,
			ExpectTime:    parseOrFail("2020-02-20 13:03:10"),
		},
		{
			Name:          "add_10_seconds",
			InputTime:     parseOrFail("2020-02-20 12:03:10"),
			InputDuration: 10 * time.Second,
			ExpectTime:    parseOrFail("2020-02-20 12:03:20"),
		},
	}

	for _, tt := range testsCases {
		t.Run(tt.Name, func(t *testing.T) {
			result := AddDuration(tt.InputTime, tt.InputDuration)
			assert.Equal(t, tt.ExpectTime, result)
		})
	}
}
