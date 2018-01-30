// font
package host

import (
	"unsafe"

	"github.com/JamesDunne/golang-openvg/vg"
)

type Font struct {
	vgHandle uint32

	count           int
	descenderHeight float32
	fontHeight      float32
	characterMap    []int16
	glyphAdvances   []float32
	glyphs          []uint32
}

func NewSansFont() *Font {
	count := font_sans_glyphCount

	f := &Font{
		vgHandle:        vg.CreateFont(int32(count)),
		count:           count,
		fontHeight:      font_sans_font_height,
		descenderHeight: font_sans_descender_height,
		characterMap:    font_sans_characterMap[:],
		glyphAdvances:   font_sans_glyphAdvances[:],
		glyphs:          make([]uint32, count),
	}

	for i := 0; i < count; i++ {
		pi := font_sans_glyphPointIndices[i] * 2
		if pi >= int32(len(font_sans_glyphPoints)) {
			continue
		}
		p := unsafe.Pointer(&font_sans_glyphPoints[pi])
		instructions := (*uint8)(&font_sans_glyphInstructions[font_sans_glyphInstructionIndices[i]])
		ic := int32(font_sans_glyphInstructionCounts[i])

		path := vg.CreatePath(
			0, /* VG_PATH_FORMAT_STANDARD */
			vg.PathDatatypeF,
			1.0, 0.0,
			0, 0,
			uint32(vg.PathCapabilityAll))
		f.glyphs[i] = path

		if ic != 0 {
			vg.AppendPathData(path, ic, instructions, p)
			if vg.GetError() != vg.NoError {
				panic("VG error!")
			}
			vg.SetGlyphToPath(f.vgHandle, uint32(i), path, vg.False, [2]float32{0, 0}, [2]float32{f.glyphAdvances[i], 0.0})
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

		xx += size * float32(f.glyphAdvances[glyph])
	}

	vg.LoadMatrix(&mm[0])
}
