package clocks

import (
	"time"

	"github.com/senorprogrammer/wtf/wtf"
)

type Clock struct {
	Label    string
	Location *time.Location
}

func NewClock(label string, timeLoc *time.Location) Clock {
	clock := Clock{
		Label:    label,
		Location: timeLoc,
	}

	return clock
}

func (clock *Clock) Date() string {
	return clock.LocalTime().Format(wtf.SimpleDateFormat)
}

func (clock *Clock) LocalTime() time.Time {
	return time.Now().In(clock.Location)
}

func (clock *Clock) Time() string {
	return clock.LocalTime().Format(wtf.SimpleTimeFormat)
}
