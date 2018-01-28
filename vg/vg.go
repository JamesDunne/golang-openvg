package vg

//#cgo LDFLAGS: -lAmanithVG
//#include "VG/openvg.h"
import "C"

import "unsafe"

type BooleanEnum int32
const (
	False BooleanEnum = 0
	True BooleanEnum = 1
	BooleanForceSize BooleanEnum = 2147483647
)

type ErrorCodeEnum int32
const (
	NoError ErrorCodeEnum = 0
	BadHandleError ErrorCodeEnum = 4096
	IllegalArgumentError ErrorCodeEnum = 4097
	OutOfMemoryError ErrorCodeEnum = 4098
	PathCapabilityError ErrorCodeEnum = 4099
	UnsupportedImageFormatError ErrorCodeEnum = 4100
	UnsupportedPathFormatError ErrorCodeEnum = 4101
	ImageInUseError ErrorCodeEnum = 4102
	NoContextError ErrorCodeEnum = 4103
	ErrorCodeForceSize ErrorCodeEnum = 2147483647
)

type ParamTypeEnum int32
const (
	MatrixMode ParamTypeEnum = 4352
	FillRule ParamTypeEnum = 4353
	ImageQuality ParamTypeEnum = 4354
	RenderingQuality ParamTypeEnum = 4355
	BlendMode ParamTypeEnum = 4356
	ImageMode ParamTypeEnum = 4357
	ScissorRects ParamTypeEnum = 4358
	ColorTransform ParamTypeEnum = 4464
	ColorTransformValues ParamTypeEnum = 4465
	StrokeLineWidth ParamTypeEnum = 4368
	StrokeCapStyle ParamTypeEnum = 4369
	StrokeJoinStyle ParamTypeEnum = 4370
	StrokeMiterLimit ParamTypeEnum = 4371
	StrokeDashPattern ParamTypeEnum = 4372
	StrokeDashPhase ParamTypeEnum = 4373
	StrokeDashPhaseReset ParamTypeEnum = 4374
	TileFillColor ParamTypeEnum = 4384
	ClearColor ParamTypeEnum = 4385
	GlyphOrigin ParamTypeEnum = 4386
	Masking ParamTypeEnum = 4400
	Scissoring ParamTypeEnum = 4401
	PixelLayout ParamTypeEnum = 4416
	ScreenLayout ParamTypeEnum = 4417
	FilterFormatLinear ParamTypeEnum = 4432
	FilterFormatPremultiplied ParamTypeEnum = 4433
	FilterChannelMask ParamTypeEnum = 4434
	MaxScissorRects ParamTypeEnum = 4448
	MaxDashCount ParamTypeEnum = 4449
	MaxKernelSize ParamTypeEnum = 4450
	MaxSeparableKernelSize ParamTypeEnum = 4451
	MaxColorRampStops ParamTypeEnum = 4452
	MaxImageWidth ParamTypeEnum = 4453
	MaxImageHeight ParamTypeEnum = 4454
	MaxImagePixels ParamTypeEnum = 4455
	MaxImageBytes ParamTypeEnum = 4456
	MaxFloat ParamTypeEnum = 4457
	MaxGaussianStdDeviation ParamTypeEnum = 4458
	ParamTypeForceSize ParamTypeEnum = 2147483647
)

type RenderingQualityEnum int32
const (
	RenderingQualityNonantialiased RenderingQualityEnum = 4608
	RenderingQualityFaster RenderingQualityEnum = 4609
	RenderingQualityBetter RenderingQualityEnum = 4610
	RenderingQualityForceSize RenderingQualityEnum = 2147483647
)

type PixelLayoutEnum int32
const (
	PixelLayoutUnknown PixelLayoutEnum = 4864
	PixelLayoutRgbVertical PixelLayoutEnum = 4865
	PixelLayoutBgrVertical PixelLayoutEnum = 4866
	PixelLayoutRgbHorizontal PixelLayoutEnum = 4867
	PixelLayoutBgrHorizontal PixelLayoutEnum = 4868
	PixelLayoutForceSize PixelLayoutEnum = 2147483647
)

type MatrixModeEnum int32
const (
	MatrixPathUserToSurface MatrixModeEnum = 5120
	MatrixImageUserToSurface MatrixModeEnum = 5121
	MatrixFillPaintToUser MatrixModeEnum = 5122
	MatrixStrokePaintToUser MatrixModeEnum = 5123
	MatrixGlyphUserToSurface MatrixModeEnum = 5124
	MatrixModeForceSize MatrixModeEnum = 2147483647
)

type MaskOperationEnum int32
const (
	ClearMask MaskOperationEnum = 5376
	FillMask MaskOperationEnum = 5377
	SetMask MaskOperationEnum = 5378
	UnionMask MaskOperationEnum = 5379
	IntersectMask MaskOperationEnum = 5380
	SubtractMask MaskOperationEnum = 5381
	MaskOperationForceSize MaskOperationEnum = 2147483647
)

type PathDatatypeEnum int32
const (
	PathDatatypeS8 PathDatatypeEnum = 0
	PathDatatypeS16 PathDatatypeEnum = 1
	PathDatatypeS32 PathDatatypeEnum = 2
	PathDatatypeF PathDatatypeEnum = 3
	PathDatatypeForceSize PathDatatypeEnum = 2147483647
)

type PathAbsRelEnum int32
const (
	Absolute PathAbsRelEnum = 0
	Relative PathAbsRelEnum = 1
	PathAbsRelForceSize PathAbsRelEnum = 2147483647
)

type PathSegmentEnum int32
const (
	ClosePath PathSegmentEnum = 0
	MoveTo PathSegmentEnum = 2
	LineTo PathSegmentEnum = 4
	HlineTo PathSegmentEnum = 6
	VlineTo PathSegmentEnum = 8
	QuadTo PathSegmentEnum = 10
	CubicTo PathSegmentEnum = 12
	SquadTo PathSegmentEnum = 14
	ScubicTo PathSegmentEnum = 16
	SccwarcTo PathSegmentEnum = 18
	ScwarcTo PathSegmentEnum = 20
	LccwarcTo PathSegmentEnum = 22
	LcwarcTo PathSegmentEnum = 24
	PathSegmentForceSize PathSegmentEnum = 2147483647
)

type PathCommandEnum int32
const (
	MoveToAbs PathCommandEnum = 2
	MoveToRel PathCommandEnum = 3
	LineToAbs PathCommandEnum = 4
	LineToRel PathCommandEnum = 5
	HlineToAbs PathCommandEnum = 6
	HlineToRel PathCommandEnum = 7
	VlineToAbs PathCommandEnum = 8
	VlineToRel PathCommandEnum = 9
	QuadToAbs PathCommandEnum = 10
	QuadToRel PathCommandEnum = 11
	CubicToAbs PathCommandEnum = 12
	CubicToRel PathCommandEnum = 13
	SquadToAbs PathCommandEnum = 14
	SquadToRel PathCommandEnum = 15
	ScubicToAbs PathCommandEnum = 16
	ScubicToRel PathCommandEnum = 17
	SccwarcToAbs PathCommandEnum = 18
	SccwarcToRel PathCommandEnum = 19
	ScwarcToAbs PathCommandEnum = 20
	ScwarcToRel PathCommandEnum = 21
	LccwarcToAbs PathCommandEnum = 22
	LccwarcToRel PathCommandEnum = 23
	LcwarcToAbs PathCommandEnum = 24
	LcwarcToRel PathCommandEnum = 25
	PathCommandForceSize PathCommandEnum = 2147483647
)

type PathCapabilitiesEnum int32
const (
	PathCapabilityAppendFrom PathCapabilitiesEnum = 1
	PathCapabilityAppendTo PathCapabilitiesEnum = 2
	PathCapabilityModify PathCapabilitiesEnum = 4
	PathCapabilityTransformFrom PathCapabilitiesEnum = 8
	PathCapabilityTransformTo PathCapabilitiesEnum = 16
	PathCapabilityInterpolateFrom PathCapabilitiesEnum = 32
	PathCapabilityInterpolateTo PathCapabilitiesEnum = 64
	PathCapabilityPathLength PathCapabilitiesEnum = 128
	PathCapabilityPointAlongPath PathCapabilitiesEnum = 256
	PathCapabilityTangentAlongPath PathCapabilitiesEnum = 512
	PathCapabilityPathBounds PathCapabilitiesEnum = 1024
	PathCapabilityPathTransformedBounds PathCapabilitiesEnum = 2048
	PathCapabilityAll PathCapabilitiesEnum = 4095
	PathCapabilitiesForceSize PathCapabilitiesEnum = 2147483647
)

type PathParamTypeEnum int32
const (
	PathFormat PathParamTypeEnum = 5632
	PathDatatype PathParamTypeEnum = 5633
	PathScale PathParamTypeEnum = 5634
	PathBias PathParamTypeEnum = 5635
	PathNumSegments PathParamTypeEnum = 5636
	PathNumCoords PathParamTypeEnum = 5637
	PathParamTypeForceSize PathParamTypeEnum = 2147483647
)

type CapStyleEnum int32
const (
	CapButt CapStyleEnum = 5888
	CapRound CapStyleEnum = 5889
	CapSquare CapStyleEnum = 5890
	CapStyleForceSize CapStyleEnum = 2147483647
)

type JoinStyleEnum int32
const (
	JoinMiter JoinStyleEnum = 6144
	JoinRound JoinStyleEnum = 6145
	JoinBevel JoinStyleEnum = 6146
	JoinStyleForceSize JoinStyleEnum = 2147483647
)

type FillRuleEnum int32
const (
	EvenOdd FillRuleEnum = 6400
	NonZero FillRuleEnum = 6401
	FillRuleForceSize FillRuleEnum = 2147483647
)

type PaintModeEnum int32
const (
	StrokePath PaintModeEnum = 1
	FillPath PaintModeEnum = 2
	PaintModeForceSize PaintModeEnum = 2147483647
)

type PaintParamTypeEnum int32
const (
	PaintType PaintParamTypeEnum = 6656
	PaintColor PaintParamTypeEnum = 6657
	PaintColorRampSpreadMode PaintParamTypeEnum = 6658
	PaintColorRampPremultiplied PaintParamTypeEnum = 6663
	PaintColorRampStops PaintParamTypeEnum = 6659
	PaintLinearGradient PaintParamTypeEnum = 6660
	PaintRadialGradient PaintParamTypeEnum = 6661
	PaintPatternTilingMode PaintParamTypeEnum = 6662
	PaintParamTypeForceSize PaintParamTypeEnum = 2147483647
)

type PaintTypeEnum int32
const (
	PaintTypeColor PaintTypeEnum = 6912
	PaintTypeLinearGradient PaintTypeEnum = 6913
	PaintTypeRadialGradient PaintTypeEnum = 6914
	PaintTypePattern PaintTypeEnum = 6915
	PaintTypeForceSize PaintTypeEnum = 2147483647
)

type ColorRampSpreadModeEnum int32
const (
	ColorRampSpreadPad ColorRampSpreadModeEnum = 7168
	ColorRampSpreadRepeat ColorRampSpreadModeEnum = 7169
	ColorRampSpreadReflect ColorRampSpreadModeEnum = 7170
	ColorRampSpreadModeForceSize ColorRampSpreadModeEnum = 2147483647
)

type TilingModeEnum int32
const (
	TileFill TilingModeEnum = 7424
	TilePad TilingModeEnum = 7425
	TileRepeat TilingModeEnum = 7426
	TileReflect TilingModeEnum = 7427
	TilingModeForceSize TilingModeEnum = 2147483647
)

type ImageFormatEnum int32
const (
	Srgbx8888 ImageFormatEnum = 0
	Srgba8888 ImageFormatEnum = 1
	Srgba8888Pre ImageFormatEnum = 2
	Srgb565 ImageFormatEnum = 3
	Srgba5551 ImageFormatEnum = 4
	Srgba4444 ImageFormatEnum = 5
	Sl8 ImageFormatEnum = 6
	Lrgbx8888 ImageFormatEnum = 7
	Lrgba8888 ImageFormatEnum = 8
	Lrgba8888Pre ImageFormatEnum = 9
	Ll8 ImageFormatEnum = 10
	A8 ImageFormatEnum = 11
	Bw1 ImageFormatEnum = 12
	A1 ImageFormatEnum = 13
	A4 ImageFormatEnum = 14
	Sxrgb8888 ImageFormatEnum = 64
	Sargb8888 ImageFormatEnum = 65
	Sargb8888Pre ImageFormatEnum = 66
	Sargb1555 ImageFormatEnum = 68
	Sargb4444 ImageFormatEnum = 69
	Lxrgb8888 ImageFormatEnum = 71
	Largb8888 ImageFormatEnum = 72
	Largb8888Pre ImageFormatEnum = 73
	Sbgrx8888 ImageFormatEnum = 128
	Sbgra8888 ImageFormatEnum = 129
	Sbgra8888Pre ImageFormatEnum = 130
	Sbgr565 ImageFormatEnum = 131
	Sbgra5551 ImageFormatEnum = 132
	Sbgra4444 ImageFormatEnum = 133
	Lbgrx8888 ImageFormatEnum = 135
	Lbgra8888 ImageFormatEnum = 136
	Lbgra8888Pre ImageFormatEnum = 137
	Sxbgr8888 ImageFormatEnum = 192
	Sabgr8888 ImageFormatEnum = 193
	Sabgr8888Pre ImageFormatEnum = 194
	Sabgr1555 ImageFormatEnum = 196
	Sabgr4444 ImageFormatEnum = 197
	Lxbgr8888 ImageFormatEnum = 199
	Labgr8888 ImageFormatEnum = 200
	Labgr8888Pre ImageFormatEnum = 201
	ImageFormatForceSize ImageFormatEnum = 2147483647
)

type ImageQualityEnum int32
const (
	ImageQualityNonantialiased ImageQualityEnum = 1
	ImageQualityFaster ImageQualityEnum = 2
	ImageQualityBetter ImageQualityEnum = 4
	ImageQualityForceSize ImageQualityEnum = 2147483647
)

type ImageParamTypeEnum int32
const (
	ImageFormat ImageParamTypeEnum = 7680
	ImageWidth ImageParamTypeEnum = 7681
	ImageHeight ImageParamTypeEnum = 7682
	ImageParamTypeForceSize ImageParamTypeEnum = 2147483647
)

type ImageModeEnum int32
const (
	DrawImageNormal ImageModeEnum = 7936
	DrawImageMultiply ImageModeEnum = 7937
	DrawImageStencil ImageModeEnum = 7938
	ImageModeForceSize ImageModeEnum = 2147483647
)

type ImageChannelEnum int32
const (
	Red ImageChannelEnum = 8
	Green ImageChannelEnum = 4
	Blue ImageChannelEnum = 2
	Alpha ImageChannelEnum = 1
	ImageChannelForceSize ImageChannelEnum = 2147483647
)

type BlendModeEnum int32
const (
	BlendSrc BlendModeEnum = 8192
	BlendSrcOver BlendModeEnum = 8193
	BlendDstOver BlendModeEnum = 8194
	BlendSrcIn BlendModeEnum = 8195
	BlendDstIn BlendModeEnum = 8196
	BlendMultiply BlendModeEnum = 8197
	BlendScreen BlendModeEnum = 8198
	BlendDarken BlendModeEnum = 8199
	BlendLighten BlendModeEnum = 8200
	BlendAdditive BlendModeEnum = 8201
	BlendModeForceSize BlendModeEnum = 2147483647
)

type FontParamTypeEnum int32
const (
	FontNumGlyphs FontParamTypeEnum = 12032
	FontParamTypeForceSize FontParamTypeEnum = 2147483647
)

type HardwareQueryTypeEnum int32
const (
	ImageFormatQuery HardwareQueryTypeEnum = 8448
	PathDatatypeQuery HardwareQueryTypeEnum = 8449
	HardwareQueryTypeForceSize HardwareQueryTypeEnum = 2147483647
)

type HardwareQueryResultEnum int32
const (
	HardwareAccelerated HardwareQueryResultEnum = 8704
	HardwareUnaccelerated HardwareQueryResultEnum = 8705
	HardwareQueryResultForceSize HardwareQueryResultEnum = 2147483647
)

type StringIDEnum int32
const (
	Vendor StringIDEnum = 8960
	Renderer StringIDEnum = 8961
	Version StringIDEnum = 8962
	Extensions StringIDEnum = 8963
	StringIdForceSize StringIDEnum = 2147483647
)

func GetError(
) ErrorCodeEnum {
	ret := C.vgGetError(
	)
	return (ErrorCodeEnum)(ret)
}

func Flush(
) {
	C.vgFlush(
	)
}

func Finish(
) {
	C.vgFinish(
	)
}

func Setf(
	_type ParamTypeEnum,
	value float32,
) {
	C.vgSetf(
		(C.VGParamType)(_type),
		(C.VGfloat)(value),
	)
}

func Seti(
	_type ParamTypeEnum,
	value int32,
) {
	C.vgSeti(
		(C.VGParamType)(_type),
		(C.VGint)(value),
	)
}

func Setfv(
	_type ParamTypeEnum,
	count int32,
	values *float32,
) {
	C.vgSetfv(
		(C.VGParamType)(_type),
		(C.VGint)(count),
		(*C.VGfloat)(values),
	)
}

func Setiv(
	_type ParamTypeEnum,
	count int32,
	values *int32,
) {
	C.vgSetiv(
		(C.VGParamType)(_type),
		(C.VGint)(count),
		(*C.VGint)(values),
	)
}

func Getf(
	_type ParamTypeEnum,
) float32 {
	ret := C.vgGetf(
		(C.VGParamType)(_type),
	)
	return (float32)(ret)
}

func Geti(
	_type ParamTypeEnum,
) int32 {
	ret := C.vgGeti(
		(C.VGParamType)(_type),
	)
	return (int32)(ret)
}

func GetVectorSize(
	_type ParamTypeEnum,
) int32 {
	ret := C.vgGetVectorSize(
		(C.VGParamType)(_type),
	)
	return (int32)(ret)
}

func Getfv(
	_type ParamTypeEnum,
	count int32,
	values *float32,
) {
	C.vgGetfv(
		(C.VGParamType)(_type),
		(C.VGint)(count),
		(*C.VGfloat)(values),
	)
}

func Getiv(
	_type ParamTypeEnum,
	count int32,
	values *int32,
) {
	C.vgGetiv(
		(C.VGParamType)(_type),
		(C.VGint)(count),
		(*C.VGint)(values),
	)
}

func SetParameterf(
	object uint32,
	paramType int32,
	value float32,
) {
	C.vgSetParameterf(
		(C.VGHandle)(object),
		(C.VGint)(paramType),
		(C.VGfloat)(value),
	)
}

func SetParameteri(
	object uint32,
	paramType int32,
	value int32,
) {
	C.vgSetParameteri(
		(C.VGHandle)(object),
		(C.VGint)(paramType),
		(C.VGint)(value),
	)
}

func SetParameterfv(
	object uint32,
	paramType int32,
	count int32,
	values *float32,
) {
	C.vgSetParameterfv(
		(C.VGHandle)(object),
		(C.VGint)(paramType),
		(C.VGint)(count),
		(*C.VGfloat)(values),
	)
}

func SetParameteriv(
	object uint32,
	paramType int32,
	count int32,
	values *int32,
) {
	C.vgSetParameteriv(
		(C.VGHandle)(object),
		(C.VGint)(paramType),
		(C.VGint)(count),
		(*C.VGint)(values),
	)
}

func GetParameterf(
	object uint32,
	paramType int32,
) float32 {
	ret := C.vgGetParameterf(
		(C.VGHandle)(object),
		(C.VGint)(paramType),
	)
	return (float32)(ret)
}

func GetParameteri(
	object uint32,
	paramType int32,
) int32 {
	ret := C.vgGetParameteri(
		(C.VGHandle)(object),
		(C.VGint)(paramType),
	)
	return (int32)(ret)
}

func GetParameterVectorSize(
	object uint32,
	paramType int32,
) int32 {
	ret := C.vgGetParameterVectorSize(
		(C.VGHandle)(object),
		(C.VGint)(paramType),
	)
	return (int32)(ret)
}

func GetParameterfv(
	object uint32,
	paramType int32,
	count int32,
	values *float32,
) {
	C.vgGetParameterfv(
		(C.VGHandle)(object),
		(C.VGint)(paramType),
		(C.VGint)(count),
		(*C.VGfloat)(values),
	)
}

func GetParameteriv(
	object uint32,
	paramType int32,
	count int32,
	values *int32,
) {
	C.vgGetParameteriv(
		(C.VGHandle)(object),
		(C.VGint)(paramType),
		(C.VGint)(count),
		(*C.VGint)(values),
	)
}

func LoadIdentity(
) {
	C.vgLoadIdentity(
	)
}

func LoadMatrix(
	m *float32,
) {
	C.vgLoadMatrix(
		(*C.VGfloat)(m),
	)
}

func GetMatrix(
	m *float32,
) {
	C.vgGetMatrix(
		(*C.VGfloat)(m),
	)
}

func MultMatrix(
	m *float32,
) {
	C.vgMultMatrix(
		(*C.VGfloat)(m),
	)
}

func Translate(
	tx float32,
	ty float32,
) {
	C.vgTranslate(
		(C.VGfloat)(tx),
		(C.VGfloat)(ty),
	)
}

func Scale(
	sx float32,
	sy float32,
) {
	C.vgScale(
		(C.VGfloat)(sx),
		(C.VGfloat)(sy),
	)
}

func Shear(
	shx float32,
	shy float32,
) {
	C.vgShear(
		(C.VGfloat)(shx),
		(C.VGfloat)(shy),
	)
}

func Rotate(
	angle float32,
) {
	C.vgRotate(
		(C.VGfloat)(angle),
	)
}

func Mask(
	mask uint32,
	operation MaskOperationEnum,
	x int32,
	y int32,
	width int32,
	height int32,
) {
	C.vgMask(
		(C.VGHandle)(mask),
		(C.VGMaskOperation)(operation),
		(C.VGint)(x),
		(C.VGint)(y),
		(C.VGint)(width),
		(C.VGint)(height),
	)
}

func RenderToMask(
	path uint32,
	paintModes uint32,
	operation MaskOperationEnum,
) {
	C.vgRenderToMask(
		(C.VGPath)(path),
		(C.VGbitfield)(paintModes),
		(C.VGMaskOperation)(operation),
	)
}

func CreateMaskLayer(
	width int32,
	height int32,
) uint32 {
	ret := C.vgCreateMaskLayer(
		(C.VGint)(width),
		(C.VGint)(height),
	)
	return (uint32)(ret)
}

func DestroyMaskLayer(
	maskLayer uint32,
) {
	C.vgDestroyMaskLayer(
		(C.VGMaskLayer)(maskLayer),
	)
}

func FillMaskLayer(
	maskLayer uint32,
	x int32,
	y int32,
	width int32,
	height int32,
	value float32,
) {
	C.vgFillMaskLayer(
		(C.VGMaskLayer)(maskLayer),
		(C.VGint)(x),
		(C.VGint)(y),
		(C.VGint)(width),
		(C.VGint)(height),
		(C.VGfloat)(value),
	)
}

func CopyMask(
	maskLayer uint32,
	dx int32,
	dy int32,
	sx int32,
	sy int32,
	width int32,
	height int32,
) {
	C.vgCopyMask(
		(C.VGMaskLayer)(maskLayer),
		(C.VGint)(dx),
		(C.VGint)(dy),
		(C.VGint)(sx),
		(C.VGint)(sy),
		(C.VGint)(width),
		(C.VGint)(height),
	)
}

func Clear(
	x int32,
	y int32,
	width int32,
	height int32,
) {
	C.vgClear(
		(C.VGint)(x),
		(C.VGint)(y),
		(C.VGint)(width),
		(C.VGint)(height),
	)
}

func CreatePath(
	pathFormat int32,
	datatype PathDatatypeEnum,
	scale float32,
	bias float32,
	segmentCapacityHint int32,
	coordCapacityHint int32,
	capabilities uint32,
) uint32 {
	ret := C.vgCreatePath(
		(C.VGint)(pathFormat),
		(C.VGPathDatatype)(datatype),
		(C.VGfloat)(scale),
		(C.VGfloat)(bias),
		(C.VGint)(segmentCapacityHint),
		(C.VGint)(coordCapacityHint),
		(C.VGbitfield)(capabilities),
	)
	return (uint32)(ret)
}

func ClearPath(
	path uint32,
	capabilities uint32,
) {
	C.vgClearPath(
		(C.VGPath)(path),
		(C.VGbitfield)(capabilities),
	)
}

func DestroyPath(
	path uint32,
) {
	C.vgDestroyPath(
		(C.VGPath)(path),
	)
}

func RemovePathCapabilities(
	path uint32,
	capabilities uint32,
) {
	C.vgRemovePathCapabilities(
		(C.VGPath)(path),
		(C.VGbitfield)(capabilities),
	)
}

func GetPathCapabilities(
	path uint32,
) uint32 {
	ret := C.vgGetPathCapabilities(
		(C.VGPath)(path),
	)
	return (uint32)(ret)
}

func AppendPath(
	dstPath uint32,
	srcPath uint32,
) {
	C.vgAppendPath(
		(C.VGPath)(dstPath),
		(C.VGPath)(srcPath),
	)
}

func AppendPathData(
	dstPath uint32,
	numSegments int32,
	pathSegments *uint8,
	pathData *byte,
) {
	C.vgAppendPathData(
		(C.VGPath)(dstPath),
		(C.VGint)(numSegments),
		(*C.VGubyte)(pathSegments),
		unsafe.Pointer(pathData),
	)
}

func ModifyPathCoords(
	dstPath uint32,
	startIndex int32,
	numSegments int32,
	pathData *byte,
) {
	C.vgModifyPathCoords(
		(C.VGPath)(dstPath),
		(C.VGint)(startIndex),
		(C.VGint)(numSegments),
		unsafe.Pointer(pathData),
	)
}

func TransformPath(
	dstPath uint32,
	srcPath uint32,
) {
	C.vgTransformPath(
		(C.VGPath)(dstPath),
		(C.VGPath)(srcPath),
	)
}

func InterpolatePath(
	dstPath uint32,
	startPath uint32,
	endPath uint32,
	amount float32,
) BooleanEnum {
	ret := C.vgInterpolatePath(
		(C.VGPath)(dstPath),
		(C.VGPath)(startPath),
		(C.VGPath)(endPath),
		(C.VGfloat)(amount),
	)
	return (BooleanEnum)(ret)
}

func PathLength(
	path uint32,
	startSegment int32,
	numSegments int32,
) float32 {
	ret := C.vgPathLength(
		(C.VGPath)(path),
		(C.VGint)(startSegment),
		(C.VGint)(numSegments),
	)
	return (float32)(ret)
}

func PointAlongPath(
	path uint32,
	startSegment int32,
	numSegments int32,
	distance float32,
	x *float32,
	y *float32,
	tangentX *float32,
	tangentY *float32,
) {
	C.vgPointAlongPath(
		(C.VGPath)(path),
		(C.VGint)(startSegment),
		(C.VGint)(numSegments),
		(C.VGfloat)(distance),
		(*C.VGfloat)(x),
		(*C.VGfloat)(y),
		(*C.VGfloat)(tangentX),
		(*C.VGfloat)(tangentY),
	)
}

func PathBounds(
	path uint32,
	minX *float32,
	minY *float32,
	width *float32,
	height *float32,
) {
	C.vgPathBounds(
		(C.VGPath)(path),
		(*C.VGfloat)(minX),
		(*C.VGfloat)(minY),
		(*C.VGfloat)(width),
		(*C.VGfloat)(height),
	)
}

func PathTransformedBounds(
	path uint32,
	minX *float32,
	minY *float32,
	width *float32,
	height *float32,
) {
	C.vgPathTransformedBounds(
		(C.VGPath)(path),
		(*C.VGfloat)(minX),
		(*C.VGfloat)(minY),
		(*C.VGfloat)(width),
		(*C.VGfloat)(height),
	)
}

func DrawPath(
	path uint32,
	paintModes uint32,
) {
	C.vgDrawPath(
		(C.VGPath)(path),
		(C.VGbitfield)(paintModes),
	)
}

func CreatePaint(
) uint32 {
	ret := C.vgCreatePaint(
	)
	return (uint32)(ret)
}

func DestroyPaint(
	paint uint32,
) {
	C.vgDestroyPaint(
		(C.VGPaint)(paint),
	)
}

func SetPaint(
	paint uint32,
	paintModes uint32,
) {
	C.vgSetPaint(
		(C.VGPaint)(paint),
		(C.VGbitfield)(paintModes),
	)
}

func GetPaint(
	paintMode PaintModeEnum,
) uint32 {
	ret := C.vgGetPaint(
		(C.VGPaintMode)(paintMode),
	)
	return (uint32)(ret)
}

func SetColor(
	paint uint32,
	rgba uint32,
) {
	C.vgSetColor(
		(C.VGPaint)(paint),
		(C.VGuint)(rgba),
	)
}

func GetColor(
	paint uint32,
) uint32 {
	ret := C.vgGetColor(
		(C.VGPaint)(paint),
	)
	return (uint32)(ret)
}

func PaintPattern(
	paint uint32,
	pattern uint32,
) {
	C.vgPaintPattern(
		(C.VGPaint)(paint),
		(C.VGImage)(pattern),
	)
}

func CreateImage(
	format ImageFormatEnum,
	width int32,
	height int32,
	allowedQuality uint32,
) uint32 {
	ret := C.vgCreateImage(
		(C.VGImageFormat)(format),
		(C.VGint)(width),
		(C.VGint)(height),
		(C.VGbitfield)(allowedQuality),
	)
	return (uint32)(ret)
}

func DestroyImage(
	image uint32,
) {
	C.vgDestroyImage(
		(C.VGImage)(image),
	)
}

func ClearImage(
	image uint32,
	x int32,
	y int32,
	width int32,
	height int32,
) {
	C.vgClearImage(
		(C.VGImage)(image),
		(C.VGint)(x),
		(C.VGint)(y),
		(C.VGint)(width),
		(C.VGint)(height),
	)
}

func ImageSubData(
	image uint32,
	data *byte,
	dataStride int32,
	dataFormat ImageFormatEnum,
	x int32,
	y int32,
	width int32,
	height int32,
) {
	C.vgImageSubData(
		(C.VGImage)(image),
		unsafe.Pointer(data),
		(C.VGint)(dataStride),
		(C.VGImageFormat)(dataFormat),
		(C.VGint)(x),
		(C.VGint)(y),
		(C.VGint)(width),
		(C.VGint)(height),
	)
}

func GetImageSubData(
	image uint32,
	data *byte,
	dataStride int32,
	dataFormat ImageFormatEnum,
	x int32,
	y int32,
	width int32,
	height int32,
) {
	C.vgGetImageSubData(
		(C.VGImage)(image),
		unsafe.Pointer(data),
		(C.VGint)(dataStride),
		(C.VGImageFormat)(dataFormat),
		(C.VGint)(x),
		(C.VGint)(y),
		(C.VGint)(width),
		(C.VGint)(height),
	)
}

func ChildImage(
	parent uint32,
	x int32,
	y int32,
	width int32,
	height int32,
) uint32 {
	ret := C.vgChildImage(
		(C.VGImage)(parent),
		(C.VGint)(x),
		(C.VGint)(y),
		(C.VGint)(width),
		(C.VGint)(height),
	)
	return (uint32)(ret)
}

func GetParent(
	image uint32,
) uint32 {
	ret := C.vgGetParent(
		(C.VGImage)(image),
	)
	return (uint32)(ret)
}

func CopyImage(
	dst uint32,
	dx int32,
	dy int32,
	src uint32,
	sx int32,
	sy int32,
	width int32,
	height int32,
	dither BooleanEnum,
) {
	C.vgCopyImage(
		(C.VGImage)(dst),
		(C.VGint)(dx),
		(C.VGint)(dy),
		(C.VGImage)(src),
		(C.VGint)(sx),
		(C.VGint)(sy),
		(C.VGint)(width),
		(C.VGint)(height),
		(C.VGboolean)(dither),
	)
}

func DrawImage(
	image uint32,
) {
	C.vgDrawImage(
		(C.VGImage)(image),
	)
}

func SetPixels(
	dx int32,
	dy int32,
	src uint32,
	sx int32,
	sy int32,
	width int32,
	height int32,
) {
	C.vgSetPixels(
		(C.VGint)(dx),
		(C.VGint)(dy),
		(C.VGImage)(src),
		(C.VGint)(sx),
		(C.VGint)(sy),
		(C.VGint)(width),
		(C.VGint)(height),
	)
}

func WritePixels(
	data *byte,
	dataStride int32,
	dataFormat ImageFormatEnum,
	dx int32,
	dy int32,
	width int32,
	height int32,
) {
	C.vgWritePixels(
		unsafe.Pointer(data),
		(C.VGint)(dataStride),
		(C.VGImageFormat)(dataFormat),
		(C.VGint)(dx),
		(C.VGint)(dy),
		(C.VGint)(width),
		(C.VGint)(height),
	)
}

func GetPixels(
	dst uint32,
	dx int32,
	dy int32,
	sx int32,
	sy int32,
	width int32,
	height int32,
) {
	C.vgGetPixels(
		(C.VGImage)(dst),
		(C.VGint)(dx),
		(C.VGint)(dy),
		(C.VGint)(sx),
		(C.VGint)(sy),
		(C.VGint)(width),
		(C.VGint)(height),
	)
}

func ReadPixels(
	data *byte,
	dataStride int32,
	dataFormat ImageFormatEnum,
	sx int32,
	sy int32,
	width int32,
	height int32,
) {
	C.vgReadPixels(
		unsafe.Pointer(data),
		(C.VGint)(dataStride),
		(C.VGImageFormat)(dataFormat),
		(C.VGint)(sx),
		(C.VGint)(sy),
		(C.VGint)(width),
		(C.VGint)(height),
	)
}

func CopyPixels(
	dx int32,
	dy int32,
	sx int32,
	sy int32,
	width int32,
	height int32,
) {
	C.vgCopyPixels(
		(C.VGint)(dx),
		(C.VGint)(dy),
		(C.VGint)(sx),
		(C.VGint)(sy),
		(C.VGint)(width),
		(C.VGint)(height),
	)
}

func CreateFont(
	glyphCapacityHint int32,
) uint32 {
	ret := C.vgCreateFont(
		(C.VGint)(glyphCapacityHint),
	)
	return (uint32)(ret)
}

func DestroyFont(
	font uint32,
) {
	C.vgDestroyFont(
		(C.VGFont)(font),
	)
}

func SetGlyphToPath(
	font uint32,
	glyphIndex uint32,
	path uint32,
	isHinted BooleanEnum,
	glyphOrigin [2]float32,
	escapement [2]float32,
) {
	C.vgSetGlyphToPath(
		(C.VGFont)(font),
		(C.VGuint)(glyphIndex),
		(C.VGPath)(path),
		(C.VGboolean)(isHinted),
		(*C.VGfloat)(&glyphOrigin[0]),
		(*C.VGfloat)(&escapement[0]),
	)
}

func SetGlyphToImage(
	font uint32,
	glyphIndex uint32,
	image uint32,
	glyphOrigin [2]float32,
	escapement [2]float32,
) {
	C.vgSetGlyphToImage(
		(C.VGFont)(font),
		(C.VGuint)(glyphIndex),
		(C.VGImage)(image),
		(*C.VGfloat)(&glyphOrigin[0]),
		(*C.VGfloat)(&escapement[0]),
	)
}

func ClearGlyph(
	font uint32,
	glyphIndex uint32,
) {
	C.vgClearGlyph(
		(C.VGFont)(font),
		(C.VGuint)(glyphIndex),
	)
}

func DrawGlyph(
	font uint32,
	glyphIndex uint32,
	paintModes uint32,
	allowAutoHinting BooleanEnum,
) {
	C.vgDrawGlyph(
		(C.VGFont)(font),
		(C.VGuint)(glyphIndex),
		(C.VGbitfield)(paintModes),
		(C.VGboolean)(allowAutoHinting),
	)
}

func DrawGlyphs(
	font uint32,
	glyphCount int32,
	glyphIndices *uint32,
	adjustments_x *float32,
	adjustments_y *float32,
	paintModes uint32,
	allowAutoHinting BooleanEnum,
) {
	C.vgDrawGlyphs(
		(C.VGFont)(font),
		(C.VGint)(glyphCount),
		(*C.VGuint)(glyphIndices),
		(*C.VGfloat)(adjustments_x),
		(*C.VGfloat)(adjustments_y),
		(C.VGbitfield)(paintModes),
		(C.VGboolean)(allowAutoHinting),
	)
}

func ColorMatrix(
	dst uint32,
	src uint32,
	matrix *float32,
) {
	C.vgColorMatrix(
		(C.VGImage)(dst),
		(C.VGImage)(src),
		(*C.VGfloat)(matrix),
	)
}

func Convolve(
	dst uint32,
	src uint32,
	kernelWidth int32,
	kernelHeight int32,
	shiftX int32,
	shiftY int32,
	kernel *int16,
	scale float32,
	bias float32,
	tilingMode TilingModeEnum,
) {
	C.vgConvolve(
		(C.VGImage)(dst),
		(C.VGImage)(src),
		(C.VGint)(kernelWidth),
		(C.VGint)(kernelHeight),
		(C.VGint)(shiftX),
		(C.VGint)(shiftY),
		(*C.VGshort)(kernel),
		(C.VGfloat)(scale),
		(C.VGfloat)(bias),
		(C.VGTilingMode)(tilingMode),
	)
}

func SeparableConvolve(
	dst uint32,
	src uint32,
	kernelWidth int32,
	kernelHeight int32,
	shiftX int32,
	shiftY int32,
	kernelX *int16,
	kernelY *int16,
	scale float32,
	bias float32,
	tilingMode TilingModeEnum,
) {
	C.vgSeparableConvolve(
		(C.VGImage)(dst),
		(C.VGImage)(src),
		(C.VGint)(kernelWidth),
		(C.VGint)(kernelHeight),
		(C.VGint)(shiftX),
		(C.VGint)(shiftY),
		(*C.VGshort)(kernelX),
		(*C.VGshort)(kernelY),
		(C.VGfloat)(scale),
		(C.VGfloat)(bias),
		(C.VGTilingMode)(tilingMode),
	)
}

func GaussianBlur(
	dst uint32,
	src uint32,
	stdDeviationX float32,
	stdDeviationY float32,
	tilingMode TilingModeEnum,
) {
	C.vgGaussianBlur(
		(C.VGImage)(dst),
		(C.VGImage)(src),
		(C.VGfloat)(stdDeviationX),
		(C.VGfloat)(stdDeviationY),
		(C.VGTilingMode)(tilingMode),
	)
}

func Lookup(
	dst uint32,
	src uint32,
	redLUT *uint8,
	greenLUT *uint8,
	blueLUT *uint8,
	alphaLUT *uint8,
	outputLinear BooleanEnum,
	outputPremultiplied BooleanEnum,
) {
	C.vgLookup(
		(C.VGImage)(dst),
		(C.VGImage)(src),
		(*C.VGubyte)(redLUT),
		(*C.VGubyte)(greenLUT),
		(*C.VGubyte)(blueLUT),
		(*C.VGubyte)(alphaLUT),
		(C.VGboolean)(outputLinear),
		(C.VGboolean)(outputPremultiplied),
	)
}

func LookupSingle(
	dst uint32,
	src uint32,
	lookupTable *uint32,
	sourceChannel ImageChannelEnum,
	outputLinear BooleanEnum,
	outputPremultiplied BooleanEnum,
) {
	C.vgLookupSingle(
		(C.VGImage)(dst),
		(C.VGImage)(src),
		(*C.VGuint)(lookupTable),
		(C.VGImageChannel)(sourceChannel),
		(C.VGboolean)(outputLinear),
		(C.VGboolean)(outputPremultiplied),
	)
}

func HardwareQuery(
	key HardwareQueryTypeEnum,
	setting int32,
) HardwareQueryResultEnum {
	ret := C.vgHardwareQuery(
		(C.VGHardwareQueryType)(key),
		(C.VGint)(setting),
	)
	return (HardwareQueryResultEnum)(ret)
}

func GetString(
	name StringIDEnum,
) *uint8 {
	ret := C.vgGetString(
		(C.VGStringID)(name),
	)
	return (*uint8)(ret)
}
