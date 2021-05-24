package light

import (
	"github.com/andresvie/gorillatracer/geometry"
	"github.com/andresvie/gorillatracer/utils"
)

type AmbientLight struct {
	Intensity utils.REAL
}

func (a *AmbientLight) CalculateIntensity(_ *geometry.Hit) utils.REAL {
	return a.Intensity
}
