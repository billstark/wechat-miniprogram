package endpoints

import (
	"context"
	"time"

	"github.com/go-kit/kit/endpoint"
	"github.com/go-kit/kit/log"

	"wechat-miniprogram/services"
	"wechat-miniprogram/services/helper"
)

const (
	// ServiceDetailInfoRetrieve defines detail info retrieve service type
	ServiceDetailInfoRetrieve = "detail_info_retrieve"

	logErrorTag    = "error"
	logTimeTag     = "took"
	logEndpointTag = "endpoint"
	logParamsTag   = "params"
)

// MakeRetrieveEndpoint makes retrieve endpoint for different types of serviecs
func MakeRetrieveEndpoint(logger log.Logger, service services.Service, serviceType string) endpoint.Endpoint {

	// Returns an endpoint (basically, an enpoint is a place to deal with request)
	// For here we just pass it to services to do it.
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		var err error

		// Before return, log status
		defer func(start time.Time) {
			logger.Log(
				logErrorTag, err,
				logTimeTag, time.Since(start),
				logEndpointTag, serviceType,
				logParamsTag, helper.ObjToString(request),
			)
		}(time.Now())

		// Pass job to services
		result, err := service.Retrieve(ctx, request)
		return result, err
	}
}
