package main

import (
	"flag"
	"fmt"
	"log"
	"runtime"

	"github.com/paaff/PASO/client"
	"github.com/paaff/PASO/config"
	"github.com/paaff/PASO/server"
)

var serverConfPath = "./serverconf.json"
var clientConfPath = "./clientconf.json"

func main() {
	// Check what instance should be launched.
	launchFlag := flag.String("launch", "client", "Flag that describes whether its a server or a client instance.")
	flag.Parse()

	if runtime.GOOS == "linux" {
		fmt.Println("Unix type OS detected")
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
	} else {
		log.Fatal("This is not a linux system, which is intended for this project at the moment.")
	}

}
