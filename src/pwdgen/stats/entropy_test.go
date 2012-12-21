package stats

import (
    "testing"
    "math"
)

func TestEntropyOf(t *testing.T) {
    var e float64
    // there's no Round function in the math package? Must be kidding...
    if e = EntropyOf(10,10); math.Floor(e) != 33 {
        t.Errorf("expected entropy to be about 33, but got %f", e)
    }
}

