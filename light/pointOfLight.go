package light

import (
	"math"

	"github.com/andresvie/gorillatracer/geometry"
	"github.com/andresvie/gorillatracer/ray"
	"github.com/andresvie/gorillatracer/utils"
	"github.com/andresvie/gorillatracer/vector"
)

type PointOfLight struct {
	Intensity utils.REAL
	Point     *vector.Vector
}

func (p *PointOfLight) CreateShadowRay(hit *geometry.Hit) (*ray.Ray, utils.REAL) {
	r := &ray.Ray{Origin: hit.InterceptionPoint, Direction: p.Point.Add(hit.InterceptionPoint.Negate())}
	return r, 1.0
}

func (p *PointOfLight) CalculateIntensity(hit *geometry.Hit) utils.REAL {
	direction := p.Point.Add(hit.InterceptionPoint.Negate()).Normal()
	diffuseIntensity := utils.Clamp(direction.Dot(hit.Normal), 0, 1)
	specularIntensity := utils.Clamp(direction.Reflect(hit.Normal).Dot(hit.View), 0, 1)
	specularIntensity = utils.REAL(math.Pow(float64(specularIntensity), float64(hit.Specular)))
	intensity := (diffuseIntensity + specularIntensity) * p.Intensity
	return intensity
}
