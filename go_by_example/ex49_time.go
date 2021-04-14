// go has extensive support for times and durations
package main

import (
	"fmt"
	"time"
)

func main() {
	p := fmt.Println

	// get the current time
	now := time.Now()
	p(now)

	// can build a time struct by providing year, month, day etc
	// times are always associated with a Location (time zone)
	then := time.Date(
		2009, 11, 17, 20, 34, 58, 651378237, time.UTC)
	p(then)

	// access the various components as expected
	p(then.Year())
	p(then.Month())
	p(then.Day())
	p(then.Hour())
	p(then.Minute())
	p(then.Second())
	p(then.Nanosecond())
	p(then.Location())

	// monday-sunday weekday
	p(then.Weekday())

	// compare two times
	p(then.Before(now))
	p(then.After(now))
	p(then.Equal(now))

	// Sub returns a Duration, an interval between two times
	diff := now.Sub(then)
	p(diff)

	// compute the length of the duration in various units
	p(diff.Hours())
	p(diff.Minutes())
	p(diff.Seconds())
	p(diff.Nanoseconds())

	// can use Add to advance a time, or with '-' to move backwards
	p(then.Add(diff))
	p(then.Add(-diff))
}
