package cli

// the command line interface.

import (
	"fmt"
	"os"
	"strconv"
)

var commandLine = os.Args[0]

type Cli struct {
	// basic options
	PwLength int // number of characters in the final password
	PwNum    int // number of passwords to print (0 prints as much as possible)

	// select generator
	Secure bool // true: use completely random symbols, false: memorable

	// select symbol sets and filters
	UseNumbers  bool // true: must use numbers, false: don't care
	UseNoVowels bool // true: don't use vowels, false: don't care
	UseSpecials bool // true: use specials, false: don't care
	UseAmbigous bool // true: no letters with might be confused, false: don't care
	UseCapitals bool // true: must use a capital, false: don't care

	// output options
	PrintOnePerLine bool // print one password per line
	PrintColumns    bool // print passwords in columns

	// options controlling what we do
	ShowHelp     bool // show help and exit
	ShowCharsets bool // show charsets and exit

	// the arguments, unparsed
	Args []string
}

// default password length
const DefaultPwLength = 8

// ------------------------------------------------------------------

// Constructor
func createCli() Cli {
	var retval Cli
	retval.PwLength = DefaultPwLength
	retval.PwNum = 0

	retval.Secure = true

	retval.UseNumbers = false
	retval.UseNoVowels = false
	retval.UseSpecials = false
	retval.UseAmbigous = false
	retval.UseCapitals = false

	retval.PrintOnePerLine = false
	retval.PrintColumns = true

	retval.ShowHelp = false
	retval.ShowCharsets = false

	return retval
}

// Default Constructor
func NewCli() Cli {
	retval := createCli()
	retval.Args = os.Args
	return retval
}

func ParseNewCli() (Cli, error) {
	retval := NewCli()
	err := retval.Parse()
	return retval, err
}

func NewCliFromString(args []string) Cli {
	retval := createCli()
	retval.Args = args
	return retval
}

func ParseNewCliFromString(args []string) (Cli, error) {
	retval := NewCliFromString(args)
	err := retval.Parse()
	return retval, err
}

// parse the command line stored in this Cli object.
// If there was a problem, return an error object,
// otherwise the caller can assume anything was ok.
func (c *Cli) Parse() error {
	var idx int
	var val string

	for idx = 1; idx < len(c.Args); idx++ {
		val = c.Args[idx]
		switch {
		case val == "-l":
			idx++
			val2 := c.Args[idx]
			if ival64, err := strconv.ParseInt(val2, 10, 32); err == nil {
				c.PwLength = int(ival64)
			} else {
				return fmt.Errorf("not a password length: \"%s\" (%s)", val2, err)
			}
		case val == "-0" || val == "--no-numerals":
			c.UseNumbers = false
		case val == "-n" || val == "--numerals":
			c.UseNumbers = true
		case val == "-1":
			c.PrintOnePerLine = true
		case val == "-A" || val == "--no-capitalize":
			c.UseCapitals = false
		case val == "--capitalize":
			c.UseCapitals = true
		case val == "-a" || val == "--alt-phonics":
			// ignore
		case val == "-B" || val == "--ambigous":
			c.UseAmbigous = true
		case val == "-C":
			c.PrintColumns = true
		case val == "-N" || val == "--num-password":
			idx++
			val2 := c.Args[idx]
			if ival64, err := strconv.ParseInt(val2, 10, 32); err == nil {
				c.PwNum = int(ival64)
			} else {
				return fmt.Errorf("not a number of passwords: \"%s\" (%s)", val2, err)
			}
		case val == "-H" || val == "--sha1":
			// ignore
		case val == "-h" || val == "--help":
			c.ShowHelp = true
		case val == "-s" || val == "--secure":
			c.Secure = true
		case val == "-v" || val == "--no-vowels":
			c.UseNoVowels = true
		case val == "-y" || val == "--symbols":
			c.UseSpecials = true
		default:
			return fmt.Errorf("unknwon argument \"%s\" at %d", val, idx)
		}
	}
	return nil
}
