// font
package vgui

import (
	"unicode/utf8"
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

		if ic != 0 {
			vg.AppendPathData(path, ic, instructions, p)
			if vg.GetError() != vg.NoError {
				panic("VG error!")
			}
			// remove "editing" capabilities, so that OpenVG driver can try to free some memory
			vg.RemovePathCapabilities(
				path,
				uint32(vg.PathCapabilityAppendFrom|vg.PathCapabilityAppendTo|
					vg.PathCapabilityModify|vg.PathCapabilityTransformFrom|
					vg.PathCapabilityTransformTo|vg.PathCapabilityInterpolateFrom|
					vg.PathCapabilityInterpolateTo),
			)
			if vg.GetError() != vg.NoError {
				panic("VG error!")
			}
		}

		// Create OpenVG font glyph:
		vg.SetGlyphToPath(f.vgHandle, uint32(i), path, vg.False, [2]float32{0, 0}, [2]float32{f.glyphAdvances[i], 0.0})
		if vg.GetError() != vg.NoError {
			panic("VG error!")
		}

		vg.DestroyPath(path)
		if vg.GetError() != vg.NoError {
			panic("VG error!")
		}
	}

	return f
}

type RenderedText struct {
	String string

	font    *Font
	runes   []uint32
	adjustX []float32
	adjustY []float32
	width   float32
}

func NewText(s string, f *Font) *RenderedText {
	t := &RenderedText{
		String: s,
		font:   f,
	}

	t.runes = make([]uint32, utf8.RuneCountInString(s))
	t.adjustX = make([]float32, len(t.runes))
	t.adjustY = make([]float32, len(t.runes))

	t.width = float32(0)
	for i, character := range s {
		glyph := f.characterMap[character]
		if glyph == -1 {
			continue
		}

		t.runes[i] = uint32(glyph)

		// TODO: kerning
		t.adjustX[i] = 0
		t.adjustY[i] = 0

		t.width += f.glyphAdvances[glyph]
	}

	return t
}

func TextRendered(t *RenderedText, x, y, size float32, align Alignment) {
	f := t.font

	if align&AlignCenter == AlignCenter {
		x -= t.width * size * 0.5
	} else if align&AlignRight == AlignRight {
		x -= t.width * size
	}
	y += f.descenderHeight * size
	if align&AlignMiddle == AlignMiddle {
		y += (f.fontHeight - f.descenderHeight) * size * 0.5
	} else if align&AlignBottom == AlignBottom {
		y += (f.fontHeight - f.descenderHeight) * size
	}

	mm := vg.Geti(vg.MatrixMode)
	if mm != int32(vg.MatrixGlyphUserToSurface) {
		vg.Seti(vg.MatrixMode, int32(vg.MatrixGlyphUserToSurface))
	}

	vg.Seti(vg.FillRule, int32(vg.NonZero))
	vg.Setfv(vg.GlyphOrigin, 2, &[]float32{0, 0}[0])
	vg.LoadIdentity()
	vg.Translate(x, y)
	vg.Scale(size, size)
	vg.DrawGlyphs(f.vgHandle, int32(len(t.runes)), &t.runes[0], &t.adjustX[0], &t.adjustY[0], uint32(vg.FillPath), vg.False)

	if mm != int32(vg.MatrixGlyphUserToSurface) {
		vg.Seti(vg.MatrixMode, mm)
	}
}

func Text(s string, x, y, size float32, align Alignment, f *Font) {
	t := NewText(s, f)
	TextRendered(t, x, y, size, align)
}
