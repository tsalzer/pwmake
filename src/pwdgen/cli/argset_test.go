package cli

import (
    "testing"
)

func TestNewArgSet(t *testing.T) {
    s := NewArgSet()
    if len(s.Args) != 0 {
        t.Errorf("got some argument at construction time: %s", s.Args)
    }
    if s.Len() != len(s.Args) {
        t.Errorf("ArgSet#Len() reports %d, whicle len(Args) report %d", s.Len(), len(s.Args))
    }
}

func TestAddBoolean(t *testing.T) {
    s := NewArgSet()
    s.AddBoolean("-N", "--numerical", "use numerical symbols", "UseNumerical", true)

    if len(s.Args) != 1 {
        t.Errorf("adding a string value failed")
    }

    if s.Len() != len(s.Args) {
        t.Errorf("ArgSet#Len() reports %d, whicle len(Args) report %d", s.Len(), len(s.Args))
    }
}


