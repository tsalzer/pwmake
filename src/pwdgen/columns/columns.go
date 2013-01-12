package columns

// How many passwords of a given length can I display with a given screen size?
// All bets are off if the password is longer than the screen is wide, but we
// still try to be reasonable.
// Regardless of the password length, we can always print at least one password.


// simple calculation: How many columns fit into a row?
// TODO: consider that the final column need no additional space
func CalcNumColumns(colw int, ssize winsize) int {
    combined := colw + 1    // 1 space between the columns
    num := int(ssize.ws_col) / combined
    return num
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

