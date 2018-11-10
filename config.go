package oanda

import (
	"fmt"

	"github.com/spf13/viper"
)

func GetToken() (name, token string) {
	viper.SetConfigName("oanda")
	viper.AddConfigPath("$HOME/.config/")
	viper.AddConfigPath(".")
	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}
	name = viper.GetString("name")
	token = viper.GetString("token")
	return
}
