package light

import (
	"math"

	"github.com/andresvie/gorillatracer/geometry"
	"github.com/andresvie/gorillatracer/ray"
	"github.com/andresvie/gorillatracer/utils"
	"github.com/andresvie/gorillatracer/vector"
)

type DirectionalLight struct {
	Intensity utils.REAL
	Direction vector.Vector
}

func (d *DirectionalLight) CreateShadowRay(hit *geometry.Hit) (*ray.Ray, utils.REAL) {
	r := &ray.Ray{Origin: hit.InterceptionPoint, Direction: d.Direction.Negate()}
	return r, utils.REAL(math.Inf(1))
}

func (d *DirectionalLight) CalculateIntensity(hit *geometry.Hit) utils.REAL {
	direction := d.Direction
	diffuseIntensity := utils.Clamp(direction.Dot(hit.Normal), 0, 1)
	specularIntensity := utils.Clamp(direction.Reflect(hit.Normal).Dot(hit.View), 0, 1)
	specularIntensity = utils.REAL(math.Pow(float64(specularIntensity), float64(hit.Specular)))
	return (diffuseIntensity + specularIntensity) * d.Intensity
}
