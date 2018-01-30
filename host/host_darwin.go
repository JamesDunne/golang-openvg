//+build darwin

package host

/*
#cgo CFLAGS: -x objective-c -Iinclude
#cgo LDFLAGS: -framework Cocoa -framework OpenGL -framework QuartzCore -Llib/macosx/ub/gle/standalone -lAmanithVG

#include "macosx/main.m"
*/
import "C"

import "unsafe"

var (
	Width  int
	Height int

	app  uintptr
	view uintptr
)

func Init(width, height int) bool {
	if C.hostInit(C.int(width), C.int(height), unsafe.Pointer(&app), unsafe.Pointer(&view)) == 0 {
		return false
	}

	Width = width
	Height = height
	return true
}

func Destroy() {

}

func PollEvent() bool {
	return int32(C.hostPollEvent(unsafe.Pointer(app), unsafe.Pointer(view))) == 0
}
