// main
package main

import (
	"fmt"

	vgf "github.com/JamesDunne/golang-openvg"
	vg "github.com/JamesDunne/golang-openvg/vg"
)

func main() {
	fmt.Println("Hello World!")
	vgf.Run(800, 480)
	vg.Setfv(vg.ClearColor, 4, &[]float32{1.0, 0.0, 0.0, 1.0}[0])
	vg.Clear(0, 0, 800, 480)
}
