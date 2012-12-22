package symbol

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
func NewMultiSetFromDefaults(names []string) *MultiSet {
    symsets := make([]*SymbolSet, len(names))
    for k,v := range(names) {
        symsets[k],_ = GetSymbolSet(v)
    }
    return NewMultiSet(symsets)
}

// get the number of symbols in all the symbol sets.
func (p *MultiSet) Len() int {
    retval := 0
    for _,v := range(p.symsets) {
        retval += v.Len()
    }
    return retval
}

