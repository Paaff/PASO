package main

import (
	"fmt"

	"github.com/paaff/PASO/server"
	"github.com/spf13/viper"
)

func main() {
	fmt.Println("Hello World! Get ready for some Auto Presence Detection.")

	// Check what instance should be launched.
	commandFlags()

	// Starting the server.
	// TODO: Should this be started as a "go" function - Async
	server.Start()
}

func commandFlags() {
	// TODO: Use viper to check the command flags and then choose which loadConfig should be ran.
}

// Loads a config
func loadConfig(version string) {
	viper.SetConfigName(version)
	viper.SetConfigType("json")
	viper.AddConfigPath("./config")
	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("Fatal error config file: %s", err))
	}

	// Config loaded properly - Do the correct stuff.

}
