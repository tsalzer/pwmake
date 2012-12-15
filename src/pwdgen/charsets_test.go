/*
 * Tests for charsets
 */

package pwdgen

import (
    "testing"
)

// When we ask for a random char, we want:
// a) get a single char exclusively from the charset
// b) after many calls, each char from the charset must be returned
// We consider 1000 calls for a charset of 10 characters as many.
func TestRandomChar(t *testing.T) {
    const numsamples = 1000
    cs := NewCharset("0123456789")
    var char string
    results := map[string] int {
        "0" : 0,
        "1" : 0,
        "2" : 0,
        "3" : 0,
        "4" : 0,
        "5" : 0,
        "6" : 0,
        "7" : 0,
        "8" : 0,
        "9" : 0,

    }

    // run the sampling
    for i := 0; i < numsamples; i++ {
        char = cs.RandomChar()

        // check for single charater
        if len(char) != 1 {
            t.Errorf("in iteration %d: expected single char, got value \"%s\" with length %d",
                i, char, len(char))
        }

        // count
        if _, ok := results[char]; ok {
            results[char]++
        }  else {
            t.Errorf("in iteration %d: got value \"%s\" which is not in the charset",
                i, char)
        }

    }

    // check the result map
    for key, value := range results {
        if value == 0 {
            t.Errorf("character %s was never picked", key)
        }
    }
}

