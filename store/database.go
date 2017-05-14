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
	ValidClientsMap.Set("24:DA:9B:BB:EE:2B", Client{"Peter Fischer", []string{"PermOpenA, PermViewA"}})
}

func populateProjects() {
	Projects.Add(Project{ProjectName: "Project A", Content: "This is the information regarding project A and its content.",
		RequiredPermissions: []string{"PermOpenA, PermViewA"},
	})
	Projects.Add(Project{ProjectName: "Project B", Content: "This is the information regarding project B and its content.",
		RequiredPermissions: []string{"PermOpenB, PermViewB"},
	})
	Projects.Add(Project{ProjectName: "Project C", Content: "This is the information regarding project C and its content.",
		RequiredPermissions: []string{"PermOpenC, PermViewC"},
	})
}
