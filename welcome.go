/*
 * Even Less Than Hello World.
 */

package main

import (
    "os"
    "fmt"
    "log"
    "flag"
)

var flagLog bool

func usage() {
    if flagLog {
        log.Printf(`Yes yes, you are welcome. Happy now?`)
    } else {
        fmt.Fprintf(os.Stderr, `Welcome.
`)
    }
}

func init() {
    flag.BoolVar(&flagLog, "log", false, "log something else")
    flag.Parse()
}

func main() {
    usage()
}

