/*
 * Simple Password Generator.
 */

package pwdgen

import (
//    "crypto/rand"
//    "math/big"
    "math/rand"
    "time"
)

type PwdGen struct {
    charset string
    length int
}

func init() {
    rand.Seed(time.Now().UTC().UnixNano())
}

func NewPwdGen(charset string, length int) *PwdGen {
    retval := new(PwdGen)
    retval.charset = charset
    retval.length = length
    return retval
}

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

