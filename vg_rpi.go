//+build rpi

package vg

//#cgo CFLAGS: -I/opt/vc/include
//#cgo LDFLAGS: -L/opt/vc/lib -lbcm_host -lGLESv2 -lEGL
//#include "vg_rpi.h"
import "C"

import "fmt"
import "unsafe"

type VG struct {
	Width  uint32
	Height uint32

	display uintptr
	surface uintptr
}

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

func Create(width, height uint32) (*VG, error) {
	vg := &VG{Width: width, Height: height}

	if retval := C.create(
		(*C.uint32_t)(unsafe.Pointer(&vg.Width)),
		(*C.uint32_t)(unsafe.Pointer(&vg.Height)),
		(*C.EGLDisplay)(unsafe.Pointer(&vg.display)),
		(*C.EGLSurface)(unsafe.Pointer(&vg.surface)),
	); retval != C.EGL_SUCCESS {
		return nil, EGLerror(retval)
	}

	return vg, nil
}

func (vg *VG) SwapBuffers() error {
	if retval := C.eglSwapBuffers(C.EGLDisplay(vg.display), C.EGLSurface(vg.surface)); retval != C.EGL_SUCCESS {
		return EGLerror(retval)
	}
	return nil
}
