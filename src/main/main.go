/*
 * main CLI
 */

// The main package, used to run the CLI.
package main

import (
	"fmt"
	"os"
	"pwdgen"
	"pwdgen/cli"
	"pwdgen/symbol"
	"pwdgen/screen"
)

var cmdline cli.Cli

// command line parser
func init() {
	cmdline = cli.NewCli()
}

// print a screen of passwords
func PrintScreen() error {
	pwlen := cmdline.PwLength
	screen := screen.DefaultWinSize()

	fn := func() (string, error) {
		var pwd string
		var err error
		if pwd, err = pwdgen.GeneratePassword(pwlen); err != nil {
			return pwd, err
		}
		return pwd, nil
	}
	return screen.PrintPasswords(pwlen, fn)
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

	if gen, err = pwdgen.NewPwdGen(symset, cmdline.PwLength); err != nil {
		return "", err
	}
	return gen.String(), nil
}

// main.
// This is, you know, main.
func main() {
	if err := cmdline.Parse(); err != nil {
		fmt.Printf("argument error: %s\n", err)
	} else {
		if err := PrintScreen(); err != nil {
			fmt.Printf("Problem generating password: %s\n", err)
			os.Exit(1)
		}
	}
}
