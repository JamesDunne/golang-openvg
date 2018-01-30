// font
package host

//#include "DejaVuSans.h"
import "C"

import (
	"unsafe"

	"github.com/JamesDunne/golang-openvg/vg"
)

type Font struct {
	count           int
	descenderHeight C.int
	fontHeight      C.int
	characterMap    []C.short
	glyphAdvances   []C.int
	glyphs          []uint32
}

func NewSansFont() *Font {
	count := int(C.DejaVuSans_glyphCount)

	f := &Font{
		count:           count,
		fontHeight:      C.DejaVuSans_font_height,
		descenderHeight: C.DejaVuSans_descender_height,
		characterMap:    make([]C.short, 500),
		glyphAdvances:   make([]C.int, count),
		glyphs:          make([]uint32, count),
	}

	n := 0
	n = copy(f.characterMap, C.DejaVuSans_characterMap[:])
	if n != 500 {
		panic("n != 500")
	}
	n = copy(f.glyphAdvances, C.DejaVuSans_glyphAdvances[:])
	if n != count {
		panic("n != count")
	}

	for i := 0; i < count; i++ {
		p := unsafe.Pointer(&C.DejaVuSans_glyphPoints[C.DejaVuSans_glyphPointIndices[i]*2])
		instructions := (*uint8)(&C.DejaVuSans_glyphInstructions[C.DejaVuSans_glyphInstructionIndices[i]])
		ic := int32(C.DejaVuSans_glyphInstructionCounts[i])

		path := vg.CreatePath(
			0, /* VG_PATH_FORMAT_STANDARD */
			vg.PathDatatypeS32,
			1.0/65536.0, 0.0,
			0, 0,
			uint32(vg.PathCapabilityAll))
		f.glyphs[i] = path

		if ic != 0 {
			vg.AppendPathData(path, ic, instructions, p)
			if vg.GetError() != vg.NoError {
				panic("error!")
			}
		}
	}

	return f
}

func Text(s string, f *Font) {
	size := float32(22.0)
	x := float32(0)
	y := float32(0)
	xx := x
	var mm [9]float32

	vg.GetMatrix(&mm[0])

	for _, character := range s {
		glyph := f.characterMap[character]
		if glyph == -1 {
			continue
		}

		mat := [9]float32{
			size, 0, 0,
			0, size, 0,
			xx, y, 1.0,
		}
		vg.LoadMatrix(&mm[0])
		vg.MultMatrix(&mat[0])
		_ = mat

		//		vg.DrawPath(f.glyphs[glyph], uint32(vg.FillPath))
		vg.DrawPath(f.glyphs[glyph], uint32(vg.StrokePath))
		if vg.GetError() != vg.NoError {
			panic("error!")
		}

		xx += size * float32(f.glyphAdvances[glyph]) / 65536.0
	}

	vg.LoadMatrix(&mm[0])
}
