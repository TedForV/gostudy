package flyweight

import "time"

type Team struct {
	ID             uint64
	Name           string
	Shield         []byte
	Players        []Player
	HistoricalData []HistoricalData
}

const (
	TEAM_A = iota
	TEAM_B
)

type Player struct {
	Name         string
	Surname      string
	PreviousTeam uint64
	Photo        []byte
}

type HistoricalData struct {
	Year          uint8
	LeagueResults []Match
}

type Match struct {
	Date          time.Time
	VisitorID     uint64
	LocalId       uint64
	LocalScore    byte
	VisitorScore  byte
	LocalShoots   uint64
	VisitorShoots uint64
}

type teamFlyweightFacotry struct {
	createdTeams map[string]*Team
}

func (t *teamFlyweightFacotry) GetTeam(name string) *Team {
	return nil
}

func (t *teamFlyweightFacotry) GetNumberOfObjects() int {
	return 0
}

func (t *teamFlyweightFacotry) GetTeam(teamID int) *Team {
	if t.createdTeams[teamID] != nil {
		return t.createdTeams[teamID]
	}
	team := getTeamFactory(teamID)
	t.createdTeams[teamID] = &team

	return t.createdTeams[teamID]
}
