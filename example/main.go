// main
package main

import (
	"fmt"
	"runtime"

	"github.com/JamesDunne/golang-openvg/host"
	"github.com/JamesDunne/golang-openvg/vg"
)

func draw(width, height int32) {
	vg.Setfv(vg.ClearColor, 4, &[]float32{0.0, 1.0, 0.0, 1.0}[0])
	vg.Clear(0, 0, 800, 480)

}

func main() {
	runtime.LockOSThread()
	fmt.Println("Hello World!")
	h := host.NewHost(800, 480)
	h.UseDrawFunc(draw)
	for h.PollEvent() {
	}
}
