package intervalgebra

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
