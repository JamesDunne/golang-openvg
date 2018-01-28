// main
package main

import (
	"fmt"

	"github.com/JamesDunne/golang-openvg/host"
	"github.com/JamesDunne/golang-openvg/vg"
)

func main() {
	fmt.Println("Hello World!")
	host.Run(800, 480)
	vg.Setfv(vg.ClearColor, 4, &[]float32{1.0, 0.0, 0.0, 1.0}[0])
	vg.Clear(0, 0, 800, 480)
}
