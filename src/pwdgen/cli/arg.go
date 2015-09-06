package cli

import (
    "strconv"
    "fmt"
    "reflect"
)


type Arg struct {
    ShortName       string      // short name, may be empty
    LongName        string      // long name, may be empty
    ValueType       interface{} // the required value type
    DefaultValue    interface{} // the default value (may be nil)
    Description     string      // description, used by usage function
    FieldName       string      // name of the field to update
    GetValue        func(string) (interface{}, error)   // function to extract value (may be nil)
}

// ------------------------------------------------------------------

// generic constructor.
// fills in the stuff we need for every argument.
func genericArg(short, long, usage, field string) (Arg, error) {
    var retval Arg
    retval.ShortName = short
    retval.LongName = long
    retval.Description = usage
    retval.FieldName = field

    if retval.ShortName == "" && retval.LongName == "" {
        return retval, fmt.Errorf("invalid argument: an argument must have at least one name")
    }
    return retval, nil
}

// a Boolean argument.
// The value defines what will be written into the field when the
// argument matches.
func BooleanArg(short, long, usage, field string, value bool) (Arg, error) {
    retval, err := genericArg(short, long, usage, field)
    return retval, err
}

// an Int argument.
// Needs a default value.
func IntArg(short, long, usage, field string, value int) (Arg, error) {
    retval, err := genericArg(short, long, usage, field)
    retval.ValueType = reflect.Int
    retval.GetValue = extractInt
    return retval, err
}

// a String argument.
// Needs a default value.
func StringArg(short, long, usage, field string, value string) (Arg, error) {
    retval, err := genericArg(short, long, usage, field)
    retval.ValueType = reflect.String
    retval.GetValue = extractString
    return retval, err
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

