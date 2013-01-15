package columns

import (
    "fmt"
)

// Print generated passwords to this screen.
func (ws winsize) PrintPasswords(pwlen int, fn func() (string, error)) error {
    num := ws.CalcPasswordsPerScreen(pwlen)
    cols := ws.CalcNumColumns(pwlen)
    col := 0

    var pwd string
    var err error

    for i := 0; i < num ; i++ {
        if pwd, err = fn(); err == nil {
            if col + 1 < cols {
                fmt.Printf("%s ", pwd)
                col++
            } else {
                fmt.Printf("%s\n", pwd)
                col = 0
            }
        } else {
            return err
        }
    }
    return nil
}

// How many passwords of a given length can I display with a given screen size?
// This comes down to two questions: How many passwords can be displayed per
// line, and how many lines are there?
// All bets are off if the password is longer than the screen is wide, but we
// still try to be reasonable.
// Regardless of the password length, we can always print at least one password.
func (ws winsize) CalcPasswordsPerScreen(pwlen int) int {
    retval := 1 // we will always cram in at least one password
    if pwlen <= int(ws.ws_col) {
        // we can fit at least one password into a line.
        perline := ws.CalcNumColumns(pwlen)
        retval = perline * int(ws.ws_row)
    } else {
        // a single password will use at least the whole line:
        if chunks := int(ws.ws_row) / ws.CalcLinesPerPassword(pwlen) ; chunks > 0 {
            return chunks
        }
	}
    return retval
}

// How many passwords fit completely into a single line, considering 1 space between
// two passwords?
// If the password is longer than the linne, this will return 0.
func (ws winsize) CalcNumColumns(pwlen int) int {
    screencols := int(ws.ws_col)
    cols  := screencols / (pwlen + 1)
    spare := screencols % (pwlen + 1)
    //gaps  := cols - 1
    if spare >= pwlen {
        return cols + 1
    }
    return cols
}
// How many lines will this password take on this screen?
// In almost all cases (screen cols <= password length) this is 1, otherwise
// the number of rows.
func (ws winsize) CalcLinesPerPassword(pwlen int) int {
    screencols := int(ws.ws_col)
    if pwlen <= screencols {
        return 1
    }
    rows := pwlen / screencols
    if mod := pwlen % screencols; mod > 0 {
        return rows + 1
    }
    return rows
}


// build a line of columns using the given generator function.
func (ws winsize) BuildColumns(colw int, fn func() string) string {
    retval := ""
    num := ws.CalcNumColumns(colw)
    lastcol := num - 1
    for i := 0; i < num; i++ {
        retval += fn()
        if i < lastcol {
            retval += " "
        }
    }
    return retval
}

// build a number of lines.
func (ws winsize) BuildScreen(colw int, fn func() string) []string {
    retval := make([]string, ws.ws_row)
    for i := 0; i < int(ws.ws_row); i++ {
        retval[i] = ws.BuildColumns(colw, fn)
    }
    return retval
}

