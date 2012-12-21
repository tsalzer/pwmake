/*
 * Symbol Sets
 */

package symbol

import (
    "bytes"
    "fmt"
)

// A symbol set.
// An unordered set of symbols.
// Ideally, this would be something like this:
//   type SymbolSet struct {
//       symbols map[*Symbol] struct{}
//   }
// Unfortunately, we need to impose some kind of order to the set, since we want
// to be able to pick a random symbol from it, which is most easy if the symbols
// are stored in a vector (array) of known size, so we can simply pick the symbol
// stored at a given index. Which allows us to utilize the known random functions.
// Therefor, the symbols are stored in an array:
type SymbolSet struct {
    symbols []*Symbol
    fnRandom func(int) int
}

// a map of default symbol sets.
var symbolsets = make(map[string] *SymbolSet)
// initialize all charsets we know
func InitializeSymbolSets() {
    symbolsets["alpha"],_     = NewSymbolSetFromString("abcdefghijklmnopqrstuvwxyz")
    symbolsets["ALPHA"],_     = NewSymbolSetFromString("ABCDEFGHIJKLMNOPQRSTUVWXYZ")
    symbolsets["num"],_       = NewSymbolSetFromString("0123456789")
    symbolsets["specials"],_  = NewSymbolSetFromString(",.!?#@/+-*$%&()")
}
// get the charset with a given name.
func GetSymbolSet(name string) (*SymbolSet, error) {
    var retval *SymbolSet
    var ok bool
    if retval,ok = symbolsets[name]; ok == false {
        return nil, fmt.Errorf("no symbol set with name %s known", name)
    }
    return retval, nil
}

//////////////////////////////////////////////////////////////////////

func NewSymbolSet() *SymbolSet {
    retval := new(SymbolSet)
    //retval.symbols = make(map[*Symbol] struct{})
    retval.symbols = []*Symbol{}
    // by leaving fnRandom = nil, we'll use the default randomizer
    return retval
}

// put a single symbol into a symbol set.
// returns an error if the symbol already is in the set.
func (p *SymbolSet) Put(symbol *Symbol) error {
    if p.symbols == nil {
        panic("Something strange happened: Got a SymboSet with nil symbols.")
    }

    // if _, ok := p.symbols[symbol]; ok == true {
    //     return fmt.Errorf("There already is a symbol \"%s\" in this symbol set", symbol)
    // }
    // p.symbols[symbol] = struct{}{}
    //p.symbols = append(p.symbols, symbol)

    for _,v := range(p.symbols) {
        if v == symbol {
            return fmt.Errorf("There already is a symbol \"%s\" in this symbol set", symbol)
        }
    }
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

// call the given function with every symbol in the set.
// cancels operation as soon as the function returns an error.
func (p *SymbolSet) Each(fn func(s *Symbol) error) error {
    var err error
    for _, symbol := range(p.symbols) {
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

// get a single random symbol from the set.
func (p *SymbolSet) RandomSymbol() *Symbol {
    maxidx := len(p.symbols)
    if p.fnRandom == nil && fnDefaultRandom == nil{
        panic("no randomize function configured")
    }
    if p.fnRandom != nil {
        return p.symbols[p.fnRandom(maxidx)]
    }
    return p.symbols[fnDefaultRandom(maxidx)]
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

