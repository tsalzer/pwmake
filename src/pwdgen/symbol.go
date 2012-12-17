/*
 * Symbol
 */

package pwdgen

import (
    "bytes"
//    "fmt"
)

// A symbol.
// Basically, a symbol is a string. In most cases, its a single rune
// string, but the intention is to support symbol sets consisting of
// symbols with longer string representations (like syllables).
type Symbol struct {
    chars string
}

// A symbol set.
// An unordered set of symbols.
type SymbolSet struct {
    symbols []*Symbol
}

// Constructor.
// Creates a single symbol from the given string. It does *not* break
// up the string into runes!
func NewSymbol(chars string) *Symbol {
    retval := new(Symbol)
    retval.chars = chars
    return retval
}

// create a new SymbolSet from symbols.
func NewSymbolSet(symbols []*Symbol) (*SymbolSet, error) {
    retval := new(SymbolSet)
    var err error
    //return retval, nil
    for _, val := range(symbols) {
        if e := retval.Put(val); e != nil {
            err = e
        }
    }
    return retval, err
}

// put a single symbol into a symbol set.
// returns an error if the symbol already is in the set.
func (p *SymbolSet) Put(symbol *Symbol) error {
    p.symbols = append(p.symbols, symbol)
    return nil
}

// get the number of symbols in the symbol set.
func (p *SymbolSet) Len() int {
    return len(p.symbols)
}

// convert the symbol set into a single string.
func (p *SymbolSet) String() string {
    var buffer bytes.Buffer
    for _,v := range(p.symbols) {
        buffer.WriteString(v.String())
    }
    return buffer.String()
}

// create a new SymbolSet from a single string.
// This will create lots of single-rune-Symbols. If any symbol is
// used more than once, there will be an error. However, the symbol
// set might still be useable.
func NewSymbolSetFromString(chars string) (*SymbolSet, error) {
    return NewSymbolSet( SymbolsFromString( chars ) );
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

