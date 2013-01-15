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
        if num := screen.CalcPasswordsPerScreen(pwlen); num != expected {
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
        if num := screen.CalcNumColumns(pwlen); num != expected {
            t.Errorf("expected %d columns of %d chars to fit in %d, but got %d",
                expected, pwlen, screen.ws_col, num)
        }
    }

    tester( 6, 11)
    tester( 8,  9)
    tester(10,  7)
    tester(70,  1)
    tester(80,  1)
    tester(90,  0)

    screen.ws_col = 10
    tester(2, 3)
    tester(3, 2)
    tester(5, 1)

    screen.ws_col = 11
    tester(2, 4)
    tester(3, 3)
    tester(5, 2)
}

func TestCalcLinesPerPassword(t *testing.T) {
    screen := DefaultWinSize()
    if screen.String() != "[24, 80]" {
        t.Fatalf("expected screen size to be [24, 80], but got %s", screen)
    }

    tester := func(pwlen, expected int) {
        if num := screen.CalcLinesPerPassword(pwlen); num != expected {
            t.Errorf("expected %d rows for password of %d chars on %s but got %d",
                expected, pwlen, screen, num)
        }
    }

    tester( 6, 1)
    tester( 8, 1)
    tester(10, 1)
    tester(70, 1)
    tester(80, 1)
    tester(90, 2)
}

