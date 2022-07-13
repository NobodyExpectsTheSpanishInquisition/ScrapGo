package src

import (
	"testing"
	"time"
)

func Test_Analyze_ShouldReturnFalse_WhenProvidedDataDoNotContainExpectedDateRange(t *testing.T) {
	dateFrom := time.Date(2022, 5, 1, 12, 0, 0, 0, time.UTC)
	dateTo := time.Date(2022, 6, 1, 12, 0, 0, 0, time.UTC)

	result := analyze(getOfferAvailabilities(), dateFrom, dateTo)

	if true == result {
		t.Fail()
	}
}

func Test_Analyze_ShouldReturnTrue_WhenProvidedDataContainsExpectedDateRange(t *testing.T) {
	dateFrom := time.Date(2022, 3, 2, 12, 0, 0, 0, time.UTC)
	dateTo := time.Date(2022, 3, 10, 12, 0, 0, 0, time.UTC)

	result := analyze(getOfferAvailabilities(), dateFrom, dateTo)

	if false == result {
		t.Fail()
	}
}

func getOfferAvailabilities() offerAvailabilities {
	return offerAvailabilities{
		[]dateRange{
			{
				From: time.Date(2022, 1, 1, 12, 0, 0, 0, time.UTC),
				To:   time.Date(2022, 2, 1, 12, 0, 0, 0, time.UTC),
			},
			{
				From: time.Date(2022, 3, 1, 12, 0, 0, 0, time.UTC),
				To:   time.Date(2022, 4, 1, 12, 0, 0, 0, time.UTC),
			},
		},
	}
}
