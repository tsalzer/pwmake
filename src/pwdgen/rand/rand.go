/*
 * Randomizer
 */

// The random package for the password generator.
// This package contains the random function used for generating passwords.
package rand

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

// default random function.
func DefaultRandom(maxval int) int {
	return rand.Intn(maxval)
}
