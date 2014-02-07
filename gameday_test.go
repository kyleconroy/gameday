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

	s, err := LoadSchedule(handle)

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

func TestGamecenter(t *testing.T) {
	handle, err := os.Open("fixtures/2013_07_14_minmlb_nyamlb_1_gamecenter.xml")

	if err != nil {
		t.Fatal(err)
	}

	gc, err := LoadGamecenter(handle)

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
}
