package record

import "time"

// TansRecord wrapps all db columns
type TansRecord struct {
	RecordID    int       `db:"id"`
	GroupID     int       `db:"g_id"`
	Date        time.Time `db:"day"`
	Payer       string    `db:"payer"`
	Spliters    []string  `db:"spliters"`
	Amount      float32   `db:"pay_amount"`
	Description string    `db:"description"`
	UpdatedAt   time.Time `db:"updated_at"`
	DeletedAt   time.Time `db:"deleted_at"`
}
