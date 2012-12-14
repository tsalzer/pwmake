/*
 * Simple Password Generator.
 */

// password generator package.
package pwdgen

import (
//    "crypto/rand"
//    "math/big"
    "math/rand"
    "time"
    "fmt"
)

type PwdGen struct {
    charset *Charset
    length int
}

// The init call will seed the random number generator.
// Maybe we can replace this later with a better, cryptographically
// strong random number generator.
func init() {
    // initialize random seed
    rand.Seed(time.Now().UTC().UnixNano())

    // initialize charsets
    InitializeCharsets()
}

// Constructor for Password Generators.
func NewPwdGen(charset *Charset, length int) (*PwdGen, error) {
    if length < 1 {
        return nil, fmt.Errorf("the minimum length of a password is 1, you provided %d.", length)
    }
    retval := new(PwdGen)
    retval.charset = charset
    retval.length = length
    return retval, nil
}

// Generate a fresh password string.
func (p *PwdGen) String() string {
    retval := ""
    for c := 0; c < p.length; c++ {
        retval += p.charset.RandomChar()
    }
    return retval
}

