package intervalgebra

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
