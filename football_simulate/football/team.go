package football

type Team struct {
	MemberNum     int               //成员数量
	Members       map[string]Player //成员
	Coach         string            //教练
	Goalkeeper    Player            //守门员
	HoldingMember Player            //持球
	Square        Square            //自己方场地
	Score         int
}

func NewTeam() Team {

	return Team{}
}

func (self *Team) AddMember(player Player) {
	if self.Members == nil {
		self.Members = make(map[string]Player)
	}
	self.Members[player.Name] = player
}

func (self *Team) DelMember(playerName string) {
	delete(self.Members, playerName)
}

func (self *Team) SetGoalkeeper(player Player) {
	self.Goalkeeper = player
}

func (self *Team) SetSquare(square Square) {
	self.Square = square
}

func (self *Team) SetCoach(coachName string) {
	self.Coach = coachName
}
