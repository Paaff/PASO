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
			fmt.Println("shit went down")
		}

		client := store.Client{
			Name:        newInfo.Name,
			Permissions: []store.Permission{},
		}

		// Go through the given projects and set the view Permission
		for _, p := range newInfo.Projects {
			project, ok := store.Projects.Get(p)
			if ok {
				client.Permissions = append(client.Permissions, project.GetPerm("View"))
			}
		}

		store.ValidClientsMap.Set(newInfo.Address, client)

		// Check that the adding of a valid client went okay
		_, ok := store.ValidClientsMap.Get(newInfo.Address)
		if !ok {
			fmt.Fprint(w, http.StatusConflict)
		}
		fmt.Fprint(w, http.StatusCreated)
	}

}
