//+build windows

package host

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

void
keyDown(const unsigned short k)
{
}

void
keyUp(const unsigned short k)
{
}

*/
import "C"
