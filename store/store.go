package store

import (
	"reflect"
	"sync"
)

// CollectedBlueData holds all the current clients detected
var CollectedBlueData ClientDataMap

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
	Address   string
	Class     string
	Timestamp string
}

// ValidClientsMap is the global map to view the valid clients
var ValidClientsMap ValidClients

// ValidClients is a map that holds the predetermined information about which clients (names of them)
// and their respective BT Address.
type ValidClients struct {
	items map[string]Client
}

// NewValidClientsMap initializes the map
func (vm *ValidClients) NewValidClientsMap() {
	vm.items = make(map[string]Client)
}

// Set - Selfexplanatory
func (vm *ValidClients) Set(key string, value Client) {
	vm.items[key] = value
}

// Get - Selfexplanatory
func (vm *ValidClients) Get(key string) (Client, bool) {
	value, ok := vm.items[key]
	return value, ok
}

// Client is a struct for the predetermined users in the system.
type Client struct {
	Name        string
	Permissions []string
}

// Projects holds a list of Project structs with the purpose of demoing and testing.
var Projects ProjectsList

// ProjectsList is a wrapper for the slice of Project.
type ProjectsList struct {
	elements []Project
}

// NewProjectsList creates a new slice of projects with length 0.
func (p *ProjectsList) NewProjectsList() {
	p.elements = make([]Project, 0)
}

// Add - Selfexplanatory
func (p *ProjectsList) Add(elem Project) {
	p.elements = append(p.elements, elem)
}

// Remove - Selfexplanatory
func (p *ProjectsList) Remove(elem Project) {
	var elemI int
	for i := 0; i < len(p.elements); i++ {
		if reflect.DeepEqual(p.elements[i], elem) {
			elemI = i
			break
		}
	}
	p.elements[elemI] = p.elements[len(p.elements)-1]
	p.elements = p.elements[:len(p.elements)-1]
}

// Project is a struct containing the content of a project (displayed) and a list of the required permissions to be fulfilled before this can be displayed.
type Project struct {
	ProjectName         string
	Content             string
	Members             []string
	RequiredPermissions []string
}
