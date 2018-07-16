package servicemodels

type RecordRetrieveParams struct {
	Type    string
	HostID  string `json:"host_id"`
	GuestID string `json:"guest_id"`
	GroupID int    `json:"group_id"`
}
