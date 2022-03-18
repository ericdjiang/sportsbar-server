package game

type Game struct {
	GameID               int
	Season               int
	Status               string
	Day                  string
	DateTime             string
	AwayTeam             string
	HomeTeam             string
	AwayTeamID           int
	HomeTeamID           int
	AwayTeamScore        int
	HomeTeamScore        int
	Updated              string
	Quarter              string
	TimeRemainingMinutes int
	TimeRemainingSeconds int
}

type Games []Game
