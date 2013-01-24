package cli

// the command line interface.

import (
    "os"
    "fmt"
    "strconv"
)

var commandLine = os.Args[0]

type Cli struct {
    PwLength int
    ShowCharsets bool
    Args []string
}

// default password length
const DefaultPwLength = 8

// ------------------------------------------------------------------

// Constructor
func createCli() Cli {
    var retval Cli
    retval.PwLength = DefaultPwLength
    return retval
}

// Default Constructor
func NewCli() Cli {
    retval := createCli()
    retval.Args = os.Args
    return retval
}

func NewCliFromString(args []string) Cli {
    retval := createCli()
    retval.Args = args
    return retval
}

// parse the command line stored in this Cli object.
// If there was a problem, return an error object,
// otherwise the caller can assume anything was ok.
func (c *Cli) Parse() error {
    var idx int
    var val string

    for idx = 1; idx < len(c.Args); idx++ {
        val = c.Args[idx]
        switch val {
        case "-l":
            idx++
            val2 := c.Args[idx]
            if ival64,err := strconv.ParseInt(val2, 10, 32); err == nil {
                c.PwLength = int(ival64)
            } else {
                return fmt.Errorf("not a password length: \"%s\" (%s)", val2, err)
            }
        default:
            return fmt.Errorf("unknwon argument \"%s\" at %d", val, idx)
        }
    }
    return nil
}

