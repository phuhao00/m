package football

import (
	"football_simulate/geomtry/solid"
)

type Position int

const (
	CentreForward            Position = iota + 1 //中锋
	Striker                                      //前锋
	LeftWinger                                   //左边锋
	RightWinger                                  //右边锋
	LeftAttackMidfielder                         //左进攻中场
	RightAttackMidfielder                        //右进攻中场
	LeftMidfielder                               //左中场
	RightMidfielder                              //右中场
	LeftDefensiveMidfielder                      //左防守中场
	RightDefensiveMidfielder                     //右防守中场
	CenterBack                                   //中后卫
	Sweeper                                      //清道夫
	Fullback                                     //边后卫
	Wingback                                     //翼卫
	DefensiveMidfielder                          //防守中场或后腰
	CentralMidfielder                            //正中场或中前卫
	AttackingMidfielder                          //前腰，或称为攻击型中场
	Goalkeeper                                   //守门员
)

type Player struct {
	Position      Position //职位
	Number        int      //编号
	Name          string   //名字
	Age           int      //年龄
	SelfTeam      *Team    //球队
	EnemyTeam     *Team
	RedCardNum    int          //红牌数
	YellowCardNum int          //黄牌数
	Location      solid.Vector //坐标
	Mark          int          //信号标记
}

//
func NewPlayer(position Position, number int, location solid.Vector) Player {

	return Player{
		Position: position,
		Number:   number,
		Location: location,
	}
}

func (self *Player) LookAt(vector solid.Vector) solid.Vector {
	return self.Location.LookAt(vector, true)
}

//
func (self *Player) GetRangeTeamMembers(Radius float64) {

}

//获取最近的球员
func (self *Player) GetNearestMember(team Team) Player {
	min := 999999.0
	var nearestTeamMember Player
	for _, player := range team.Members {
		distance := self.Location.GetDistance(player.Location)
		if min > distance {
			min = distance
			nearestTeamMember = player
		}
	}
	return nearestTeamMember
}

//是否越位
func (self *Player) IsOffside() (Offside bool) {
	self2LineDistance := self.Location.GetLineDistance([]solid.Vector{self.EnemyTeam.Square.A, self.EnemyTeam.Square.B})
	for _, player := range self.EnemyTeam.Members {
		if self2LineDistance < player.Location.GetLineDistance([]solid.Vector{self.SelfTeam.Square.A, self.SelfTeam.Square.B}) {
			return true
		}
	}
	return false
}

//
func (self *Player) Run() {

	for {

		select {}
	}
}
