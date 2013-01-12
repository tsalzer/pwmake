/*
 * command line utils
 */

package termcap

import (
    "syscall"
    "unsafe"
)

const (
    tiocgwinsz = 0x5413
)

type WinSize struct {
    Rows, Cols uint16
    XPixel, YPixel uint16
}

func GetTermSize() (ws WinSize) {
    syscall.Syscall(syscall.SYS_IOCTL,
        uintptr(0), uintptr(tiocgwinsz),
        uintptr(unsafe.Pointer(&ws)))
    return
}
