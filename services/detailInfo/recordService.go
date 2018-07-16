package records

import (
	"context"

	"wechat-miniprogram/datastore"
	"wechat-miniprogram/services"

	recordStoreModels "wechat-miniprogram/datastore/record/storemodels"
	"wechat-miniprogram/services/detailInfo/servicemodels"

	serviceErr "wechat-miniprogram/services/errors"
	"wechat-miniprogram/services/helper"
)

// DetailInfoService is service for Detailed info
// Used for two situations:
// 1. records page between users
// 2. records page inside a group
type DetailInfoService struct {
	// GroupStore  datastore.Store
	RecordStore datastore.Store
}

// NewDetailInfoService is the constructor
func NewDetailInfoService(recordStore datastore.Store) services.Service {
	return DetailInfoService{recordStore}
}

// Retrieve retrieves detailed infos
// compulsary param: host_id
// possible params:
// - guest_id
// - group_id
func (s DetailInfoService) Retrieve(_ context.Context, args interface{}) (interface{}, error) {
	infoRetrieveParams, ok := args.(servicemodels.DetailRetrieveParams)
	if !ok {
		return nil, serviceErr.ErrIncorrectParamsFormat
	}

	// If has group id, need to get group info
	// Else first
	records, err := s.RecordStore.Retrieve(infoRetrieveParams)
	if err != nil {
		return nil, err
	}
	castedRecords := records.([]*recordStoreModels.RecordRetrieveResult)
	return helper.GenerateDetailBetweenUsers(castedRecords, infoRetrieveParams.HostID, infoRetrieveParams.GuestID)
}

// Create not in used yet
func (s DetailInfoService) Create(ctx context.Context, args interface{}) (interface{}, error) {
	return nil, nil
}

// Update not in used yet
func (s DetailInfoService) Update(ctx context.Context, args interface{}) (interface{}, error) {
	return nil, nil
}

// Delete not in used yet
func (s DetailInfoService) Delete(ctx context.Context, args interface{}) (interface{}, error) {
	return nil, nil
}
