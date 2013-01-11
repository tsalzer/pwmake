package columns

// simple calculation: How many columns fit into a row?
// TODO: consider that the final column need no additional space
func CalcNumColumns(colw int, scrnw int) int {
    combined := colw + 1    // 1 space between the columns
    num := scrnw / combined
    return num
}

// build a line of columns using the given generator function.
func BuildColumns(colw int, scrnw int, fn func() string) string {
    retval := ""
    num := CalcNumColumns(colw, scrnw)
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
func BuildScreen(colw int, scrnw int, scrnh int, fn func() string) []string {
    retval := make([]string, scrnh)
    for i := 0; i < scrnh; i++ {
        retval[i] = BuildColumns(colw, scrnw, fn)
    }
    return retval
}

