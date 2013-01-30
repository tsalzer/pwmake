package cli

import (
    "strconv"
    "fmt"
    "reflect"
)


type ArgTypes struct {
}

type Arg struct {
    ShortName       string
    LongName        string
    ValueType       interface{}
    DefaultValue    string
    Description     string
    FieldName       string
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
// The value defines what will be written into the field when the
// argument matches.
func Boolean(short, long, usage, field string, value bool) Arg {
    retval := genericArg(short, long, usage, field)
    return retval
}

// an Int argument.
// Needs a default value.
func Int(short, long, usage, field string, value int) Arg {
    retval := genericArg(short, long, usage, field)
    retval.ValueType = reflect.Int
    retval.GetValue = extractInt
    return retval
}

// a String argument.
// Needs a default value.
func String(short, long, usage, field string, value string) Arg {
    retval := genericArg(short, long, usage, field)
    retval.ValueType = reflect.String
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
func (a *Arg) MatchShift(args []string) ([]string, interface{}, error) {
    if a.Match(args[0]) {
        shaved := args[1:]  // shave off the first string

        if a.GetValue != nil {
            // we expect a value
            if len(shaved) > 0 {
                if val,err := a.GetValue(shaved[0]); err == nil {
                    return shaved[1:], val, nil
                } else {
                    return args, nil, err
                }
            } else {
               return args, nil, fmt.Errorf("no value given for option %s", a.ShortName)
            }
        } else {
            // this is just a flag
            return shaved, nil, nil
        }
    }
    return args, nil, fmt.Errorf("option %s not recognized", args[0])
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

