// Package gigasecond
package gigasecond

import "time"

// AddGigasecond function
func AddGigasecond(t time.Time) time.Time {
	return t.Add(time.Second * 1e9)
}
