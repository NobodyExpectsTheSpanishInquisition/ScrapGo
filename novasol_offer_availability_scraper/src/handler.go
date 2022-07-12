package src

import (
	"encoding/json"
	"time"
)

type command struct {
	url      string
	dateFrom time.Time
	dateTo   time.Time
}

func newCommand(url string, dateFrom time.Time, dateTo time.Time) command {
	return command{url: url, dateFrom: dateFrom, dateTo: dateTo}
}

type offerAvailabilities struct {
	DatesRanges []dateRange
}

type dateRange struct {
	From time.Time `json:"from"`
	To   time.Time `json:"to"`
}

func (d *dateRange) UnmarshalJSON(bytes []byte) error {
	var err error
	var data struct {
		From string `json:"from"`
		To   string `json:"to"`
	}

	err = json.Unmarshal(bytes, &data)

	from, err := time.Parse("20060102", data.From)
	to, err := time.Parse("20060102", data.To)

	d.From = from
	d.To = to

	return err
}

func handle(c command) (response, error) {
	var err error
	availabilities, err := scrap(c.url)

	if err != nil {
		return newResponse(false), err
	}

	response := analyze(availabilities, c.dateFrom, c.dateTo)

	return newResponse(response), err
}

func analyze(availabilities offerAvailabilities, dateFrom time.Time, dateTo time.Time) bool {
	available := false
	datesRanges := availabilities.DatesRanges

	for _, datesRange := range datesRanges {
		if dateFrom.After(datesRange.From) && dateTo.Before(dateTo) {
			available = true
			break
		}
	}

	return available
}
