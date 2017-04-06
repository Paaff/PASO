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

func main() {
	// Check what instance should be launched.
	launchFlag := flag.String("launch", "client", "Flag that describes whether its a server or a client instance.")
	flag.Parse()

	if runtime.GOOS == "linux" {
		fmt.Println("Unix type OS detected")
		if *launchFlag == "server" {
			// TODO: Should this be started as a "go" function - Async
			// Why should it be that?
			c := config.LoadConfig(*launchFlag)
			server.Start()

		} else if *launchFlag == "client" {
			c := config.LoadConfig(*launchFlag)
			client.Start()

		} else {
			log.Fatal("Launch Flag is either 'server' or 'client'")
		}
	}

}
