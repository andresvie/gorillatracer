package main

import (
	"fmt"
	"io"
	"os"

	"github.com/andresvie/gorillatracer/ray"
	"github.com/andresvie/gorillatracer/vector"
)

var origin = &vector.Vector{X: 0.0, Y: 0.0, Z: 0.0, W: 0.0}
var lowerLeftConner = &vector.Vector{X: -2.0, Y: -1.0, Z: -1.0, W: 0.0}
var horizontal = &vector.Vector{X: 4.0, Y: 0.0, Z: 0.0, W: 0.0}
var vertical = &vector.Vector{X: 0.0, Y: 2.0, Z: 0.0, W: 0.0}

func main() {
	width := 200
	height := 200

	f, err := os.Create("/Users/lordoftherootland/dat2.ppm")
	defer f.Close()
	if err != nil {
		panic(err)
	}
	fmt.Fprintln(f, "P3")
	fmt.Fprintln(f, width, " ", height)
	fmt.Fprintln(f, 255)
	for i := height; i >= 0; i-- {
		for j := 0; j < width; j++ {
			v := float32(i) / float32(height)
			u := float32(j) / float32(width)
			r := getRayFromUV(u, v)
			color := getColor(r)
			writeColor(f, color)
		}
	}
}

func getRayFromUV(u float32, v float32) *ray.Ray {
	direction := lowerLeftConner.Add(horizontal.Scale(u)).Add(vertical.Scale(v))
	return &ray.Ray{Origin: origin, Direction: direction}
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
	ir := int(color.X * 255.99)
	ig := int(color.Y * 255.99)
	ib := int(color.Z * 255.99)
	fmt.Fprintln(w, ir, ig, ib)
}
