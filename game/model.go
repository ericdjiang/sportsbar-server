package game

import (
	"fmt"
	"time"
)

type SportsDataDateTime struct {
	time.Time
}

func (self *SportsDataDateTime) UnmarshalJSON(b []byte) (err error) {
	s := string(b)
	s = s[1 : len(s)-1]
	s = s + "Z"
	t, err := time.Parse(time.RFC3339, s)
	if err != nil {
		fmt.Println(err)
		return err
	}
	t = t.AddDate(0, 0, 1)
	self.Time = t
	return
}

type Game struct {
	GameID               int
	Season               int
	Status               string
	Day                  SportsDataDateTime
	DateTime             SportsDataDateTime
	AwayTeam             string
	HomeTeam             string
	AwayTeamID           int
	HomeTeamID           int
	AwayTeamScore        int
	HomeTeamScore        int
	Updated              SportsDataDateTime
	Quarter              string
	TimeRemainingMinutes int
	TimeRemainingSeconds int
	PointSpread          float32
	OverUnder            float32
	AwayTeamMoneyLine    int
	HomeTeamMoneyLine    int
}

type Games []Game
