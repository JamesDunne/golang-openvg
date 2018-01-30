package vgui

import (
	"math"
	"unicode/utf8"

	"github.com/JamesDunne/golang-openvg/vg"
	"github.com/JamesDunne/golang-openvg/vgu"
)

type Alignment int32

const (
	// Horizontal alignment:
	AlignLeft Alignment = 1 << iota
	AlignCenter
	AlignRight
	// Vertical alignment:
	AlignTop
	AlignMiddle
	AlignBottom
)

type Color [4]uint8

func RGB(r, g, b uint8) Color {
	return Color([4]uint8{r, g, b, 255})
}

func RGBA(r, g, b, a uint8) Color {
	return Color([4]uint8{r, g, b, a})
}

func (c *Color) UInt32() uint32 {
	return uint32(c[0])<<24 | uint32(c[1])<<16 | uint32(c[2])<<8 | uint32(c[3])
}

type UIPalette [5]Color
type PaletteIndex int

var (
	// Palettes are ordered from darkest to lightest shades:
	// This palette from https://flatuicolors.com/
	//palette0 = UIPalette{
	//	nvg.RGB(44, 62, 80),    // midnight blue
	//	nvg.RGB(52, 73, 94),    // wet asphalt
	//	nvg.RGB(127, 140, 141), // asbestos
	//	nvg.RGB(149, 165, 166), // concrete
	//}
	// http://flatcolors.net/palette/724-flat-design-blue
	//palette = UIPalette{
	//	nvg.RGB(0x00, 0x00, 0x40),
	//	nvg.RGB(0x00, 0x4C, 0x79),
	//	nvg.RGB(0x00, 0x5F, 0x97),
	//	nvg.RGB(0x3C, 0x8A, 0xB8),
	//	nvg.RGB(0xE8, 0xF4, 0xFA),
	//}
	// http://flatcolors.net/palette/462-flat-existence#
	//palette = UIPalette{
	//	nvg.RGB(0x13, 0x1A, 0x1E),
	//	nvg.RGB(0x11, 0x36, 0xC7),
	//	nvg.RGB(0x1C, 0x57, 0xE1),
	//	nvg.RGB(0x59, 0x7D, 0xF7),
	//	nvg.RGB(0x77, 0x9B, 0xF0),
	//}
	palette = UIPalette{
		RGB(0x00, 0x00, 0x00),
		RGB(0x02, 0x05, 0x24),
		RGB(0x0D, 0x28, 0x4F),
		RGB(0x71, 0x78, 0x9F),
		RGB(0xb4, 0xb7, 0xF7),
	}
)

type Touch struct {
	Point
	ID int32
}

type UI struct {
	p UIPalette
	w Window

	font        *Font
	fillPaint   uint32
	strokePaint uint32
	path        uint32

	saved struct {
		strokeWidth      float32
		strokeMiterLimit float32
		lineCap          int32
	}

	Touches []Touch
}

func NewUI() *UI {
	ui := &UI{
		p:       palette,
		Touches: make([]Touch, 10),
	}

	return ui
}

func (u *UI) SetWindow(w Window) {
	u.w = w
}

func (u *UI) Window() Window {
	return u.w
}

func (ui *UI) Init() {
	ui.font = NewSansFont()

	ui.fillPaint = vg.CreatePaint()
	vg.SetParameteri(ui.fillPaint, int32(vg.PaintType), int32(vg.PaintTypeColor))
	vg.SetParameterfv(ui.fillPaint, int32(vg.PaintColor), 4, &[]float32{1.0, 1.0, 1.0, 1.0}[0])
	vg.SetPaint(ui.fillPaint, uint32(vg.FillPath))

	ui.strokePaint = vg.CreatePaint()
	vg.SetParameteri(ui.strokePaint, int32(vg.PaintType), int32(vg.PaintTypeColor))
	vg.SetParameterfv(ui.strokePaint, int32(vg.PaintColor), 4, &[]float32{1.0, 1.0, 1.0, 1.0}[0])
	vg.SetPaint(ui.strokePaint, uint32(vg.StrokePath))

	clearColor := &[]float32{0.0, 0.0, 0.0, 1.0}[0]

	// set some default parameters for the OpenVG context
	vg.Seti(vg.FillRule, int32(vg.EvenOdd))
	//vg.Seti(vg.FillRule, int32(vg.NonZero))

	vg.Setfv(vg.ClearColor, 4, clearColor)
	vg.Setf(vg.StrokeLineWidth, 0.125)
	vg.Seti(vg.StrokeCapStyle, int32(vg.CapButt))
	vg.Seti(vg.StrokeJoinStyle, int32(vg.JoinBevel))
	vg.Seti(vg.RenderingQuality, int32(vg.RenderingQualityBetter))
	vg.Seti(vg.BlendMode, int32(vg.BlendSrc))
}

func (ui *UI) BeginFrame() {
	vg.LoadIdentity()
	vg.Clear(0, 0, int32(ui.w.W), int32(ui.w.H))
}

func (ui *UI) EndFrame() {
}

func (u *UI) FontFace(name string) {
	// No-op
}

type PreparedText struct {
	String string

	font    *Font
	runes   []uint32
	adjustX []float32
	adjustY []float32
	width   float32
}

func (u *UI) Palette(p PaletteIndex) Color {
	return u.p[p]
}

func (u *UI) Save() {
	u.saved.strokeWidth = vg.Getf(vg.StrokeLineWidth)
	u.saved.strokeMiterLimit = vg.Getf(vg.StrokeMiterLimit)
	u.saved.lineCap = vg.Geti(vg.StrokeCapStyle)
}

func (u *UI) Restore() {
	vg.Setf(vg.StrokeLineWidth, u.saved.strokeWidth)
	vg.Setf(vg.StrokeMiterLimit, u.saved.strokeMiterLimit)
	vg.Seti(vg.StrokeCapStyle, u.saved.lineCap)
}

func (u *UI) MiterLimit(limit float32) {
	vg.Setf(vg.StrokeMiterLimit, limit)
}

func (u *UI) LineCap(style vg.CapStyleEnum) {
	vg.Seti(vg.StrokeCapStyle, int32(style))
}

func (u *UI) FillColor(c Color) {
	vg.SetColor(u.fillPaint, c.UInt32())
}

func (u *UI) StrokeColor(c Color) {
	vg.SetColor(u.strokePaint, c.UInt32())
}

func (u *UI) StrokeWidth(size float32) {
	vg.Setf(vg.StrokeLineWidth, size)
}

func (u *UI) BeginPath() {
	if u.path != 0 {
		vg.DestroyPath(u.path)
	}
	u.path = vg.CreatePath(0, vg.PathDatatypeF, 1.0, 0.0, 0, 0, uint32(vg.PathCapabilityAll))
}

func (u *UI) Fill() {
	// TODO: figure this out
	vg.DrawPath(u.path, uint32(vg.FillPath))
}

func (u *UI) Stroke() {
	vg.DrawPath(u.path, uint32(vg.StrokePath))
}

func (u *UI) y(y float32) float32 {
	return u.w.H - y
}

func (u *UI) Rect(w Window) {
	vgu.Rect(u.path, w.X, u.y(w.Y), w.W-1, w.H-1)
}

func (u *UI) RoundedRect(w Window, radius float32) {
	vgu.RoundRect(u.path, w.X, u.y(w.Y+w.H), w.W-1, w.H-1, radius*2, radius*2)
}

func (u *UI) Circle(p Point, r float32) {
	vgu.Ellipse(u.path, p.X, u.y(p.Y), r*2, r*2)
}

func (ui *UI) PrepareText(s string) *PreparedText {
	f := ui.font

	t := &PreparedText{
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

var glyphOrigin = []float32{0, 0}

func (u *UI) TextPoint(p Point, size float32, align Alignment, t *PreparedText) {
	f := t.font

	x := p.X
	y := u.y(p.Y + size)

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
	vg.Setfv(vg.GlyphOrigin, 2, &glyphOrigin[0])
	vg.LoadIdentity()
	vg.Translate(x, y)
	vg.Scale(size, size)
	vg.DrawGlyphs(f.vgHandle, int32(len(t.runes)), &t.runes[0], &t.adjustX[0], &t.adjustY[0], uint32(vg.FillPath), vg.False)

	if mm != int32(vg.MatrixGlyphUserToSurface) {
		vg.Seti(vg.MatrixMode, mm)
	}
}

//func Text(s string, x, y, size float32, align Alignment, f *Font) {
//	t := NewText(s, f)
//	TextRendered(t, x, y, size, align)
//}

func (u *UI) Text(w Window, size float32, align Alignment, t *PreparedText) {
	u.TextPoint(w.AlignedPoint(align), size, align, t)
}

// Angles in radians, 0 is horizontal extending right.
func (u *UI) Arc(p Point, r, a0, a1 float32) {
	vgu.Arc(u.path, p.X, u.y(p.Y), r*2, r*2, a0*-180/math.Pi, (a1-a0)*-180/math.Pi, vgu.ArcOpen)
}
