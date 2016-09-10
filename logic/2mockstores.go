package logic

import (
	"database/sql"
)

// User MockDB for injecting into tests; both logic and controller.
type MockDB struct {
	GetMock    func(dest interface{}, q string, args ...interface{}) error
	SelectMock func(dest interface{}, q string, args ...interface{}) error
	ExecMock   func(q string, args ...interface{}) (sql.Result, error)
}

func (m MockDB) Get(dest interface{}, q string, args ...interface{}) error {
	return m.GetMock(dest, q, args...)
}
func (m MockDB) Select(dest interface{}, q string, args ...interface{}) error {
	return m.SelectMock(dest, q, args...)
}
func (m MockDB) Exec(q string, args ...interface{}) (sql.Result, error) {
	return m.ExecMock(q, args...)
}

// MockLgc is for injecting into your controller for testing how the controller
// responds to the logic return values.
type MockLgc struct {
	CurrentUserMock func() *User
}

func (m *MockLgc) CurrentUser() *User {
	return m.CurrentUserMock()
}
