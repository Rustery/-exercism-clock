// Package clock is Exercism.io exercise
package clock

import (
	"fmt"
	"math"
)

// Clock is a type that stores passed hours and minutes
type Clock struct {
	h, m int
}

// New method generates a new clock instance from hours and minutes params
func New(h, m int) Clock {
	hours, minutes := generate(h*60 + m)
	return Clock{hours, minutes}
}

func (c Clock) String() string {
	return fmt.Sprintf("%02d:%02d", c.h, c.m)
}

// Add method adds minutes to clock instance
func (c Clock) Add(m int) Clock {
	c.h, c.m = generate(c.h*60 + c.m + m)
	return c
}

// Subtract method subtracts minutes from clock instance
func (c Clock) Subtract(m int) Clock {
	c.h, c.m = generate(c.h*60 + c.m - m)
	return c
}

func generate(initialMinutes int) (h, m int) {
	days := float64(initialMinutes) / 1440
	clearDays := math.Floor(days)
	hours := (days - clearDays) * 24
	clearHours := math.Floor(hours)
	minutes := (hours - clearHours) * 60
	clearMinutes := math.Round(minutes)
	return int(clearHours), int(clearMinutes)
}
