package store

import (
	"sync"
)

// CollectedClients holds all the current clients detected
var CollectedClients ClientDataSlice

// ClientDataSlice is a Global slice designed to hold the Bluedata items.
type ClientDataSlice struct {
	sync.RWMutex
	items []BlueData
}

// ClientDataItem is of type BlueData
type ClientDataItem struct {
	Index int
	Value BlueData
}

// Append function will acquire a lock on the slice, append and release the lock.
func (cs *ClientDataSlice) Append(item BlueData) {
	cs.Lock()
	defer cs.Unlock()
	cs.items = append(cs.items, item)
}

// ReadHej function will acquire a read lock and return the slice
func (cs *ClientDataSlice) Read() []BlueData {
	cs.RLock()
	defer cs.RUnlock()
	return cs.items
}

// BlueData - Bluetooth data.
type BlueData struct {
	Name      string
	Bdaddress string
	Class     string
	Timestamp string
}
