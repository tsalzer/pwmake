package symbol

import (
    "testing"
)

func TestMultiSet(t *testing.T) {
    ms := NewMultiSetFromDefaults( []string{"alpha", "ALPHA", "num"} )
    if ms.symsets == nil {
        t.Fatalf("generated multiset does not contain any symbol set")
    }
    if num := len(ms.symsets); num != 3 {
        t.Errorf("expected 3 symbol sets generated multiset, got %d", num)
    }
}

func TestMultiSetLen(t *testing.T) {
    ms := NewMultiSetFromDefaults( []string{"alpha", "ALPHA", "num"} )
    // 26 + 26 + 10 = 62
    if numsyms := ms.Len(); numsyms != 62 {
        t.Errorf("expected multiset to hold 63 symbold, got %d", numsyms)
    }
}

