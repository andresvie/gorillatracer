package camera

import (
	"math"

	"github.com/andresvie/gorillatracer/ray"
	"github.com/andresvie/gorillatracer/utils"
	"github.com/andresvie/gorillatracer/vector"
)

type Camera struct {
	Origin     vector.Vector
	Width      utils.REAL
	Height     utils.REAL
	ImageRatio utils.REAL
	Fov        utils.REAL
}

func CreateCamera(origin vector.Vector, width, imageRatio, fovRadians utils.REAL) *Camera {
	camera := &Camera{Origin: origin}
	camera.Width = width
	camera.Height = width / imageRatio
	camera.ImageRatio = imageRatio
	camera.Fov = utils.REAL(math.Tan(float64(fovRadians) / 2.0))
	return camera
}

func (c *Camera) CalculatePixelRay(x, y int) *ray.Ray {
	u := (utils.REAL(x) + 0.5) / c.Width
	v := (utils.REAL(y) + 0.5) / c.Height
	ndx := (2 * u) - 1.0
	ndy := 1 - (2 * v)
	px := ndx * c.ImageRatio * c.Fov
	py := ndy * c.Fov
	direction := vector.Vector{X: px, Y: py, Z: -1.0, W: 1.0}
	return &ray.Ray{Origin: &c.Origin, Direction: &direction}
}
