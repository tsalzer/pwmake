package symbol

import (
    "fmt"
)

// A set of SymbolSets.
type MultiSet struct {
    symsets []*SymbolSet
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

