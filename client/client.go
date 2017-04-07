package client

import "github.com/paaff/PASO/config"

// Start - Global function to start the client.
func Start(conf *config.Config) {
	InitCRabbit(conf)

}
