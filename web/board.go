package web

import (
	"fmt"
	"net/http"
)

// InitBoard will start the web server and host index.html.
func InitBoard() {
	http.Handle("/", http.FileServer(http.Dir("/home/hub/freeboard")))
	err := http.ListenAndServe("192.168.0.109:3000", nil)
	if err != nil {
		fmt.Println(err)
	}
}
