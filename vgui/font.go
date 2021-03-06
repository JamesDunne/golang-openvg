// font
package vgui

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
