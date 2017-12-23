package common

import (
	"math"
	"testing"
)

const (
	epsilon = 1e-10
)

func numsClose(n1, n2 float64) bool {
	return math.Abs(n1-n2) < epsilon
}

func radsClose(r1, r2 Radians) bool {
	return numsClose(float64(r1), float64(r2))
}

func vectorsClose(v1, v2 Vector) bool {
	return v1.Sub(v2).Length() < epsilon
}

func pointsClose(p1, p2 Point) bool {
	dx := p1.X - p2.X
	dy := p1.Y - p2.Y
	dist := math.Sqrt(dx*dx + dy*dy)
	return dist < epsilon
}

////////////////////////////////////////////////////////////////////////////////

const (
	deg45 = 45.0
	rad45 = deg45 * (2 * math.Pi / 360)
)

func TestDegToRad(t *testing.T) {
	if rad := DegToRad(deg45); rad != rad45 {
		t.Errorf("DegToRad(%v) = %v, want %v", deg45, rad, rad45)
	}
}

func TestRadToDeg(t *testing.T) {
	if deg := RadToDeg(rad45); deg != deg45 {
		t.Errorf("TestRadToDeg(%v) = %v, want %v", rad45, deg, deg45)
	}
}

func TestNormalizeAngle(t *testing.T) {
	tests := []struct {
		input Degrees
		want  Degrees
	}{
		{-5, 355},
		{0, 0},
		{180, 180},
		{360, 0},
		{365, 5},
		{725, 5},
	}

	for _, test := range tests {
		if got := NormalizeAngleD(test.input); got != test.want {
			t.Errorf("NormalizeAngleD(%v) = %v, want %v",
				test.input, got, test.want)
		}
	}
}

////////////////////////////////////////////////////////////////////////////////

func TestVectorString(t *testing.T) {
	tests := []struct {
		input Vector
		want  string
	}{
		{Vector{0, 0, 0}, "(0.00, 0.00, 0.00)"},
		{Vector{1, 2, -3}, "(1.00, 2.00, -3.00)"},
		{Vector{1.2, 2.45, -3.67}, "(1.20, 2.45, -3.67)"},
		{Vector{1.20, 2.451, 3.456}, "(1.20, 2.45, 3.46)"},
	}

	for _, test := range tests {
		if got := test.input.String(); got != test.want {
			t.Errorf("{%v, %v, %v}.String() = %q, want %q",
				test.input.X, test.input.Y, test.input.Z,
				got, test.want)
		}
	}
}

func TestVectorDot(t *testing.T) {
	tests := []Vector{
		Vector{2, -3, 4},
	}

	for _, test := range tests {
		got := test.Dot()
		want := test.X*test.X + test.Y*test.Y + test.Z*test.Z
		if !numsClose(got, want) {
			t.Errorf("%v.Dot() = %v, want %v", test, got, want)
		}
	}
}

func TestVectorLength(t *testing.T) {
	tests := []Vector{
		Vector{2, -3, 4},
	}

	for _, test := range tests {
		got := test.Length()
		want := math.Sqrt(test.X*test.X + test.Y*test.Y + test.Z*test.Z)
		if !numsClose(got, want) {
			t.Errorf("%v.Length() = %v, want %v", test, got, want)
		}
	}
}

func TestVectorAdd(t *testing.T) {
	type Input struct {
		vec1 Vector
		vec2 Vector
	}

	tests := []struct {
		input Input
		want  Vector
	}{
		{
			Input{
				Vector{1, 2, 3},
				Vector{4, 5, 6},
			},
			Vector{X: 5, Y: 7, Z: 9},
		},
	}

	for _, test := range tests {
		got := test.input.vec1.Add(test.input.vec2)
		if !vectorsClose(got, test.want) {
			t.Errorf("%v.Add(%v) = %v, want %v",
				test.input.vec1, test.input.vec2, got, test.want)
		}
	}
}

func TestVectorSub(t *testing.T) {
	type Input struct {
		vec1 Vector
		vec2 Vector
	}

	tests := []struct {
		input Input
		want  Vector
	}{
		{
			Input{
				Vector{4, 2, 7},
				Vector{1, 5, 6},
			},
			Vector{X: 3, Y: -3, Z: 1},
		},
	}

	for _, test := range tests {
		got := test.input.vec1.Sub(test.input.vec2)
		if !vectorsClose(got, test.want) {
			t.Errorf("%v.Sub(%v) = %v, want %v",
				test.input.vec1, test.input.vec2, got, test.want)
		}
	}
}

func TestVectorMult(t *testing.T) {
	type Input struct {
		vec    Vector
		scalar float64
	}

	tests := []struct {
		input Input
		want  Vector
	}{
		{
			Input{
				Vector{1, 2, 3},
				2,
			},
			Vector{2, 4, 6},
		},
	}

	for _, test := range tests {
		got := test.input.vec.Mult(test.input.scalar)
		if !vectorsClose(got, test.want) {
			t.Errorf("%v.Mult(%v) = %v, want %v",
				test.input.vec, test.input.scalar, got, test.want)
		}
	}
}

func TestVectorDiv(t *testing.T) {
	type Input struct {
		vec    Vector
		scalar float64
	}

	tests := []struct {
		input Input
		want  Vector
	}{
		{
			Input{
				Vector{1, 2, 3},
				2,
			},
			Vector{0.5, 1, 1.5},
		},
	}

	for _, test := range tests {
		got := test.input.vec.Div(test.input.scalar)
		if !vectorsClose(got, test.want) {
			t.Errorf("%v.Div(%v) = %v, want %v",
				test.input.vec, test.input.scalar, got, test.want)
		}
	}
}

func TestVectorRotateZ(t *testing.T) {
	type Input struct {
		vec   Vector
		angle Radians
	}

	tests := []struct {
		input Input
		want  Vector
	}{
		{
			Input{
				Vector{1, 0, 0},
				math.Pi / 2,
			},
			Vector{0, 1, 0},
		},
	}

	for _, test := range tests {
		got := test.input.vec.RotateZ(test.input.angle)
		if !vectorsClose(got, test.want) {
			t.Errorf("%v.RotateZRad(%v) = %v, want %v (delta: %v)",
				test.input.vec, test.input.angle, got, test.want)
		}
	}
}

func TestVectorAngle(t *testing.T) {
	tests := []struct {
		input Vector
		want  Radians
	}{
		{Vector{0, 0, 0}, 0},
		{Vector{1, 0, 0}, 0},
		{Vector{0, 1, 0}, math.Pi / 2},
	}

	for _, test := range tests {
		if got := test.input.Angle(); !radsClose(got, test.want) {
			t.Errorf("%v.Angle() = %v, want %v", test.input, got, test.want)
		}
	}
}

func TestVectorUnit(t *testing.T) {
	unitLenFor111 := 1.0 / math.Sqrt(3)

	tests := []struct {
		input Vector
		want  Vector
	}{
		{Vector{0, 0, 0}, Vector{1, 0, 0}},
		{Vector{1, 0, 0}, Vector{1, 0, 0}},
		{Vector{6, 0, 0}, Vector{1, 0, 0}},
		{Vector{1, 1, 1}, Vector{unitLenFor111, unitLenFor111, unitLenFor111}},
		{Vector{9, 9, 9}, Vector{unitLenFor111, unitLenFor111, unitLenFor111}},
	}

	for _, test := range tests {
		if got := test.input.Unit(); !vectorsClose(got, test.want) {
			t.Errorf("%v.Unit() = %v, want %v", test.input, got, test.want)
		}
	}
}

////////////////////////////////////////////////////////////////////////////////

func TestRayIntersectWithCircle(t *testing.T) {
	type Circle struct {
		pos    Point
		radius float64
	}

	type Input struct {
		r Ray
		c Circle
	}

	type Result struct {
		p  Point
		ok bool
	}

	unitLenFor11 := 1.0 / math.Sqrt(2)

	tests := []struct {
		input Input
		want  Result
	}{
		{
			Input{
				Ray{
					Vector{0, 0, 0},
					Vector{2, 2, 2},
				},
				Circle{Point{0, 0}, 1},
			},
			Result{Point{unitLenFor11, unitLenFor11}, true},
		},
		{
			Input{
				Ray{
					Vector{0, 0, 0},
					Vector{2, 2, 2},
				},
				Circle{Point{0, 0}, 3},
			},
			Result{Point{3 * unitLenFor11, 3 * unitLenFor11}, true},
		},
		{
			Input{
				Ray{
					Vector{-3, -3, 0},
					Vector{-2, -2, 0},
				},
				Circle{Point{0, 0}, 1},
			},
			Result{Point{-unitLenFor11, -unitLenFor11}, true},
		},
		{
			Input{
				Ray{
					Vector{4, -6, 0},
					Vector{4, -5, 0},
				},
				Circle{Point{1, 0}, 1},
			},
			Result{Point{0, 0}, false},
		},
	}

	for _, test := range tests {
		got_p, got_ok := test.input.r.IntersectWithCircle(
			test.input.c.pos, test.input.c.radius)
		if (got_ok != test.want.ok) || !pointsClose(got_p, test.want.p) {
			t.Errorf("%v.IntersectWithCircle(%v, %v) = (%v, %v), want (%v, %v)",
				test.input.r, test.input.c.pos, test.input.c.radius,
				got_p, got_ok, test.want.p, test.want.ok)
		}
	}
}

func TestRayIntersectWithSegment(t *testing.T) {
	type Segment struct {
		p1, p2 Point
	}

	type Input struct {
		r   Ray
		seg Segment
	}

	type Result struct {
		p  Point
		ok bool
	}

	//unitLenFor11 := 1.0 / math.Sqrt(2)

	tests := []struct {
		input Input
		want  Result
	}{
		{
			Input{
				Ray{
					Vector{0, 0, 0},
					Vector{1, 1, 0},
				},
				Segment{
					Point{1, 3},
					Point{3, 1},
				},
			},
			Result{Point{2, 2}, true},
		},
		{
			Input{
				Ray{
					Vector{0, 0, 0},
					Vector{1, 1, 0},
				},
				Segment{
					Point{0, 4},
					Point{1, 3},
				},
			},
			Result{Point{0, 0}, false},
		},
	}

	for _, test := range tests {
		got_p, got_ok := test.input.r.IntersectWithSegment(
			test.input.seg.p1, test.input.seg.p2)
		if (got_ok != test.want.ok) || !pointsClose(got_p, test.want.p) {
			t.Errorf("%v.IntersectWithSegment(%v, %v) = (%v, %v), want (%v, %v)",
				test.input.r, test.input.seg.p1, test.input.seg.p2,
				got_p, got_ok, test.want.p, test.want.ok)
		}
	}
}
