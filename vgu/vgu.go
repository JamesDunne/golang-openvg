package vgu

//#include "VG/vgu.h"
import "C"

import "github.com/JamesDunne/golang-openvg/vg"

type ErrorCodeEnum int32

const (
	NoError              ErrorCodeEnum = 0
	BadHandleError       ErrorCodeEnum = 61440
	IllegalArgumentError ErrorCodeEnum = 61441
	OutOfMemoryError     ErrorCodeEnum = 61442
	PathCapabilityError  ErrorCodeEnum = 61443
	BadWarpError         ErrorCodeEnum = 61444
	ErrorCodeForceSize   ErrorCodeEnum = 2147483647
)

type ArcTypeEnum int32

const (
	ArcOpen          ArcTypeEnum = 61696
	ArcChord         ArcTypeEnum = 61697
	ArcPie           ArcTypeEnum = 61698
	ArcTypeForceSize ArcTypeEnum = 2147483647
)

func Line(
	path uint32,
	x0 float32,
	y0 float32,
	x1 float32,
	y1 float32,
) ErrorCodeEnum {
	ret := C.vguLine(
		(C.VGPath)(path),
		(C.VGfloat)(x0),
		(C.VGfloat)(y0),
		(C.VGfloat)(x1),
		(C.VGfloat)(y1),
	)
	return (ErrorCodeEnum)(ret)
}

func Polygon(
	path uint32,
	points *float32,
	count int32,
	closed vg.BooleanEnum,
) ErrorCodeEnum {
	ret := C.vguPolygon(
		(C.VGPath)(path),
		(*C.VGfloat)(points),
		(C.VGint)(count),
		(C.VGboolean)(closed),
	)
	return (ErrorCodeEnum)(ret)
}

func Rect(
	path uint32,
	x float32,
	y float32,
	width float32,
	height float32,
) ErrorCodeEnum {
	ret := C.vguRect(
		(C.VGPath)(path),
		(C.VGfloat)(x),
		(C.VGfloat)(y),
		(C.VGfloat)(width),
		(C.VGfloat)(height),
	)
	return (ErrorCodeEnum)(ret)
}

func RoundRect(
	path uint32,
	x float32,
	y float32,
	width float32,
	height float32,
	arcWidth float32,
	arcHeight float32,
) ErrorCodeEnum {
	ret := C.vguRoundRect(
		(C.VGPath)(path),
		(C.VGfloat)(x),
		(C.VGfloat)(y),
		(C.VGfloat)(width),
		(C.VGfloat)(height),
		(C.VGfloat)(arcWidth),
		(C.VGfloat)(arcHeight),
	)
	return (ErrorCodeEnum)(ret)
}

func Ellipse(
	path uint32,
	cx float32,
	cy float32,
	width float32,
	height float32,
) ErrorCodeEnum {
	ret := C.vguEllipse(
		(C.VGPath)(path),
		(C.VGfloat)(cx),
		(C.VGfloat)(cy),
		(C.VGfloat)(width),
		(C.VGfloat)(height),
	)
	return (ErrorCodeEnum)(ret)
}

func Arc(
	path uint32,
	x float32,
	y float32,
	width float32,
	height float32,
	startAngle float32,
	angleExtent float32,
	arcType ArcTypeEnum,
) ErrorCodeEnum {
	ret := C.vguArc(
		(C.VGPath)(path),
		(C.VGfloat)(x),
		(C.VGfloat)(y),
		(C.VGfloat)(width),
		(C.VGfloat)(height),
		(C.VGfloat)(startAngle),
		(C.VGfloat)(angleExtent),
		(C.VGUArcType)(arcType),
	)
	return (ErrorCodeEnum)(ret)
}

func ComputeWarpQuadToSquare(
	sx0 float32,
	sy0 float32,
	sx1 float32,
	sy1 float32,
	sx2 float32,
	sy2 float32,
	sx3 float32,
	sy3 float32,
	matrix *float32,
) ErrorCodeEnum {
	ret := C.vguComputeWarpQuadToSquare(
		(C.VGfloat)(sx0),
		(C.VGfloat)(sy0),
		(C.VGfloat)(sx1),
		(C.VGfloat)(sy1),
		(C.VGfloat)(sx2),
		(C.VGfloat)(sy2),
		(C.VGfloat)(sx3),
		(C.VGfloat)(sy3),
		(*C.VGfloat)(matrix),
	)
	return (ErrorCodeEnum)(ret)
}

func ComputeWarpSquareToQuad(
	dx0 float32,
	dy0 float32,
	dx1 float32,
	dy1 float32,
	dx2 float32,
	dy2 float32,
	dx3 float32,
	dy3 float32,
	matrix *float32,
) ErrorCodeEnum {
	ret := C.vguComputeWarpSquareToQuad(
		(C.VGfloat)(dx0),
		(C.VGfloat)(dy0),
		(C.VGfloat)(dx1),
		(C.VGfloat)(dy1),
		(C.VGfloat)(dx2),
		(C.VGfloat)(dy2),
		(C.VGfloat)(dx3),
		(C.VGfloat)(dy3),
		(*C.VGfloat)(matrix),
	)
	return (ErrorCodeEnum)(ret)
}

func ComputeWarpQuadToQuad(
	dx0 float32,
	dy0 float32,
	dx1 float32,
	dy1 float32,
	dx2 float32,
	dy2 float32,
	dx3 float32,
	dy3 float32,
	sx0 float32,
	sy0 float32,
	sx1 float32,
	sy1 float32,
	sx2 float32,
	sy2 float32,
	sx3 float32,
	sy3 float32,
	matrix *float32,
) ErrorCodeEnum {
	ret := C.vguComputeWarpQuadToQuad(
		(C.VGfloat)(dx0),
		(C.VGfloat)(dy0),
		(C.VGfloat)(dx1),
		(C.VGfloat)(dy1),
		(C.VGfloat)(dx2),
		(C.VGfloat)(dy2),
		(C.VGfloat)(dx3),
		(C.VGfloat)(dy3),
		(C.VGfloat)(sx0),
		(C.VGfloat)(sy0),
		(C.VGfloat)(sx1),
		(C.VGfloat)(sy1),
		(C.VGfloat)(sx2),
		(C.VGfloat)(sy2),
		(C.VGfloat)(sx3),
		(C.VGfloat)(sy3),
		(*C.VGfloat)(matrix),
	)
	return (ErrorCodeEnum)(ret)
}
