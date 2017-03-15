package server

// Start - Global function to start the server.
func Start() {

}

type sensorData struct {
	timestamp float64
}

type client struct {
	id string
	sensorData
}
