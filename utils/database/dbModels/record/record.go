package record

import (
	"time"

	"github.com/lib/pq"
)

// TansRecord wrapps all db columns
type TansRecord struct {
	RecordID    int            `db:"id"`
	GroupID     int            `db:"g_id"`
	Date        time.Time      `db:"day"`
	Payer       string         `db:"payer"`
	Spliters    pq.StringArray `db:"spliters"`
	Amount      float32        `db:"pay_amount"`
	Description string         `db:"description"`
	UpdatedAt   time.Time      `db:"updated_at"`
	Deleted     bool           `db:"deleted"`
}
