package main

import (
	"flag"
	"fmt"
	"io"
	"math"
	"os"

	"github.com/andresvie/gorillatracer/camera"
	"github.com/andresvie/gorillatracer/geometry"
	"github.com/andresvie/gorillatracer/light"
	"github.com/andresvie/gorillatracer/ray"
	"github.com/andresvie/gorillatracer/utils"
	"github.com/andresvie/gorillatracer/vector"
)

var origin = &vector.Vector{X: 0.0, Y: 0.0, Z: 0.0, W: 0.0}

func main() {
	width := 200
	ratio := 16.0 / 9.0
	camera := camera.CreateCamera(*origin, utils.REAL(width), utils.REAL(ratio), math.Pi/4)
	sphere := &geometry.Sphere{Color: vector.CreateColor(1.0, 0.0, 0.0), Radius: 0.2, Center: &vector.Vector{X: 0.0, Y: 0.0, Z: 2.0, W: 1}}
	pointOfLight := &light.PointOfLight{Point: origin, Intensity: 0.2}
	flag.Parse()
	fmt.Printf("file name %v\n", flag.Arg(0))
	f, err := os.Create(flag.Arg(0))
	defer f.Close()
	if err != nil {
		panic(err)
	}
	fmt.Fprintln(f, "P3")
	fmt.Fprintln(f, int(camera.Width), " ", int(camera.Height))
	fmt.Fprintln(f, 255)
	for i := int(camera.Height); i >= 0; i-- {
		for j := 0; j < int(camera.Width); j++ {
			r := camera.CalculatePixelRay(j, i)
			hit := sphere.InterceptRay(r, utils.REAL(math.Inf(1)), true)
			color := sphere.Color
			intensity := utils.REAL(1)
			if !hit.Collide {
				color = getColor(r)
			} else {
				intensity = pointOfLight.CalculateIntensity(&hit)
			}
			color = color.Scale(intensity)
			writeColor(f, color)
		}
	}
}

func getColor(r *ray.Ray) *vector.Vector {
	dir := r.Direction.Normal()
	white := &vector.Vector{X: 1.0, Y: 1.0, Z: 1.0, W: 0.0}
	lightBlue := &vector.Vector{X: 0.5, Y: 0.7, Z: 1.0, W: 0.0}
	t := 0.5 * (dir.Y + 1.0)
	whiteColor := white.Scale(1.0 - t)
	lightBlueColor := lightBlue.Scale(t)
	return whiteColor.Add(lightBlueColor)
}

func writeColor(w io.Writer, color *vector.Vector) {
	ir := math.Abs(float64(color.X * 255.999))
	ig := math.Abs(float64(color.Y * 255.999))
	ib := math.Abs(float64(color.Z * 255.999))
	fmt.Fprintln(w, int(ir), int(ig), int(ib))
	fmt.Printf("%d %d %d\n", int(ir), int(ig), int(ib))
}
