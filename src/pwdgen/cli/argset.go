package cli

type ArgSet struct {
    Args []*Arg
}

func NewArgSet() ArgSet {
    var retval ArgSet
    //retval.Args = make([]*Args)
    return retval
}

func (s *ArgSet) Len() int {
    return len(s.Args)
}


// ------------------------------------------------------------------


// Add a string option to the argument set.
func (s *ArgSet) AddInt(short, long, usage, field string, value int) error {
    if arg, err := IntArg(short, long, usage, field, value); err != nil {
        return err
    } else {
        s.Args = append(s.Args, &arg)
    }
    return nil
}

// Add a string option to the argument set.
func (s *ArgSet) AddString(short, long, usage, field, value string) error {
    if arg, err := StringArg(short, long, usage, field, value); err != nil {
        return err
    } else {
        s.Args = append(s.Args, &arg)
    }
    return nil
}

// Add a boolean flag option to the argument set.
func (s *ArgSet) AddBoolean(short, long, usage, field string, value bool) error {
    if arg, err := BooleanArg(short, long, usage, field, value); err != nil {
        return err
    } else {
        s.Args = append(s.Args, &arg)
    }
    return nil
}


