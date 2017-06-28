package store

// InitDB - Initializes the Demo Database.
func InitDB() {
	ValidClientsMap = ClientsMap{}
	ValidClientsMap.NewClientsMap()
	//populateValidClients()
	demoValidClients()
	CollectedBlueData = BlueDataMap{}
	CollectedBlueData.NewBlueDataMap()

	Projects = ProjectsList{}
	demoProjects()
	//populateProjects()
}
func demoValidClients() {
	// Scenario 1
	ValidClientsMap.Set("Alice1_ID", Client{"Alice",
		[]Permission{
			Permission{Perm: "scenario1_ViewA", PermType: "View"},
			Permission{Perm: "scenario1_OpenA", PermType: "Open"},
			Permission{Perm: "scenario1_ViewB", PermType: "View"},
			Permission{Perm: "scenario1_OpenB", PermType: "Open"}}})

	ValidClientsMap.Set("Bob1_ID", Client{"Bob",
		[]Permission{
			Permission{Perm: "scenario1_OpenB", PermType: "Open"},
			Permission{Perm: "scenario1_ViewB", PermType: "View"}}})

	// Scenario 2
	ValidClientsMap.Set("Alice2_ID", Client{"Alice",
		[]Permission{
			Permission{Perm: "scenario2_ViewA", PermType: "View"},
			Permission{Perm: "scenario2_OpenA", PermType: "Open"}}})

	ValidClientsMap.Set("Bob2_ID", Client{"Bob",
		[]Permission{
			Permission{Perm: "scenario2_ViewA", PermType: "View"}}})

	// Scenario 3
	ValidClientsMap.Set("Alice3_ID", Client{"Alice",
		[]Permission{
			Permission{Perm: "scenario3_ViewA", PermType: "View"},
			Permission{Perm: "scenario3_OpenA", PermType: "Open"},
			Permission{Perm: "scenario3_ViewB", PermType: "View"},
			Permission{Perm: "scenario3_OpenB", PermType: "Open"}}})

	ValidClientsMap.Set("Bob3_ID", Client{"Bob",
		[]Permission{
			Permission{Perm: "scenario3_ViewA", PermType: "View"},
			Permission{Perm: "scenario3_OpenA", PermType: "Open"},
			Permission{Perm: "scenario3_ViewB", PermType: "View"},
			Permission{Perm: "scenario3_OpenB", PermType: "Open"}}})
}

func demoProjects() {

	// Scenario 1
	scenario1ProjectA := Project{
		ProjectName:         "Project A",
		Content:             "This is the information regarding project A and its content. (User Scenario 1)",
		Members:             []string{"Alice"},
		RequiredPermissions: []Permission{Permission{Perm: "scenario1_ViewA", PermType: "View"}, Permission{Perm: "scenario1_OpenA", PermType: "Open"}},
	}
	scenario1ProjectB := Project{
		ProjectName:         "Project B",
		Content:             "This is the information regarding project B and its content. (User Scenario 1)",
		Members:             []string{"Alice", "Bob"},
		RequiredPermissions: []Permission{Permission{Perm: "scenario1_ViewB", PermType: "View"}, Permission{Perm: "scenario1_OpenB", PermType: "Open"}},
	}

	Projects.Add(scenario1ProjectA)
	Projects.Add(scenario1ProjectB)

	// Scenario 2
	scenario2ProjectA := Project{
		ProjectName:         "Project A",
		Content:             "This is the information regarding project A and its content. (User Scenario 2)",
		Members:             []string{"Alice", "Bob"},
		RequiredPermissions: []Permission{Permission{Perm: "scenario2_ViewA", PermType: "View"}, Permission{Perm: "scenario2_OpenA", PermType: "Open"}},
	}
	Projects.Add(scenario2ProjectA)

	// Scenario 3
	scenario3ProjectA := Project{
		ProjectName:         "Project A",
		Content:             "This is the information regarding project A and its content.",
		Members:             []string{"Alice", "Bob"},
		RequiredPermissions: []Permission{Permission{Perm: "scenario3_ViewA", PermType: "View"}, Permission{Perm: "scenario3_OpenA", PermType: "Open"}},
	}
	scenario3ProjectB := Project{
		ProjectName:         "Project B",
		Content:             "This is the information regarding project B and its content.",
		Members:             []string{"Alice", "Bob"},
		RequiredPermissions: []Permission{Permission{Perm: "scenario3_ViewB", PermType: "View"}, Permission{Perm: "scenario3_OpenB", PermType: "Open"}},
	}

	Projects.Add(scenario3ProjectA)
	Projects.Add(scenario3ProjectB)

}

func populateValidClients() {
	ValidClientsMap.Set("24:DA:9B:BB:EE:2B", Client{"Peter Fischer",
		[]Permission{
			Permission{Perm: "ViewA", PermType: "View"},
			Permission{Perm: "OpenA", PermType: "Open"},
			Permission{Perm: "ViewC", PermType: "View"},
			Permission{Perm: "OpenC", PermType: "Open"}}})

	ValidClientsMap.Set("20:7D:74:0B:B6:BB",
		Client{
			"Alberto Lafuente",
			[]Permission{Permission{}}})

	ValidClientsMap.Set("54:9B:12:D2:09:4C",
		Client{
			"Mathias Mortensen",
			[]Permission{
				Permission{Perm: "ViewA", PermType: "View"},
				Permission{Perm: "ViewB", PermType: "View"}}})
}

func populateProjects() {
	Projects.Add(Project{
		ProjectName:         "Project A",
		Content:             "This is the information regarding project A and its content.",
		Members:             []string{"Peter Fischer", "Mathias Mortensen"},
		RequiredPermissions: []Permission{Permission{Perm: "ViewA", PermType: "View"}, Permission{Perm: "OpenA", PermType: "Open"}},
	})
	Projects.Add(Project{
		ProjectName:         "Project B",
		Content:             "This is the information regarding project B and its content.",
		Members:             []string{"Mathias Mortensen", "Alberto Lafuente"},
		RequiredPermissions: []Permission{Permission{Perm: "ViewB", PermType: "View"}, Permission{Perm: "OpenB", PermType: "Open"}},
	})
	Projects.Add(Project{
		ProjectName:         "Project C",
		Content:             "This is the information regarding project C and its content.",
		Members:             []string{"Peter Fischer"},
		RequiredPermissions: []Permission{Permission{Perm: "ViewC", PermType: "View"}, Permission{Perm: "OpenC", PermType: "Open"}},
	})
}
