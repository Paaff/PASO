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
	mux.HandleFunc("/data", retrieveBTData)

	err := http.ListenAndServe(":3000", nil)
	if err != nil {
		fmt.Println(err)
	}
}

func retrieveBTData(w http.ResponseWriter, r *http.Request) {
	Counter++
	fmt.Fprintf(w, "Hi there, I love %v!", Counter)
}
