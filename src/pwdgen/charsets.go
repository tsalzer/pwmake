/*
 * Charsets
 */

package pwdgen

import (
    "fmt"
    "math/rand"
)

// A Charset. Basically, a string.
type Charset struct {
    chars   string
}

// a map of possible charsets.
var charsets = make(map[string] *Charset)

// create a new charset.
func NewCharset(chars string) *Charset {
    retval := new(Charset)
    retval.chars = chars
    return retval
}

// initialize all charsets we know
func InitializeCharsets() {
    charsets["alpha"] = NewCharset("abcdefghijklmnopqrstuvwxyz")
    charsets["ALPHA"] = NewCharset("ABCDEFGHIJKLMNOPQRSTUVWXYZ")
    charsets["num"]   = NewCharset("0123456789")
    charsets["specials"] = NewCharset(",.!?#@/+-*$%&()")
}

// get the charset with a given name.
func GetCharset(name string) (*Charset, error) {
    retval := charsets[name]
    if retval == nil {
        return nil, fmt.Errorf("no charset with name %s known", name)
    }
    return retval, nil
}


// A way to get all the charsets built into the application.
func GetCharsetsNames() []string {
    return []string{"alpha", "ALPHA", "num", "specials"}
}


// Generate a fresh password string.
func (p *Charset) RandomChar() string {
    maxidx := len(p.chars)
    idx := rand.Intn(maxidx)
    return string(p.chars[idx])
}

