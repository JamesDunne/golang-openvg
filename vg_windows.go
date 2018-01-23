//+build windows

package main

/*
#cgo CFLAGS: -Iinclude
#cgo LDFLAGS: -mwindows -s -Llib/win/x86_64/gle/standalone -lAmanithVG -lopengl32

#include "win/main.c"

void
appInit(const VGint surfaceWidth, const VGint surfaceHeight)
{

}

void
appDestroy(void)
{

}

void
appResize(const VGint surfaceWidth, const VGint surfaceHeight)
{

}

void
appDraw(const VGint surfaceWidth, const VGint surfaceHeight)
{

}

// handle mouse events
void
mouseLeftButtonDown(const VGint x, const VGint y)
{

}

void
mouseLeftButtonUp(const VGint x, const VGint y)
{

}

void
mouseRightButtonDown(const VGint x, const VGint y)
{
}

void
mouseRightButtonUp(const VGint x, const VGint y)
{
}


void
mouseMove(const VGint x, const VGint y)
{
}


// handle touch events
void
touchDoubleTap(const VGint x, const VGint y)
{
}

void
touchPinch(const VGfloat deltaScl)
{
}

void
touchRotate(const VGfloat deltaRot)
{
}

*/
import "C"

type VG struct {
	Width  uint32
	Height uint32

	display uintptr
	surface uintptr
}

func main() {
	C.vgMain(800, 480)
}
