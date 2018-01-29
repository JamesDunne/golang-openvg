// main
package main

import (
	"fmt"

	"github.com/JamesDunne/golang-openvg/host"
	"github.com/JamesDunne/golang-openvg/vg"
)

// Rendering is done on a specific OS thread managed by host:
func draw(width, height int32) {
	vg.Setfv(vg.ClearColor, 4, &[]float32{0.0, 0.0, 1.0, 1.0}[0])
	vg.Clear(0, 0, width, height)
}

func main() {
	fmt.Println("Hello World!")

	// Initialize the host to display OpenVG graphics at 800x480 resolution:
	host.Init(800, 480)
	// Supply Go function callback for rendering:
	host.DrawFunc = draw

	// Event polling with idle loop:
	for host.PollEvent() {
	}
}
