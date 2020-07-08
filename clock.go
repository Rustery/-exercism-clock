// Package clock is Exercism.io exercise
package clock

import (
	"fmt"
	"math"
)

type Clock struct {
	h, m int
}

func New(h, m int) Clock {
	hours, minutes := generate(h*60 + m)
	return Clock{hours, minutes}
}

func (c Clock) String() string {
	return fmt.Sprintf("%02d:%02d", c.h, c.m)
}

func (c Clock) Add(m int) Clock {
	c.h, c.m = generate(c.h*60 + c.m + m)
	return c
}

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
