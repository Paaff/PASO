package main

import (
	"flag"
	"fmt"

	"github.com/paaff/PASO/client"
	"github.com/spf13/viper"
)

func main() {
	// Check what instance should be launched.
	//launchFlag := flag.String("launch", "client", "Flag that describes whether its a server or a client instance.")
	flag.Parse()
	//loadConfig(*launchFlag)
	client.Start()
	// if runtime.GOOS == "linux" {
	// 	fmt.Println("Unix type OS detected")
	// 	if *launchFlag == "server" {
	// 		// Starting the server.
	// 		// TODO: Should this be started as a "go" function - Async
	// 		server.Start()
	//
	// 	} else {
	// 		client.Start()
	// 	}
	// }

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

}
