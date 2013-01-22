package screen

import (
    "testing"
    "fmt"
)

func TestSendPasswordToFunc(t *testing.T) {
    screen := DefaultWinSize()
    if screen.String() != "[24, 80]" {
        t.Fatalf("expected screen size to be [24, 80], but got %s", screen)
    }

    tcase := ""
    fnPwd := func() (string,error) { return "12345678",nil }

    tester := func(num int, pwlen int, expected []string) {
        count := 0
        explen := len(expected)

        fnCb := func(pwd string) error {
            // check we're not exceeding the number of expected calls
            count++
            if count > explen {
                return fmt.Errorf("%s: expected %s passwords, but got %s as number %d",
                    tcase, explen, pwd, count)
            }

            // check if the string looks as expected
            exppwd := expected[count - 1]
            if pwd != exppwd {
                return fmt.Errorf("%s: at %d of %d in %s: got \"%s\" but expected \"%s\"",
                    tcase, count, num, screen, pwd, exppwd)
            }
            // all is well
            return nil
        }

        if err := screen.sendPasswordToFunc(num, pwlen, fnPwd , fnCb); err != nil {
            t.Errorf("sendPasswordToFunc returned error: %s", err)
        } else {
            if count != explen {
                t.Errorf("got %d passwords but expected %d", count, explen)
            }
        }
    }

    // regular 80x24
    tcase = "regular 80x24 screen"
    tester(1, 8, []string{"12345678\n"})
    tester(2, 8, []string{"12345678 ", "12345678\n"})

    // 16 cols, only room for one password:
    // 8(pw) + 1(space) + 8(pw) = 17
    tcase = "16col screen, single pwd only"
    screen.ws_col = 16
    tester(2, 8, []string{"12345678\n", "12345678\n"})
    tester(3, 8, []string{"12345678\n", "12345678\n", "12345678\n"})
    tester(4, 8, []string{"12345678\n", "12345678\n", "12345678\n", "12345678\n"})

    // 17 cols, room for exactly 2 passwords
    tcase = "17col screen, two pwds plus one space"
    screen.ws_col = 17
    tester(2, 8, []string{"12345678 ", "12345678\n"})
    tester(3, 8, []string{"12345678 ", "12345678\n", "12345678\n"})
    tester(4, 8, []string{"12345678 ", "12345678\n", "12345678 ", "12345678\n"})

    // 18 cols, plenty of room for 2 passwords
    tcase = "18col screen, two pwds plus two spaces"
    screen.ws_col = 18
    tester(2, 8, []string{"12345678 ", "12345678\n"})
    tester(3, 8, []string{"12345678 ", "12345678\n", "12345678\n"})
    tester(4, 8, []string{"12345678 ", "12345678\n", "12345678 ", "12345678\n"})
}

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

