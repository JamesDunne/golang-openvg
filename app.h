/****************************************************************************
** Copyright (C) 2004-2017 Mazatech S.r.l. All rights reserved.
**
** This file is part of AmanithVG software, an OpenVG implementation.
**
** Khronos and OpenVG are trademarks of The Khronos Group Inc.
** OpenGL is a registered trademark and OpenGL ES is a trademark of
** Silicon Graphics, Inc.
**
** This file is provided AS IS with NO WARRANTY OF ANY KIND, INCLUDING THE
** WARRANTY OF DESIGN, MERCHANTABILITY AND FITNESS FOR A PARTICULAR PURPOSE.
**
** For any information, please contact info@mazatech.com
**
****************************************************************************/

#ifndef APP_H
#define APP_H

#include <VG/openvg.h>
#include <VG/vgu.h>
#include <VG/vgext.h>

void appInit(const VGint surfaceWidth, const VGint surfaceHeight);
void appDestroy(void);
void appResize(const VGint surfaceWidth, const VGint surfaceHeight);
void appDraw(const VGint surfaceWidth, const VGint surfaceHeight);

// handle mouse events
void mouseLeftButtonDown(const VGint x, const VGint y);
void mouseLeftButtonUp(const VGint x, const VGint y);
void mouseRightButtonDown(const VGint x, const VGint y);
void mouseRightButtonUp(const VGint x, const VGint y);
void mouseMove(const VGint x, const VGint y);

// handle keyboard events
void keyDown(const unsigned short k);
void keyUp(const unsigned short k);

#endif