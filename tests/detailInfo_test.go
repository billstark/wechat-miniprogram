package tests

import (
	"encoding/json"
	"fmt"
	"net/http"
	"testing"
	"wechat-miniprogram/utils/testutils"
)

func TestGetNoDetailInfo(t *testing.T) {
	clearTables()
	req, _ := http.NewRequest("GET", "/records/user/12345u/23456u", nil)
	response := testutils.ExecuteRequest(req, app)
	testutils.CheckResponseCode(t, http.StatusOK, response.Code)

	var result map[string]interface{}
	json.Unmarshal(response.Body.Bytes(), &result)

	hostID := result["host_id"].(string)
	guestID := result["guest_id"].(string)
	records := result["records"].([]interface{})

	if hostID != "12345u" || guestID != "23456u" {
		t.Errorf("Expect host id to be '12345u', guest id to be 23456u, got %s and %s", hostID, guestID)
	}

	if len(records) != 0 {
		t.Errorf("Expect no record, got %s", fmt.Sprintln(records))
	}
}

func clearTables() {
	app.DB.Exec(testutils.RecordTableClear)
	app.DB.Exec(testutils.UserTableClear)
	app.DB.Exec(testutils.GroupTableClear)
}

func InsertUsers() {
	app.DB.Exec(testutils.InsertUserA)
	app.DB.Exec(testutils.InsertUserB)
	app.DB.Exec(testutils.InsertUserC)
}

func InsertGroups() {
	app.DB.Exec(testutils.InsertGroupA)
}

func InsertRecords() {
	app.DB.Exec(testutils.InsertRecordOne)
	app.DB.Exec(testutils.InsertRecordTwo)
	app.DB.Exec(testutils.InsertRecordThree)
}
