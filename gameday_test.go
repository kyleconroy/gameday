package gameday

import (
	"os"
	"reflect"
	"testing"
)

func check(t *testing.T, a interface{}, b interface{}) {
	if !reflect.DeepEqual(a, b) {
		t.Errorf("'%+v' != '%+v'", a, b)
	}
}

func TestSchedule(t *testing.T) {
	handle, err := os.Open("fixtures/2013_07_13_epg.xml")

	if err != nil {
		t.Fatal(err)
	}

	var s Schedule

	err = Load(handle, &s)

	if err != nil {
		t.Fatal(err)
	}

	check(t, s.Date, "20130714")
	check(t, s.DisplayTimeZone, "ET")
	check(t, len(s.Games), 15)

	game := s.Games[0]

	check(t, game.Id, "2013/07/14/kcamlb-clemlb-1")
	check(t, game.Venue, "Progressive Field")
	check(t, game.Key, "348159")
	check(t, game.Timezone, "ET")
	check(t, game.AwayTeamId, Royals)
	check(t, game.AwayTeamCode, "kca")
	check(t, game.HomeTeamId, Indians)
	check(t, game.HomeTeamCode, "cle")
}

func TestBoxscore(t *testing.T) {
	handle, err := os.Open("fixtures/2013_07_14_minmlb_nyamlb_1_boxscore.xml")

	if err != nil {
		t.Fatal(err)
	}

	var box BoxScore

	err = Load(handle, &box)

	if err != nil {
		t.Fatal(err)
	}

	check(t, box.GameId, "2013/10/08/oakmlb-detmlb-1")
	check(t, box.HomeId, Tigers)
	check(t, box.AwayId, Athletics)
	check(t, len(box.Pitching), 2)

	pitching := box.Pitching[0]

	check(t, len(pitching.Pitchers), 5)

	pitcher := pitching.Pitchers[0]

	check(t, pitcher.Name, "Straily")
	check(t, pitcher.Position, "P")
	check(t, pitcher.Outs, 18)
}

func TestLinescore(t *testing.T) {
	handle, err := os.Open("fixtures/2013_07_14_minmlb_nyamlb_1_linescore.xml")

	if err != nil {
		t.Fatal(err)
	}

	var game Game

	err = Load(handle, &game)

	if err != nil {
		t.Fatal(err)
	}

	check(t, game.Id, "2013/07/14/minmlb-nyamlb-1")
	check(t, len(game.Lines), 9)

	score := game.Lines[1]

	check(t, score.Inning, "2")
	check(t, score.Home, "0")
	check(t, score.Away, "2")
}

func TestGamecenter(t *testing.T) {
	handle, err := os.Open("fixtures/2013_07_14_minmlb_nyamlb_1_gamecenter.xml")

	if err != nil {
		t.Fatal(err)
	}

	var gc Gamecenter

	err = Load(handle, &gc)

	if err != nil {
		t.Fatal(err)
	}

	check(t, gc.Id, "2013_07_14_minmlb_nyamlb_1")
	check(t, gc.VenueShort, "Yankee Stadium")
	check(t, gc.VenueShort, "Yankee Stadium")
	check(t, gc.HomeBroadcast.Radio, "WCBS 880, WADO 1280")
	check(t, gc.HomeBroadcast.TV, "YES")
	check(t, gc.AwayBroadcast.Radio, "96.3 K-TWIN, TIBN, BOB FM")
	check(t, gc.AwayBroadcast.TV, "FS-N")
	check(t, gc.Status, "F")
	check(t, gc.StartTime, "1:05")
	check(t, gc.AMPM, "pm")
	check(t, gc.Type, "R")
	check(t, gc.League, "AL")

	check(t, gc.HomePitcher.FirstName, "CC")
	check(t, gc.HomeRecap.Url, "/mlb/gameday/index.jsp?gid=2013_07_14_minmlb_nyamlb_1&mode=recap&c_id=nyy")
}
