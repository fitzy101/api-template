package logic

import (
	"database/sql"
)

// DB is required when interfacing with the database.
type DB interface {
	Get(dest interface{}, q string, args ...interface{}) error
	Select(dest interface{}, q string, args ...interface{}) error
	Exec(q string, args ...interface{}) (sql.Result, error)
}

// Utilities is an interface that is used to store the utility functions within
// the logic layer. Having this as an interface allows you to test the logic
// layer when calls to utility fail/pass/return certain values.
type Utilities interface {
}

// Lgc is going to hold all of the functions that are called from the controller,
// and functions that need to be unit tested.
type Lgc struct {
	User *User // This will be the user that has been safely authenticated.
}
