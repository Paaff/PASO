package web

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/paaff/PASO/store"
)

type newClient struct {
	Address  string
	Name     string
	Projects []string
}

// InitBoard will start the web server and host index.html.
func InitBoard() {
	webDir := "./web"
	http.Handle("/", http.FileServer(http.Dir(webDir)))

	mux := http.NewServeMux()
	http.Handle("/api/", http.StripPrefix("/api", mux))
	mux.HandleFunc("/data", RetrieveBTData)
	mux.HandleFunc("/projects", RetrieveProjects)
	mux.HandleFunc("/users", UserHandle)

	err := http.ListenAndServe(":3000", nil)
	if err != nil {
		fmt.Println(err)
	}
}

// RetrieveBTData will provide bluetooth data gathered from the system.
func RetrieveBTData(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(store.CollectedBlueData.GetAsSlice())
}

// RetrieveProjects provides the valid projects based on the currently discovered valid clients
func RetrieveProjects(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(store.Projects.GetValidProjects())
}

// UserHandle - Takes an unknown detected BlueData and adds it as a valid client in the system, together
// with the appropiate permissions chosen.
func UserHandle(w http.ResponseWriter, r *http.Request) {

	if r.Method == "POST" {
		var newInfo newClient
		dec := json.NewDecoder(r.Body)
		err := dec.Decode(&newInfo)
		if err != nil {
			fmt.Print(err)
		}
		fmt.Println(newInfo)

		client := store.Client{
			Name:        newInfo.Name,
			Permissions: []store.Permission{},
		}

		// Go through the given projects and set the view Permission
		for _, p := range newInfo.Projects {
			project, ok := store.Projects.Get(p)
			if ok {
				fmt.Printf("This is the project: %v and this is the project we got %v", p, project)
				client.Permissions = append(client.Permissions, project.GetPerm("View"))
				fmt.Printf("After appending the permission on: %v", client.Permissions)
			}
		}

		store.ValidClientsMap.Set(newInfo.Address, client)

		// Check that the adding of a valid client went okay
		c, ok := store.ValidClientsMap.Get(newInfo.Address)
		if !ok {
			fmt.Fprint(w, http.StatusConflict)
		}
		fmt.Printf("We've set the new client and now we try to get it again and see it: %v", c)
		fmt.Fprint(w, http.StatusCreated)
	}

}
