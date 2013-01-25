package screen

import (
    "fmt"
)


// simply print a string.
// we need this in different places, here.
func simplePrint(str string) error {
    fmt.Printf(str)
    return nil
}

// send a password to a given function.
// We use this to test the PrintPasswordsToScreen method.
func (ws winsize) sendPasswordsToFunc(num int, pwlen int,
        fnGenPwd func() (string, error),
        fnPrint  func(string) error) error {
    var err error
    var dcol int
    var pwd string
    var result string

    col := 0
    wscol := int(ws.ws_col)

    for count := 0; count < num; count ++ {
        if pwd,err = fnGenPwd(); err != nil {
            // an error, return immediately
            return err
        } else {
            // not an error
            col += pwlen + 1
            dcol = col
            // fmt.Printf("col=%d", col)
            if (count + 1 == num) || (col + pwlen > wscol) {
                col = 0
                // fmt.Printf(" reset\n")
                result = fmt.Sprintf("%s\n", pwd)
            } else {
                // fmt.Printf(" ")
                result = fmt.Sprintf("%s ", pwd)
            }

            if err = fnPrint(result); err != nil {
                return fmt.Errorf("count=%d, num=%d, col=%d, col+pwlen=%d, ws_col=%d - %s",
                    count, num, dcol, dcol + pwlen, ws.ws_col, err)
            }
        }
    }
    return nil
}

// Print generated passwords to this screen.
func (ws winsize) sendPasswordsToGrid(pwlen int,
        fnGenPw func() (string, error),
        fnPrintPw func(string) error) error {
    num := ws.CalcPasswordsPerScreen(pwlen)
    cols := ws.CalcNumColumns(pwlen)
    col := 0

    var pwd string
    var err error
    var pwprint string

    for i := 0; i < num ; i++ {
        if pwd, err = fnGenPw(); err == nil {
            if col + 1 < cols {
                pwprint = fmt.Sprintf("%s ", pwd)
                col++
            } else {
                pwprint = fmt.Sprintf("%s\n", pwd)
                col = 0
            }
            if err = fnPrintPw(pwprint); err != nil {
                return err
            }
        } else {
            return err
        }
    }
    return nil
}

// Print generated passwords to this screen.
func (ws winsize) PrintPasswords(pwlen int, fn func() (string, error)) error {
    return ws.sendPasswordsToFunc(ws.CalcPasswordsPerScreen(pwlen), pwlen, fn, simplePrint)
}

// Print NUM generated passwords to this screen.
func (ws winsize) PrintNumPasswords(num int, pwlen int, fn func() (string, error)) error {
    return ws.sendPasswordsToFunc(num, pwlen, fn, simplePrint)
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

