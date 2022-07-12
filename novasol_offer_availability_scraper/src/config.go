package src

import (
	"github.com/spf13/viper"
	"time"
)

func LoadEnv() error {
	var err error

	viper.SetConfigFile("./novasol_offer_availability_scraper/.env")
	err = viper.ReadInConfig()

	return err
}

func getCallendarOfferUrl() string {
	return viper.GetString("OFFER_CALENDAR_URL")
}

func getDateFromToCheck() time.Time {
	return viper.GetTime("OFFER_DATE_FROM")
}

func getDateToToCheck() time.Time {
	return viper.GetTime("OFFER_DATE_TO")
}
