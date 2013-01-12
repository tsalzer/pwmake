package columns

import (
    "fmt"
)

type winsize struct {
    ws_row, ws_col uint16
    ws_xpixel, ws_ypixe uint16
}

func NewWinSize(ws_row, ws_col uint16) winsize {
    var retval winsize
    retval.ws_row = ws_row
    retval.ws_col = ws_col
    return retval
}

func DefaultWinSize() winsize {
    return NewWinSize(24, 80)
}

func (ws winsize) String() string {
    return fmt.Sprintf("[%d, %d]", ws.ws_row, ws.ws_col)
}

