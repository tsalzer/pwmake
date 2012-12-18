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

// Splits a string up in runes, generate a symbol from each rune.
// This will return an array containing a Symbol for each rune in the
// string. This function does not check for duplicate symbols.
func SymbolsFromString(chars string) []*Symbol {
    retval := make([]*Symbol, len(chars))
    for i,v := range(chars) {
        retval[i] = NewSymbol(string(v))
    }
    return retval
}

// Stringer
func (p *Symbol) String() string {
    return p.chars
}

