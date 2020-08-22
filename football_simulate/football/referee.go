package football

import "football_simulate/geomtry/solid"

//裁判
type Referee struct {
	Name        string       //名字
	Age         int          //年龄
	Nationality string       //国籍
	Rank        int          //职位
	Location    solid.Vector //当前位置
}
