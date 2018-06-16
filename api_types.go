package footballdata

import "time"

type FixtureStatus string

const (
	FixtureStatus_Scheduled FixtureStatus = "SCHEDULED"
	FixtureStatus_Timed     FixtureStatus = "TIMED"
	FixtureStatus_Postponed FixtureStatus = "POSTPONED"
	FixtureStatus_InPlay    FixtureStatus = "IN_PLAY"
	FixtureStatus_Canceled  FixtureStatus = "CANCELED"
	FixtureStatus_Finished  FixtureStatus = "FINISHED"
)

// Describes the venue.
type Venue string

const (
	// The home venue.
	Venue_Home Venue = "home"
	// The away venue.
	Venue_Away Venue = "away"
)

// DEPRECATED.
//
// Contains the list of soccer seasons returned by the API.
type SoccerSeasonList []SoccerSeason

// DEPRECATED.
//
// Contains information about a soccer season.
type SoccerSeason struct {
	Id                uint64
	Caption           string
	League            string
	Year              string
	CurrentMatchday   uint16
	NumberOfMatchdays uint16
	NumberOfTeams     uint16
	NumberOfGames     uint16
	LastUpdated       time.Time
}

// Contains the list of competitions returned by the API.
type CompetitionList []Competition

// Contains information about a competition.
type Competition struct {
	Id                uint64
	Caption           string
	League            string
	Year              string
	CurrentMatchday   uint16
	NumberOfMatchdays uint16
	NumberOfTeams     uint16
	NumberOfGames     uint16
	LastUpdated       time.Time
}

// Contains the fixture and the head to head information delivered by the API
// for a wanted fixture.
type FixtureResponse struct {
	Fixture   Fixture
	Head2Head Head2Head
}

// Contains head to head information.
type Head2Head struct {
	Count               uint64
	TimeFrameStart      time.Time
	TimeFrameEnd        time.Time
	HomeTeamWins        uint64
	AwayTeamWins        uint64
	Draws               uint64
	LastHomeWinHomeTeam *Fixture
	LastWinHomeTeam     *Fixture
	LastAwayWinAwayTeam *Fixture
	Fixtures            []Fixture
}

// Contains information about a fixture.
type Fixture struct {
	Date         time.Time
	Status       FixtureStatus
	Matchday     uint16
	HomeTeamName string
	AwayTeamName string
	Result       FixtureResult

	//HomeTeamId   uint64
	//AwayTeamId   uint64
}

// Contains a list of fixtures.
type FixtureList struct {
	Count    uint64
	Fixtures []Fixture
}

// Contains information about a list of fixtures as returned by the API.
type FixturesResponse struct {
	FixtureList

	TimeFrameStart time.Time
	TimeFrameEnd   time.Time
}

// Contains information about a fixture's results.
type FixtureResult struct {
	FixtureScore

	HalfTime        *FixtureScore
	ExtraTime       *FixtureScore
	PenaltyShootout *FixtureScore
}

type FixtureScore struct {
	GoalsHomeTeam uint16
	GoalsAwayTeam uint16
}

// Contains a list of teams.
type TeamList struct {
	Count uint64
	Teams []Team
}

// Contains information about a team.
type Team struct {
	Id               uint64
	Name             string
	ShortName        string
	SquadMarketValue string
	CrestUrl         string
}

// Contains a list of players.
type PlayerList struct {
	Count   uint64
	Players []Player
}

// Contains information about a player.
type Player struct {
	Id            uint64
	Name          string
	Position      string
	JerseyNumber  uint8
	DateOfBirth   time.Time
	Nationality   string
	ContractUntil time.Time
	MarketValue   string
}

// Contains the league table for a season.
type LeagueTable struct {
	LeagueCaption string
	Matchday      uint16
	Standing      []TeamLeagueStatisticsInStanding
	Standings     map[string][]TeamLeagueStatisticsInStandings
}

// Contains statistical information about a team's performance in a league.
// TODO - Introduce minified variant of this struct.
type TeamLeagueStatistics struct {
	CrestURI       string
	Draws          uint16
	GoalDifference int16
	Goals          uint16
	GoalsAgainst   uint16
	Losses         uint16
	PlayedGames    uint8
	Points         uint16
	Wins           uint16
	Home           ShortTeamLeagueStatistics
	Away           ShortTeamLeagueStatistics
}

// Contains statistical information as a variant for the "Standing" field.
type TeamLeagueStatisticsInStanding struct {
	TeamLeagueStatistics
	TeamName string
	Position uint8
}

// Contains statistical information as a variant for the "Standings" field.
type TeamLeagueStatisticsInStandings struct {
	TeamLeagueStatistics
	Team string
	Rank uint8
}

type ShortTeamLeagueStatistics struct {
	Draws        uint16
	Goals        uint16
	GoalsAgainst uint16
	Losses       uint16
	Wins         uint16
}
