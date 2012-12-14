/*
 * main CLI
 */

package main

import (
    "flag"
    "fmt"
    "os"
    "pwdgen"
)

var flagLength int
var gen *pwdgen.PwdGen
var genErr error

func init() {
    flag.IntVar(&flagLength, "l", 10, "length of the password to generate")
    flag.Parse()
    gen, genErr = pwdgen.NewPwdGen("abcdefghijklmnopqrstuvwxyz", flagLength)

    /*
    if gen, genErr = pwdgen.NewPwdGen("abcdefghijklmnopqrstuvwxyz", flagLength); err != nil {
        panic(err)
    }
    */
}

func main() {
    if gen == nil {
        fmt.Printf("%s\n", genErr)
        os.Exit(1)
    } else {
        fmt.Printf("%s\n", gen)
    }
}

