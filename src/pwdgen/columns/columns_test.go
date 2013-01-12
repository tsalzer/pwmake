package columns

import (
    "testing"
)

func TestCalcNumColumns(t *testing.T) {
    screensize := DefaultWinSize()
    colw := 8
    expected := 8

    if num := CalcNumColumns(colw, screensize); num != expected {
        t.Errorf("expected %d columns, but got %d", expected, num)
    }
}


func TestBuildColumns(t *testing.T) {
    screensize := DefaultWinSize()
    colw := 8
    fn := func() string { return "12345678" }
    expected := "12345678 12345678 12345678 12345678 12345678 12345678 12345678 12345678"

    if line := BuildColumns(colw, screensize, fn); line != expected {
        t.Errorf("expected line\n\"%s\"\nbut got\n\"%s\"", expected, line)
    }
}

func TestBuildScreen(t *testing.T) {
    screensize := DefaultWinSize()
    colw := 8
    fn := func() string { return "12345678" }
    expected := "12345678 12345678 12345678 12345678 12345678 12345678 12345678 12345678"

    count := 0
    lines := BuildScreen(colw, screensize, fn)

    for _,line := range lines {
        count++
        if line != expected {
            t.Errorf("line %d: expected line\n\"%s\"\nbut got\n\"%s\"", count, expected, line)
        }
    }

    if count != int(screensize.ws_row) {
        t.Errorf("expected %d lines, but got %d", screensize.ws_row, count)
    }
}

