/*
 * Even Less Than Hello World.
 */

package main

import (
    "os"
    "fmt"
)

func usage() {
    fmt.Fprintf(os.Stderr, `Welcome.
`)
}

func main() {
    usage()
}

