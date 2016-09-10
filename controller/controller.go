package controller

import (
	"fmt"
	"net/http"

	"github.com/jmoiron/sqlx"
)

var (
	db *sqlx.DB
)

// Set the database instance to the database connected to.
func SetDB(conn *sqlx.DB) {
	db = conn
	return
}

// StatusCheck is an unauthenticated route indicating the API has booted and
// has successfully established a DB connection.
func StatusCheck() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "API is ready to serve.\n")
		return
	})
}
