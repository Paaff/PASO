package config

import (
	"fmt"

	"github.com/spf13/viper"
)

// Loads a config
func loadConfig(version string) {
	viper.SetConfigName(version)
	viper.SetConfigType("json")
	viper.AddConfigPath("./config")
	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("Fatal error config file: %s", err))
	}

}
