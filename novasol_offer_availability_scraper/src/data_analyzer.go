package src

import "time"

func analyze(availabilities offerAvailabilities, dateFrom time.Time, dateTo time.Time) bool {
	available := false
	datesRanges := availabilities.DatesRanges

	for _, datesRange := range datesRanges {
		if dateFrom.After(datesRange.From) && dateTo.Before(datesRange.To) {
			available = true
			break
		}
	}

	return available
}
