package test

import (
	"context"
	"fmt"
	"football_simulate/football"
	"reflect"
	"testing"
)

func TestT(t *testing.T) {
	x := football.EnemyTeamBreakTheRules
	fmt.Printf("%b\n", x)
	fmt.Println(reflect.TypeOf(x))
	y := 1 ^ x
	fmt.Printf("%b\n", y)
}

type CC struct {
	C context.Context
}

func TestContext(t *testing.T) {
	CC := CC{}
	CC.C = context.Background()
	fmt.Println(&CC.C)
	VTexx(&CC)
	fmt.Println(CC.C)
	val := CC.C.Value("hey")
	fmt.Println(val)
}

func VTexx(cc *CC) {

	cc.C = context.WithValue(cc.C, "hey", "val")

	fmt.Println(&cc.C)

}

func VV() {

}
