package cli

import (
	"testing"
)

func TestNewArgSet(t *testing.T) {
	s := NewArgSet()
	if len(s.Args) != 0 {
		t.Errorf("got some argument at construction time: %s", s.Args)
	}
	if s.Len() != len(s.Args) {
		t.Errorf("ArgSet#Len() reports %d, whicle len(Args) report %d", s.Len(), len(s.Args))
	}
}

func TestAddBoolean(t *testing.T) {
	s := NewArgSet()
	if err := s.AddBoolean("-N", "--numerical", "use numerical symbols", "UseNumerical", true); err != nil {
		t.Errorf("failed to add a bool option: %s", err)
	}

	if len(s.Args) != 1 {
		t.Errorf("adding a boolean value failed")
	}

	if s.Len() != len(s.Args) {
		t.Errorf("ArgSet#Len() reports %d, whicle len(Args) report %d", s.Len(), len(s.Args))
	}
}

func TestAddInt(t *testing.T) {
	s := NewArgSet()
	if err := s.AddInt("-l", "--length", "generate passwords of the given length", "PwLength", 8); err != nil {
		t.Errorf("failed to add an int option: %s", err)
	}

	if len(s.Args) != 1 {
		t.Errorf("adding an int value failed")
	}

	if s.Len() != len(s.Args) {
		t.Errorf("ArgSet#Len() reports %d, whicle len(Args) report %d", s.Len(), len(s.Args))
	}
}
