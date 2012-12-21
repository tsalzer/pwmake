package symbol

import (
    "testing"
    "pwdgen/rand"
)

// init function for tests.
// TODO: find out if this is really only used by the tests
func init() {
    Initialize(rand.DefaultRandom)
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




// When we ask for a random char, we want:
// a) get a single char exclusively from the charset
// b) after many calls, each char from the charset must be returned
// We consider 1000 calls for a charset of 10 characters as many.
func TestRandomSymbol(t *testing.T) {
    const numsamples = 1000
    const charsetstring = "0123456789"

    results := make(map[string] int)
    for _,v := range charsetstring {
        results[string(v)] = 0
    }

    symset,_ := NewSymbolSetFromString(charsetstring)
    var char string

    // run the sampling
    for i := 0; i < numsamples; i++ {
        char = symset.RandomSymbol().String()

        // check for single charater
        if len(char) != 1 {
            t.Errorf("in iteration %d: expected single char, got value \"%s\" with length %d",
                i, char, len(char))
        }

        // count
        if _, ok := results[char]; ok {
            results[char]++
        }  else {
            t.Errorf("in iteration %d: got value \"%s\" which is not in the charset",
                i, char)
        }

    }

    // check the result map
    for key, value := range results {
        if value == 0 {
            t.Errorf("character %s was never picked", key)
        }
    }
}
