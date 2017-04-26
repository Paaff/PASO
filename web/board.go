package web

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// InitBoard will start the web server and host index.html.
func InitBoard() {
	webDir := "./web"
	http.Handle("/", http.FileServer(http.Dir(webDir)))

	mux := http.NewServeMux()
	http.Handle("/api/", http.StripPrefix("/api", mux))
	mux.HandleFunc("/data", RetrieveBTData)

	err := http.ListenAndServe(":3000", nil)
	if err != nil {
		fmt.Println(err)
	}
}

type s struct {
	Name      string
	Address   string
	Class     string
	Timestamp string
}

// RetrieveBTData will provide bluetooth data gathered from the system.
func RetrieveBTData(w http.ResponseWriter, r *http.Request) {
	test := s{"testName", "testAddress", "testClass", "testTime"}
	json.NewEncoder(w).Encode(test)

}
