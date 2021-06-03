package light

import (
	"github.com/andresvie/gorillatracer/geometry"
	"github.com/andresvie/gorillatracer/ray"
	"github.com/andresvie/gorillatracer/utils"
)

type AmbientLight struct {
	Intensity utils.REAL
}

func (_ *AmbientLight) CreateShadowRay(_ *geometry.Hit) (*ray.Ray, utils.REAL) {
	return nil, 1.0
}
func (a *AmbientLight) CalculateIntensity(_ *geometry.Hit) utils.REAL {
	return a.Intensity
}
