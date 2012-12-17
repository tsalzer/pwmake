package pwdgen

import (
    "testing"
)


// test the simple system String function.
func TestSystemString(t *testing.T) {
    s := NewSymbol("a")
    if (s.String() != "a") {
        t.Errorf("symbol \"%s\" was not stringified correctly, returns \"%s\"",
            "a", s)
    }
}

