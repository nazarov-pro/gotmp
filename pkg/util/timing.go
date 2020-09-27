package util

import (
	"time"
)

// Duration - measures duration
func Duration(invocation time.Time) time.Duration {
	return time.Since(invocation)
}
