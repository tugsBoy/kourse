package api

import (
	"net/http"
)

func main() {
	http.HandleFunc("/workflows/create", handlers.workflowCreateHandler)

	http.ListenAndServe(":5000", nil)
}
