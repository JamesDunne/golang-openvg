//+build rpi

package host

//#cgo CFLAGS: -I/opt/vc/include
//#cgo LDFLAGS: -L/opt/vc/lib -lbcm_host -lGLESv2 -lEGL
//#include "rpi/main.c"
import "C"

import "fmt"
import "unsafe"

var (
	Width  int32
	Height int32

	display uintptr
	surface uintptr
)

type EGLerror C.EGLint

func (e EGLerror) Error() string {
	switch e {
	case C.EGL_SUCCESS:
		return "No error"
	case C.EGL_NOT_INITIALIZED:
		return "EGL not initialized or failed to initialize"
	case C.EGL_BAD_ACCESS:
		return "Resource inaccessible"
	case C.EGL_BAD_ALLOC:
		return "Cannot allocate resources"
	case C.EGL_BAD_ATTRIBUTE:
		return "Unrecognized attribute or attribute value"
	case C.EGL_BAD_CONTEXT:
		return "Invalid EGL context"
	case C.EGL_BAD_CONFIG:
		return "Invalid EGL frame buffer configuration"
	case C.EGL_BAD_CURRENT_SURFACE:
		return "Current surface is no longer valid"
	case C.EGL_BAD_DISPLAY:
		return "Invalid EGL display"
	case C.EGL_BAD_SURFACE:
		return "Invalid surface"
	case C.EGL_BAD_MATCH:
		return "Inconsistent arguments"
	case C.EGL_BAD_PARAMETER:
		return "Invalid argument"
	case C.EGL_BAD_NATIVE_PIXMAP:
		return "Invalid native pixmap"
	case C.EGL_BAD_NATIVE_WINDOW:
		return "Invalid native window"
	case C.EGL_CONTEXT_LOST:
		return "Context lost"
	default:
		return fmt.Sprintf("Unknown error 0x%04x", uint32(e))
	}
}

func Init(width, height int) bool {
	w32 := uint32(0)
	h32 := uint32(0)

	if retval := C.create(
		(*C.uint32_t)(unsafe.Pointer(&w32)),
		(*C.uint32_t)(unsafe.Pointer(&h32)),
		(*C.EGLDisplay)(unsafe.Pointer(&display)),
		(*C.EGLSurface)(unsafe.Pointer(&surface)),
	); retval != C.EGL_SUCCESS {
		panic(EGLerror(retval))
	}

	// TODO: respect passed-in width,height.
	Width = int32(w32)
	Height = int32(h32)

	if cb := InitFunc; cb != nil {
		cb(Width, Height)
	}

	return true
}

func Destroy() int {
	//return int(C.hostDestroy())
	return 0
}

func PollEvent() bool {
	//return int32(C.hostPollEvent()) == 0

	if cb := DrawFunc; cb != nil {
		cb(Width, Height)
	}
	SwapBuffers()

	return true
}

func SwapBuffers() error {
	if retval := C.eglSwapBuffers(C.EGLDisplay(display), C.EGLSurface(surface)); retval != C.EGL_SUCCESS {
		return EGLerror(retval)
	}
	return nil
}
