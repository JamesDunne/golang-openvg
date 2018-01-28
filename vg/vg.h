VGErrorCode vgGetError();

void vgFlush();
void vgFinish();

/* Getters and Setters */
void vgSetf (VGParamType type, VGfloat value);
void vgSeti (VGParamType type, VGint value);
void vgSetfv(VGParamType type, VGint count, const VGfloat * values);
void vgSetiv(VGParamType type, VGint count, const VGint * values);

VGfloat vgGetf(VGParamType type);
VGint vgGeti(VGParamType type);
VGint vgGetVectorSize(VGParamType type);
void vgGetfv(VGParamType type, VGint count, VGfloat * values);
void vgGetiv(VGParamType type, VGint count, VGint * values);

void vgSetParameterf(VGHandle object, VGint paramType, VGfloat value);
void vgSetParameteri(VGHandle object, VGint paramType, VGint value);
void vgSetParameterfv(VGHandle object, VGint paramType, VGint count, const VGfloat * values);
void vgSetParameteriv(VGHandle object, VGint paramType, VGint count, const VGint * values);

VGfloat vgGetParameterf(VGHandle object, VGint paramType);
VGint vgGetParameteri(VGHandle object, VGint paramType);
VGint vgGetParameterVectorSize(VGHandle object, VGint paramType);
void vgGetParameterfv(VGHandle object, VGint paramType, VGint count, VGfloat * values);
void vgGetParameteriv(VGHandle object, VGint paramType, VGint count, VGint * values);

/* Matrix Manipulation */
void vgLoadIdentity();
void vgLoadMatrix(const VGfloat * m);
void vgGetMatrix(VGfloat * m);
void vgMultMatrix(const VGfloat * m);
void vgTranslate(VGfloat tx, VGfloat ty);
void vgScale(VGfloat sx, VGfloat sy);
void vgShear(VGfloat shx, VGfloat shy);
void vgRotate(VGfloat angle);

/* Masking and Clearing */
void vgMask(VGHandle mask, VGMaskOperation operation, VGint x, VGint y, VGint width, VGint height);
void vgRenderToMask(VGPath path, VGbitfield paintModes, VGMaskOperation operation);
VGMaskLayer vgCreateMaskLayer(VGint width, VGint height);
void vgDestroyMaskLayer(VGMaskLayer maskLayer);
void vgFillMaskLayer(VGMaskLayer maskLayer, VGint x, VGint y, VGint width, VGint height, VGfloat value);
void vgCopyMask(VGMaskLayer maskLayer, VGint dx, VGint dy, VGint sx, VGint sy, VGint width, VGint height);
void vgClear(VGint x, VGint y, VGint width, VGint height);

/* Paths */
VGPath vgCreatePath(VGint pathFormat, VGPathDatatype datatype, VGfloat scale, VGfloat bias, VGint segmentCapacityHint, VGint coordCapacityHint, VGbitfield capabilities);
void vgClearPath(VGPath path, VGbitfield capabilities);
void vgDestroyPath(VGPath path);
void vgRemovePathCapabilities(VGPath path, VGbitfield capabilities);
VGbitfield vgGetPathCapabilities(VGPath path);
void vgAppendPath(VGPath dstPath, VGPath srcPath);
void vgAppendPathData(VGPath dstPath, VGint numSegments, const VGubyte * pathSegments, const void * pathData);
void vgModifyPathCoords(VGPath dstPath, VGint startIndex, VGint numSegments, const void * pathData);
void vgTransformPath(VGPath dstPath, VGPath srcPath);
VGboolean vgInterpolatePath(VGPath dstPath, VGPath startPath, VGPath endPath, VGfloat amount);
VGfloat vgPathLength(VGPath path, VGint startSegment, VGint numSegments);
void vgPointAlongPath(VGPath path, VGint startSegment, VGint numSegments, VGfloat distance, VGfloat * x, VGfloat * y, VGfloat * tangentX, VGfloat * tangentY);
void vgPathBounds(VGPath path, VGfloat * minX, VGfloat * minY, VGfloat * width, VGfloat * height);
void vgPathTransformedBounds(VGPath path, VGfloat * minX, VGfloat * minY, VGfloat * width, VGfloat * height);
void vgDrawPath(VGPath path, VGbitfield paintModes);

/* Paint */
VGPaint vgCreatePaint();
void vgDestroyPaint(VGPaint paint);
void vgSetPaint(VGPaint paint, VGbitfield paintModes);
VGPaint vgGetPaint(VGPaintMode paintMode);
void vgSetColor(VGPaint paint, VGuint rgba);
VGuint vgGetColor(VGPaint paint);
void vgPaintPattern(VGPaint paint, VGImage pattern);

/* Images */
VGImage vgCreateImage(VGImageFormat format, VGint width, VGint height, VGbitfield allowedQuality);
void vgDestroyImage(VGImage image);
void vgClearImage(VGImage image, VGint x, VGint y, VGint width, VGint height);
void vgImageSubData(VGImage image, const void * data, VGint dataStride, VGImageFormat dataFormat, VGint x, VGint y, VGint width, VGint height);
void vgGetImageSubData(VGImage image, void * data, VGint dataStride, VGImageFormat dataFormat, VGint x, VGint y, VGint width, VGint height);
VGImage vgChildImage(VGImage parent, VGint x, VGint y, VGint width, VGint height);
VGImage vgGetParent(VGImage image); 
void vgCopyImage(VGImage dst, VGint dx, VGint dy, VGImage src, VGint sx, VGint sy, VGint width, VGint height, VGboolean dither);
void vgDrawImage(VGImage image);
void vgSetPixels(VGint dx, VGint dy, VGImage src, VGint sx, VGint sy, VGint width, VGint height);
void vgWritePixels(const void * data, VGint dataStride, VGImageFormat dataFormat, VGint dx, VGint dy, VGint width, VGint height);
void vgGetPixels(VGImage dst, VGint dx, VGint dy, VGint sx, VGint sy, VGint width, VGint height);
void vgReadPixels(void * data, VGint dataStride, VGImageFormat dataFormat, VGint sx, VGint sy, VGint width, VGint height);
void vgCopyPixels(VGint dx, VGint dy, VGint sx, VGint sy, VGint width, VGint height);

/* Text */
VGFont vgCreateFont(VGint glyphCapacityHint);
void vgDestroyFont(VGFont font);
void vgSetGlyphToPath(VGFont font, VGuint glyphIndex, VGPath path, VGboolean isHinted, const VGfloat glyphOrigin [2], const VGfloat escapement[2]);
void vgSetGlyphToImage(VGFont font, VGuint glyphIndex, VGImage image, const VGfloat glyphOrigin [2], const VGfloat escapement[2]);
void vgClearGlyph(VGFont font,VGuint glyphIndex);
void vgDrawGlyph(VGFont font, VGuint glyphIndex, VGbitfield paintModes, VGboolean allowAutoHinting);
void vgDrawGlyphs(VGFont font, VGint glyphCount, const VGuint *glyphIndices, const VGfloat *adjustments_x, const VGfloat *adjustments_y, VGbitfield paintModes, VGboolean allowAutoHinting);

/* Image Filters */
void vgColorMatrix(VGImage dst, VGImage src, const VGfloat * matrix);
void vgConvolve(VGImage dst, VGImage src, VGint kernelWidth, VGint kernelHeight, VGint shiftX, VGint shiftY, const VGshort * kernel, VGfloat scale, VGfloat bias, VGTilingMode tilingMode);
void vgSeparableConvolve(VGImage dst, VGImage src, VGint kernelWidth, VGint kernelHeight, VGint shiftX, VGint shiftY, const VGshort * kernelX, const VGshort * kernelY, VGfloat scale, VGfloat bias, VGTilingMode tilingMode);
void vgGaussianBlur(VGImage dst, VGImage src, VGfloat stdDeviationX, VGfloat stdDeviationY, VGTilingMode tilingMode);
void vgLookup(VGImage dst, VGImage src, const VGubyte * redLUT, const VGubyte * greenLUT, const VGubyte * blueLUT, const VGubyte * alphaLUT, VGboolean outputLinear, VGboolean outputPremultiplied);
void vgLookupSingle(VGImage dst, VGImage src, const VGuint * lookupTable, VGImageChannel sourceChannel, VGboolean outputLinear, VGboolean outputPremultiplied);

/* Hardware Queries */
VGHardwareQueryResult vgHardwareQuery(VGHardwareQueryType key, VGint setting);

/* Renderer and Extension Information */
const VGubyte * vgGetString(VGStringID name);
