package gameday

import (
	"encoding/xml"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"time"
)

const (
	Angels    = "108"
	Astros    = "117"
	Athletics = "133"
	BlueJays  = "141"
	Braves    = "144"
	Brewers   = "158"
	Cardinals = "138"
	Cubs      = "112"
	Dbacks    = "109"
	Dodgers   = "119"
	Giants    = "137"
	Indians   = "114"
	Mariners  = "136"
	Marlins   = "146"
	Mets      = "121"
	Nationals = "120"
	Orioles   = "110"
	Padres    = "135"
	Phillies  = "143"
	Pirates   = "134"
	Rangers   = "140"
	Rays      = "139"
	RedSox    = "111"
	Reds      = "113"
	Rockies   = "115"
	Royals    = "118"
	Tigers    = "116"
	Twins     = "142"
	WhiteSox  = "145"
	Yankees   = "147"
)
const (
	GamedayHostname = "http://gd2.mlb.com"
	GamedayBaseUrl  = "http://gd2.mlb.com/components/game/mlb"
)

type BoxLineScore struct {
	Inning string `xml:"inning,attr"`
	Home   string `xml:"home,attr"`
	Away   string `xml:"away,attr"`
}

type Pitcher struct {
	Id             string  `xml:"id,attr"`
	Name           string  `xml:"name,attr"`
	InningsPitched int     `xml:"`
	Hits           int     `xml:"h,attr"`
	Runs           int     `xml:"r,attr"`
	EarnedRuns     int     `xml:"er,attr"`
	Walks          int     `xml:"bb,attr"`
	Strikeouts     int     `xml:"so,attr"`
	Pitches        int     `xml:"np,attr"`
	Position       string  `xml:"pos,attr"`
	Outs           int     `xml:"out,attr"`
	Strikes        int     `xml:"s,attr"`
	Win            bool    `xml:"win,attr"`
	Loss           bool    `xml:"loss,attr"`
	Note           string  `xml:"note,attr"`
	ERA            float64 `xml:"era,attr"`
}

type BoxPitching struct {
	InningsPitched int       `xml:"`
	Hits           int       `xml:"h,attr"`
	Runs           int       `xml:"r,attr"`
	EarnedRuns     int       `xml:"er,attr"`
	Walks          int       `xml:"bb,attr"`
	Strikeouts     int       `xml:"so,attr"`
	Homeruns       int       `xml:"hr,attr"`
	Outs           int       `xml:"outs,attr"`
	ERA            float64   `xml:"era,attr"`
	Pitchers       []Pitcher `xml:"pitcher"`
}

type Batter struct {
	Id             string  `xml:"id,attr"`
	Name           string  `xml:"name,attr"`
	AtBats         int     `xml:"ab,attr"`
	Runs           int     `xml:"r,attr"`
	Hits           int     `xml:"h,attr"`
	Doutbles       int     `xml:"d,attr"`
	Triples        int     `xml:"t,attr"`
	HomeRuns       int     `xml:"hr,attr"`
	RBIs           int     `xml:"rbi,attr"`
	Walks          int     `xml:"bb,attr"`
	Strikeouts     int     `xml:"so,attr"`
	BattingAverage float64 `xml:"avg,attr"`
	Position       string  `xml:"pos,attr"`
}

type BoxBatting struct {
	Flag           string   `xml:"team_flag,attr"`
	AtBats         int      `xml:"ab,attr"`
	Runs           int      `xml:"r,attr"`
	Hits           int      `xml:"h,attr"`
	Doutbles       int      `xml:"d,attr"`
	Triples        int      `xml:"t,attr"`
	HomeRuns       int      `xml:"hr,attr"`
	RBIs           int      `xml:"rbi,attr"`
	Walks          int      `xml:"bb,attr"`
	StrikeOuts     int      `xml:"so,attr"`
	BattingAverage float64  `xml:"avg,attr"`
	Batters        []Batter `xml:"batter"`
}

type BoxScore struct {
	GameId   string         `xml:"game_id,attr"`
	GameKey  string         `xml:"game_pk,attr"`
	HomeId   string         `xml:"home_id,attr"`
	AwayId   string         `xml:"away_id,attr"`
	Lines    []BoxLineScore `xml:"linescore>inning_line_score"`
	Pitching []BoxPitching  `xml:"pitching"`
	Batting  []BoxBatting   `xml:"batting"`
}

type GameLineScore struct {
	Inning string `xml:"inning,attr"`
	Home   string `xml:"home_inning_runs,attr"`
	Away   string `xml:"away_inning_runs,attr"`
}

type Score struct {
	Team        string `xml:"team,attr"`
	Description string `xml:"description,attr"`
	AwayScore   int    `xml:"away_score,attr"`
	HomeScore   int    `xml:"home_score"`
}

type Summary struct {
	Inning string  `xml:"inning,attr"`
	Plays  []Score `xml:"score"`
}

type Game struct {
	AwayAPMP     string          `xml:"away_ampm,attr"`
	AwayLoss     string          `xml:"away_loss,attr"`
	AwayTeamCity string          `xml:"away_team_city,attr"`
	AwayTeamCode string          `xml:"away_code,attr"`
	AwayTeamId   string          `xml:"away_team_id,attr"`
	AwayTeamName string          `xml:"away_team_name,attr"`
	AwayTime     string          `xml:"away_time,attr"`
	AwayTimezone string          `xml:"away_time_zone,attr"`
	AwayWin      string          `xml:"away_win,attr"`
	HomeAMPM     string          `xml:"home_ampm,attr"`
	HomeLoss     string          `xml:"home_loss,attr"`
	HomeTeamCity string          `xml:"home_team_city,attr"`
	HomeTeamCode string          `xml:"home_code,attr"`
	HomeTeamId   string          `xml:"home_team_id,attr"`
	HomeTeamName string          `xml:"home_team_name,attr"`
	HomeTime     string          `xml:"home_time,attr"`
	HomeTimezone string          `xml:"home_time_zone,attr"`
	HomeWin      string          `xml:"home_win,attr"`
	Id           string          `xml:"id,attr"`
	Key          string          `xml:"game_pk,attr"`
	Lines        []GameLineScore `xml:"linescore"`
	Timezone     string          `xml:"time_zone,attr"`
	Venue        string          `xml:"venue,attr"`
	Weather      Weather         `xml:"weather"`
	Summaries    []Summary       `xml:"scoring-summary"`
	DataDir      string          `xml:"game_data_directory,attr"`
}

func (g *Game) loadFile(path string, val interface{}) error {
	if g.DataDir == "" {
		return fmt.Errorf("The DataDir is empty, so now URL can be created")
	}

	url := fmt.Sprintf(GamedayHostname + g.DataDir + "/" + path)

	resp, err := http.Get(url)

	if err != nil {
		return err
	}

	err = Load(resp.Body, val)

	if err != nil {
		return err
	}

	return nil
}

func (g *Game) GetGameCenter() (GameCenter, error) {
	var gc GameCenter
	err := g.loadFile("gamecenter.xml", &gc)
	return gc, err
}

func (g *Game) GetBoxScore() (BoxScore, error) {
	var box BoxScore
	err := g.loadFile("boxscore.xml", &box)
	return box, err
}

func (g *Game) GetSummaries() ([]Summary, error) {
	var gd Gameday
	err := g.loadFile("gameday_Syn.xml", &gd)
	return gd.Game.Summaries, err
}

func (g *Game) GetWeather() (Weather, error) {
	var game Game
	err := g.loadFile("plays.xml", &game)
	return game.Weather, err
}

type Weather struct {
	Wind        string `xml:"wind,attr"`
	Temperature string `xml:"temp,attr"`
	Conditions  string `xml:"condition,attr"`
}

type Schedule struct {
	Games           []Game `xml:"game"`
	Date            string `xml:"date,attr"`
	DisplayTimeZone string `xml:"display_time_zone,attr"`
}

type Broadcast struct {
	Radio string `xml:"radio"`
	TV    string `xml:"tv"`
}

type Story struct {
	Headline string `xml:"headline"`
	Blurb    string `xml:"blurb"`
	Url      string `xml:"url"`
}

type GCPitcher struct {
	Id         string  `xml:"player_id"`
	FirstName  string  `xml:"useName"`
	LastName   string  `xml:"lastName"`
	RosterName string  `xml:"rosterDisplayName"`
	Number     int     `xml:"number"`
	Hand       string  `xml:"throwinghand"`
	Wins       int     `xml:"wins"`
	Lossers    int     `xml:"losses"`
	StrikeOuts int     `xml:"so"`
	ERA        float64 `xml:"era"`
	Report     string  `xml:"report"`
}

type GameCenter struct {
	Id            string    `xml:"id,attr"`
	Status        string    `xml:"status,attr"`
	StartTime     string    `xml:"start_time,attr"`
	AMPM          string    `xml:"ampm,attr"`
	Timezone      string    `xml:"time_zone,attr"`
	Type          string    `xml:"type,attr"`
	League        string    `xml:"league,attr"`
	VenueShort    string    `xml:"venueShort"`
	VenueLong     string    `xml:"venueLong"`
	HomeBroadcast Broadcast `xml:"broadcast>home"`
	HomePitcher   GCPitcher `xml:"probables>home"`
	HomePreview   Story     `xml:"previews>home"`
	HomeRecap     Story     `xml:"recaps>home"`
	AwayBroadcast Broadcast `xml:"broadcast>away"`
	AwayPitcher   GCPitcher `xml:"probables>away"`
	AwayRecap     Story     `xml:"recaps>away"`
	MLBPreview    Story     `xml:"preview>mlb"`
	MLBWrap       Story     `xml:"wrap>mlb"`
}

type Gameday struct {
	Game Game `xml:"game"`
}

func Load(reader io.Reader, anything interface{}) error {
	blob, err := ioutil.ReadAll(reader)
	if err != nil {
		return err
	}
	err = xml.Unmarshal(blob, anything)
	if err != nil {
		return err
	}
	return nil
}

func TeamSubreddit(id string) string {
	subs := map[string]string{
		Angels:    "/r/AngelsBaseball",
		Astros:    "/r/Astros",
		Athletics: "/r/oaklandathletics",
		BlueJays:  "/r/TorontoBlueJays",
		Braves:    "/r/Braves",
		Brewers:   "/r/Brewers",
		Cardinals: "/r/Cardinals",
		Cubs:      "/r/Cubs",
		Dbacks:    "/r/azdiamondbacks",
		Dodgers:   "/r/Dodgers",
		Giants:    "/r/SFGiants",
		Indians:   "/r/WahoosTipi",
		Mariners:  "/r/Mariners",
		Marlins:   "/r/letsgofish",
		Mets:      "/r/NewYorkMets",
		Nationals: "/r/Nationals",
		Orioles:   "/r/Orioles",
		Padres:    "/r/Padres",
		Phillies:  "/r/Phillies",
		Pirates:   "/r/Buccos",
		Rangers:   "/r/TexasRangers",
		Rays:      "/r/TampaBayRays",
		RedSox:    "/r/RedSox",
		Reds:      "/r/Reds",
		Rockies:   "/r/ColoradoRockies",
		Royals:    "/r/KCRoyals",
		Tigers:    "/r/MotorCityKitties",
		Twins:     "/r/MinnesotaTwins",
		WhiteSox:  "/r/WhiteSox",
		Yankees:   "/r/NYYankees",
	}
	return subs[id]
}

func GetGame(id string, t time.Time) (Game, error) {
	y, m, d := t.Date()

	url := fmt.Sprintf(GamedayBaseUrl+"/year_%d/month_%02d/day_%02d/epg.xml", y, m, d)

	resp, err := http.Get(url)

	if err != nil {
		return Game{}, err
	}

	var s Schedule

	err = Load(resp.Body, &s)

	if err != nil {
		return Game{}, err
	}

	for _, game := range s.Games {
		if game.HomeTeamId == id || game.AwayTeamId == id {
			return game, nil
		}
	}

	return Game{}, fmt.Errorf("Team %s doesn't have a game on %s", id, t)
}
