package util

import (
	"testing"
	"time"
)

func TestDuration(t *testing.T) {
	now := time.Now()
	defer func() {
		duration := Duration(now)
		if duration.Nanoseconds() > 0 {
			t.Logf("Execution time is %d ns.", duration.Nanoseconds())
		} else {
			t.Errorf("Something went wrong, duration is %d ns.", duration.Nanoseconds())
		}
	}()

	var factorial int64 = 1
	for i := 1; i <= 100; i++ {
		factorial = factorial * int64(i)
	}
}
