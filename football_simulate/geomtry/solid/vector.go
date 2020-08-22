package solid

import (
	"math"
	"math/rand"
)

type Vector struct {
	X, Y, Z float64
}

func (v Vector) RandomAdd()Vector {

	return 	Vector{
		Y: v.Y,
		X: 	float64(v.X)*rand.Float64(),
		Z: 	float64(v.Z)*rand.Float64(),
	}

}
// 返回向量a->b
func AimTo(a, b Spot) Vector {
	return Vector{b.I - a.I, b.J - a.J, b.K - a.K}
}

func (v Vector) ToSpot() Spot {
	return Spot{
		v.X, v.Y, v.Z,
	}
}

// 返回向量的模。
func (v Vector) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y + v.Z*v.Z)
}

// 返回同向的单位向量；如果向量的模为零，返回该向量。
func (v Vector) Unit() Vector {
	r := v.Abs()
	return Vector{v.X / r, v.Y / r, v.Z / r}
}

// 返回两个向量的和
func (v Vector) Add(u Vector) Vector {
	return Vector{v.X + u.X, v.Y + u.Y, v.Z + u.Z}
}

// 返回两个向量的差
func (v Vector) Dec(u Vector) Vector {
	return Vector{v.X - u.X, v.Y - u.Y, v.Z - u.Z}
}

// 返回系数和向量的乘积
func (v Vector) Mul(k float64) Vector {
	return Vector{v.X * k, v.Y * k, v.Z * k}
}

//转单位向量
func (self Vector) ToIdentityVector() Vector {
	distance := math.Sqrt(math.Pow(self.X, 2) + math.Pow(self.Y, 2) + math.Pow(self.Z, 2))
	identityVector := Vector{self.X / distance, self.Y / distance, self.Z / distance}
	return identityVector
}

//转单位向量忽略y轴
func (self Vector) ToIdentityVectorIgnoreY() Vector {
	distance := math.Sqrt(math.Pow(self.X, 2) + 0 + math.Pow(self.Z, 2))
	identityVector := Vector{self.X / distance, 0, self.Z / distance}
	return identityVector
}

//朝向
func (self Vector) LookAt(target Vector, ignoreY bool) Vector {
	direction := Vector{self.X - target.X, self.Y - target.Y, self.Z - target.Z}
	if ignoreY {
		return direction.ToIdentityVectorIgnoreY()
	}
	return direction.ToIdentityVector()
}

//获取两点间的距离
func (self *Vector) GetDistance(target Vector) float64 {
	tmpVector := Vector{self.X - target.X, self.Y - target.Y, self.Z - target.Z}
	distance := math.Sqrt(math.Pow(tmpVector.X, 2) + math.Pow(tmpVector.Y, 2) + math.Pow(tmpVector.Z, 2))
	return distance
}

//获取点到直线的距离
func (self *Vector) GetLineDistance(line []Vector) float64 {
	lineEdge := Vector{line[1].X - line[0].X, line[1].Y - line[0].Y, line[1].Z - line[0].Z}
	point2EdgePoint := Vector{self.X - line[0].X, self.Y - line[0].Y, self.Z - line[0].Z}
	distanceLineEdge := line[0].GetDistance(line[1])
	distancePoint2EdgePoint := line[0].GetDistance(*self)

	lineEdgeXpoint2EdgePoint := lineEdge.X*point2EdgePoint.X + lineEdge.Y*point2EdgePoint.Y + lineEdge.Z*point2EdgePoint.Z

	Cos := lineEdgeXpoint2EdgePoint / (distanceLineEdge * distancePoint2EdgePoint)
	ShadowDistance := distancePoint2EdgePoint * Cos
	normalized := point2EdgePoint.ToIdentityVector()
	shadow := Vector{ShadowDistance * normalized.X, ShadowDistance * normalized.Y, ShadowDistance * normalized.Z}
	distance := shadow.GetDistance(*self)
	return distance
}

// 返回两个向量的内积
func InnerProduct(p, q Vector) float64 {
	return p.X*q.X + p.Y*q.Y + p.Z*q.Z
}

// 返回两个向量的外积
func OuterProduct(p, q Vector) Vector {
	return Vector{p.Y*q.Z - q.Y*p.Z, q.X*p.Z - p.X*q.Z, p.X*q.Y - q.X*p.Y}
}
