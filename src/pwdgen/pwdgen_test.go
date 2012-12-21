/*
 * Tests for PwdGen
 */

package pwdgen

import (
    "testing"
    "pwdgen/symbol"
)

func TestStringLength(t *testing.T) {
    const testlen = 30
    sym,_ := symbol.NewSymbolSetFromString("a")
    gen, _ := NewPwdGen(sym, testlen)
    if (gen == nil) {
        t.Errorf("failed to create PwdGen instance")
    }

    pwd := gen.String()
    if len(pwd) != testlen {
        t.Errorf("expected password length %d, but got %d (%s)", testlen, len(pwd), pwd)
    }
}

