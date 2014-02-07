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

type Gamecenter struct {
        Id string `xml:"id,attr"`
        VenueShort string `xml:"venueShort"`
        VenueLong string `xml:"venueLong"`
        HomeBroadcast Broadcast `xml:"broadcast>home"`
        AwayBroadcast Broadcast `xml:"broadcast>away"`
}

func LoadSchedule(reader io.Reader) (Schedule, error) {
	var schedule Schedule

	blob, err := ioutil.ReadAll(reader)

	if err != nil {
		return schedule, err
	}

	err = xml.Unmarshal(blob, &schedule)

	if err != nil {
		return schedule, err
	}

	return schedule, nil
}

func LoadGamecenter(reader io.Reader) (Gamecenter, error) {
	var center Gamecenter

	blob, err := ioutil.ReadAll(reader)

	if err != nil {
		return center, err
	}

	err = xml.Unmarshal(blob, &center)

	if err != nil {
		return center, err
	}

	return center, nil
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
