package user

import (
	mStore "wechat-miniprogram/datastore"
	storeErr "wechat-miniprogram/datastore/error"
	converter "wechat-miniprogram/datastore/helper"
	userStoreModels "wechat-miniprogram/datastore/user/storemodels"
	mDB "wechat-miniprogram/utils/database"
	userDBModels "wechat-miniprogram/utils/database/dbModels/user"
)

// Defines related SQL queries
// Currently the table is quite small and simple, therefore no need for update
// Allow user to permenently remove him/herself
const (
	sqlInsertUser      = `INSERT INTO MUser VALUES ($1)`
	sqlSelectUser      = `SELECT * FROM User`
	sqlSelectUserWhere = ` WHERE w_id = $1`
	sqlDeleteUserWhere = `DELETE FROM MUser WHERE w_id = $1`
)

// Store defines user store struct
type Store struct {
	DB mDB.Database
}

// NewUserStore is the constructor of Store
func NewUserStore(db mDB.Database) mStore.Store {
	return Store{db}
}

// Create creates a new user record, return errors if any
func (s Store) Create(args interface{}) (interface{}, error) {

	// try to user converter to convert service model to store model
	converted, err := converter.ServiceToStore(args)
	if err != nil {
		return nil, err
	}

	// Then cast it to correct data type
	casted, ok := converted.(userStoreModels.CreateParams)
	if !ok {
		return nil, storeErr.ErrConverterError
	}

	// execute insertion
	_, err = s.DB.Exec(sqlInsertUser, casted.WechatID)
	if err != nil {
		return nil, err
	}

	return casted.WechatID, nil
}

// Retrieve retrieves users. Since currently it does not make sense getting one user,
// we just assume retrieving all users first
func (s Store) Retrieve(_ interface{}) (interface{}, error) {
	var dbUsers []*userDBModels.Account
	var storeUsers []*userStoreModels.UserRecord
	err := s.DB.SelectMany(&dbUsers, sqlSelectUser)
	if err != nil {
		return nil, err
	}

	for _, dbUser := range dbUsers {
		storeUsers = append(storeUsers, converter.DBUserInfoToStore(*dbUser))
	}
	return storeUsers, nil
}

// Update updates a user records. Currently no need for update
func (s Store) Update(_ string, _ interface{}) (interface{}, error) {
	return nil, nil
}

// Delete deletes a user record.
func (s Store) Delete(id string, _ interface{}) (interface{}, error) {
	_, err := s.DB.Exec(sqlDeleteUserWhere, id)
	if err != nil {
		return nil, err
	}
	return nil, nil
}
