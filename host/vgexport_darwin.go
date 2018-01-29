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
	if cb := initFunc; cb != nil {
		cb(int32(surfaceWidth), int32(surfaceHeight))
	}
}

//export appDestroy
func appDestroy() {
	if cb := destroyFunc; cb != nil {
		cb()
	}
}

//export appResize
func appResize(surfaceWidth, surfaceHeight C.VGint) {
	if cb := resizeFunc; cb != nil {
		cb(int32(surfaceWidth), int32(surfaceHeight))
	}
}

//export appDraw
func appDraw(surfaceWidth, surfaceHeight C.VGint) {
	if cb := drawFunc; cb != nil {
		cb(int32(surfaceWidth), int32(surfaceHeight))
	}
}

// handle mouse events

//export mouseLeftButtonDown
func mouseLeftButtonDown(x, y C.VGint) {

}

//export mouseLeftButtonUp
func mouseLeftButtonUp(x, y C.VGint) {

}

//export mouseRightButtonDown
func mouseRightButtonDown(x, y C.VGint) {

}

//export mouseRightButtonUp
func mouseRightButtonUp(x, y C.VGint) {

}

//export mouseMove
func mouseMove(x, y C.VGint) {

}

//export keyDown
func keyDown(k uint16) {

}

//export keyUp
func keyUp(k uint16) {

}

type SurfaceCallbackFunc func(width, height int32)
type VoidCallbackFunc func()
type PositionCallbackFunc func(x, y int32)

var initFunc SurfaceCallbackFunc = nil
var destroyFunc VoidCallbackFunc = nil
var resizeFunc SurfaceCallbackFunc = nil
var drawFunc SurfaceCallbackFunc = nil
var mouseLeftButtonDownFunc PositionCallbackFunc = nil
var mouseLeftButtonUpFunc PositionCallbackFunc = nil
var mouseRightButtonDownFunc PositionCallbackFunc = nil
var mouseRightButtonUpFunc PositionCallbackFunc = nil
var mouseMoveFunc PositionCallbackFunc = nil

func UseDrawFunc(draw SurfaceCallbackFunc) {
	drawFunc = draw
}
