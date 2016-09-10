package logic

import (
	"os/exec"
	"strings"
)

// CurrentUser returns the current user that was set after authenticating. This
// is set for every request.
func (l Lgc) CurrentUser() *User {
	return l.User
}

/*
  genID returns a nicely formatted UUID from the uuidgen tool.
*/
func (Lgc) genID() (string, error) {
	f, err := exec.Command("uuidgen").Output()

	// Trim the resulting whitespace (if any) from the uuidgen command.
	uid := string(f)
	uid = strings.TrimSpace(uid)
	return uid, err
}
