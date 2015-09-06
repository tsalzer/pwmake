package symbol

import (
	"testing"
)

func TestMultiSet(t *testing.T) {
	ms, _ := NewMultiSetFromDefaults([]string{"alpha", "ALPHA", "num"})
	if ms.symsets == nil {
		t.Fatalf("generated multiset does not contain any symbol set")
	}
	if num := len(ms.symsets); num != 3 {
		t.Errorf("expected 3 symbol sets generated multiset, got %d", num)
	}
}

func TestMultiSetLen(t *testing.T) {
	ms, _ := NewMultiSetFromDefaults([]string{"alpha", "ALPHA", "num"})
	// 26 + 26 + 10 = 62
	if numsyms := ms.Len(); numsyms != 62 {
		t.Errorf("expected multiset to hold 63 symbold, got %d", numsyms)
	}
}

func TestMultiSetErrors(t *testing.T) {
	ms, err := NewMultiSetFromDefaults([]string{"alpha", "beta", "num"})
	if err == nil {
		t.Errorf("failed to flag missing symbol set beta")
	}

	if len(ms.symsets) != 2 {
		t.Errorf("unexpected number of symbol sets: expected 2, got %d", len(ms.symsets))
	}
}

func TestGetContainingSet(t *testing.T) {
	ms, symsets := generateMultisetTestSet()
	specials, _ := GetSymbolSet("specials")

	// positive checks
	for name, symset := range symsets {
		symset.Each(func(symbol *Symbol) error {
			if s, err := ms.GetContainingSet(symbol); s != symset || err != nil {
				t.Errorf("symbol %s should be located in symbol set %s, but was in %s (error was: %s)", symbol, name, s, err)
			}
			return nil // we want to check all symbols, not break after the first error
		})
	}

	// negative checks
	specials.Each(func(symbol *Symbol) error {
		if s, err := ms.GetContainingSet(symbol); s != nil || err == nil {
			t.Errorf("symbol %s should not be in any set, but was located in set %v (error was: %s)", s, err)
		}
		return nil // we want to check all symbols
	})
}

// Helpers

func generateMultisetTestSet() (*MultiSet, map[string]*SymbolSet) {
	symsets := make(map[string]*SymbolSet)
	setnames := []string{"alpha", "ALPHA", "num"}
	ms, _ := NewMultiSetFromDefaults(setnames)

	for _, name := range setnames {
		symsets[name], _ = GetSymbolSet(name)
	}

	return ms, symsets
}

// --- BENCHMARKS ---------------------------------------------------

func BenchmarkGetContainingSet(b *testing.B) {
	b.StopTimer()
	ms, _ := generateMultisetTestSet()
	symbol := NewSymbol("m")
	b.StartTimer()

	// only the positive checks
	for i := 0; i < b.N; i++ {
		if _, err := ms.GetContainingSet(symbol); err != nil {
			b.Errorf("symbol %s was not found: %s", symbol, err)
		}
	}
}

// TODO: implement this test
// make sure the random set selection selects each symbol set the multiset consists of.
//func TestMultiSetRandomSetSelection(t *testing.T) {
//	const numsamples = 1000
//	const setnames = []string{"alpha", "ALPHA", "num"}
//	ms,_ := NewMultiSetFromDefaults( setnames )
//
//	m := map[string] int{
//		"alpha" : 0,
//		"ALPHA"	: 0,
//		"num"	: 0,
//	}
//
//	for i := 0; i < numsamples; i++ {
//		sym := ms.RandomSymbol()
//		for k,v := range(m){
//			if ms.symsets[k].Contains(sym) {
//				v++
//			}
//		}
//	}
//}
