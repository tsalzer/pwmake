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

// func TestNewSymbolSet(t *testing.T) {
//     t.Errorf("test not implemented")
// }

func TestEachSymbol(t *testing.T) {
    const cs = "0123456789"
    var symset *SymbolSet
    var err error
    if symset,err = NewSymbolSetFromString(cs); err != nil {
        if (symset != nil) {
            t.Errorf("error while creating symbol set from \"%s\": %s", cs, err)
        } else {
            t.Fatalf("fatal error while creating symbol set from \"%s\": %s", cs, err)
        }
    }


    // initialize the map
    m := make(map[string] int)
    for _,v := range(cs) { m[string(v)] = 0 }

    // run the Each function, which should increment the value 
    symset.Each(func(s *Symbol) error {
        if v, ok := m[s.String()]; ok == true {
            m[s.String()] = v + 1
        } else {
            t.Errorf("encountered unknown symbol \"%s\"", s.String())
        }
        return nil
    })

    // finally, check that each entry in the map has value=1
    for k,v := range(m) {
        if v != 1 {
            if v == 0 {
                t.Errorf("symbol \"%s\" was expected, but never visited.", k)
            } else {
                t.Errorf("symbol \"%s\" was visited %d times instead of once.", k, v)
            }
        }
    }
}

func TestContains(t *testing.T) {
    const cs = "01234567890"
    var symset *SymbolSet
    var err error

    if symset, err = NewSymbolSetFromString(cs); err != nil{
        t.Errorf("error while creating symbol set from \"%s\": %s", cs, err)
        return
    }


    if symset.ContainsString("0") == false {
        t.Errorf("symbol set \"%s\" claims not to contain string \"%s\"", symset, "0")
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

