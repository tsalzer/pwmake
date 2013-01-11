package columns

import (
    "testing"
)

func TestCalcNumColumns(t *testing.T) {
    colw := 8
    scrnw := 80
    expected := 8

    if num := CalcNumColumns(colw, scrnw); num != expected {
        t.Errorf("expected %d columns, but got %d", expected, num)
    }
}


func TestBuildColumns(t *testing.T) {
    colw := 8
    scrnw := 80
    fn := func() string { return "12345678" }
    expected := "12345678 12345678 12345678 12345678 12345678 12345678 12345678 12345678"

    if line := BuildColumns(colw, scrnw, fn); line != expected {
        t.Errorf("expected line\n\"%s\"\nbut got\n\"%s\"", expected, line)
    }
}

func TestBuildScreen(t *testing.T) {
    colw := 8
    scrnw := 80
    scrnh := 24
    fn := func() string { return "12345678" }
    expected := "12345678 12345678 12345678 12345678 12345678 12345678 12345678 12345678"

    count := 0
    lines := BuildScreen(colw, scrnw, scrnh, fn)

    for _,line := range lines {
        count++
        if line != expected {
            t.Errorf("line %d: expected line\n\"%s\"\nbut got\n\"%s\"", count, expected, line)
        }
    }

    if count != scrnh {
        t.Errorf("expected %d lines, but got %d", scrnh, count)
    }
}

