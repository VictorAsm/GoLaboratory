package linear

import (
	"math"
)

type Vec3 struct {
	X, Y, Z float64
}

func (v Vec3) Add(u Vec3) Vec3 {
	return Vec3{v.X + u.X, v.Y + u.Y, v.Z + u.Z}
}

func (v Vec3) Sub(u Vec3) Vec3 {
	return Vec3{v.X - u.X, v.Y - u.Y, v.Z - u.Z}
}

func (v Vec3) Dot(u Vec3) float64 {
	return v.X*u.X + v.Y*u.Y + v.Z*u.Z
}

func (v Vec3) Cross(u Vec3) Vec3 {
	return Vec3{
		v.Y*u.Z - v.Z*u.Y,
		v.X*u.Z - v.Z*u.X,
		v.X*u.Y - v.Y*u.X}
}

func (v Vec3) Magnitude() float64 {
	return math.Sqrt(v.Dot(v))
}

func (v Vec3) Normalize() Vec3 {
	mag := v.Magnitude()
	if mag == 0 {
		return Vec3{0, 0, 0}
	}
	return Vec3{v.X / mag, v.Y / mag, v.Z / mag}
}

func (v Vec3) Scale(a float64) Vec3 {
	return Vec3{a * v.X, a * v.Y, a * v.Z}
}

func (v Vec3) Lerp(u Vec3, t float64) Vec3 {
	return Vec3{
		v.X + t*(u.X-v.X),
		v.Y + t*(u.Y-v.Y),
		v.Z + t*(u.Z-v.Z),
	}
}

func (v Vec3) Reflect(n Vec3) Vec3 {
	n = n.Normalize()
	dot := v.Dot(n)
	return v.Sub(n.Scale(2 * dot))
}

func (v Vec3) LengthSquared() float64 {
	return v.X*v.X + v.Y*v.Y + v.Z*v.Z
}
