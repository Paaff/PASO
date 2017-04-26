package store

import (
	"sync"
)

// CollectedClients holds all the current clients detected
var CollectedClients ClientDataMap

// ClientDataMap is a Global slice designed to hold the Bluedata items.
type ClientDataMap struct {
	sync.RWMutex
	items map[string]BlueData
}

// NewCollectedClientsMap initializes the map
func (cdm *ClientDataMap) NewCollectedClientsMap() {
	cdm.items = make(map[string]BlueData)
}

// Set function will acquire a lock on the slice, append and release the lock.
func (cdm *ClientDataMap) Set(key string, value BlueData) {
	cdm.Lock()
	defer cdm.Unlock()
	cdm.items[key] = value

}

// Get function will acquire a read lock and return the slice
func (cdm *ClientDataMap) Get(key string) (BlueData, bool) {
	cdm.RLock()
	defer cdm.RUnlock()
	value, ok := cdm.items[key]
	return value, ok
}

// GetAsSlice will pull each value from the map and return it as a slice of BlueData
func (cdm *ClientDataMap) GetAsSlice() []BlueData {
	var result []BlueData
	cdm.RLock()
	defer cdm.RUnlock()

	for _, v := range cdm.items {
		result = append(result, v)
	}
	return result
}

// BlueData - Bluetooth data.
type BlueData struct {
	Name      string
	Address   string
	Class     string
	Timestamp string
}

// ValidClientsMap is the global map to view the valid clients
var ValidClientsMap ValidClients

// ValidClients is a map that holds the predetermined information about which clients (names of them)
// and their respective BT Address.
type ValidClients struct {
	items map[string]string
}

// NewValidClientsMap initializes the map
func (vm *ValidClients) NewValidClientsMap() {
	vm.items = make(map[string]string)
}

// Set - Selfexplanatory
func (vm *ValidClients) Set(key string, value string) {
	vm.items[key] = value
}

// Get - Selfexplanatory
func (vm *ValidClients) Get(key string) (string, bool) {
	value, ok := vm.items[key]
	return value, ok
}
