package clock

import "fmt"

type Clock struct {
	m int
	h int
}

func New(h, m int) Clock {
	c := Clock{m: m, h: h}
	return normalize(&c)
}

func (c Clock) Add(m int) Clock {
	c.m += m
	if !isClockValid(c) {
		return normalize(&c)
	}
	return c
}

func (c Clock) Subtract(m int) Clock {
	c.m -= m
	if !isClockValid(c) {
		return normalize(&c)
	}
	return c
}

func (c Clock) String() string {
	return fmt.Sprintf("%02d:%02d", c.h, c.m)
}

func normalize(clock *Clock) Clock {
	if isClockValid(*clock) {
		return *clock
	}

	if clock.h > 23 {
		clock.h = clock.h % 24
	} else if clock.h < 0 {
		clock.h = clock.h%24 + 24
	}

	if clock.m > 59 {
		clock.h += 1
		clock.m -= 60
	} else if clock.m < 0 {
		clock.h -= 1
		clock.m += 60
	}

	if !isClockValid(*clock) {
		return normalize(clock)
	}

	return *clock

}

func isClockValid(c Clock) bool {
	if c.m >= 0 && c.m < 60 && c.h >= 0 && c.h < 24 {
		return true
	}
	return false
}
