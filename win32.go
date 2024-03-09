package streamd

import (
	"errors"
	"os/exec"
	"syscall"
	"unsafe"
)

var (
	user32DLL = syscall.NewLazyDLL("user32.dll")
	// BOOL GetCursorPos([out] LPPOINT lpPoint)
	getCursorPos = user32DLL.NewProc("GetCursorPos")
)

type LPPoint struct {
	x int32
	y int32
}

// GetCursorPos Retrieves the position of the mouse cursor, in screen coordinates.
func GetCursorPos() (LPPoint, error) {
	lpPoint := LPPoint{}

	success, _, err := getCursorPos.Call(uintptr(unsafe.Pointer(&lpPoint)))
	successValue := (uint64)(success)

	if successValue == 0 {
		err = errors.New("GetCursorPos returned 0")
		return lpPoint, err
	}

	return lpPoint, err
}

func OpenInBrowser(url string) error {
	err := exec.Command("rundll32", "url.dll,FileProtocolHandler", url).Start()

	return err
}
