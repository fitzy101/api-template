package controller

import (
	"fmt"
	"net/http"
)

// StatusCheck is an unauthenticated route indicating the API has booted and
// has successfully established a DB connection.
func StatusCheck() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "API is ready to serve.\n")
		return
	})
}
