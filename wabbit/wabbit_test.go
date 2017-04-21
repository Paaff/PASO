package wabbit

import "testing"

func TestCreateDialPath(t *testing.T) {
	// Test data
	username := "username"
	pass := "pass"
	address := "192.168.0.150"
	port := "8080"

	expected := "amqp://username:pass@192.168.0.150:8080/"
	actual := createDialPath(username, pass, address, port)
	if expected != actual {
		t.Errorf("Expected: %s\nActual: %s", expected, actual)
	}

}
