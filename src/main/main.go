/*
 * main CLI
 */

// The main package, used to run the CLI.
package main

import (
	"flag"
	"fmt"
	"os"
	"pwdgen"
	"pwdgen/symbol"
	"pwdgen/columns"
)

var flagLength int
var flagShowCharsets bool

// command line parser
func init() {
	flag.IntVar(&flagLength, "l", 10, "length of the password to generate")
	flag.Parse()
}

// print a screen of passwords
func PrintScreen() error {
	pwlen := flagLength
	screen := columns.DefaultWinSize()

	fn := func() (string, error) {
		var pwd string
		var err error
		if pwd, err = pwdgen.GeneratePassword(pwlen); err != nil {
			return pwd, err
		}
		return pwd, nil
	}
	return screen.PrintPasswords(pwlen, fn)
	// lines := screen.BuildScreen(flagLength, fn)

	// for _,line := range(lines) {
	//     fmt.Printf("%s\n", line)
	// }
	// return nil
}

// print ia single password.
// This will use GeneratePassword to generate a password.
// Any error from the password generator will be relayed here and can be printed
// to the user.
func PrintPassword() error {
	if pwd, err := pwdgen.GeneratePassword(flagLength); err != nil {
		return err
	} else {
		fmt.Printf("%s\n", pwd)
	}
	return nil
}

// generate the password,
// This will generate a password with the given specifications, and return it.
// Any error from the password generator will be relayed here and can be printed
// to the user.
func GeneratePassword() (string, error) {
	var gen *pwdgen.PwdGen
	var symset *symbol.MultiSet
	var err error

	if symset, err = symbol.NewMultiSetFromDefaults([]string{"alpha", "ALPHA", "num"}); err != nil {
		return "", err
	}

	if gen, err = pwdgen.NewPwdGen(symset, flagLength); err != nil {
		return "", err
	}
	return gen.String(), nil
}

// main.
// This is, you know, main.
func main() {
	if err := PrintScreen(); err != nil {
		fmt.Printf("Problem generating password: %s\n", err)
		os.Exit(1)
	}

	// if err := PrintPassword() ; err != nil {
	//     fmt.Printf("%s\n", err)
	//     os.Exit(1)
	// }
}
