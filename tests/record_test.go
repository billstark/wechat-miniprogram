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

func TestGetDetailedInfo(t *testing.T) {
	clearTables()
	insertUsers()
	insertGroups()
	insertRecords()
	req, _ := http.NewRequest("GET", "/records/user/12345u/23456u", nil)
	response := testutils.ExecuteRequest(req, app)
	testutils.CheckResponseCode(t, http.StatusOK, response.Code)

	var result map[string]interface{}
	json.Unmarshal(response.Body.Bytes(), &result)

	fmt.Println(result)

	hostID := result["host_id"].(string)
	guestID := result["guest_id"].(string)
	records := result["records"].([]interface{})

	if hostID != "12345u" || guestID != "23456u" {
		t.Errorf("Expect host id to be '12345u', guest id to be 23456u, got %s and %s", hostID, guestID)
	}

	if len(records) == 0 {
		t.Errorf("Expect to have records, got 0 record")
	}

	var recordOne map[string]interface{}
	marshedRecordOne, _ := json.Marshal(records[0])
	json.Unmarshal(marshedRecordOne, &recordOne)

	var recordTwo map[string]interface{}
	marshedRecordTwo, _ := json.Marshal(records[1])
	json.Unmarshal(marshedRecordTwo, &recordTwo)

	if int(recordOne["id"].(float64)) != 1 ||
		int(recordOne["group_id"].(float64)) != 1 ||
		recordOne["date"].(string) != "2018-07-14T00:00:00Z" ||
		recordOne["amount"].(float64) != 50 ||
		recordOne["description"].(string) != "dinner" {
		t.Errorf("Expect record to be id: 1, group_id: 1, date: 2018-07-14T00:00:00Z, amount: 50, description: dinner, got %d, %d, %s, %f, %s", int(recordOne["id"].(float64)), int(recordOne["group_id"].(float64)), recordOne["date"].(string), recordOne["amount"].(float64), recordOne["description"].(string))
	}

	if int(recordTwo["id"].(float64)) != 3 ||
		int(recordTwo["group_id"].(float64)) != 1 ||
		recordTwo["date"].(string) != "2016-07-14T00:00:00Z" ||
		recordTwo["amount"].(float64) != -30 ||
		recordTwo["description"].(string) != "settle" {
		t.Errorf("Expect record to be id: 3, group_id: 1, date: 2016-07-14T00:00:00Z, amount: -30, description: settle, got %d, %d, %s, %f, %s", int(recordTwo["id"].(float64)), int(recordTwo["group_id"].(float64)), recordTwo["date"].(string), recordTwo["amount"].(float64), recordTwo["description"].(string))
	}
}

func clearTables() {
	app.DB.Exec(testutils.RecordTableClear)
	app.DB.Exec(testutils.UserTableClear)
	app.DB.Exec(testutils.GroupTableClear)
	app.DB.Exec(testutils.ResetUserSeq)
	app.DB.Exec(testutils.ResetGroupSeq)
	app.DB.Exec(testutils.ResetRecordSeq)
}

func insertUsers() {
	app.DB.Exec(testutils.InsertUserA)
	app.DB.Exec(testutils.InsertUserB)
	app.DB.Exec(testutils.InsertUserC)
}

func insertGroups() {
	app.DB.Exec(testutils.InsertGroupA)
}

func insertRecords() {
	app.DB.Exec(testutils.InsertRecordOne)
	app.DB.Exec(testutils.InsertRecordTwo)
	app.DB.Exec(testutils.InsertRecordThree)
}
