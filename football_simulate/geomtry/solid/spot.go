package solid

// 表示一个点。
type Spot struct {
	I, J, K float64
}

// 移动一个点
func (p Spot) Move(v Vector) Spot {
	return Spot{p.I + v.X, p.J + v.Y, p.K + v.Z}
}

func (p Spot) ToVector() Vector {
	return Vector{
		p.I,
		p.J,
		p.K,
	}
}
