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
    const charsetstring = "0123456789"

    results := make(map[string] int)
    for _,v := range charsetstring {
        results[string(v)] = 0
    }

    cs := NewCharset(charsetstring)
    var char string

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

