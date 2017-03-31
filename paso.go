package main

import (
	"flag"

	"github.com/paaff/PASO/client"
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
