package pwdgen

import (
    "bytes"
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
    var buffer bytes.Buffer
    var symset *SymbolSet
    var err error
    if symset,err = NewSymbolSetFromString(cs); err != nil {
        if (symset != nil) {
            t.Errorf("error while creating symbol set from \"%s\": %s", cs, err)
        } else {
            t.Fatalf("fatal error while creating symbol set from \"%s\": %s", cs, err)
        }
    }

    symset.Each(func(s *Symbol) error {
        buffer.WriteString(s.String())
        return nil
    })
    order := buffer.String()
    if order != cs {
        // the order is not really important, but now we must make sure
        // each rune of cs is in the string order
        // TODO: this still needs to be implemented.
        t.Errorf("unexpected behavior of Each: visited in order %s", order)
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

