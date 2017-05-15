package store

import (
	"reflect"
	"testing"
)

/*******************
 Testing BlueDataMap
*******************/

func TestNewBlueDataMap(t *testing.T) {
	// Create the struct. The items fields should be nil then.
	testingMap := BlueDataMap{}

	if len(testingMap.items) != 0 {
		t.Errorf("Items field in the BlueDataMap struct should be empty when initializing this way. item: %v \n", testingMap.items)
	}

	testingMap.NewBlueDataMap()
	testingMap.Set("Test", BlueData{})
	if len(testingMap.items) == 0 {
		t.Error("The NewBlueDataMap method should have initialized the items field.\n")
	}
}

func TestSetBlueDataMap(t *testing.T) {
	testingMap := BlueDataMap{}
	testingMap.NewBlueDataMap()
	v, ok := testingMap.Get("test")
	if ok != false && !reflect.DeepEqual(v, BlueData{}) {
		t.Errorf("There should be no key-value pair with key test in the map. Map: %v, value: %v", testingMap.items, v)
	}

	testingMap.Set("test", BlueData{"Address", "Class", "Time"})
	v, ok = testingMap.Get("test")
	if ok != true && reflect.DeepEqual(v, BlueData{"Address", "Class", "Time"}) {
		t.Errorf("There should've been a key, bool: %v, and a value: %v in the map: %v", ok, v, testingMap.items)
	}

}

func TestBlueDataMapGetAsSlice(t *testing.T) {
	testingMap := BlueDataMap{}
	testingMap.NewBlueDataMap()

	slice := testingMap.GetAsSlice()
	if len(slice) != 0 {
		t.Errorf("The slice should be empty, slice: %v", slice)
	}

	testingMap.Set("test1", BlueData{"Address1", "Class1", "Time1"})
	testingMap.Set("test2", BlueData{"Address2", "Class2", "Time2"})
	testingMap.Set("test3", BlueData{"Address3", "Class3", "Time3"})

	slice = testingMap.GetAsSlice()
	if len(slice) != 3 {
		t.Errorf("The slice should be of length 3 as we set 3 different keys and BlueData structs, slice: %v", slice)
	}
}

/*******************
 Testing ClientsMap
*******************/

func TestNewClientMap(t *testing.T) {
	testingMap := ClientsMap{}
	if len(testingMap.items) != 0 {
		t.Errorf("Items field in the ClientsMap struct should be empty when initializing this way. item: %v \n", testingMap.items)
	}

	testingMap.NewClientsMap()
	testingMap.Set("Test", Client{})
	if len(testingMap.items) == 0 {
		t.Error("The NewClientsMap method should have initialized the items field.\n")
	}
}

func TestSetAndGetClientsMap(t *testing.T) {
	testingMap := ClientsMap{}
	testingMap.NewClientsMap()
	v, ok := testingMap.Get("test")
	if ok != false && !reflect.DeepEqual(v, Client{}) {
		t.Errorf("There should be no key-value pair with key test in the map. Map: %v, value: %v", testingMap.items, v)
	}

	testingMap.Set("test", Client{"Name", []Permission{Permission{}}})
	v, ok = testingMap.Get("test")
	if ok != true && reflect.DeepEqual(v, BlueData{"Address", "Class", "Time"}) {
		t.Errorf("There should've been a key, bool: %v, and a value: %v in the map: %v", ok, v, testingMap.items)
	}
}

/*******************
 Testing Client
*******************/

func TestClientContainsPerm(t *testing.T) {
	perm1 := Permission{"perm1", "type1"}
	perm2 := Permission{"perm2", "type2"}
	perm3 := Permission{"perm3", "type1"}

	testClient := Client{"TestName1", []Permission{}}
	ok := testClient.ContainsPerm(perm1)
	if ok != false {
		t.Errorf("the perm should not exist in the client struct right now, ok: %v", ok)
	}

	testClient = Client{"TestName2", []Permission{perm1, perm2, perm3}}

	ok = testClient.ContainsPerm(perm1)
	if ok != true {
		t.Errorf("Perm1 should've been found and bool: %v, should be true", ok)
	}
}
