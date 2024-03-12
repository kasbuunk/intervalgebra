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

func TestPrecedes(t *testing.T) {
	interval := New(time.Now(), time.Second)
	other := New(time.Now().Add(time.Minute), time.Second)
	if !interval.Precedes(other) {
		t.Error("first should precede second")
	}
	if other.Precedes(interval) {
		t.Error("second should not precede first")
	}
}
