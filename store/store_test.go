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

/*******************
 Testing ProjectList
*******************/

func TestAddToAndRemoveFromProjectList(t *testing.T) {
	testingProjects := ProjectsList{}
	if len(testingProjects.elements) != 0 {
		t.Errorf("The list length should be 0, as we havent added anything yet. List: %v", testingProjects.elements)
	}
	p1 := Project{"Project1", ".", []string{"member1"}, []Permission{Permission{"perm1", "permtype1"}}}
	p2 := Project{"Project2", ".", []string{"member2"}, []Permission{Permission{"perm2", "permtype2"}}}

	testingProjects.Add(p1)
	testingProjects.Add(p2)
	if len(testingProjects.elements) != 2 {
		t.Errorf("The list length should be 2, as we added two Project structs to it. List: %v", testingProjects.elements)
	}

	contains := testingProjects.Contains(p1)
	if !contains {
		t.Errorf("It should contain: %v,\n the projectslist: %v", p1, testingProjects.elements)
	}

	ok := testingProjects.Remove(p1)
	if ok != true {
		t.Errorf("Project 1 should've been removed. ok: %v", ok)
	}

	contains = testingProjects.Contains(p1)
	if contains != false {
		t.Errorf("It should no longer contain: %v", p1)
	}

	ok = testingProjects.Remove(p2)
	if ok != true {
		t.Errorf("Project 2 should've been removed. ok: %v", ok)
	}
	contains = testingProjects.Contains(p2)
	if contains != false {
		t.Errorf("It should no longer contain: %v", p2)
	}

	p3 := Project{"Project3", ".", []string{"member3"}, []Permission{Permission{"perm3", "permtype3"}}}
	ok = testingProjects.Remove(p3)
	if ok != false {
		t.Errorf("Project 3 cant be removed been removed, because it is not added, ok: %v", ok)
	}

	contains = testingProjects.Contains(p3)
	if contains != false {
		t.Errorf("It never contained: %v", p3)
	}
}

func TestGetValidProjects(t *testing.T) {
	InitDB()

	// No data should be in CollectedBlueData yet.
	shouldBeEmpty := Projects.GetValidProjects()

	if len(shouldBeEmpty) != 0 {
		t.Errorf("As there are no CollectedBlueData yet, the valid projects should be empty. Projects: %v", shouldBeEmpty)
	}

	projectA := Project{
		ProjectName:         "Project A",
		Content:             "This is the information regarding project A and its content.",
		Members:             []string{"Peter Fischer", "Mathias Mortensen"},
		RequiredPermissions: []Permission{Permission{Perm: "ViewA", PermType: "View"}, Permission{Perm: "OpenA", PermType: "Open"}},
	}

	projectC := Project{
		ProjectName:         "Project C",
		Content:             "This is the information regarding project C and its content.",
		Members:             []string{"Peter Fischer"},
		RequiredPermissions: []Permission{Permission{Perm: "ViewC", PermType: "View"}, Permission{Perm: "OpenC", PermType: "Open"}},
	}

	expected := []Project{projectA, projectC}

	// Adding BlueData
	CollectedBlueData.Set("24:DA:9B:BB:EE:2B", BlueData{"24:DA:9B:BB:EE:2B", "Smartphone", "Time 12345"})

	actual := Projects.GetValidProjects()

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("Expected Projectslist is not as the actual. Expected: %v\n, Actual: %v\n", expected, actual)
	}

	// Removing projectC from expected and adding Mathias Mortensen to CollectedBlueData

	expected = []Project{projectA}

	CollectedBlueData.Set("54:9B:12:D2:09:4C", BlueData{"54:9B:12:D2:09:4C", "Smartphone", "Time Mathias"})

	actual = Projects.GetValidProjects()

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("Expected Projectslist is not as the actual. \nExpected: %v\n, Actual: %v\n", expected, actual)
	}

}
