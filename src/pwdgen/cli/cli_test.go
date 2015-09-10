package cli

import (
	"reflect"
	"testing"
)

// ------------------------------------------------------------------

func concatStringSlices(old1, old2 []string) []string {
	newslice := make([]string, len(old1)+len(old2))
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
	for _, arg := range args {
		cliLoader(t, []string{arg}, cb)
	}
}

func cliCheckBool(t *testing.T, args []string, name string, expval bool) {
	cliLoaderSingles(t, args, func(c Cli, a []string) {
		ps := reflect.ValueOf(&c) // pointer to struct
		s := ps.Elem()            // struct
		field := s.FieldByName(name)
		if field.IsValid() {
			val := field.Bool()
			if val != expval {
				t.Errorf("expected %s to set %s to %s, got %s", a, name, expval, val)
			}
		} else {
			t.Errorf("field %s not found, cannot test this", name)
		}
	})

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
	loader([]string{"-l", "10"}, func() {
		if cl.PwLength != 10 {
			t.Errorf("expected %s to set password length to 10, got %d", args, cl.PwLength)
		}
	})

	cliCheckBool(t, []string{"-0", "--no-numerals"}, "UseNumbers", false)
	cliCheckBool(t, []string{"-1"}, "PrintOnePerLine", true)
	cliCheckBool(t, []string{"-n", "--numerals"}, "UseNumbers", true)
	cliCheckBool(t, []string{"--capitalize"}, "UseCapitals", true)
	cliCheckBool(t, []string{"-A", "--no-capitalize"}, "UseCapitals", false)
	cliCheckBool(t, []string{"-B", "--ambigous"}, "UseAmbigous", true)
	cliCheckBool(t, []string{"-C"}, "PrintColumns", true)
	cliCheckBool(t, []string{"-s", "--secure"}, "Secure", true)
	cliCheckBool(t, []string{"-v", "--no-vowels"}, "UseNoVowels", true)
	cliCheckBool(t, []string{"-y", "--symbols"}, "UseSpecials", true)
	cliCheckBool(t, []string{"-h", "--help"}, "ShowHelp", true)

}

func TestPositionalArguments(t *testing.T) {
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

	loader([]string{"10"}, func() {
		if (cl.PwLength != 10) {
			t.Errorf("expected pw_length to set password length to 10, got %d", cl.PwLength)
		}
	})

	loader([]string{"10", "4"}, func() {
		if (cl.PwLength != 10) {
			t.Errorf("expected pw_length to set password length to 10, got %d", cl.PwLength)
		}

		if (cl.PwNum != 4) {
			t.Errorf("expected num_pw to set number of passwords  to 4, got %d", cl.PwNum)
		}
	})

	loader([]string{"-A", "10", "4"}, func() {
		if (cl.PwLength != 10) {
			t.Errorf("expected pw_length to set password length to 9, got %d", cl.PwLength)
		}

		if (cl.PwNum != 4) {
			t.Errorf("expected num_pw to set number of passwords  to 4, got %d", cl.PwNum)
		}

		if (cl.UseCapitals != false) {
			t.Errorf("expected -A to set UseCapitals to false, got true")
		}
	})

}
