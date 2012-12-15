/*
 * Randomizer
 */

package pwdgen

import (
//    "crypto/rand"
//    "math/big"
    "math/rand"
    "time"
)

func InitializeRandomizer() {
    // initialize random seed
    rand.Seed(time.Now().UTC().UnixNano())
}

// get a random rune from a given string.
func RandomRune(charset string) rune {
    var retval rune

    maxidx := len(charset)
    idx := rand.Intn(maxidx)
    retval = rune(charset[idx])

    return retval
}



