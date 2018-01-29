// main
package main

import (
	"fmt"

	"github.com/JamesDunne/golang-openvg/host"
	"github.com/JamesDunne/golang-openvg/vg"
)

func initVG(width, height int32) {
	clearColor := &[]float32{0.0, 0.0, 1.0, 1.0}[0]
	tileColor := &[]float32{0.0, 1.0, 0.0, 1.0}[0]

	// set some default parameters for the OpenVG context
	vg.Seti(vg.FillRule, int32(vg.EvenOdd))
	vg.Setfv(vg.ClearColor, 4, clearColor)
	vg.Setfv(vg.TileFillColor, 4, tileColor)
	vg.Setf(vg.StrokeLineWidth, 20.0)
	vg.Seti(vg.StrokeCapStyle, int32(vg.CapButt))
	vg.Seti(vg.StrokeJoinStyle, int32(vg.JoinBevel))
	vg.Seti(vg.RenderingQuality, int32(vg.RenderingQualityBetter))
	vg.Seti(vg.BlendMode, int32(vg.BlendSrc))
}

// Rendering is done on a specific OS thread managed by host:
func drawVG(width, height int32) {
	vg.Clear(0, 0, width, height)

	//color := vg.CreatePaint()
	//vg.SetPaint(color, uint32(vg.FillPath|vg.StrokePath))
}

func main() {
	fmt.Println("Hello World!")

	// Supply Go function callbacks:
	host.InitFunc = initVG
	host.DrawFunc = drawVG

	// Initialize the host to display OpenVG graphics at 800x480 resolution:
	host.Init(800, 480)

	// Event polling with idle loop:
	for host.PollEvent() {
	}
}
