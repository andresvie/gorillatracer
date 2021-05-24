package light

import (
	"github.com/andresvie/gorillatracer/geometry"
	"github.com/andresvie/gorillatracer/utils"
	"github.com/andresvie/gorillatracer/vector"
)

type PointOfLight struct {
	Intensity utils.REAL
	Point     *vector.Vector
}

func (p *PointOfLight) CalculateIntensity(hit *geometry.Hit) utils.REAL {
	direction := p.Point.Add(hit.InterceptionPoint.Negate()).Normal()
	intensity := direction.Dot(hit.Normal)
	return utils.Clamp(intensity, 0, 1) * intensity
}
