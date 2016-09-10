package main

import (
	// Local packages
	"github.com/fitzy101/api-template/config"
	"github.com/fitzy101/api-template/controller"
	"github.com/fitzy101/api-template/logic"

	// External packages
	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	"github.com/jmoiron/sqlx"

	// Native packages
	"fmt"
	"net/http"
	"time"
)

func main() {
	// Connect to the database firstly.
	db := connDB()

	// Register the database with the controller.
	controller.SetDB(db)

	// Create the logic struct to be passed to the controller.
	ls := logic.Lgc{}

	// We can boot the api now after setting it up, and accept requests.
	server := &http.Server{
		Addr:           ":3000",
		Handler:        NewRouter(db, ls),
		ReadTimeout:    30 * time.Second,
		WriteTimeout:   30 * time.Second,
		MaxHeaderBytes: 1 << 20, // Equivalent to 1 mb.
	}
	fmt.Println("API Template is ready to roll.")
	if err := server.ListenAndServe(); err != nil {
		panic(err)
	}
}

/*
  connDB connects to the database using the configuration values, and returns
  the database instance connection. This will panic if the connection does
  not succeed, so the API will not go live without a database connection.
*/
func connDB() *sqlx.DB {
	fmt.Println("Attempting to connect to DB")
	u := config.DB_USERNAME()
	url := fmt.Sprintf("%v:%v", config.DB_HOST(), config.DB_PORT())
	p := config.DB_PASSWORD()
	n := config.DB_SCHEMA()

	// The API should not boot if the config values were not found.
	if u == "" || url == "" || n == "" || p == "" {
		panic("Database config values not found.")
	}

	// Connect to the database and ping, confirming connection was made.
	// Panic if the connection was not successful.
	connString := fmt.Sprintf("%s:%s@tcp(%s)/%s", u, p, url, n)
	db, err := sqlx.Connect("mysql", connString)
	if err != nil {
		panic(err)
	}
	db.SetMaxIdleConns(10)
	db.SetMaxOpenConns(60)
	fmt.Println("Connected to DB succesfully")
	return db
}

/*
  NewRouter is the function that handles all of the routing for the
  API. Each request is authenticated before serving.
  This returns the http handler for the http server.
*/
func NewRouter(db *sqlx.DB, ls logic.Lgc) http.Handler {
	// Create the handler function to serve the requests, and return
	// to the server.
	return mware(db, ls)
}

// mware is used to route the request depending on the method
// of authentication. Depending on your auth system this will obviously differ.
// This will eventually pass the request to the respective router for
// the api/spa after determining where the request came from.
func mware(db logic.DB, ls logic.Lgc) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		// We need to set the JSON content type otherwise
		// Go will automatically append a text/plain content type
		// header to the request
		w.Header().Set("Content-Type", "application/json")

		// If this is an options header, return before
		// setting up any routers.
		if r.Method == "OPTIONS" {
			// This by default returns a 200.
			return
		}

		// Here is where you would add your authentication steps.
		// Within this is where you might set request context variables,
		// such as the current user.
		var authed bool

		// Assuming the user is authenticated, we can continue with
		// routing the request. If the auth failed or no auth was
		// included, route to the list of authentication exceptions
		// instead.
		if authed {
			router := auth(db, ls)
			router.ServeHTTP(w, r)
		} else {
			// The user is not authenticated. Check the
			// request against the list of exceptions. Return a 401 if
			// the route didn't find a match in the exception list.
			router := exceptions(db, ls)
			router.ServeHTTP(w, r)
		}
	})
}

// auth is a router that assumes the current user is authenticated.
func auth(db logic.DB, ls logic.Lgc) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		// The request is ready to be served.
		router := router(db, ls)
		router.ServeHTTP(w, r)
	})
}

// notFound is returned when the requested url was not matched
// on any of the other routers. This is the generic 404 response.
func notFound() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, "{\"error\":\"Route does not exist - %v %v\"}", r.Method, r.URL.Path)
		return
	})
}

// exceptions is the handler for the list of excepted, unauthorised routes.
func exceptions(db logic.DB, ls logic.Lgc) http.Handler {
	r := mux.NewRouter()
	r.NotFoundHandler = notFound()
	r.Handle("/status", controller.StatusCheck())
	return r
}

// router creates a mux that handles all of the incoming authenticated requests.
func router(db logic.DB, ls logic.Lgc) http.Handler {
	r := mux.NewRouter()
	r.NotFoundHandler = notFound()
	return r
}
