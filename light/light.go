package light

import (
	"fmt"

	"github.com/andresvie/gorillatracer/geometry"
	"github.com/andresvie/gorillatracer/ray"
	"github.com/andresvie/gorillatracer/utils"
)

type Light interface {
	CalculateIntensity(hit *geometry.Hit) utils.REAL
	CreateShadowRay(hit *geometry.Hit) (*ray.Ray, utils.REAL)
}

func IntegrateLight(lights []Light, objects []geometry.Geometry, hit *geometry.Hit) utils.REAL {
	intensity := utils.REAL(0.0)
	for _, light := range lights {
		r, maxValue := light.CreateShadowRay(hit)
		if r == nil {
			intensity += light.CalculateIntensity(hit)
			continue
		}
		if isPixelInTheShadow(r, maxValue, objects, hit) {
			continue
		}
		intensity += light.CalculateIntensity(hit)
	}
	return utils.Clamp(intensity, 0, 1)
}

func isPixelInTheShadow(r *ray.Ray, maxInterval utils.REAL, objects []geometry.Geometry, objectHit *geometry.Hit) bool {
	for _, object := range objects {
		hit := object.InterceptRay(r, maxInterval, 0.001)
		if object == objectHit.Object {
			continue
		}
		if hit.Collide {
			fmt.Printf("currentObjectColor(%v) interval(%v) ShadowTestingColor(%v) ray origin(%v) ray direction(%v)\n", objectHit.Object.GetColor(), hit.Interval, object.GetColor(), r.Origin, r.Direction)
			return true
		}
	}
	return false
}
