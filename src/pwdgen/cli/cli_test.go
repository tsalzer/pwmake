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

