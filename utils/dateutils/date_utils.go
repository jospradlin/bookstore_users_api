package dateutils

import (
	"time"
)

const (
	apiDateLayout = "2006-01-02T15:04:05Z"
)

// GetNow returns the current time in UTC as a Time structure
func GetNow() time.Time {
	return time.Now().UTC()
}

// GetNowString returns the current time in UTC as a string
func GetNowString() string {
	return GetNow() .Format(apiDateLayout)
}