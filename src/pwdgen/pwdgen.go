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
    charset string
    length int
}

// The init call will seed the random number generator.
// Maybe we can replace this later with a better, cryptographically
// strong random number generator.
func init() {
    rand.Seed(time.Now().UTC().UnixNano())
}

// Constructor for Password Generators.
func NewPwdGen(charset string, length int) (*PwdGen, error) {
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
    //maxidx := big.NewInt(int64(len(p.charset)))A
    maxidx := len(p.charset)
    var idx int

    for c := 0; c < p.length; c++ {
        idx = rand.Intn(maxidx)
        retval += string(p.charset[idx])
    }
    return retval
}

