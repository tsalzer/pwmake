package cli

import (
    "testing"
)

// ------------------------------------------------------------------

func concatStringSlices(old1, old2 []string) []string {
    newslice := make([]string, len(old1) + len(old2))
    copy(newslice, old1)
    copy(newslice[len(old1):], old2)
    return newslice
}

func cliLoader(t *testing.T, args []string, cb func(c Cli, a []string)) {
    fullargs := concatStringSlices([]string{"x"}, args)
    if cl, err := ParseNewCliFromString(fullargs); err != nil {
        t.Errorf("failed to parse %s, results in error: %s", args, err)
    } else {
        cb(cl, fullargs)
    }
}

func cliLoaderSingle(t *testing.T, arg string, cb func(c Cli, a []string)) {
    cliLoader(t, []string{arg}, cb)
}

func cliLoaderSingles(t *testing.T, args []string, cb func(c Cli, a []string)) {
    for _, arg := range(args) {
        cliLoader(t, []string{arg}, cb)
    }
}

// ------------------------------------------------------------------

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


    cliLoaderSingles(t, []string{"-0", "--no-numerals"}, func(c Cli, a []string){
        if c.UseNumbers != false {
            t.Errorf("expected %s to set UseNumbers to false, got %s", a, c.UseNumbers)
        } })

    cliLoaderSingles(t, []string{"-n", "--numerals"}, func(c Cli, a []string){
        if c.UseNumbers != true {
            t.Errorf("expected %s to set UseNumbers to true, got %s", a, c.UseNumbers)
        } })



    cliLoaderSingles(t, []string{"-h", "--help"}, func(c Cli, a []string){
        if c.ShowHelp != true {
            t.Errorf("expected %s to set ShowHelp to true, got %s", a, c.ShowHelp)
        } })
}

