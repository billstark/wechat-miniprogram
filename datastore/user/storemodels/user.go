package storemodels

// CreateParams wraps all the params required to create
// a new user
type CreateParams struct {
	WechatID string
}

// UserRecord wraps all the info for a user record
type UserRecord struct {
	WechatID string
}
