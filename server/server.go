package server

import "fmt"

// Start - Global function to start the server.
func Start() {
	fmt.Println("Shit son")
}

type sensorData struct {
	timestamp float64
}

type client struct {
	id string
	sensorData
}
