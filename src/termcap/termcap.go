/*
 * command line utils
 */

package termcap

import (
	"os"
	"errors"
    "syscall"
    "unsafe"
)

type WinSize struct {
    Row, Col uint16
    XPixel, YPixel uint16
}

func GetTermSize() (*WinSize, error) {
    ws := &WinSize{}
    _, _, errno := syscall.Syscall(syscall.SYS_IOCTL,
        uintptr(os.Stdin.Fd()), uintptr(syscall.TIOCGWINSZ),
        uintptr(unsafe.Pointer(&ws)))

    if errno != 0 {
        return nil, errors.New(errno.Error())
    }

    return ws, nil
}
