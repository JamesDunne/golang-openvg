package vgui

import (
	"math"

	"github.com/JamesDunne/golang-openvg/vg"
)

const pad = 2
const size = 28.0
const round = 4.0

func (ui *UI) isTouched(w Window) *Touch {
	for i := range ui.Touches {
		t := &ui.Touches[i]

		// Skip released touch points:
		if t.ID <= 0 {
			continue
		}

		p := Point{t.X, t.Y}
		if w.IsPointInside(p) {
			return t
		}
	}
	return nil
}

func (ui *UI) Button(w Window, depressed bool, label string) *Touch {
	touched := ui.isTouched(w)
	c1, c2, c3 := PaletteIndex(1), PaletteIndex(2), PaletteIndex(3)
	if touched != nil || depressed {
		c1, c2, c3 = 4, 3, 2
	}

	ui.StrokeColor(ui.Palette(c2))
	ui.FillColor(ui.Palette(c3))
	ui.BeginPath()
	ui.RoundedRect(w, round)
	ui.Stroke()
	ui.Fill()

	ui.FillColor(ui.Palette(c1))
	ui.Text(w, size, AlignCenter|AlignMiddle, label)

	return touched
}

func (ui *UI) Pane(w Window) {
	ui.BeginPath()
	ui.RoundedRect(w, round)
	ui.Stroke()
}

func (ui *UI) Label(w Window, string string, align Alignment) {
	ui.BeginPath()
	ui.RoundedRect(w, round)
	ui.StrokeColor(ui.Palette(3))
	ui.FillColor(ui.Palette(1))
	ui.Fill()
	ui.Stroke()

	lblText := w.Inner(pad*2, 0, pad*2, 0)
	ui.FillColor(ui.Palette(4))
	ui.Text(lblText, size, align, string)
}

func (ui *UI) Dial(w Window, label string, value float32, valueStr string) (Window, *Touch) {
	top, bottom := w.SplitH(size + 8)
	w, bottom = bottom.SplitH(bottom.H - (size + 8))

	c := w.AlignedPoint(AlignCenter | AlignMiddle)
	r := w.RadiusMin()

	// Labels on top and bottom:
	top = top.Inner(w.W*0.5-r, 0, w.W*0.5-r, 0)
	bottom = bottom.Inner(w.W*0.5-r, 0, w.W*0.5-r, 0)

	ui.Label(top, valueStr, AlignCenter|AlignMiddle)
	ui.Label(bottom, label, AlignCenter|AlignMiddle)

	arcWidth := r * 0.05

	// Clamp value between 0 and 1:
	clampedValue := value
	if clampedValue > 1.0 {
		clampedValue = 1.0
	} else if clampedValue < 0.0 {
		clampedValue = 0.0
	}

	ui.Save()

	// Filled center:
	ui.BeginPath()
	ui.Circle(c, r-arcWidth)
	ui.FillColor(ui.Palette(2))
	ui.Fill()

	// Highlighted arc:
	ui.BeginPath()
	ui.Arc(c, r-arcWidth*0.5, math.Pi*0.8, math.Pi*0.8+math.Pi*1.4*clampedValue)
	ui.StrokeWidth(arcWidth)
	ui.MiterLimit(1.0)
	ui.LineCap(vg.CapSquare)
	ui.StrokeColor(ui.Palette(3))
	ui.Stroke()

	ui.Restore()

	return w, ui.isTouched(w)
}
