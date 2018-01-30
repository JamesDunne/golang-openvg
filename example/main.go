// main
package main

import (
	"runtime"

	"github.com/JamesDunne/golang-openvg/host"
	"github.com/JamesDunne/golang-openvg/vgui"
)

var f *vgui.Font
var ui *vgui.UI

var amps [2]*vgui.PreparedText
var (
	ptSong      *vgui.PreparedText
	ptAcoustic  *vgui.PreparedText
	ptDirty     *vgui.PreparedText
	ptClean     *vgui.PreparedText
	ptGain      *vgui.PreparedText
	ptVolume    *vgui.PreparedText
	ptGainStr   *vgui.PreparedText
	ptVolumeStr *vgui.PreparedText
)

func initVG(width, height int32) {
	ui = vgui.NewUI()

	ui.Init()
	ui.SetWindow(vgui.NewWindow(0, 0, float32(width), float32(height)))

	amps = [2]*vgui.PreparedText{ui.PrepareText("MG"), ui.PrepareText("JD")}

	ptSong = ui.PrepareText("Trippin on a Hole in a Paper Heart")
	ptAcoustic = ui.PrepareText("acoustic")
	ptDirty = ui.PrepareText("dirty")
	ptClean = ui.PrepareText("clean")
	ptGain = ui.PrepareText("Gain")
	ptVolume = ui.PrepareText("Volume")
	ptGainStr = ui.PrepareText("0.68")
	ptVolumeStr = ui.PrepareText("0 dB")
}

const size = 16

// Rendering is done on a specific OS thread managed by host:
func drawVG(width, height int32) {
	// Update display window:
	ui.SetWindow(vgui.NewWindow(0, 0, float32(width), float32(height)))
	w := ui.Window()

	ui.BeginFrame()

	ui.FillColor(ui.Palette(0))
	ui.BeginPath()
	ui.Rect(w)
	ui.Fill()

	top, bottom := w.SplitH(size + 8)

	ui.Label(top, ptSong, vgui.AlignLeft|vgui.AlignTop)

	// Split screen for MG v JD:
	mg, jd := bottom.SplitV(bottom.W * 0.5)

	drawAmp := func(w vgui.Window, ampNo int) {
		ui.StrokeWidth(1.0)
		ui.StrokeColor(ui.Palette(3))
		ui.Pane(w)

		// Amp label at top center:
		label, w := w.SplitH(size + 8)
		ui.FillColor(ui.Palette(4))
		ui.Text(label, size, vgui.AlignCenter|vgui.AlignTop, amps[ampNo])

		// Tri-state buttons:
		top, bottom := w.SplitH(size + 16)
		btnHeight := top.W * 0.33333333
		btnDirty, top := top.SplitV(btnHeight)
		btnClean, btnAcoustic := top.SplitV(btnHeight)

		if t := ui.Button(btnDirty, true, "dirty"); t != nil {

		}
		ui.Button(btnClean, false, "clean")
		ui.Button(btnAcoustic, false, "acoustic")

		// FX toggles:
		fxWidth := bottom.W / 5.0
		mid, bottom := bottom.SplitH(bottom.H - (size + 16))
		fxNames := [...]string{"pit1", "rtr1", "phr1", "cho1", "dly1"}
		for i := 0; i < 5; i++ {
			var btnFX vgui.Window
			btnFX, bottom = bottom.SplitV(fxWidth)
			ui.Button(btnFX, false, fxNames[i])
		}

		ui.StrokeColor(ui.Palette(3))
		ui.Pane(mid)

		gain, volume := mid.SplitV(mid.W * 0.5)
		g := float32(96) / 127.0
		ui.Dial(gain, ptGain, g, ptGainStr)
		_, _ = gain, g
		v := float32(98) / 127.0
		ui.Dial(volume, ptVolume, v, ptVolumeStr)
		_, _ = volume, v
	}
	drawAmp(mg, 0)
	drawAmp(jd, 1)

	// Draw touch points:
	for _, tp := range ui.Touches {
		if tp.ID <= 0 {
			continue
		}

		ui.FillColor(vgui.RGBA(255, 255, 255, 160))
		ui.BeginPath()
		ui.Circle(tp.Point, 15.0)
		ui.Fill()
	}

	ui.EndFrame()
}

func main() {
	runtime.LockOSThread()

	// Supply Go function callbacks:
	host.InitFunc = initVG
	host.DrawFunc = drawVG

	// Initialize the host to display OpenVG graphics at 800x480 resolution:
	if !host.Init(800, 480) {
		return
	}

	// Event polling with idle loop:
	for host.PollEvent() {
	}

	host.Destroy()
}
