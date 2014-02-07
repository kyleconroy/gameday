package gameday

import (
	"encoding/xml"
	"io"
	"io/ioutil"
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

type LineScore struct {
        Inning int `xml:"inning,attr"`
        Home int `xml:"home_inning_runs,attr"`
        Away int `xml:"away_inning_runs,attr"`
}

type Game struct {
	Id           string `xml:"id,attr"`
	Venue        string `xml:"venue,attr"`
	Key          string `xml:"game_pk,attr"`
	Timezone     string `xml:"time_zone,attr"`
	AwayTeamCode string `xml:"away_code,attr"`
	AwayTeamId   string `xml:"away_team_id,attr"`
	AwayTeamName string `xml:"away_team_name,attr"`
	HomeTeamCode string `xml:"home_code,attr"`
	HomeTeamId   string `xml:"home_team_id,attr"`
	HomeTeamName string `xml:"home_team_name,attr"`
    Lines []LineScore `xml:"linescore"`
}

type Schedule struct {
	Games           []Game `xml:"game"`
	Date            string `xml:"date,attr"`
	DisplayTimeZone string `xml:"display_time_zone,attr"`
}

type Broadcast struct {
        Radio string `xml:"radio"`
        TV string `xml:"tv"`
}

type Story struct {
        Headline string `xml:"headline"`
        Blurb string `xml:"blurb"`
        Url string `xml:"url"`
}

type Pitcher struct {
        Id string `xml:"player_id"`
        FirstName string `xml:"useName"`
        LastName string `xml:"lastName"`
        RosterName string `xml:"rosterDisplayName"`
        Number int `xml:"number"`
        Hand string `xml:"throwinghand"`
        Wins int `xml:"wins"`
        Lossers int `xml:"losses"`
        StrikeOuts int `xml:"so"`
        ERA float64 `xml:"era"`
        Report string `xml:"report"`
}

type Gamecenter struct {
        Id string `xml:"id,attr"`
        Status string `xml:"status,attr"`
        StartTime string `xml:"start_time,attr"`
        Meridiem string `xml:"ampm,attr"`
        Timezone string `xml:"time_zone,attr"`
        Type string `xml:"type,attr"`
        League string `xml:"league,attr"`
        VenueShort string `xml:"venueShort"`
        VenueLong string `xml:"venueLong"`
        HomeBroadcast Broadcast `xml:"broadcast>home"`
        HomePitcher Pitcher `xml:"probables>home"`
        HomePreview Story `xml:"previews>home"`
        HomeRecap Story `xml:"recaps>home"`
        AwayBroadcast Broadcast `xml:"broadcast>away"`
        AwayPitcher Pitcher `xml:"probables>away"`
        AwayRecap Story `xml:"recaps>away"`
        MLBPreview Story `xml:"preview>mlb"`
        MLBWrap Story `xml:"wrap>mlb"`
}

func Load(reader io.Reader, anything interface{}) (error) {
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
