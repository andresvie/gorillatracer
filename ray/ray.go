package ray

import (
	"github.com/andresvie/gorillatracer/utils"
	"github.com/andresvie/gorillatracer/vector"
)

// Ray useless comment
type Ray struct {
	Origin    *vector.Vector
	Direction *vector.Vector
}

func (r *Ray) PointAt(t utils.REAL) *vector.Vector {
	scaleVector := r.Direction.Scale(t)
	newPoint := r.Origin.Add(scaleVector)
	return &vector.Vector{X: newPoint.X, Y: newPoint.Y, Z: newPoint.Z, W: 0.0}
}

func (r *Ray) View() *vector.Vector {
	return r.Direction.Negate().Normal()
}
