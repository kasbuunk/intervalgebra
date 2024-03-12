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

func (i Interval) equalStart(other Interval) bool {
	return i.Start() == other.Start()
}

func (i Interval) equalEnd(other Interval) bool {
	return i.End() == other.End()
}

func (i Interval) startsEarlier(other Interval) bool {
	return i.Start().Before(other.Start())
}

func (i Interval) endsBeforeOtherStart(other Interval) bool {
	return i.End().Before(other.Start())
}

func (i Interval) endsEarlier(other Interval) bool {
	return i.End().Before(other.End())
}

func (i Interval) endEqualsOtherStart(other Interval) bool {
	return i.End() == other.Start()
}

func (i Interval) Precedes(other Interval) bool {
	return i.endsBeforeOtherStart(other)
}

func (i Interval) IsPrecededBy(other Interval) bool {
	return other.Precedes(i)
}

func (i Interval) Meets(other Interval) bool {
	return i.endEqualsOtherStart(other)
}

func (i Interval) IsMetBy(other Interval) bool {
	return other.Meets(i)
}

func (i Interval) OverlapsWith(other Interval) bool {
	return i.startsEarlier(other) && i.endsEarlier(other)
}

func (i Interval) IsOverlappedBy(other Interval) bool {
	return other.OverlapsWith(i)
}

func (i Interval) Starts(other Interval) bool {
	return i.equalStart(other)
}

func (i Interval) IsStartedBy(other Interval) bool {
	return other.Starts(i)
}

func (i Interval) During(other Interval) bool {
	return other.startsEarlier(i) && i.endsEarlier(other)
}

func (i Interval) Contains(other Interval) bool {
	return other.During(i)
}

func (i Interval) Finishes(other Interval) bool {
	return i.equalEnd(other) && other.startsEarlier(i)
}

func (i Interval) IsFinishedBy(other Interval) bool {
	return other.Finishes(i)
}

func (i Interval) Equals(other Interval) bool {
	return i.equalStart(other) && i.equalEnd(other)
}

type Relation func(Interval) bool

type Precedes Relation
type IsPrecededBy Relation
type Meets Relation
type IsMetBy Relation
type OverlapsWith Relation
type IsOverlappedBy Relation
type Starts Relation
type IsStartedBy Relation
type During Relation
type Contains Relation
type Finishes Relation
type IsFinishedBy Relation
type Equals Relation
