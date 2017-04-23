package web

import (
	"fmt"
	"net/http"
)

// Counter test
var Counter int

// InitBoard will start the web server and host index.html.
func InitBoard() {
	Counter = 0
	freeboard := "/home/hub/freeboard"
	http.Handle("/", http.FileServer(http.Dir(freeboard)))

	mux := http.NewServeMux()
	http.Handle("/api/", http.StripPrefix("/api", mux))
	mux.HandleFunc("/data", RetrieveBTData)

	err := http.ListenAndServe("192.168.0.109:3000", nil)
	if err != nil {
		fmt.Println(err)
	}
}

// RetrieveBTData will provide bluetooth data gathered from the system.
func RetrieveBTData(w http.ResponseWriter, r *http.Request) {
	Counter++
	fmt.Fprintf(w, "This is the %v time the endpoint has been visited", Counter)
}
