package vector

import (
	"math"

	"github.com/andresvie/gorillatracer/utils"
)

type Vector struct {
	X utils.REAL
	Y utils.REAL
	Z utils.REAL
	W utils.REAL
}

func (v *Vector) Reflect(n *Vector) *Vector {
	p := n.Scale(2 * v.Dot(n))
	return p.Add(v.Negate()).Normal()
}

func CreateColor(red, green, blue utils.REAL) *Vector {
	return &Vector{X: red, Y: green, Z: blue}
}
func (v *Vector) Add(b *Vector) *Vector {
	return &Vector{v.X + b.X, v.Y + b.Y, v.Z + b.Z, v.W}
}

func (v *Vector) Negate() *Vector {
	return v.Scale(-1.0)
}

func (v *Vector) Scale(factor utils.REAL) *Vector {
	return &Vector{v.X * factor, v.Y * factor, v.Z * factor, v.W}
}

func (v *Vector) Dot(b *Vector) utils.REAL {
	dX := v.X * b.X
	dY := v.Y * b.Y
	dZ := v.Z * b.Z
	return dX + dY + dZ
}

func (v *Vector) Normal() *Vector {
	len := utils.REAL(v.Length())
	return &Vector{v.X / len, v.Y / len, v.Z / len, 1.0}
}

func (v *Vector) Length() utils.REAL {
	square := float64(v.Dot(v))
	return utils.REAL(math.Sqrt(square))
}

func (v *Vector) Cross(q *Vector) *Vector {
	i := v.Y*q.Z - v.Z*q.Y
	j := -(v.X*q.Z - v.Z*q.X)
	k := (v.X*q.Y - v.Y*q.X)
	return &Vector{i, j, k, 1.0}
}
