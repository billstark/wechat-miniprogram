package helper

import (
	storeErr "wechat-miniprogram/datastore/error"

	detailInfoServiceModels "wechat-miniprogram/services/detailInfo/servicemodels"

	recordStoreModels "wechat-miniprogram/datastore/record/storemodels"
	userStoreModels "wechat-miniprogram/datastore/user/storemodels"
	recordDBModels "wechat-miniprogram/utils/database/dbModels/record"
	userDBModels "wechat-miniprogram/utils/database/dbModels/user"
)

// ServiceToStore converts service models to store models
func ServiceToStore(serviceModel interface{}) (interface{}, error) {
	switch serviceModel.(type) {
	case detailInfoServiceModels.DetailRetrieveParams:
		recordRetrieval := serviceModel.(detailInfoServiceModels.DetailRetrieveParams)
		return recordRetrieveServiceToStore(recordRetrieval), nil
	default:
		return nil, storeErr.ErrUnrecognizedServiceModel
	}
}

func recordRetrieveServiceToStore(serviceModel detailInfoServiceModels.DetailRetrieveParams) recordStoreModels.RecordRetrieveParams {
	return recordStoreModels.RecordRetrieveParams{
		HostID:  serviceModel.HostID,
		GuestID: serviceModel.GuestID,
		GroupID: serviceModel.GroupID,
	}
}

// DBRecordsToStore converts db record models to store record models
func DBRecordsToStore(dbRecords []recordDBModels.TansRecord) []*recordStoreModels.RecordRetrieveResult {
	result := make([]*recordStoreModels.RecordRetrieveResult, 0)
	for _, dbRecord := range dbRecords {
		temp := recordStoreModels.RecordRetrieveResult{
			RecordID:    dbRecord.RecordID,
			GroupID:     dbRecord.GroupID,
			Date:        dbRecord.Date,
			Payer:       dbRecord.Payer,
			Spliters:    dbRecord.Spliters,
			Amount:      dbRecord.Amount,
			Description: dbRecord.Description,
		}
		result = append(result, &temp)
	}
	return result
}

// DBUserInfoToStore converts db user models to store user models
func DBUserInfoToStore(dbUser userDBModels.Account) *userStoreModels.UserRecord {
	return &userStoreModels.UserRecord{
		WechatID: dbUser.WechatID,
	}
}
