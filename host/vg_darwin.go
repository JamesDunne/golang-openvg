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

func Init(width, height int) {
	C.hostInit(C.int(width), C.int(height), unsafe.Pointer(&app), unsafe.Pointer(&view))

	Width = width
	Height = height
}

func PollEvent() bool {
	return int32(C.hostPollEvent(unsafe.Pointer(app), unsafe.Pointer(view))) == 0
}
