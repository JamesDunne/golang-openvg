//+build darwin

package host

/*
#cgo CFLAGS: -x objective-c -Iinclude
#cgo LDFLAGS: -framework Cocoa -framework OpenGL -framework QuartzCore -Llib/macosx/ub/gle/standalone -lAmanithVG

#include "app.h"
*/
import "C"

//export appInit
func appInit(surfaceWidth, surfaceHeight C.VGint) {

}

//export appDestroy
func appDestroy() {

}

//export appResize
func appResize(surfaceWidth, surfaceHeight C.VGint) {

}

//export appDraw
func appDraw(width, height int32) {
	if drawFunc != nil {
		drawFunc(width, height)
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

type DrawFunc func(int32, int32)

var drawFunc DrawFunc = nil

func (h *Host) UseDrawFunc(draw DrawFunc) {
	drawFunc = draw
}
