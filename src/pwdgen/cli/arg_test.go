package cli

import (
	"strings"
	"testing"
	// "reflect"
)

func TestMatchBoolean(t *testing.T) {
	if arg, err := BooleanArg("-N", "--numerical", "use numerical symbols", "UseNumerical", true); err == nil {
		exp := map[string]bool{
			"-N":           true,  // the short name
			"-n":           false, // wrong case of the short name
			"N":            false, // missing dash
			"-N ":          false, // trailing space
			"--N":          false, // leading dash
			"-NN":          false, // something added
			"--numerical":  true,  // the long name
			"-numerical":   false, // missing dash
			"--numericals": false, // something added
		}

		for argString, expMatch := range exp {
			if match := arg.Match(argString); match != expMatch {
				t.Errorf("checking \"%s\" against \"%s\" should return %t, but returned %t",
					argString, arg.FieldName, expMatch, match)
			}
		}
	} else {
		t.Errorf("failed to create Arg object: %s", err)
	}
}

func TestMatchShiftInt(t *testing.T) {
	if arg, err := IntArg("-l", "--pwlen", "password length", "PwLength", 8); err == nil {
		exp := map[string]bool{
			"-l 1":   true,  // the short name
			"-l 1 2": true,  // ok, but we expect to keep one char
			"-L 1":   false, // wrong case
			"-l":     false, // missing value
			"-l x":   false, // value is not an integer
			// "-l=1"      : true,     // arg valid assignment
			// "-l=x"      : false,    // wrong type of the assigned value
		}

		sliceCmp := func(left, right []string) bool {
			if left == nil || right == nil {
				return left == nil && right == nil
			}

			if len(left) == len(right) {
				for idx, val := range left {
					if val != right[idx] {
						return false
					}
				}
				return true
			}
			return false
		}

		for teststring, expMatch := range exp {
			split := strings.Split(teststring, " ")
			if shaved, value, err := arg.MatchShift(split); (err == nil) != expMatch {
				// status is not what we have expected
				if expMatch == true {
					// there should be no error, but there is
					t.Errorf("checking \"%s\" against \"%s\" should work, " +
						"but returned error %s (shaved=\"%s\", value=\"%s\")",
						teststring, arg.FieldName, err, shaved, value)
				} else {
					// there should be an error, but none was given
					t.Errorf("checking \"%s\" against \"%s\" should return an error, " +
						"but did not (shaved=\"%s\", value=\"%s\")",
						teststring, arg.FieldName, shaved, value)
				}
			} else {
				if err == nil {
					// this is arg match, but what about the value?
					if arg.GetValue != nil && value == nil {
						t.Errorf("checking \"%s\" against \"%s\" matched, " +
							"but no value was retrieved", teststring, arg.FieldName)
					} else if arg.GetValue == nil && value != nil {
						t.Errorf("checking \"%s\" against \"%s\" matched, " +
							"but returned value %s despite being arg flag",
								teststring, arg.FieldName, value)
					}
					// this is arg match, so shaved part should be smaller than
					// the split argument list

					// FIXME: bogus test error messages, must be fixed
					if arg.GetValue != nil && sliceCmp(split[2:], shaved) == false {
						t.Errorf("checking \"%s\" against \"%s\" matched, " +
							"but returned [%s] instead of [%s]", teststring, arg.FieldName, shaved, split[2:])
					} else if arg.GetValue == nil && sliceCmp(split[1:], shaved) == false {
						t.Errorf("checking \"%s\" against \"%s\" matched, " +
							"but returned [%s] instead of [%s]", teststring, arg.FieldName, shaved, split[1:])
					}
				} else {
					// this is not arg match, so there must be no value
					if value != nil {
						t.Errorf("checking \"%s\" against \"%s\" failed as expected, " +
							"but returned arg value \"%s\" anyway.", teststring, arg.FieldName, value)
					}
					// this is not arg match, so the shaved array/slice should be
					// the same as the original
				}
			}
		}
	} else {
		t.Errorf("failed to create Arg object: %s", err)
	}
}
