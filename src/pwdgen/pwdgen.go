/*
 * Simple Password Generator.
 */

// password generator package.
package pwdgen

import (
    "fmt"
    "bytes"
)

type PwdGen struct {
    symbols *SymbolSet
    length int
}

// The init call will seed the random number generator.
// Maybe we can replace this later with a better, cryptographically
// strong random number generator.
func init() {
    // initialize random seed
    InitializeRandomizer()

    // initialize charsets
    InitializeCharsets()
    InitializeSymbolSets()
}

// Constructor for Password Generators.
func NewPwdGen(symbols *SymbolSet, length int) (*PwdGen, error) {
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

