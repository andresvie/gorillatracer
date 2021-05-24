package light

import (
	"github.com/andresvie/gorillatracer/geometry"
	"github.com/andresvie/gorillatracer/utils"
	"github.com/andresvie/gorillatracer/vector"
)

type DirectionalLight struct {
	Intensity utils.REAL
	Direction vector.Vector
}

func (d *DirectionalLight) CalculateIntensity(hit *geometry.Hit) utils.REAL {
	intensity := d.Direction.Normal().Dot(hit.Normal)
	return utils.Clamp(intensity, 0, 1) * intensity
}
