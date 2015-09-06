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
)

var flagLength int
var flagShowCharsets bool

// command line parser
func init() {
	flag.IntVar(&flagLength, "l", 10, "length of the password to generate")
	flag.Parse()
}

// print the password.
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
func _GeneratePassword() (string, error) {
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
	if err := PrintPassword(); err != nil {
		fmt.Printf("%s\n", err)
		os.Exit(1)
	}
}
