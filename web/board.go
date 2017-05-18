package web

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/paaff/PASO/store"
)

// InitBoard will start the web server and host index.html.
func InitBoard() {
	webDir := "./web"
	http.Handle("/", http.FileServer(http.Dir(webDir)))

	mux := http.NewServeMux()
	http.Handle("/api/", http.StripPrefix("/api", mux))
	mux.HandleFunc("/data", RetrieveBTData)
	mux.HandleFunc("/projects", RetrieveProjects)

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

// AddClient - Takes an unknown detected BlueData and adds it as a valid client in the system, together
// with the appropiate permissions chosen.
func AddClient(w http.ResponseWriter, r *http.Request) {
	// Parse the form data
	err := r.ParseForm()
	if err != nil {
		log.Fatal(err)
	}

	// Use the form data to create a new ValidClient addition
	address := r.FormValue("address")
	name := r.FormValue("name")
	perms := r.Form["permissions"]

	var permissions []store.Permission
	for _, perm := range perms {
		splitted := strings.Split(perm, ",")
		p := splitted[0]
		t := splitted[1]
		permissions = append(permissions, store.Permission{
			Perm: p, PermType: t,
		})
	}

	newClient := store.Client{
		Name:        name,
		Permissions: permissions,
	}

	store.ValidClientsMap.Set(address, newClient)

	// Check that the adding of a valid client went okay
	_, ok := store.ValidClientsMap.Get(address)
	if !ok {
		fmt.Fprint(w, http.StatusConflict)
	}
	fmt.Fprint(w, http.StatusCreated)

}
