package storemodels

import "time"

// RecordCreateParams wraps params used to create a new record
type RecordCreateParams struct {
	GroupID     int
	Date        time.Time
	PayerID     string
	Spliters    []string
	Amount      float32
	Description string
	UpdatedAt   time.Time
	DeletedAt   time.Time
}

// RecordUpdateParams wraps params used to update a record
type RecordUpdateParams struct {
	Date        time.Time
	PayerID     string
	Spliters    []string
	Amount      float32
	Description string
	UpdatedAt   time.Time
	DeletedAt   time.Time
}

// RecordRetrieveParams wraps params for record retrieval
type RecordRetrieveParams struct {
	HostID  string
	GuestID string
	GroupID int
}

// RecordRetrieveResult wraps record retrieval info
type RecordRetrieveResult struct {
	RecordID    int
	GroupID     int
	Date        time.Time
	Payer       string
	Spliters    []string
	Amount      float32
	Description string
}
