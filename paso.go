package main

import (
	"flag"
	"log"

	"github.com/paaff/PASO/client"
	"github.com/paaff/PASO/config"
	"github.com/paaff/PASO/server"
)

var serverConfPath = "./config/serverconf.json"
var clientConfPath = "./config/clientconf.json"

func main() {
	// Check what instance should be launched.
	launchFlag := flag.String("launch", "client", "Flag that describes whether its a server or a client instance.")
	flag.Parse()

	if *launchFlag == "server" {
		// TODO: Should this be started as a "go" function - Async
		// Why should it be that?
		c, err := config.LoadConfig(serverConfPath)
		if err != nil {
			log.Fatal(err)
		}
		server.Start(c)

	} else if *launchFlag == "client" {
		c, err := config.LoadConfig(clientConfPath)
		if err != nil {
			log.Fatal(err)
		}
		client.Start(c)

	} else {
		log.Fatal("Launch Flag is either 'server' or 'client'")
	}
}

func initLaunchFlags() {

}

func isUnix(runtimeOS string) bool {
	return runtimeOS == "linux"
}
