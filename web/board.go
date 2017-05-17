package web

import (
	"encoding/json"
	"fmt"
	"net/http"

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

	r.Form
}
