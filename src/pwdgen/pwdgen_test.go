/*
 * Tests for PwdGen
 */

package pwdgen

import (
    "testing"
)

func TestStringLength(t *testing.T) {
    const testlen = 30
    cs := NewCharset("a")
    gen, _ := NewPwdGen(cs, testlen)
    if (gen == nil) {
        t.Errorf("failed to create PwdGen instance")
    }

    pwd := gen.String()
    if len(pwd) != testlen {
        t.Errorf("expected password length %d, but got %d (%s)", testlen, len(pwd), pwd)
    }
}
