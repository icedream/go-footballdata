package footballdata

import "time"

type FixtureStatus string

const (
	FixtureStatus_Timed    FixtureStatus = "TIMED"
	FixtureStatus_InPlay   FixtureStatus = "IN_PLAY"
	FixtureStatus_Finished FixtureStatus = "FINISHED"
)

// Describes the venue.
type Venue string

const (
	// The home venue.
	Venue_Home Venue = "home"
	// The away venue.
	Venue_Away Venue = "away"
)

// Contains the list of soccer seasons returned by the API.
type SoccerSeasonList []SoccerSeason

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
