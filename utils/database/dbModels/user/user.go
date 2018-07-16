package user

// Account wraps user table from db
type Account struct {
	ID       int    `db:"id"`
	WechatID string `db:"w_id"`
}
