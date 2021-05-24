package geometry

import (
	"github.com/andresvie/gorillatracer/ray"
	"github.com/andresvie/gorillatracer/utils"
	"github.com/andresvie/gorillatracer/vector"
)

type Hit struct {
	Collide           bool
	Interval          utils.REAL
	InterceptionPoint *vector.Vector
	Normal            *vector.Vector
}
type Geometry interface {
	InterceptRay(r *ray.Ray, depth utils.REAL, depthTesting bool) Hit
}
