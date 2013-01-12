package columns

import (
    "testing"
)

func TestCalcPasswordsPerScreen(t *testing.T) {
    screen := DefaultWinSize()
    if screen.String() != "[24, 80]" {
        t.Fatalf("expected screen size to be [24, 80], but got %s", screen)
    }

    tester := func(pwlen int, expected int) {
        if num := CalcPasswordsPerScreen(pwlen, screen); num != expected {
            t.Errorf("expected %d passwords of %d characters to fit on %s, but got %d",
                expected, pwlen, screen, num)
        }
    }

    tester( 6, 11 * 24)
    tester( 8,  9 * 24)
    tester(10,  7 * 24)
    tester(70,  1 * 24)
    tester(80,  1 * 24)
    tester(90,  1 * 12)

    screen.ws_col = 10
    tester(2, 3 * 24)
    tester(3, 2 * 24)
    tester(5, 1 * 24)

    screen.ws_col = 11
    tester(2, 4 * 24)
    tester(3, 3 * 24)
    tester(5, 2 * 24)
}

func TestCalcNumColumns(t *testing.T) {
    screen := DefaultWinSize()
    if screen.String() != "[24, 80]" {
        t.Fatalf("expected screen size to be [24, 80], but got %s", screen)
    }

    tester := func(pwlen, expected int) {
        if num := CalcNumColumns(pwlen, screen); num != expected {
            t.Errorf("expected %d columns of %d chars to fit in %d, but got %d",
                expected, pwlen, screen.ws_col, num)
        }
    }

    tester( 6, 11)
    tester( 8,  9)
    tester(10,  7)
    tester(70,  1)
    tester(80,  1)
    tester(90,  1)

    screen.ws_col = 10
    tester(2, 3)
    tester(3, 2)
    tester(5, 1)

    screen.ws_col = 11
    tester(2, 4)
    tester(3, 3)
    tester(5, 2)
}


func TestBuildColumns(t *testing.T) {
    screen := DefaultWinSize()
    if screen.String() != "[24, 80]" {
        t.Fatalf("expected screen size to be [24, 80], but got %s", screen)
    }

    pwlen := 8
    fn := func() string { return "12345678" }
    expected := "12345678 12345678 12345678 12345678 12345678 12345678 12345678 12345678 12345678"

    if line := BuildColumns(pwlen, screen, fn); line != expected {
        t.Errorf("expected line\n\"%s\"\nbut got\n\"%s\"", expected, line)
    }
}

func TestBuildScreen(t *testing.T) {
    screen := DefaultWinSize()
    if screen.String() != "[24, 80]" {
        t.Fatalf("expected screen size to be [24, 80], but got %s", screen)
    }

    pwlen := 8
    fn := func() string { return "12345678" }
    expected := "12345678 12345678 12345678 12345678 12345678 12345678 12345678 12345678 12345678"

    count := 0
    lines := BuildScreen(pwlen, screen, fn)

    for _,line := range lines {
        count++
        if line != expected {
            t.Errorf("line %d: expected line\n\"%s\"\nbut got\n\"%s\"", count, expected, line)
        }
    }

    if count != int(screen.ws_row) {
        t.Errorf("expected %d lines, but got %d", screen.ws_row, count)
    }
}

