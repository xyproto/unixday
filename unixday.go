package unixday

import (
	"fmt"
	"time"
)

// UNIXDay represents the number of days since 1970
type UNIXDay int64

// UNIX time does not take leap seconds into account. Neither does UNIX Days.
const secondsPerDay = 3600 * 24

// String returns the current UNIXDay as a number, in the form of a string
func (ud UNIXDay) String() string {
	return fmt.Sprintf("%d", ud)
}

// Date returns the current UNIXDay as a formatted date string
func (ud UNIXDay) Date() string {
	return time.Unix(int64(ud)*secondsPerDay, 0).UTC().Format("02.01.2006")
}

// Time returns the current UNIXDay as a time.Time
func (ud UNIXDay) Time() time.Time {
	return time.Unix(int64(ud)*secondsPerDay, 0)
}

// Now returns the current UNIXDay
func Now() UNIXDay {
	return UNIXDay(time.Now().Unix() / secondsPerDay)
}

// At returns the UNIXDay at the given point in time
func At(t time.Time) UNIXDay {
	return UNIXDay(t.Unix() / secondsPerDay)
}
