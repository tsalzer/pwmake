package symbol

import (
    "testing"
)

func TestMultiSet(t *testing.T) {
    ms,_ := NewMultiSetFromDefaults( []string{"alpha", "ALPHA", "num"} )
    if ms.symsets == nil {
        t.Fatalf("generated multiset does not contain any symbol set")
    }
    if num := len(ms.symsets); num != 3 {
        t.Errorf("expected 3 symbol sets generated multiset, got %d", num)
    }
}

func TestMultiSetLen(t *testing.T) {
    ms,_ := NewMultiSetFromDefaults( []string{"alpha", "ALPHA", "num"} )
    // 26 + 26 + 10 = 62
    if numsyms := ms.Len(); numsyms != 62 {
        t.Errorf("expected multiset to hold 63 symbold, got %d", numsyms)
    }
}

func TestMultiSetErrors(t *testing.T) {
    ms, err := NewMultiSetFromDefaults( []string{"alpha", "beta", "num"} )
    if err == nil {
        t.Errorf("failed to flag missing symbol set beta")
    }

    if len(ms.symsets) != 2 {
        t.Errorf("unexpected number of symbol sets: expected 2, got %d", len(ms.symsets))
    }
}

