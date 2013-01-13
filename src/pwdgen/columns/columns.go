package columns

// How many passwords of a given length can I display with a given screen size?
// This comes down to two questions: How many passwords can be displayed per
// line, and how many lines are there?
// All bets are off if the password is longer than the screen is wide, but we
// still try to be reasonable.
// Regardless of the password length, we can always print at least one password.
func CalcPasswordsPerScreen(pwlen int, screen winsize) int {
    retval := 1 // we will always cram in at least one password
    if pwlen <= int(screen.ws_col) {
        // we can fit at least one password into a line.
        perline := CalcNumColumns(pwlen, screen)
        retval = perline * int(screen.ws_row)
    }
    return retval
}

// How many passwords fit completely into a single line, considering 1 space between
// two passwords?
// If the password is longer than the linne, this will return 0.
func CalcNumColumns(pwlen int, ssize winsize) int {
    screencols := int(ssize.ws_col)
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
func CalcLinesPerPassword(pwlen int, ssize winsize) int {
    screencols := int(ssize.ws_col)
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
func BuildColumns(colw int, ssize winsize, fn func() string) string {
    retval := ""
    num := CalcNumColumns(colw, ssize)
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
func BuildScreen(colw int, ssize winsize, fn func() string) []string {
    retval := make([]string, ssize.ws_row)
    for i := 0; i < int(ssize.ws_row); i++ {
        retval[i] = BuildColumns(colw, ssize, fn)
    }
    return retval
}

