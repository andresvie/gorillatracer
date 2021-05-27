package geometry

import (
	"math"

	"github.com/andresvie/gorillatracer/ray"
	"github.com/andresvie/gorillatracer/utils"
	"github.com/andresvie/gorillatracer/vector"
)

type Sphere struct {
	Radius         utils.REAL
	Center         *vector.Vector
	Color          *vector.Vector
	SpecularFactor utils.REAL
}

func (s *Sphere) InterceptRay(r *ray.Ray, depth utils.REAL, depthTesting bool) Hit {
	hit := Hit{Collide: false}
	dir := r.Direction.Normal()
	rsq := s.Radius * s.Radius
	co := r.Origin.Add(s.Center.Negate())
	a := dir.Dot(dir)
	b := co.Dot(dir) * 2.0
	c := co.Dot(co) - rsq
	hit.View = r.View()
	discriminant := float64(b*b - 4*a*c)
	if discriminant < 0.0 {
		return hit
	}
	hit.Specular = s.SpecularFactor
	t1 := (-b + utils.REAL(math.Sqrt(discriminant))) / (2 * a)
	t2 := (-b - utils.REAL(math.Sqrt(discriminant))) / (2 * a)
	t := utils.REAL(math.Min(float64(t1), float64(t2)))
	if depthTesting && t < depth {
		hit.Collide = true
		hit.InterceptionPoint = r.PointAt(t)
		hit.Normal = hit.InterceptionPoint.Add(s.Center.Negate()).Normal()
		return hit
	}
	return hit
}
