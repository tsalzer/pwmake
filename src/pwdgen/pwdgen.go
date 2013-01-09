/*
 * Simple Password Generator.
 */

// password generator package.
package pwdgen

import (
    "fmt"
    "bytes"
    "pwdgen/symbol"
    "pwdgen/rand"
)

type PwdGen struct {
    symbols symbol.RandomSymboler
    length int
}

// The init call will seed the random number generator.
// Maybe we can replace this later with a better, cryptographically
// strong random number generator.
func init() {
    // initialize random seed
    rand.InitializeRandomizer()

    // initialize symbols
    symbol.Initialize(rand.DefaultRandom)
}

// Constructor for Password Generators.
func NewPwdGen(symbols symbol.RandomSymboler, length int) (*PwdGen, error) {
    if length < 1 {
        return nil, fmt.Errorf("the minimum length of a password is 1, you provided %d.", length)
    }
    retval := new(PwdGen)
    retval.symbols = symbols
    retval.length = length
    return retval, nil
}

// Generate a fresh password string.
func (p *PwdGen) String() string {
    var buffer bytes.Buffer
    for i:= 0; i < p.length; i++ {
        buffer.WriteString(p.symbols.RandomSymbol().String())
    }
    return buffer.String()
}

// generate the password,
// This will generate a password with the given specifications, and return it.
// Any error from the password generator will be relayed here and can be printed
// to the user.
func GeneratePassword(length int) (string,error) {
    var gen *PwdGen
    var symset *symbol.MultiSet
    var err error

    if symset, err = symbol.NewMultiSetFromDefaults([]string{"alpha", "ALPHA", "num"}); err != nil {
        return "",err
    }

    if gen, err = NewPwdGen(symset, length); err != nil {
       return "",err
    }
    return gen.String(),nil
}


