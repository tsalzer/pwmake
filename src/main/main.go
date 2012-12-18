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
)

var flagLength int
var flagShowCharsets bool


// command line parser
func init() {
    flag.IntVar(&flagLength, "l", 10, "length of the password to generate")
    flag.Parse()
}

// print the password.
// This will generate a password with the given specifications, and print it.
// Any error from the password generator will be relayed here and can be printed
// to the user.
func PrintPassword() error {
    var gen *pwdgen.PwdGen
    var symset *pwdgen.SymbolSet
    var err error

    if symset, err = pwdgen.GetSymbolSet("alpha"); err != nil {
        return err
    }

    if gen, err = pwdgen.NewPwdGen(symset, flagLength); err != nil {
       return err
    }
    fmt.Printf("%s\n", gen)
    return nil
}

// main.
// This is, you know, main.
func main() {
    if err := PrintPassword() ; err != nil {
        fmt.Printf("%s\n", err)
        os.Exit(1)
    }
}

