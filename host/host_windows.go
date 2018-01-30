//+build windows

package host

/*
#cgo CFLAGS: -Iinclude
#cgo LDFLAGS: -mwindows -s -Llib/win/x86_64/gle/standalone -lAmanithVG -lopengl32

#include "win/main.c"
*/
import "C"

var (
	Width  int
	Height int
)

func Init(width, height int) bool {
	if C.hostInit(C.int(width), C.int(height)) == 0 {
		return false
	}

	Width = width
	Height = height
	return true
}

func Destroy() int {
	return int(C.hostDestroy())
}

func PollEvent() bool {
	return int32(C.hostPollEvent()) == 0
}
