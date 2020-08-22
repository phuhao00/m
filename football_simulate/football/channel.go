package football

type MatchChannelType int

const (
	MatchBeginC  MatchChannelType = iota + 1 //开始
	MatchStopC                               //停止
	MatchEndC                                //结束
	MatchNormalC                             //正常消息
)

type MatchChannel struct {
	Category MatchChannelType
	Data     interface{}
}

func NewMatchChannel(Category MatchChannelType, data interface{}) MatchChannel {
	return MatchChannel{
		Category: Category,
		Data:     data,
	}
}

const (
	Min                    = iota
	EnemyTeamHold          = 1 << 1
	SelfTeamHold           = 1 << 2
	SelfTeamScore          = 1 << 3
	EnemyTeamScore         = 1 << 4
	SelfTeamPass           = 1 << 5
	EnemyTeamPass          = 1 << 6
	SelfTeamPointSphere    = 1 << 7
	EnemyTeamPointSphere   = 1 << 8
	SelfTeamCornerKick     = 1 << 9
	EnemyTeamCornerKick    = 1 << 10
	SelfTeamBreakTheRules  = 1 << 11
	EnemyTeamBreakTheRules = 1 << 12
)

//BeginChannel           chan struct{} //开始
//	StopChannel            chan struct{} //停止
//	EndChannel             chan struct{} //结束
//	SelfTeamHoldChannel    chan struct{} //己方持球
//	EnemyTeamHoldChannel   chan struct{} //敌方持球
//	SelfTeamScoreChannel   chan struct{} //己方得分
//	EnemyTeamScoreChannel  chan struct{} //敌方得分
//	SelfTeamPass           chan struct{} //己方传球
//	EnemyTeamPass          chan struct{} //敌方传球
//	SelfTeamPointSphere    chan struct{} //己方点球
//	EnemyTeamPointSphere   chan struct{} //敌方点球
//	SelfTeamCornerKick     chan struct{} //己方角球
//	EnemyTeamCornerKick    chan struct{} //敌方角球
//	SelfTeamBreakTheRules  chan struct{} //己方犯规
//	EnemyTeamBreakTheRules chan struct{} //敌方犯规
