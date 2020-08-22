package football

import (
	"errors"
	"fmt"
	"football_simulate/geomtry/solid"
)

type MatchStatus int

const (
	MatchStop MatchStatus = iota + 1
	MatchIng
)

//
type Match struct {
	Teams         map[string]Team    //队伍
	Where         string             //在哪里举办
	Name          string             //名字
	BeginTime     int64              //开始时间
	EndTime       int64              //结束时间
	Winner        string             //获胜队伍
	Looser        string             //失败队伍
	Referee       map[string]Referee //裁判
	CurrentStatus MatchStatus        //当前状态
	playground    Playground         //场地
	Ball          Ball               //球
	MatchChannel  chan MatchChannel  //新闻消息

}

func NewMatch() Match {
	return Match{}
}

func (self *Match) Run() {
	//todo 收到新闻消息
	//todo
	for {
		select {
		case ch, opened := <-self.MatchChannel:
			if opened {
				self.DealWithChannel(ch)
			}
		}
	}

}

func (self *Match) Stop() {

}

func (self *Match) Continue() {

}

func (self *Match) Begin() {

}

func (self *Match) End() {

}

func (self *Match) Save() {

}

func (self *Match) Upgrade() {

	//todo 更新match 状态数据
	//todo 更新match 未来状态数据

}

func (self *Match) DealWithChannel(Mc MatchChannel) {

	switch Mc.Category {

	case MatchNormalC:

	case MatchEndC:

	case MatchStopC:

	case MatchBeginC:

	default:

		fmt.Println("type err")

	}

}

func (self* Match)GetPlayerNextLocation(player Player,newsArr []MatchNewsTemplate) solid.Vector  {

	old:=newsArr[0]
	current:=newsArr[1]
	k:=current.Ball.Location.GetDistance(old.Ball.Location)/2.0
	camparePoint:=old.Ball.Location.ToSpot().Move(current.Ball.Location).ToVector().Unit().Mul(k)
	if camparePoint.GetDistance(player.Location) < 10 {
		if current.HolderName!=player.Name {
			return player.Location.ToSpot().Move(camparePoint).ToVector().Unit().Mul(k)
		}
	}
	return player.Location.RandomAdd()
}

//
func (self *Match) GetToClientInfo(newsArr []MatchNewsTemplate) (*MatchNewsToClientInfo, error) {
	if newsArr == nil {
		return nil, errors.New("array is nil  ")
	}
	Len := len(newsArr)
	if Len < 2 || Len > 2 {
		return nil, errors.New("length error ")
	}

	news1 := newsArr[0]
	news2 := newsArr[1]

	mNewInfo := MatchNewsToClientInfo{}
	mNewInfo.Players = make(map[string]PlayerToClient)
	for _, team := range self.Teams {
		for memberName, member := range team.Members {
			mNewInfo.Players[memberName] = PlayerToClient{member.Location,self.GetPlayerNextLocation(member,newsArr),memberName}
		}
	}

	mNewInfo.Ball = BallToClient{
		news1.Ball.Location,
		news2.Ball.Location,
	}
	return &mNewInfo, nil
}

type PlayerToClient struct {
	Current solid.Vector
	Next    solid.Vector
	Name    string
}
type BallToClient struct {
	Current solid.Vector
	Next    solid.Vector
}

//
type MatchNewsToClientInfo struct {
	Ball    BallToClient              //球
	Players map[string]PlayerToClient //
}
