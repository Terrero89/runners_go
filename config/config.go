package config

import (
	"fmt"
	"github.com/spf13/viper"
)

func InitConfig(filename string) *viper.Viper {
	config := viper.New()
	config.SetConfigFile(filename)
	config.AddConfigPath(".")
	config.AddConfigPath("$HOME")
	err := config.ReadInConfig()
	if err != nil {
		fmt.Println("Error reading/parson config file:", err)
	}
	return config

}
