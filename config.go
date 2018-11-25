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

func GetClient() *Oanda {
	name, token := GetToken()
	if name == "" || token == "" {
		panic(fmt.Sprintf("name or token empty:%s %s", name, token))
	}
	api := NewOanda(name, token)
	return api
}

func GetDebugClient() *Oanda {
	name, token := GetToken()
	if name == "" || token == "" {
		panic(fmt.Sprintf("name or token empty:%s %s", name, token))
	}
	api := NewTestOanda(name, token)
	api.SetDebug(true)
	return api
}
