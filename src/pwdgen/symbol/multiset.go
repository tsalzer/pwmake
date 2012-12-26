package symbol

import (
    "fmt"
)

// A set of SymbolSets.
type MultiSet struct {
    symsets []*SymbolSet
    fnRandom func(int) int
}

// construct a new MultiSet.
func NewMultiSet(symsets []*SymbolSet) *MultiSet {
    retval := new(MultiSet)
    retval.symsets = symsets
    return retval
}

// generate a new MultiSet from default symbol sets.
func NewMultiSetFromDefaults(names []string) (*MultiSet, error) {
    notfound := ""
    symsets := make([]*SymbolSet, 0)

    var symset *SymbolSet
    var err error

    for _,v := range(names) {
        if symset, err = GetSymbolSet(v); err != nil {
            notfound = fmt.Sprintf("%s %s", notfound, v)
        } else {
            symsets = append(symsets, symset)
        }
    }

    retval := NewMultiSet(symsets)
    err = nil

    if notfound != "" {
        err = fmt.Errorf("did not find symbol set(s): %s", notfound)
    }
    return retval, err
}

// get the number of symbols in all the symbol sets.
func (p *MultiSet) Len() int {
    retval := 0
    for _,v := range(p.symsets) {
        retval += v.Len()
    }
    return retval
}

// get a single random symbol from the multiset.
func (p *MultiSet) RandomSymbol() *Symbol {
    maxidx := len(p.symsets)
    if p.fnRandom == nil && fnDefaultRandom == nil{
        panic("no randomize function configured")
    }
    if p.fnRandom != nil {
        return p.symsets[p.fnRandom(maxidx)].RandomSymbol()
    }
    return p.symsets[fnDefaultRandom(maxidx)].RandomSymbol()
}

// get the symbol set containing the given stringer.
func (p *MultiSet) GetContainingSet(symbol fmt.Stringer) (*SymbolSet, error) {
	ch := make(chan *SymbolSet)

	// launch n goroutines
	for _,v := range(p.symsets) {
		go isContained(ch, symbol, v)
	}

	var retval *SymbolSet
	// collect n results
	for _,_ = range(p.symsets){
		if x := <-ch; x != nil && retval == nil {
			retval = x
		}
	}

	if retval == nil {
		return nil, fmt.Errorf("Symbol %s not contained in any SymbolSet of this MultiSet", symbol)
	}
	return retval, nil
}

func isContained(ch chan *SymbolSet, symbol fmt.Stringer, symset *SymbolSet){
	if symset.Contains(symbol) {
		ch <- symset
	} else {
		ch <- nil
	}
}
