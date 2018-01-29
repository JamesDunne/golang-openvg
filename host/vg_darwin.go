//+build darwin

package host

/*
#cgo CFLAGS: -x objective-c -Iinclude
#cgo LDFLAGS: -framework Cocoa -framework OpenGL -framework QuartzCore -Llib/macosx/ub/gle/standalone -lAmanithVG

#include "macosx/main.m"
*/
import "C"

import "unsafe"

type Host struct {
	Width  int
	Height int

	app  uintptr
	view uintptr
}

func NewHost(width, height int) *Host {
	var app uintptr
	var view uintptr

	C.hostInit(C.int(width), C.int(height), unsafe.Pointer(&app), unsafe.Pointer(&view))

	return &Host{
		Width:  width,
		Height: height,
		app:    app,
		view:   view,
	}
}

func (h *Host) PollEvent() bool {
	return int32(C.hostPollEvent(unsafe.Pointer(h.app), unsafe.Pointer(h.view))) == 0
}
