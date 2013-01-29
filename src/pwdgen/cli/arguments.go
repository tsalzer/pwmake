package cli

// import (
//     "fmt"
//     "reflect"
// )


type ArgTypes struct {
}

type Arg struct {
    ShortName string
    LongName string
    ValueType string
    DefaultValue string
    Description string
    FieldName string
}

type ArgSet struct {
    Args []*Arg
}


// ------------------------------------------------------------------

func NewArgSet() ArgSet {
    var retval ArgSet
    //retval.Args = make([]*Args)
    return retval
}

func (s *ArgSet) Len() int {
    return len(s.Args)
}

// ------------------------------------------------------------------

func genericArg(short, long, usage, field string) Arg {
    var retval Arg
    retval.ShortName = short
    retval.LongName = long
    retval.Description = usage
    retval.FieldName = field

    return retval
}

func Boolean(short, long, usage, field string, value bool) Arg {
    retval := genericArg(short, long, usage, field)
    return retval
}

func String(short, long, usage, field string, value string) Arg {
    retval := genericArg(short, long, usage, field)
    return retval
}


func (s *ArgSet) AddString(short, long, usage, field, value string ) {
    arg := String(short, long, usage, field, value)
    s.Args = append(s.Args, &arg)
}

func (s *ArgSet) AddBoolean(short, long, usage, field string, value bool) {
    arg := Boolean(short, long, usage, field, value)
    s.Args = append(s.Args, &arg)
}

// ------------------------------------------------------------------

// checkif the Arg matches the given string.
func (a *Arg) Match(arg string) bool {
    switch arg {
    case a.ShortName:
        return true
    case a.LongName:
        return true
    }
    return false
}

