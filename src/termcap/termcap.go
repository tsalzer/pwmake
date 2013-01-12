/*
 * command line utils
 */

package termcap

import (
	"os"
    "syscall"
    "unsafe"
)

type WinSize struct {
    Row, Col uint16
    XPixel, YPixel uint16
}

func GetTermSize() (ws WinSize) {
    syscall.Syscall(syscall.SYS_IOCTL,
        uintptr(os.Stdin.Fd()), uintptr(syscall.TIOCGWINSZ),
        uintptr(unsafe.Pointer(&ws)))
    return
}
