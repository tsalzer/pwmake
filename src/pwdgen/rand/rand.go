/*
 * Randomizer
 */

// The random package for the password generator.
// This package contains the random function used for generating passwords.
package rand

import (
	"crypto/rand"
	"math/big"
)

func InitializeRandomizer() {
	// initialize random seed
	// This was required for the math/rand approach.
}

// default random function.
func DefaultRandom(maxval int) int {
	nBig, err := rand.Int(rand.Reader, big.NewInt(int64(maxval)))
	if err != nil {
		panic("somehow cannot read random data")
	}

	return int(nBig.Int64())
}
