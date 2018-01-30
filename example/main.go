// main
package main

import (
	"runtime"

	"github.com/JamesDunne/golang-openvg/host"
	"github.com/JamesDunne/golang-openvg/vgui"
)

var f *vgui.Font
var ui *vgui.UI

func initVG(width, height int32) {
	ui = vgui.NewUI()

	ui.Init()
	ui.SetWindow(vgui.NewWindow(0, 0, float32(width), float32(height)))
}

// Rendering is done on a specific OS thread managed by host:
func drawVG(width, height int32) {
	// Update display window:
	ui.SetWindow(vgui.NewWindow(0, 0, float32(width), float32(height)))

	ui.BeginFrame()

	ui.FillColor(ui.Palette(4))
	ui.TextPoint(vgui.Point{150, 200}, 20, vgui.AlignLeft, "Tripping on a Hole in a Paper Heart")

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
