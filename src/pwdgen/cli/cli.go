package cli

type Cli struct {
    PwLength int
}

// default password length
const DefaultPwLength = 8

// Constructor
func NewCli() Cli {
    var retval Cli
    retval.PwLength = DefaultPwLength
    return retval
}

