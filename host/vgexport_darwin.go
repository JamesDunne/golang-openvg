//+build darwin

package host

/*
#cgo CFLAGS: -x objective-c -Iinclude
#cgo LDFLAGS: -framework Cocoa -framework OpenGL -framework QuartzCore -Llib/macosx/ub/gle/standalone -lAmanithVG

#include "app.h"
*/
import "C"

// Known Go callback trampoline functions:

//export appInit
func appInit(surfaceWidth, surfaceHeight C.VGint) {
	if cb := InitFunc; cb != nil {
		cb(int32(surfaceWidth), int32(surfaceHeight))
	}
}

//export appDestroy
func appDestroy() {
	if cb := DestroyFunc; cb != nil {
		cb()
	}
}

//export appResize
func appResize(surfaceWidth, surfaceHeight C.VGint) {
	if cb := ResizeFunc; cb != nil {
		cb(int32(surfaceWidth), int32(surfaceHeight))
	}
}

//export appDraw
func appDraw(surfaceWidth, surfaceHeight C.VGint) {
	if cb := DrawFunc; cb != nil {
		cb(int32(surfaceWidth), int32(surfaceHeight))
	}
}

// handle mouse events

//export mouseLeftButtonDown
func mouseLeftButtonDown(x, y C.VGint) {
	if cb := MouseLeftButtonDownFunc; cb != nil {
		cb(int32(x), int32(y))
	}
}

//export mouseLeftButtonUp
func mouseLeftButtonUp(x, y C.VGint) {
	if cb := MouseLeftButtonUpFunc; cb != nil {
		cb(int32(x), int32(y))
	}
}

//export mouseRightButtonDown
func mouseRightButtonDown(x, y C.VGint) {
	if cb := MouseRightButtonDownFunc; cb != nil {
		cb(int32(x), int32(y))
	}
}

//export mouseRightButtonUp
func mouseRightButtonUp(x, y C.VGint) {
	if cb := MouseRightButtonUpFunc; cb != nil {
		cb(int32(x), int32(y))
	}
}

//export mouseMove
func mouseMove(x, y C.VGint) {
	if cb := MouseMoveFunc; cb != nil {
		cb(int32(x), int32(y))
	}
}

//export keyDown
func keyDown(k uint16) {
	if cb := KeyDownFunc; cb != nil {
		cb(k)
	}
}

//export keyUp
func keyUp(k uint16) {
	if cb := KeyUpFunc; cb != nil {
		cb(k)
	}
}

type SurfaceCallbackFunc func(width, height int32)
type VoidCallbackFunc func()
type PositionCallbackFunc func(x, y int32)
type KeyCallbackFunc func(k uint16)

var InitFunc SurfaceCallbackFunc = nil
var DestroyFunc VoidCallbackFunc = nil
var ResizeFunc SurfaceCallbackFunc = nil
var DrawFunc SurfaceCallbackFunc = nil
var MouseLeftButtonDownFunc PositionCallbackFunc = nil
var MouseLeftButtonUpFunc PositionCallbackFunc = nil
var MouseRightButtonDownFunc PositionCallbackFunc = nil
var MouseRightButtonUpFunc PositionCallbackFunc = nil
var MouseMoveFunc PositionCallbackFunc = nil
var KeyDownFunc KeyCallbackFunc = nil
var KeyUpFunc KeyCallbackFunc = nil
