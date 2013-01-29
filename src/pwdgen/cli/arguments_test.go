package cli

import (
    "testing"
    // "reflect"
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

func TestMatch(t *testing.T) {
    a := Boolean("-N", "--numerical", "use numerical symbols", "UseNumerical", true)
    exp := map[string] bool{
        "-N"            : true,     // the short name
        "-n"            : false,    // wrong case of the short name
        "N"             : false,    // missing dash
        "-N "           : false,    // trailing space
        "--N"           : false,    // leading dash
        "-NN"           : false,    // something added
        "--numerical"   : true,     // the long name
        "-numerical"    : false,    // missing dash
        "--numericals"  : false,    // something added
    }

    for k,v := range(exp) {
        if m := a.Match(k); m != v {
            t.Errorf("checking \"%s\" against \"%s\" should return %s, but returned %s",
                k, a, v, m)
            }
    }
}

