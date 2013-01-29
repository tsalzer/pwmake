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


func (s *ArgSet) AddString(short, long, usage, field, value string ) {
    arg := String(short, long, usage, field, value)
    s.Args = append(s.Args, &arg)
}

func (s *ArgSet) AddBoolean(short, long, usage, field string, value bool) {
    arg := Boolean(short, long, usage, field, value)
    s.Args = append(s.Args, &arg)
}


