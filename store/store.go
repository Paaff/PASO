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

// Permission contains the permission and the type of permission "View" or "Open"
type Permission struct {
	Perm     string
	PermType string
}

// Client is a struct for the predetermined users in the system.
type Client struct {
	Name        string
	Permissions []Permission
}

// ContainsPerm will check the clients permissions list and return true if a matching permission is found.
func (c *Client) ContainsPerm(perm Permission) bool {
	for _, p := range c.Permissions {
		if reflect.DeepEqual(p, perm) {
			return true
		}
	}
	return false
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

	// Get the index of the desired element
	var elemI int
	for i := 0; i < len(p.elements); i++ {
		if reflect.DeepEqual(p.elements[i], elem) {
			elemI = i
			break
		}
	}
	// As we dont care about ordering we can simply take the last element in the slice and replace
	// the desired element.
	p.elements[elemI] = p.elements[len(p.elements)-1]
	p.elements = p.elements[:len(p.elements)-1]
}

// GetValidProjects provides the projects in which all the clients are fulfilling the permissions
func (p *ProjectsList) GetValidProjects() []Project {
	validProjects := make([]Project, 0)

	currentDetected := CollectedBlueData.GetAsSlice()
	if len(currentDetected) == 0 {
		return validProjects
	}

	for _, project := range p.elements {
		ok := permsFulfilled(project, currentDetected)
		if ok {
			validProjects = append(validProjects, project)
		}
	}
	return validProjects
}

func permsFulfilled(project Project, currDetected []BlueData) bool {
	for _, perm := range project.RequiredPermissions {
		if perm.PermType == "Open" {
			if ok := singleFulfilled(perm, currDetected); ok != true {
				return false
			}
		} else if perm.PermType == "View" {
			if ok := allFulfilled(perm, currDetected); ok != true {
				return false
			}
		}
	}
	return false
}

func singleFulfilled(perm Permission, currDetected []BlueData) bool {
	for _, blueData := range currDetected {
		client, okClient := ValidClientsMap.Get(blueData.Address)
		okPerm := client.ContainsPerm(perm)
		if okClient && okPerm {
			return true
		}
	}
	return false
}

func allFulfilled(perm Permission, currDetected []BlueData) bool {
	for _, blueData := range currDetected {
		client, okClient := ValidClientsMap.Get(blueData.Address)
		okPerm := client.ContainsPerm(perm)
		if okClient && !okPerm {
			return false
		}
	}
	return true
}

// Project is a struct containing the content of a project (displayed) and a list of the required permissions to be fulfilled before this can be displayed.
type Project struct {
	ProjectName         string
	Content             string
	Members             []string
	RequiredPermissions []Permission
}
