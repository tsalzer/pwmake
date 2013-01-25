package cli

import (
    "testing"
)

func concatStringSlices(old1, old2 []string) []string {
    newslice := make([]string, len(old1) + len(old2))
    copy(newslice, old1)
    copy(newslice[len(old1):], old2)
    return newslice
}
// 
// func cliLoader(t *testing.T, args []string, cb func(c Cli, a []string)) {
//     fullargs = concatStringClices([]string{"x"}, args)
//     if cl, err := ParseNewCliFromString(fullargs); err != nil {
//         t.Errorf("failed to parse %s, results in error: %s", args, err)
//     } else {
//         fn(cl, fullargs)
//     }
// }
// 
// func cliLoaderiSingle(t *testing.T, arg string, cb func(c Cli, a []string)) {
//     args = []string{arg}
//     cliLoader(args)
// }

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



    loader := func(tstargs []string, fn func()) {
        args = tstargs
        fullargs := concatStringSlices([]string{"x"}, args)
        if cl, err = ParseNewCliFromString(fullargs); err != nil {
            t.Errorf("failed to parse %s, results in error: %s", args, err)
        } else {
            fn()
        }
    }

    // password length, flags style
    loader([]string{"-l", "10"}, func(){ if cl.PwLength != 10 {
            t.Errorf("expected %s to set password length to 10, got %d", args, cl.PwLength)
        } })


    // cliLoaderSingle(t, "-0", func(c Cli) { if c.UseNumber != false {
    //     t.Errorf("expected %s to set UseNumber to false, got %s" a, c.UseNumbers)
    // }
    loader([]string{"-0"}, func(){ if cl.UseNumbers != false {
            t.Errorf("expected %s to set UseNumbers to false, got %b", args, cl.UseNumbers)
        } })
    loader([]string{"--no-numerals"}, func(){ if cl.UseNumbers != false {
            t.Errorf("expected %s to set UseNumbers to false, got %b", args, cl.UseNumbers)
        } })
    loader([]string{"-n"}, func(){ if cl.UseNumbers != true {
            t.Errorf("expected %s to set UseNumbers to true, got %b", args, cl.UseNumbers)
        } })
    loader([]string{"--numerals"}, func(){ if cl.UseNumbers != true {
            t.Errorf("expected %s to set UseNumbers to true, got %b", args, cl.UseNumbers)
        } })


    loader([]string{"-h"}, func(){ if cl.ShowHelp != true {
            t.Errorf("expected %s to set ShowHelp to true, got %b", args, cl.ShowHelp)
        } })
    loader([]string{"--help"}, func(){ if cl.ShowHelp != true {
            t.Errorf("expected %s to set ShowHelp to true, got %b", args, cl.ShowHelp)
        } })
}

