package gameday

import (
	"testing"
    "os"
)

func TestSchedule(t *testing.T) {
    handle, err := os.Open("fixtures/2013_07_13_epg.xml")

    if err != nil {
            t.Fatal(err)
    }

    s, err := LoadSchedule(handle)

    if err != nil {
            t.Fatal(err)
    }

    if s.Date != "20130714" {
            t.Errorf("Expected the date to be '20130714' not '%s'", s.Date)
    }

    if s.DisplayTimeZone != "ET" {
            t.Errorf("Expected the date to be 'ET' not '%s'", s.DisplayTimeZone)
    }
}
