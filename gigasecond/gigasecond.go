package gigasecond

import (
	"time"
)

// AddGigasecond should have a comment documenting it.
func AddGigasecond(t time.Time) time.Time {
	const GIGASECOND = 1000000000

	return t.Add(time.Second * GIGASECOND)
}
