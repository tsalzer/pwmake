package pwdgen

import (
    "testing"
)


// test the simple system String function.
func TestSystemString(t *testing.T) {
    s := NewSymbol("a")
    if (s.String() != "a") {
        t.Errorf("symbol \"%s\" was not stringified correctly, returns \"%s\"",
            "a", s)
    }
}


func TestSymbolsFromString(t *testing.T) {
    // careful not tu use the globally defined charsets!!!
    const cs = "0123456789"
    var symbols []*Symbol
    // var err error

    // if symbols, err = SymbolsFromString(cs); err != nil {
    //     t.Errorf("failed to create symbols from \"%s\": %s", cs, err)
    //     return
    // }

    symbols = SymbolsFromString(cs)

    if len(symbols) != 10 {
        t.Errorf("expected to find 10 symbols, but got %d", len(symbols))
    }
}

func TestSymbolSetFromString(t *testing.T) {
    const cs = "0123456789"
    var symset *SymbolSet
    var err error

    if symset,err = NewSymbolSetFromString(cs); err != nil {
        t.Errorf("when creating the symbol set from \"%s\": %s", cs, err)
    }
    if symset == nil {
        t.Errorf("did not get a symbol set, can't go on.")
        return
    }

    if symset.Len() != len(cs) {
        t.Errorf("symbol \"%s\" contains %d elements, expected %d (from string \"%s\")",
            symset, symset.Len(), len(cs), cs)
    }
}

