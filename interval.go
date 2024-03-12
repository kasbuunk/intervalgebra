package intervalgebra

import "time"

type Interval struct {
	start    time.Time
	duration time.Duration
}

func New(start time.Time, duration time.Duration) Interval {
	return Interval{
		start:    start,
		duration: duration,
	}
}

func (i Interval) Start() time.Time {
	return i.start
}

func (i Interval) End() time.Time {
	return i.start.Add(i.duration)
}

func (i Interval) Duration() time.Duration {
	return i.duration
}

func (i Interval) Precedes(other Interval) bool {
	return i.End().Before(other.Start())
}
