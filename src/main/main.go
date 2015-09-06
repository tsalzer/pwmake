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

// print a screen of num passwords
func PrintScreen(num int) error {
	pwlen := cmdline.PwLength
	screen := screen.DefaultWinSize()
	if num == 0 {
		num = screen.CalcPasswordsPerScreen(pwlen)
	}

	fn := func() (string, error) {
		var pwd string
		var err error
		if pwd, err = pwdgen.GeneratePassword(pwlen); err != nil {
			return pwd, err
		}
		return pwd, nil
	}
	return screen.PrintNumPasswords(num, pwlen, fn)
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
		if cmdline.ShowHelp {
			fmt.Printf("Yes, I assume there should be some help. Try -l NUM for password lengths.\n")
		} else {
			if err := PrintScreen(cmdline.PwNum); err != nil {
				fmt.Printf("Problem generating password: %s\n", err)
				os.Exit(1)
			}
		}
	}
}
