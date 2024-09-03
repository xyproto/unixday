package unixday

import (
	"fmt"
	"time"
)

const (
	secondsPerDay = 86400
	version       = "1.1.0"
)

// UnixDayToDate converts a UNIX day to a formatted date string.
func UnixDayToDate(ud int64) string {
	t := time.Unix(ud*secondsPerDay, 0).UTC()
	return t.Format("2006-01-02")
}

// CurrentUnixDay returns the current UNIX day number.
func CurrentUnixDay() int64 {
	return time.Now().UTC().Unix() / secondsPerDay
}

// Version returns the version of the package.
func Version() string {
	return version
}

// DateStringToUnixDay converts a date string to a UNIX day.
func DateStringToUnixDay(dateString string) (int64, error) {
	formats := []string{"2006-01-02", "02.01.2006", "01/02/2006"}

	for _, format := range formats {
		if t, err := time.Parse(format, dateString); err == nil {
			return t.UTC().Truncate(24*time.Hour).Unix() / secondsPerDay, nil
		}
	}

	return 0, fmt.Errorf("invalid date format: %s", dateString)
}
