package main

import (
	"github.com/spf13/viper"
	"notifier/src"
)

func main() {
	var err error

	err = loadEnv()

	n := src.NewSmptNotifier(
		viper.GetString("NOTIFIER_HOST"),
		viper.GetString("NOTIFIER_PORT"),
		viper.GetString("NOTIFIER_PASSWORD"),
	)

	email := src.NewEmail(
		viper.GetString("FROM_EMAIL"),
		viper.GetString("TO_EMAIL"),
		"Subject: PRIVATE NOTIFICATION",
		"BODY",
	)

	err = n.SendEmail("", email)

	if err != nil {
		panic(err)
	}
}

func loadEnv() error {
	var err error

	viper.SetConfigFile(".env")
	err = viper.ReadInConfig()

	return err
}
