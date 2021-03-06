/*
 * Tests for PwdGen
 */

package pwdgen

import (
	"pwdgen/symbol"
	"testing"
)

func TestStringLength(t *testing.T) {
	const testlen = 30
	sym, _ := symbol.NewSymbolSetFromString("a")
	gen, _ := NewPwdGen(sym, testlen)
	if gen == nil {
		t.Errorf("failed to create PwdGen instance")
	}

	pwd := gen.String()
	if len(pwd) != testlen {
		t.Errorf("expected password length %d, but got %d (%s)", testlen, len(pwd), pwd)
	}
}

// --- BENCHMARKS ---------------------------------------------------

func BenchmarkGeneratePassword(b *testing.B) {
	length := 10
	for i := 0; i < b.N; i++ {
		if s, err := GeneratePassword(length); err != nil || len(s) != length {
			b.Errorf("Generating password failed: Password is \"%s\" (expected %d characters), error is %s",
				s, length, err)
		}
	}
}

func BenchmarkPwdGenString(b *testing.B) {
	length := 10
	b.StopTimer()
	var gen *PwdGen
	var symset *symbol.MultiSet
	var err error

	if symset, err = symbol.NewMultiSetFromDefaults([]string{"alpha", "ALPHA", "num"}); err != nil {
		b.Fatalf("Unable to create default MultiSet: %s", err)
	}

	if gen, err = NewPwdGen(symset, length); err != nil {
		b.Fatalf("Unable to create generator: %s", err)
	}
	b.StartTimer()

	for i := 0; i < b.N; i++ {
		if s := gen.String(); len(s) != length {
			b.Errorf("generated password \"%s\" does not have the expected length of %d", s, length)
		}
	}
}
