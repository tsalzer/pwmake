/*
 * Random
 */

package rand

import (
    "math/rand"
)

// get a random rune from a given string.
func RandomRune(charset string) rune {
    var retval rune

    maxidx := len(charset)
    idx := rand.Intn(maxidx)
    retval = rune(charset[idx])

    return retval
}



