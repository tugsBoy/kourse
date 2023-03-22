package api

import (
	"io"
	"log"
	"net/http"
)

func HandleRequests() {
	route1 := func(w http.ResponseWriter, _ *http.Request) {
		io.WriteString(w, "Welcome to your workflow!\n")
	}
	http.HandleFunc("/", route1)
	log.Fatal(http.ListenAndServe(":8888", nil))
}
