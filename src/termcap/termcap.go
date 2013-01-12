/*
 * command line utils
 */

package termcap

import (
    "syscall"
    "unsafe"
)

type WinSize struct {
    Row, Col uint16
    XPixel, YPixel uint16
}

func GetTermSize() (ws WinSize) {
    syscall.Syscall(syscall.SYS_IOCTL,
        uintptr(0), uintptr(syscall.TIOCGWINSZ),
        uintptr(unsafe.Pointer(&ws)))
    return
}
