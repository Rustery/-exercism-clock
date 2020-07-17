// Package clock is Exercism.io exercise
package clock

import (
	"fmt"
)

// Clock is a type that stores passed hours and minutes
type Clock struct {
	m int
}

// New method generates a new clock instance from hours and minutes params
func New(h, m int) Clock {
	minutes := h*60 + m
	minutes %= 24 * 60
	if minutes < 0 {
		minutes += 24 * 60
	}
	return Clock{minutes}
}

// String method returns string view of a clock instance data
func (c Clock) String() string {
	return fmt.Sprintf("%02d:%02d", c.hours(), c.minutes())
}

// Add method adds minutes to clock instance
func (c Clock) Add(m int) Clock {
	c = New(c.hours(), c.minutes()+m)
	return c
}

// Subtract method subtracts minutes from clock instance
func (c Clock) Subtract(m int) Clock {
	c = New(c.hours(), c.minutes()-m)
	return c
}

// Hours method returns hours of clock instance
func (c Clock) hours() int {
	return (c.m / 60) - (c.m / (60 * 24))
}

// Minutes method returns minutes of clock instance
func (c Clock) minutes() int {
	return c.m - (c.hours() * 60)
}
