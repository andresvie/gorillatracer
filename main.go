package main

import (
	"fmt"
	"math"
	"os"

	"github.com/andresvie/gorillatracer/camera"
	"github.com/andresvie/gorillatracer/gml"
	"github.com/andresvie/gorillatracer/light"
	"github.com/andresvie/gorillatracer/utils"
	"github.com/andresvie/gorillatracer/vector"
)

var origin = &vector.Vector{X: 0.0, Y: 0.0, Z: 0.0, W: 0.0}

func main() {
	width := 800
	ratio := 16.0 / 9.0
	shouldShowUsage := len(os.Args) != 3
	if shouldShowUsage {
		fmt.Fprintf(os.Stderr, "Usage of %s test.gml output.ppm \n", os.Args[0])
		return
	}
	sceneReader, err := os.Open(os.Args[1])
	defer sceneReader.Close()
	if err != nil {
		fmt.Fprintf(os.Stderr, "please provide a valid GML file %s \nUsage of %s test.gml output.ppm \n", os.Args[1], os.Args[0])
		return
	}
	f, err := os.Create(os.Args[2])
	defer f.Close()
	if err != nil {
		fmt.Fprintf(os.Stderr, "please provide a valid PPM path file %s \nUsage of %s test.gml output.ppm \n", os.Args[2], os.Args[0])
		return
	}
	scene, err := gml.ParseGML(sceneReader)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error parsing file %s please prove a valid GML file %v \nUsage of %s test.gml output.ppm \n", os.Args[1], err, os.Args[0])
		return
	}
	camera := camera.CreateCamera(*origin, utils.REAL(width), utils.REAL(ratio), math.Pi/4)
	scene.Lights = append(scene.Lights, &light.AmbientLight{Intensity: 0.1})
	scene.Camera = camera
	scene.Render(f)
}
