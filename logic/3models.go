package logic

import ()

// User holds the information for each user.
type User struct {
	ID        string `db:"id"`
	FirstName string `db:"first_name"`
	LastName  string `db:"last_name"`
	Email     string `db:"email"`
}
