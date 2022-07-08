package src

import "github.com/spf13/viper"

func LoadEnv() error {
	var err error

	viper.SetConfigFile("./notifier/.env")
	err = viper.ReadInConfig()

	return err
}

func GetEnvironment() string {
	return viper.GetString("ENVIRONMENT")
}

func GetNotifierHost() string {
	return viper.GetString("NOTIFIER_HOST")
}

func GetNotifierPort() string {
	return viper.GetString("NOTIFIER_PORT")
}

func GetNotifierIdentity() string {
	return viper.GetString("NOTIFIER_PORT")
}

func GetSenderEmail() string {
	return viper.GetString("FROM_EMAIL")
}

func GetSenderPassword() string {
	return viper.GetString("NOTIFIER_PASSWORD")
}

func GetReceiverEmail() string {
	return viper.GetString("TO_EMAIL")
}
