package utils

import (
	"github.com/g3n/engine/math32"
	"math"
)

func SetAxisAngle(q *math32.Quaternion, axis *math32.Vector3, angle float32) {
	halfAngle := angle / 2
	s := float32(math.Sin(float64(halfAngle)))
	q.X = axis.X * s
	q.Y = axis.Y * s
	q.Z = axis.Z * s
	q.W = float32(math.Cos(float64(halfAngle)))
}

func MultiplyQuaternions(a *math32.Quaternion, b *math32.Quaternion) *math32.Quaternion {

	q := math32.Quaternion{}

	qax := a.X
	qay := a.Y
	qaz := a.Z
	qaw := a.W
	qbx := b.X
	qby := b.Y
	qbz := b.Z
	qbw := b.W

	q.X = qax*qbw + qaw*qbx + qay*qbz - qaz*qby
	q.Y = qay*qbw + qaw*qby + qaz*qbx - qax*qbz
	q.Z = qaz*qbw + qaw*qbz + qax*qby - qay*qbx
	q.W = qaw*qbw - qax*qbx - qay*qby - qaz*qbz

	return &q
}

func ApplyQuaternionToVector(q *math32.Quaternion, v *math32.Vector3) {

	x := v.X
	y := v.Y
	z := v.Z

	qx := q.X
	qy := q.Y
	qz := q.Z
	qw := q.W

	// calculate quat * vector
	ix := qw*x + qy*z - qz*y
	iy := qw*y + qz*x - qx*z
	iz := qw*z + qx*y - qy*x
	iw := -qx*x - qy*y - qz*z
	// calculate result * inverse quat
	v.X = ix*qw + iw*-qx + iy*-qz - iz*-qy
	v.Y = iy*qw + iw*-qy + iz*-qx - ix*-qz
	v.Z = iz*qw + iw*-qz + ix*-qy - iy*-qx
}

func MultiplyScalarVector(s float32, v *math32.Vector3) {
	v.X *= s
	v.Y *= s
	v.Z *= s
}

func DivideScalarVector(scalar float32, v *math32.Vector3) {

	if scalar != 0 {
		invScalar := 1 / scalar
		v.X *= invScalar
		v.Y *= invScalar
		v.Z *= invScalar
	} else {
		v.X = 0
		v.Y = 0
		v.Z = 0
	}
}

func NormalizeVector(v *math32.Vector3) {
	DivideScalarVector(v.Length(), v)
}

func NegateVector(v *math32.Vector3) {
	v.X = -v.X
	v.Y = -v.Y
	v.Z = -v.Z
}
