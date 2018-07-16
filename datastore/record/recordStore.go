package record

import (
	mStore "wechat-miniprogram/datastore"
	storeErr "wechat-miniprogram/datastore/error"
	helper "wechat-miniprogram/datastore/helper"
	recordStoreModels "wechat-miniprogram/datastore/record/storemodels"
	mDB "wechat-miniprogram/utils/database"
	recordDBModels "wechat-miniprogram/utils/database/dbModels/record"
)

const (
	nonGroupID           = -1
	nonUserID            = ""
	sqlSelectRecord      = `SELECT * FROM record`
	sqlWhereUserRelated  = ` WHERE payer = $1 OR $1=ANY(spliters)`
	sqlWhereGroupRelated = ` WHERE payer = $1 OR $1=ANY(spliters) AND g_id = $2`
	sqlWhereBetweenUsers = ` WHERE (payer = $1 AND $2=ANY(spliters)) OR (payer = $2 AND $1=ANY(spliters))`
	sqlCreateGroupRecord = `INSERT INTO record
                            (g_id, day, payer, spliters, pay_amount, description, updated_at, deleted_at)
                            VALUES($1, $2, $3, $4, $5, $6, $7, $8)`
	sqlCreateNoGroup = `INSERT INTO record
                            (day, payer, spliters, pay_amount, description, updated_at, deleted_at)
                            VALUES($1, $2, $3, $4, $5, $6, $7)`
)

// Store is an struct for record basic crud operations
type Store struct {
	DB mDB.Database
}

// NewRecordStore creates a new record store instance
func NewRecordStore(db mDB.Database) mStore.Store {
	return Store{db}
}

// Create creates a record. all arguments are included in args
func (s Store) Create(args interface{}) (interface{}, error) {
	return nil, nil
}

// Retrieve retrieves record(s).
func (s Store) Retrieve(args interface{}) (interface{}, error) {
	converted, err := helper.ServiceToStore(args)
	if err != nil {
		return nil, err
	}
	retrieveParams := converted.(recordStoreModels.RecordRetrieveParams)

	if retrieveParams.HostID == nonUserID {
		return nil, storeErr.ErrInvalidQuery
	}

	// When no group and guest id is provided (records for a user)
	if retrieveParams.GroupID == nonGroupID && retrieveParams.GuestID == nonUserID {
		return retrieveIndividualRecords(s.DB, retrieveParams.HostID)
	}

	// When group specified but no guest id (records for a user in a specific group)
	if retrieveParams.GroupID != nonGroupID && retrieveParams.GuestID == nonUserID {
		return retrieveInGroupRecords(s.DB, retrieveParams.HostID, retrieveParams.GroupID)
	}

	// When two user ids are specified (records between two specific users)
	if retrieveParams.GuestID != nonUserID {
		return retrieveBetweenUserRecords(s.DB, retrieveParams.HostID, retrieveParams.GuestID)
	}

	return nil, storeErr.ErrNotSupportedQuery
}

// Update updates a record. id should be specified
func (s Store) Update(id string, args interface{}) (interface{}, error) {
	return nil, nil
}

// Delete deletes a record. id should be specified
func (s Store) Delete(id string, args interface{}) (interface{}, error) {
	return nil, nil
}

// ================
// Helper functions
// ================

// Retrieves records that are related to a specific individual
// Returns an array of references to the records
func retrieveIndividualRecords(db mDB.Database, userID string) (interface{}, error) {
	var dbResult []recordDBModels.TansRecord
	err := db.SelectMany(&dbResult, sqlSelectRecord+sqlWhereUserRelated, userID)
	if err != nil {
		return nil, storeErr.Wrapper(err)
	}
	return helper.DBRecordsToStore(dbResult), nil
}

// Retrieves records that are related to a specific individual in a specific group
// Returns an array of references to the records
func retrieveInGroupRecords(db mDB.Database, userID string, groupID int) (interface{}, error) {
	var dbResult []recordDBModels.TansRecord
	err := db.SelectMany(&dbResult, sqlSelectRecord+sqlWhereGroupRelated, userID, groupID)
	if err != nil {
		return nil, storeErr.Wrapper(err)
	}
	return helper.DBRecordsToStore(dbResult), nil
}

// Retrieves records that are between two specific users
// Returns an array of references to the records
func retrieveBetweenUserRecords(db mDB.Database, hostID string, guestID string) (interface{}, error) {
	var dbResult []recordDBModels.TansRecord
	err := db.SelectMany(&dbResult, sqlSelectRecord+sqlWhereBetweenUsers, hostID, guestID)
	if err != nil {
		return nil, storeErr.Wrapper(err)
	}
	return helper.DBRecordsToStore(dbResult), nil
}
