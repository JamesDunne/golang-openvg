//+build rpi

package vg

//#cgo CFLAGS: -I/opt/vc/include
//#cgo LDFLAGS: -L/opt/vc/lib -lbcm_host -lGLESv2 -lEGL
//#include "vg_rpi.h"
import "C"

import "unsafe"

type VG struct {
	Width   uint32
	Height  uint32
	display uintptr
}

func Create(width, height uint32) (*VG, error) {
	vg := &VG{Width: width, Height: height}
	if retval := C.create(
		(*C.uint32_t)(unsafe.Pointer(&vg.Width)),
		(*C.uint32_t)(unsafe.Pointer(&vg.Height)),
		(*C.EGLDisplay)(unsafe.Pointer(&vg.display))); retval != 0 {
		return nil, nil
	}

	return vg, nil
}
