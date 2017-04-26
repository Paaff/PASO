package server

import "github.com/paaff/PASO/store"

// InitPolicy - Initializes the policy loading.
func initPolicy() {
	store.ValidClients = make(map[string]string)
	store.ValidClients["24:DA:9B:BB:EE:2B"] = "Peter Fischer"
}
