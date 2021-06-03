package scene

import (
	"fmt"
	"io"
	"math"

	"github.com/andresvie/gorillatracer/camera"
	"github.com/andresvie/gorillatracer/geometry"
	"github.com/andresvie/gorillatracer/light"
	"github.com/andresvie/gorillatracer/ray"
	"github.com/andresvie/gorillatracer/utils"
	"github.com/andresvie/gorillatracer/vector"
)

type Scene struct {
	Camera  *camera.Camera
	Lights  []light.Light
	Objects []geometry.Geometry
}

func (s *Scene) Render(w io.Writer) {
	camera := s.Camera
	fmt.Fprintln(w, "P3")
	fmt.Fprintln(w, int(camera.Width), " ", int(camera.Height))
	fmt.Fprintln(w, 255)
	infinity := utils.REAL(math.Inf(1))
	for i := 0; i < int(camera.Height); i++ {
		for j := 0; j < int(camera.Width); j++ {
			r := camera.CalculatePixelRay(j, i)
			hit, hitObject := s.closestObject(r, infinity)
			color := getBackGrounColor(r)
			intensity := utils.REAL(1)
			if hit.Collide {
				color = hitObject.GetColor()
				intensity = light.IntegrateLight(s.Lights, s.Objects, &hit)
			}
			color = color.Scale(intensity)
			writeColor(w, color)
		}
	}
}

func getBackGrounColor(r *ray.Ray) *vector.Vector {
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
}

func (s *Scene) closestObject(r *ray.Ray, min utils.REAL) (geometry.Hit, geometry.Geometry) {
	tMin := min
	var hit geometry.Hit
	var hitObject geometry.Geometry
	for _, object := range s.Objects {
		newHit := object.InterceptRay(r, tMin, 0.0)
		if newHit.Collide && newHit.Interval < tMin {
			hit = newHit
			tMin = hit.Interval
			hitObject = object
		}
	}
	return hit, hitObject
}
