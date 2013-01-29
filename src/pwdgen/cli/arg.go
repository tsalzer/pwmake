package cli

import (
    "strconv"
    // "fmt"
    // "reflect"
)


type ArgTypes struct {
}

type Arg struct {
    ShortName       string
    LongName        string
    ValueType       string
    DefaultValue    string
    Description     string
    FieldName       string
    RequireValue    bool
    GetValue        func(string) (interface{}, error)
}

// ------------------------------------------------------------------

// generic constructor.
// fills in the stuff we need for every argument.
func genericArg(short, long, usage, field string) Arg {
    var retval Arg
    retval.ShortName = short
    retval.LongName = long
    retval.Description = usage
    retval.FieldName = field

    return retval
}

// a Boolean argument.
// Needs a default value.
func Boolean(short, long, usage, field string, value bool) Arg {
    retval := genericArg(short, long, usage, field)
    return retval
}

// an Int argument.
// Needs a default value.
func Int(short, long, usage, field string, value int) Arg {
    retval := genericArg(short, long, usage, field)
    retval.RequireValue = true
    retval.GetValue = extractInt
    return retval
}

// a String argument.
// Needs a default value.
func String(short, long, usage, field string, value string) Arg {
    retval := genericArg(short, long, usage, field)
    retval.RequireValue = true
    retval.GetValue = extractString
    return retval
}

// ------------------------------------------------------------------

func extractInt(arg string) (interface{}, error) {
    var ival64 int64
    var err error
    if ival64,err = strconv.ParseInt(arg, 10, 32); err == nil {
        return int(ival64), nil
    }
    return nil, err
}

func extractString(arg string) (interface{}, error) {
    return arg, nil
}


// ------------------------------------------------------------------

// check if the given argument matches, then shift.
// only consider the first argument.
func (a *Arg) MatchShift(args []string) ([]string, interface{}, bool) {
    if a.Match(args[0]) {
        shaved := args[1:]  // shave off the first string

        if a.GetValue != nil {
            // we expect a value
            if len(shaved) > 0 {
                if val,err := a.GetValue(shaved[0]); err == nil {
                    return shaved[1:], val, true
                } else {
                    return args, nil, false
                }
            } else {
               return args, nil, false
            }
        } else {
            // this is just a flag
            return shaved, nil, true
        }
    }
    return args, nil, false
}

// check if the Arg matches the given string.
func (a *Arg) Match(arg string) bool {
    switch {
    case a.ShortName != "" && a.ShortName == arg:
        return true
    case a.LongName != "" && a.LongName == arg:
        return true
    }
    return false
}

