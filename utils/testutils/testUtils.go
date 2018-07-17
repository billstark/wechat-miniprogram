package testutils

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
	"wechat-miniprogram/application"
	"wechat-miniprogram/utils/database"
)

// CheckResponseCode checks response code equal
func CheckResponseCode(t *testing.T, expected, actual int) {
	if expected != actual {
		t.Errorf("Expected response code %d. Got %d\n", expected, actual)
	}
}

// ExecuteRequest executes a http request
func ExecuteRequest(req *http.Request, app application.App) *httptest.ResponseRecorder {
	rr := httptest.NewRecorder()
	app.Router.ServeHTTP(rr, req)
	return rr
}

// CreateTables creates all required tables
func CreateTables(db database.Database) {
	var err error
	_, err = db.Exec(CreateUserTable)
	_, err = db.Exec(CreateGroupTable)
	_, err = db.Exec(CreateBelongToTable)
	_, err = db.Exec(CreateRecordTable)
	_, err = db.Exec(CreateHistoryTable)
	if err != nil {
		fmt.Println(err.Error())
	}
}

// DropAllTables Drops all tables in database
func DropAllTables(db database.Database) {
	// var err error
	db.Exec(DropUserTable)
	db.Exec(DropGroupTable)
	db.Exec(DropBelongToTable)
	db.Exec(DropRecordTable)
	db.Exec(DropHistoryTable)
	// if err != nil {
	// 	fmt.Println("No db exists")
	// }
}

// InsertUsers inserts 3 test users
func InsertUsers(db database.Database) {
	var err error
	_, err = db.Exec(InsertUserA)
	_, err = db.Exec(InsertUserB)
	_, err = db.Exec(InsertUserC)
	if err != nil {
		fmt.Println(err.Error())
	}
}

// InsertGroups inserts 1 test group
func InsertGroups(db database.Database) {
	var err error
	_, err = db.Exec(InsertGroupA)
	if err != nil {
		fmt.Println(err.Error())
	}
}

// InsertRecords inserts 3 test records
func InsertRecords(db database.Database) {
	var err error
	_, err = db.Exec(InsertRecordOne)
	_, err = db.Exec(InsertRecordTwo)
	_, err = db.Exec(InsertRecordThree)
	if err != nil {
		fmt.Println(err.Error())
	}
}
