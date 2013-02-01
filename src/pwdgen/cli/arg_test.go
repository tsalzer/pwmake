package cli

import (
    "testing"
    "strings"
    // "reflect"
)


func TestMatchBoolean(t *testing.T) {
    if a,err := BooleanArg("-N", "--numerical", "use numerical symbols", "UseNumerical", true); err == nil {
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
    } else {
        t.Errorf("failed to create Arg object: %s", err)
    }
}

func TestMatchShiftInt(t *testing.T) {
    if a, err := IntArg("-l", "--pwlen", "password length", "PwLength", 8); err == nil {
        exp := map[string] bool{
            "-l 1"      : true,     // the short name
            "-l 1 2"    : true,     // ok, but we expect to keep one char
            "-L 1"      : false,    // wrong case
            "-l"        : false,    // missing value
            "-l x"      : false,    // value is not an integer
            // "-l=1"      : true,     // a valid assignment
            // "-l=x"      : false,    // wrong type of the assigned value
        }

        sliceCmp := func(left, right []string) bool{
            if left == nil || right == nil {
                return left == nil && right == nil
            }

            if len(left) == len(right) {
                for idx,val := range(left) {
                    if val != right[idx] {
                        return false
                    }
                }
                return true
            }
            return false
        }

        for k,v := range(exp) {
            splitted := strings.Split(k, " ")
            if shaved, value, err := a.MatchShift(splitted); (err == nil) != v {
                // status is not what we have expected
                if v == true {
                    // there should be no error, but there is
                    t.Errorf("checking \"%s\" against \"%s\" should work, but returned error %s (shaved=\"%s\", value=\"%s\")",
                        k, a, err, shaved, value)
                } else{
                    // there should be an error, but none was given
                    t.Errorf("checking \"%s\" against \"%s\" should return an error, but did not (shaved=\"%s\", value=\"%s\")",
                        k, a, shaved, value)
                }
            } else {
                if err == nil {
                    // this is a match, but what about the value?
                    if a.GetValue != nil && value == nil {
                        t.Errorf("checking \"%s\" against \"%s\" matched, but no value was retrieved", k, a)
                    } else if a.GetValue == nil && value != nil {
                        t.Errorf("checking \"%s\" against \"%s\" matched, but returned value %s despite being a flag", k, a)
                    }
                    // this is a match, so shaved part should be smaller than
                    // the splitted argument list
                    if a.GetValue != nil && sliceCmp(splitted[2:], shaved) == false {
                        t.Errorf("checking \"%s\" against \"%s\" matched, but returned [%s] instead of [%s]", shaved, splitted[2:])
                    } else if a.GetValue == nil && sliceCmp(splitted[1:], shaved) == false {
                        t.Errorf("checking \"%s\" against \"%s\" matched, but returned [%s] instead of [%s]", shaved, splitted[1:])
                    }
                } else {
                    // this is not a match, so there must be no value
                    if value != nil {
                        t.Errorf("checking \"%s\" against \"%s\" failed as expected, but returned a value \"%s\" anyway.", k, a, value)
                    }
                    // this is not a match, so the shaved array/slice should be
                    // the same as the original
                }
            }
        }
    } else {
        t.Errorf("failed to create Arg object: %s", err)
    }
}

