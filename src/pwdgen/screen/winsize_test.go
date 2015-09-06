package screen

import (
	"testing"
)

func TestWinsizeString(t *testing.T) {
	if s := NewWinSize(1, 2).String(); s != "[1, 2]" {
		t.Errorf("size.String() should be \"[1, 2]\" but is \"%s\"", s)
	}
	if s := NewWinSize(11, 15).String(); s != "[11, 15]" {
		t.Errorf("size.String() should be \"[11, 15]\" but is \"%s\"", s)
	}
	if s := NewWinSize(80, 24).String(); s != "[80, 24]" {
		t.Errorf("size.String() should be \"[80, 24]\" but is \"%s\"", s)
	}
}

func TestWinsizeConstructor(t *testing.T) {
	const col_max = 65535
	const row_max = 65535

	for col := 1; col < col_max; col += 64 {
		for row := 1; row < row_max; row += 64 {
			size := NewWinSize(uint16(row), uint16(col))
			if size.ws_col != uint16(col) {
				t.Fatalf("failed to set column to %d, size is %s", col, size)
			}
			if size.ws_row != uint16(row) {
				t.Fatalf("failed to set row to %d, size is %s", row, size)
			}
		}
	}
}
