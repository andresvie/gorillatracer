package vector

import "math"

type Vector struct {
	X float32
	Y float32
	Z float32
	W float32
}

func (v *Vector) Add(b *Vector) *Vector {
	return &Vector{v.X + b.X, v.Y + b.Y, v.Z + b.Z, v.W}
}

func (v *Vector) Negate() *Vector {
	return v.Scale(-1.0)
}

func (v *Vector) Scale(factor float32) *Vector {
	return &Vector{v.X * factor, v.Y * factor, v.Z * factor, v.W * factor}
}

func (v *Vector) Dot(b *Vector) float32 {
	dX := v.X * b.X
	dY := v.Y * b.Y
	dZ := v.Z * b.Z
	return dX + dY + dZ
}

func (v *Vector) Normal() *Vector {
	len := v.Length()
	return &Vector{v.X / len, v.Y / len, v.Z / len, 1.0}
}

func (v *Vector) Length() float32 {
	square := float64(v.Dot(v))
	return float32(math.Sqrt(square))
}

func (v *Vector) Cross(q *Vector) *Vector {
	i := v.Y*q.Z - v.Z*q.Y
	j := -(v.X*q.Z - v.Z*q.X)
	k := (v.X*q.Y - v.Y*q.X)
	return &Vector{i, j, k, 1.0}
}
