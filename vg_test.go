package vg

import "testing"

func TestBasicLoop(t *testing.T) {
	vg, err := Create(800, 480)
	if err != nil {
		t.Fatal(err)
	}

	for i := 0; i < 60; i++ {
		vg.SwapBuffers()
	}
}
