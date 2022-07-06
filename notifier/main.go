package main

import "github.com/spf13/viper"

func main() {
	loadEnv()
}

func loadEnv() {
	viper.SetConfigFile(".env")
	err := viper.ReadInConfig()
	if err != nil {
		return
	}
}
