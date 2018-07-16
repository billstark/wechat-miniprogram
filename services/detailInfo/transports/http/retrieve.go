package http

import (
	"context"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"

	"wechat-miniprogram/services/detailInfo/servicemodels"
	serviceErr "wechat-miniprogram/services/errors"
)

const (
	hostIDKey  = "host_id"
	guestIDKey = "guest_id"
	groupIDKey = "group_id"
	emptyID    = ""
)

// DecodeRetrieveRequest decodes a raw incoming http request to readable object for services
func DecodeRetrieveRequest(_ context.Context, req *http.Request) (interface{}, error) {

	// extracts params from path
	vars := mux.Vars(req)
	hostID := vars[hostIDKey]
	guestID := vars[guestIDKey]
	groupIDStr := vars[groupIDKey]

	// do not allow empty host id
	if hostID == emptyID {
		return nil, serviceErr.ErrInsufficientParams
	}

	// deal with group id. if no group id provided, sub with -1
	var groupID int
	if groupIDStr == emptyID {
		groupID = -1
	} else {
		groupID, _ = strconv.Atoi(groupIDStr)
	}

	// returns retrieve params
	return servicemodels.DetailRetrieveParams{
		HostID:  hostID,
		GuestID: guestID,
		GroupID: groupID}, nil
}
