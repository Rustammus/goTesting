package task7

import "time"

func AddDuration(t time.Time, d time.Duration) time.Time {
	return t.Add(d)
}
