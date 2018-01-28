//+build darwin

package host

/*
#cgo CFLAGS: -x objective-c -Iinclude
#cgo LDFLAGS: -framework Cocoa -framework OpenGL -framework QuartzCore -Llib/macosx/ub/gle/standalone -lAmanithVG

#include "macosx/main.m"

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

func Run(width, height int) {
	C.vgMain(C.int(width), C.int(height))
}
