package test

import (
	"bytes"
	"encoding/json"
	"fmt"
	"football_simulate/football"
	"football_simulate/geomtry/solid"
	"io/ioutil"
	"net/http"
	"testing"
)

//模拟发新闻端测试
func TestNews_Send(t *testing.T) {
	news := []football.MatchNewsTemplate{
		NewMatchNews(),
		NewMatchNews(),
	}
	for i := 0; i < len(news); i++ {
		data, err := json.Marshal(news[i])
		if err != nil {
			fmt.Println(err.Error())
		}
		resp, err := http.Post("127.0.0.1:8080", "application/json;charset=utf-8", bytes.NewReader(data))
		body, err := ioutil.ReadAll(resp.Body)
		fmt.Println(string(body))
		respS := football.MatchNewsResp{}
		err = json.Unmarshal(body, &respS)
	}
}

//
func NewMatchNews() football.MatchNewsTemplate {

	return football.MatchNewsTemplate{
		Ball: football.Ball{
			Location: solid.Vector{},
		},
		Where:           "",
		Name:            "",
		BeginTime:       0,
		EndTime:         0,
		Name2PlayerNews: nil,
		Referee: football.Referee{
			Name:        "",
			Age:         0,
			Nationality: "",
			Rank:        0,
			Location:    solid.Vector{},
		},
	}
}
