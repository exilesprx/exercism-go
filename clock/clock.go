package clock

import "fmt"

type Clock struct {
	minutes int
}

func New(h, m int) Clock {
	return Clock{minutes: ((h*60+m)%1440 + 1440) % 1440}
}

func (c Clock) Add(m int) Clock {
	return New(0, c.minutes+m)
}

func (c Clock) Subtract(m int) Clock {
	return New(0, c.minutes-m)
}

func (c Clock) String() string {
	return fmt.Sprintf("%.2d:%.2d", c.minutes/60, c.minutes%60)
}
