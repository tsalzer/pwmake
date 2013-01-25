package cli

import (
    "testing"
)

func TestNewCli(t *testing.T) {
    cli := NewCli()

    if cli.PwLength != DefaultPwLength {
        t.Errorf("expected default password length %d, got %d", DefaultPwLength, cli.PwLength)
    }
}

func TestParseFriendly(t *testing.T) {
    var cl Cli
    var err error
    var args []string

    loader := func(args []string, fn func()) {
        if cl, err = ParseNewCliFromString(args); err != nil {
            t.Errorf("failed to parse %s, results in error: %s", args, err)
        } else {
            fn()
        }
    }

    // password length, flags style
    loader([]string{"x", "-l", "10"}, func(){ if cl.PwLength != 10 {
            t.Errorf("expected %s to set password length to 10, got %d", args, cl.PwLength)
        } })


    loader([]string{"x", "-0"}, func(){ if cl.UseNumbers != false {
            t.Errorf("expected %s to set UseNumbers to false, got %b", args, cl.UseNumbers)
        } })
    loader([]string{"x", "--no-numerals"}, func(){ if cl.UseNumbers != false {
            t.Errorf("expected %s to set UseNumbers to false, got %b", args, cl.UseNumbers)
        } })
    loader([]string{"x", "-n"}, func(){ if cl.UseNumbers != true {
            t.Errorf("expected %s to set UseNumbers to true, got %b", args, cl.UseNumbers)
        } })
    loader([]string{"x", "--numerals"}, func(){ if cl.UseNumbers != true {
            t.Errorf("expected %s to set UseNumbers to true, got %b", args, cl.UseNumbers)
        } })


    loader([]string{"x", "-h"}, func(){ if cl.ShowHelp != true {
            t.Errorf("expected %s to set ShowHelp to true, got %b", args, cl.ShowHelp)
        } })
    loader([]string{"x", "--help"}, func(){ if cl.ShowHelp != true {
            t.Errorf("expected %s to set ShowHelp to true, got %b", args, cl.ShowHelp)
        } })
}

