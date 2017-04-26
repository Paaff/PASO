package server

import "github.com/paaff/PASO/store"

// InitPolicy - Initializes the policy loading.
func initPolicy() {
	store.ValidClientsMap = store.ValidClients{}
	store.ValidClientsMap.NewValidClientsMap()
	store.ValidClientsMap.Set("24:DA:9B:BB:EE:2B", "Peter Fischer")

	store.CollectedClients = store.ClientDataMap{}
	store.CollectedClients.NewCollectedClientsMap()

}
