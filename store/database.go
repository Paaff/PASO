package store

// InitDB - Initializes the Demo Database.
func InitDB() {
	ValidClientsMap = ValidClients{}
	ValidClientsMap.NewValidClientsMap()
	populateValidClients()

	CollectedBlueData = ClientDataMap{}
	CollectedBlueData.NewCollectedClientsMap()

	Projects = ProjectsList{}
	Projects.NewProjectsList()
	populateProjects()
}

func populateValidClients() {
	ValidClientsMap.Set("24:DA:9B:BB:EE:2B", Client{"Peter Fischer", []string{"PermOpenA", "PermViewA", "PermOpenC", "PermViewC"}})
	ValidClientsMap.Set("20:7D:74:0B:B6:BB", Client{"Alberto Lafuente", []string{""}})
	ValidClientsMap.Set("54:9B:12:D2:09:4C", Client{"Mathias Mortensen", []string{"PermViewB", "PermViewA"}})
}

func populateProjects() {
	Projects.Add(Project{ProjectName: "Project A", Content: "This is the information regarding project A and its content.",
		Members: []string{"Peter Fischer", "Mathias Mortensen"}, RequiredPermissions: []string{"PermOpenA", "PermViewA"},
	})
	Projects.Add(Project{ProjectName: "Project B", Content: "This is the information regarding project B and its content.",
		Members: []string{"Mathias Mortensen", "Alberto Lafuente"}, RequiredPermissions: []string{"PermOpenB", "PermViewB"},
	})
	Projects.Add(Project{ProjectName: "Project C", Content: "This is the information regarding project C and its content.",
		Members: []string{"Peter Fischer"}, RequiredPermissions: []string{"PermOpenC", "PermViewC"},
	})
}
