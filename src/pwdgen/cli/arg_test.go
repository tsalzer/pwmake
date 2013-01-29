package cli

import (
    "testing"
    // "reflect"
)


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

