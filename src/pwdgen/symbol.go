/*
 * Symbol
 */

package pwdgen

import (
    "bytes"
    "fmt"
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
    symbols map[*Symbol] struct{}
}

// Constructor.
// Creates a single symbol from the given string. It does *not* break
// up the string into runes!
func NewSymbol(chars string) *Symbol {
    retval := new(Symbol)
    retval.chars = chars
    return retval
}

func NewSymbolSet() *SymbolSet {
    retval := new(SymbolSet)
    retval.symbols = make(map[*Symbol] struct{})
    return retval
}

// put a single symbol into a symbol set.
// returns an error if the symbol already is in the set.
func (p *SymbolSet) Put(symbol *Symbol) error {
    if p.symbols == nil {
        panic("Something strange happened: Got a SymboSet with nil symbols.")
    }

    if _, ok := p.symbols[symbol]; ok == true {
        return fmt.Errorf("There already is a symbol \"%s\" in this symbol set", symbol)
    }
    p.symbols[symbol] = struct{}{}
    //p.symbols = append(p.symbols, symbol)
    return nil
}

// get the number of symbols in the symbol set.
func (p *SymbolSet) Len() int {
    return len(p.symbols)
}

// convert the symbol set into a single string.
func (p *SymbolSet) String() string {
    var buffer bytes.Buffer
    for v,_ := range(p.symbols) {
        buffer.WriteString(v.String())
    }
    return buffer.String()
}

// call the given function with every symbol in the set.
// cancels operation as soon as the function returns an error.
func (p *SymbolSet) Each(fn func(s *Symbol) error) error {
    var err error
    for symbol,_ := range(p.symbols) {
        if err = fn(symbol) ; err != nil {
            return err
        }
    }
    return err
}

// check if the given Stringer is conainet in the set.
func (p *SymbolSet) Contains(s fmt.Stringer) bool {
    return p.ContainsString(s.String())
}

// check if the set contains the given string.
// TODO: this is implemented in an ugly way.
func(p *SymbolSet) ContainsString(cmp string) bool {
    var err error

    err = p.Each(func(s *Symbol) error {
        if (s.chars == cmp) {
            return  fmt.Errorf("found %s", cmp)
        }
        return nil
    })
    return err != nil
}

// create a new SymbolSet from symbols.
func NewSymbolSetFromSymbols(symbols []*Symbol) (*SymbolSet, error) {
    retval := NewSymbolSet()
    if retval == nil {
        panic("did not get a symbol set from NewSymbolSet.")
    }
    var err error
    //return retval, nil
    for _, val := range(symbols) {
        if e := retval.Put(val); e != nil {
            err = e
        }
    }
    return retval, err
}

// create a new SymbolSet from a single string.
// This will create lots of single-rune-Symbols. If any symbol is
// used more than once, there will be an error. However, the symbol
// set might still be useable.
func NewSymbolSetFromString(chars string) (*SymbolSet, error) {
    return NewSymbolSetFromSymbols( SymbolsFromString( chars ) );
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

