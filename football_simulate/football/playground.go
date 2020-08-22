package football

import (
	"football_simulate/geomtry/solid"
)

type Playground struct {
	Range            Square         //球场范围
	R                Square         //球门矩形框
	B                Square         //球门矩形框
	HalfwayLine      []solid.Vector //中线
	RPenaltyShotZone Square         //点球罚球区
	BPenaltyShotZone Square         //点球罚球区
	RPenaltyArc      Circle         //
	BPenaltyArc      Circle
}

type Square struct {
	A solid.Vector
	B solid.Vector
	C solid.Vector
	D solid.Vector
}

//圆上的3点，可以确定一个圆
type Circle struct {
	A solid.Vector
	B solid.Vector
	C solid.Vector
}

func NewPlayground() Playground {

	return Playground{}
}
