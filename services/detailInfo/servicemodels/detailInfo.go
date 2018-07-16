package servicemodels

import "time"

// DetailRetrieveParams wraps the params used to retrieve detailed info
type DetailRetrieveParams struct {
	HostID  string
	GuestID string
	GroupID int
}

// DetailBetweenUsers wraps the detail info result between users
type DetailBetweenUsers struct {
	HostID  string        `json:"host_id"`
	GuestID string        `json:"guest_id"`
	Records []*UserRecord `json:"records"`
}

// UserRecord wraps the record used for detail info
// between users
type UserRecord struct {
	RecordID    int
	GroupID     int
	Date        time.Time
	Amount      float32
	Description string
}

// GroupDetails wraps the detail info result for a group
type GroupDetails struct {
	HostID    string            `json:"host_id"`
	GroupID   int               `json:"group_id"`
	GroupName string            `json:"group_name"`
	Members   map[string]string `json:"members"`
	Records   []*GroupRecord    `json:"records"`
}

// GroupRecord wraps the record used for detail info
// between user and group
type GroupRecord struct {
	RecordID    int
	Date        time.Time
	Amount      float32
	Payer       string
	Spliters    map[string]string
	Description string
}
