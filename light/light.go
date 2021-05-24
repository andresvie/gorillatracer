package light

import (
	"github.com/andresvie/gorillatracer/geometry"
	"github.com/andresvie/gorillatracer/utils"
)

type Light interface {
	CalculateIntensity(hit *geometry.Hit) utils.REAL
}

func IntegrateLight(lights []Light, hit *geometry.Hit) utils.REAL {
	intensity := utils.REAL(0.0)
	for _, light := range lights {
		intensity += light.CalculateIntensity(hit)
	}
	return utils.Clamp(intensity, 0, 1)
}
