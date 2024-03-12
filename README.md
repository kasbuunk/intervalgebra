# IntervAlgebra

A Go package that provides the `Interval` type, implementing [Allen's Interval Algebra](https://en.wikipedia.org/wiki/Allen%27s_interval_algebra).

```go
import (
    "github.com/kasbuunk/intervalgebra"
)

func main() {
    now := time.Now()

    // Construct some Intervals.
    intervalA := intervalalgebra.New(now, time.Minute)
    intervalB := intervalalgebra.New(now.Add(time.Hour), time.Minute)

    // A entirely precedes B.
    fmt.Println("A precedes B: %t", intervalA.Precedes(intervalB))
    fmt.Println("B is preceded by A: %t", intervalA.IsPrecededBy(intervalB))
}
```

Check out all thirteen algebraic possibilities. At all 'times', exactly one relation returns true, whilst all others must return false, provided that a non-zero duration was entered upon construction.
