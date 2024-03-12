package intervalgebra

import (
	"testing"
	"time"
)

func TestAttributes(t *testing.T) {
	now := time.Now()
	duration := time.Second

	interval := New(now, duration)

	if interval.Start() != now {
		t.Error("interval start should be as it was constructed")
	}
	if interval.End() != now.Add(duration) {
		t.Error("interval end should be its duration later than start")
	}
	if interval.Duration() != duration {
		t.Error("duration should be as it was constructed")
	}
}

func TestRelations(t *testing.T) {
	now := time.Now()
	intervalA := New(now, time.Minute)
	testCases := []struct {
		name     string
		a        Interval
		b        Interval
		relation Relation
	}{
		{
			name:     "APrecedesB",
			a:        intervalA,
			b:        New(now.Add(time.Hour), time.Minute),
			relation: intervalA.Precedes,
		},
		{
			name:     "AIsPrecededByB",
			a:        intervalA,
			b:        New(now.Add(-time.Hour), time.Minute),
			relation: intervalA.IsPrecededBy,
		},
		{
			name:     "AMeetsB",
			a:        intervalA,
			b:        New(now.Add(time.Minute), time.Minute),
			relation: intervalA.Meets,
		},
		{
			name:     "AIsMetByB",
			a:        intervalA,
			b:        New(now.Add(-time.Minute), time.Minute),
			relation: intervalA.IsMetBy,
		},
		{
			name:     "AOverlapsWithB",
			a:        intervalA,
			b:        New(now.Add(30*time.Second), time.Minute),
			relation: intervalA.OverlapsWith,
		},
		{
			name:     "AIsOverlappedByB",
			a:        intervalA,
			b:        New(now.Add(-30*time.Second), time.Minute),
			relation: intervalA.IsOverlappedBy,
		},
		{
			name:     "AStartsB",
			a:        intervalA,
			b:        New(now, 2*time.Minute),
			relation: intervalA.Starts,
		},
		{
			name:     "AIsStartedByB",
			a:        intervalA,
			b:        New(now, 30*time.Second),
			relation: intervalA.IsStartedBy,
		},
		{
			name:     "ADuringB",
			a:        intervalA,
			b:        New(now.Add(-time.Minute), 3*time.Minute),
			relation: intervalA.During,
		},
		{
			name:     "AContainsB",
			a:        intervalA,
			b:        New(now.Add(15*time.Second), 30*time.Second),
			relation: intervalA.Contains,
		},
		{
			name:     "AFinishesB",
			a:        intervalA,
			b:        New(now.Add(-time.Minute), 2*time.Minute),
			relation: intervalA.Finishes,
		},
		{
			name:     "AIsFinishedByB",
			a:        intervalA,
			b:        New(now.Add(30*time.Second), 30*time.Second),
			relation: intervalA.IsFinishedBy,
		},
		{
			name:     "AEqualsB",
			a:        intervalA,
			b:        New(now, time.Minute),
			relation: intervalA.Equals,
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			if !testCase.relation(testCase.b) {
				t.Errorf(
					"relation failed: \nA.start %v, A.end: %v;\nB.start %v, b.end: %v",
					testCase.a.Start(), testCase.a.End(),
					testCase.b.Start(), testCase.b.End(),
				)
			}
		})
	}
}
