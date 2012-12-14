/*
 * main CLI
 */

package main

import (
    "flag"
    "fmt"
    "pwdgen"
)

var flagDebug bool
var gen *pwdgen.PwdGen

func init() {
    flag.BoolVar(&flagDebug, "d", false, "debug mode")
    flag.Parse()
    gen = pwdgen.NewPwdGen("abcdefghijklmnopqrstuvwxyz", 10)
}

func main() {
    fmt.Printf("Will create a password for you shortly...\n")
    fmt.Printf("password: %s\n", gen)
}

