package football

import (
	"football_simulate/geomtry/solid"
	"github.com/gin-gonic/gin"
	_ "github.com/google/uuid"
	"time"
)

type PlayerNews struct {
	CurrentLocation solid.Vector //坐标
	RedCardNum      int          //红牌数
	YellowCardNum   int          //黄牌数
	Position        int          //职位
	Number          int          //编号
	Name            string       //名字
	TeamName        string       //队伍名字
}

type MatchNewsTemplate struct {
	Ball              Ball                  //球
	Where             string                //在哪里举办
	Name              string                //名字
	BeginTime         int64                 //开始时间
	EndTime           int64                 //结束时间
	//Name2PlayerNews   map[string]PlayerNews //
	Referee           Referee               //裁判
	HolderName        string                //持球
	ScoreName         string                //得分
	PassName          string                //传球
	PointName         string                //点球
	CornerKickName    string                //角球
	BreakTheRulesName string                //犯规
}

type News struct {
	Hub     *Hub
	Matches map[string]*Match
}

func NewNews() News {
	return News{
		Hub:     NewHub(),
		Matches: make(map[string]*Match),
	}
}

func (self *News) Run() {
	//
	self.InitGin()
	//todo 监听news 发布
	tickerNews := time.NewTicker(time.Second) //暂时模拟

	for {
		select {
		case <-tickerNews.C:
			self.Hub.Broadcast <- []byte("sss")

		}
	}

}

func (self *News) InitGin() {
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	router := r.Group("/")
	{
		router.POST("/news", self.handleNews)
	}
	r.Run(":8080")
}

type MatchNewsResp struct {
	Err  string
	Code int
}

func (self *News) handleNews(c *gin.Context) {

	tmp := []MatchNewsTemplate{}
	err := c.Bind(&tmp)
	if err != nil {
		c.JSON(400, []byte(err.Error()))
	}
	if self.Matches[tmp[0].Name] == nil {
		match := NewMatch()
		self.Matches[tmp[0].Name] = &match
		match.Run()
	}
	if tmp[0].EndTime > time.Now().Unix() {
		self.Matches[tmp[0].Name].MatchChannel <- MatchChannel{MatchNormalC, tmp}
	}
}
