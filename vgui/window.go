package vgui

type Point struct {
	X float32
	Y float32
}

type Window struct {
	X float32
	Y float32
	W float32
	H float32
}

func NewWindow(x, y, w, h float32) Window {
	return Window{x, y, w, h}
}

func (n Window) Inner(l, t, r, b float32) Window {
	return Window{n.X + l, n.Y + t, n.W - r - l, n.H - t - b}
}

func (n Window) SplitH(t float32) (top Window, bottom Window) {
	bottom = Window{n.X, n.Y, n.W, t - 1}
	top = Window{n.X, n.Y + t, n.W, n.H - t}
	return
}

func (n Window) SplitV(l float32) (left Window, right Window) {
	left = Window{n.X, n.Y, l - 1, n.H}
	right = Window{n.X + l, n.Y, n.W - l, n.H}
	return
}

func (w Window) AlignedPoint(align Alignment) Point {
	// Default is AlignLeft,AlignTop:
	x := w.X
	y := w.Y
	// Respect alignment flags by moving the text point within the window:
	if align&AlignCenter != 0 {
		x += w.W * 0.5
	} else if align&AlignRight != 0 {
		x += w.W
	}
	if align&AlignMiddle != 0 {
		y += w.H * 0.5
	} else if align&AlignBottom != 0 {
		y += w.H
	}
	return Point{x, y}
}

func (w Window) RadiusMin() float32 {
	wr := w.W * 0.5
	hr := w.H * 0.5
	if wr < hr {
		return wr
	}
	return hr
}

func (w Window) IsPointInside(p Point) bool {
	//fmt.Println(p, "in", w)
	return p.X >= w.X && p.X <= w.X+w.W &&
		p.Y >= w.Y && p.Y <= w.Y+w.H
}
