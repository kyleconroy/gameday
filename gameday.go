package gameday

import (
	"encoding/xml"
	"io/ioutil"
    "io"
)

type Game struct {
	Id string
}

type Schedule struct {
	Games []Game
    Date string `xml:"date,attr"`
    DisplayTimeZone string `xml:"display_time_zone,attr"`
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
