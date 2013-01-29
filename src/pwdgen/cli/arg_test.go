package cli

import (
    "testing"
    "strings"
    // "reflect"
)


func TestMatchBoolean(t *testing.T) {
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

func TestMatchShiftInt(t *testing.T) {
    a := Int("-l", "--pwlen", "password length", "PwLength", 8)
    exp := map[string] bool{
        "-l 1"      : true,     // the short name
        "-l 1 2"    : true,     // ok, but we expect to keep one char
        "-L 1"      : false,    // wrong case
        "-l"        : false,    // missing value
        "-l x"      : false,    // value is not an integer
    }

    for k,v := range(exp) {
        splitted := strings.Split(k, " ")
        if shaved, value, m := a.MatchShift(splitted); m != v {
            t.Errorf("checking \"%s\" against \"%s\" should return %s, but returned %s (shaved=\"%s\", value=\"%s\")",
                k, a, v, m, shaved, value)
            }
    }
}

