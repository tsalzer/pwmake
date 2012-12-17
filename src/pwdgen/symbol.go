/*
 * Symbol
 */

package pwdgen

// A symbol.
// Basically, a symbol is a string. In most cases, its a single rune
// string, but the intention is to support symbol sets consisting of
// symbols with longer string representations (like syllables).
type Symbol struct {
    chars string
}

// Constructor.
// Creates a single symbol from the given string. It does *not* break
// up the string into runes!
func NewSymbol(chars string) *Symbol {
    retval := new(Symbol)
    retval.chars = chars
    return retval
}

// Stringer
func (p *Symbol) String() string {
    return p.chars
}

