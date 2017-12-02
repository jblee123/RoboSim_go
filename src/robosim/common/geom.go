package common

import (
	"fmt"
	"math"
)

////////////////////////////////////////////////////////////////////////////////

type Degrees float64
type Radians float64

func DegToRad(deg Degrees) Radians {
	return Radians(deg * ((2 * math.Pi) / 360))
}

func RadToDeg(rad Radians) Degrees {
	return Degrees(rad * (360 / (2 * math.Pi)))
}

////////////////////////////////////////////////////////////////////////////////

type Point struct{ X, Y float64 }

////////////////////////////////////////////////////////////////////////////////

type Vector struct{ X, Y, Z float64 }

func (v Vector) String() string {
	return fmt.Sprintf("(%.02f, %.02f, %.02f)", v.X, v.Y, v.Z)
}

func (v Vector) Dot() float64 {
	return v.X*v.X + v.Y*v.Y + v.Z*v.Z
}

func (v Vector) Length() float64 {
	return math.Sqrt(v.Dot())
}

func (v Vector) Add(v2 Vector) Vector {
	return Vector{v.X + v2.X, v.Y + v2.Y, v.Z + v2.Z}
}

func (v Vector) Sub(v2 Vector) Vector {
	return Vector{v.X - v2.X, v.Y - v2.Y, v.Z - v2.Z}
}

func (v Vector) Mult(scalar float64) Vector {
	return Vector{v.X * scalar, v.Y * scalar, v.Z * scalar}
}

func (v Vector) Div(scalar float64) Vector {
	return Vector{v.X / scalar, v.Y / scalar, v.Z / scalar}
}

func (v Vector) RotateZ(angle Radians) Vector {
	c := math.Cos(float64(angle))
	s := math.Sin(float64(angle))
	x, y := v.X, v.Y

	v.X = x*c - y*s
	v.Y = x*s + y*c

	return v
}

func (v Vector) Angle() Radians {
	return Radians(math.Atan2(v.Y, v.X))
}

func (v Vector) Unit() Vector {
	length := v.Length()
	if length == 0 {
		return Vector{1, 0, 0}
	}
	return v.Div(length)
}

////////////////////////////////////////////////////////////////////////////////

type Ray struct {
	From Vector
	To   Vector
}

func (r Ray) IntersectWithCircle(center Point, radius float64) (Point, bool) {
	x0, y0 := r.From.X, r.From.Y
	dx, dy := r.To.X-x0, r.To.Y-y0
	a := dx*dx + dy*dy
	b := 2 * (x0*dx - center.X*dx + y0*dy - center.Y*dy)
	c := (x0*x0 + center.X*center.Y - 2*x0*center.X) +
		(y0*y0 + center.Y*center.Y - 2*y0*center.Y) - radius*radius
	quot := b*b - 4*a*c
	if quot < 0 {
		return Point{0, 0}, false
	}

	sqrt_quot := math.Sqrt(quot)
	two_a := 2 * a
	t1 := (-b + sqrt_quot) / two_a
	t2 := (-b - sqrt_quot) / two_a

	var v1, v2 Vector
	t1_ok, t2_ok := t1 >= 0, t2 >= 0
	if t1_ok {
		v1 = Vector{x0 + t1*dx, y0 + t1*dy, 0}
	}
	if t2_ok {
		v2 = Vector{x0 + t2*dx, y0 + t2*dy, 0}
	}

	var intersection Vector
	if t1_ok && t2_ok {
		if v1.Sub(r.From).Dot() < v2.Sub(r.From).Dot() {
			intersection = v1
		} else {
			intersection = v2
		}
	} else if t1_ok {
		intersection = v1
	} else {
		intersection = v2
	}

	return Point{intersection.X, intersection.Y}, true
}

func (r Ray) IntersectWithSegment(p0, p1 Point) (Point, bool) {
	x0, y0 := r.From.X, r.From.Y
	dxr, dyr := r.To.X-x0, r.To.Y-y0
	dxs, dys := p1.X-p0.X, p1.Y-p0.Y

	denom := dxs*dyr - dys*dxr
	if (denom == 0) || (dxr == 0 && dyr == 0) {
		return Point{0, 0}, false
	}

	ts := dxr*(p0.Y-y0) + dyr*(x0-p0.X)
	ts /= dxs*dyr - dys*dxr

	var tr float64
	if dxr != 0 {
		tr = (p0.X + ts*dxs - x0) / dxr
	} else {
		tr = (p0.Y + ts*dys - y0) / dyr
	}

	if (ts < 0) || (ts > 1) || (tr < 0) {
		return Point{0, 0}, false
	}

	return Point{p0.X + ts*dxs, p0.Y + ts*dys}, true
}
