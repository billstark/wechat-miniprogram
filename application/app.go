package application

import (
	"context"
	"encoding/json"
	"net/http"
	"os"

	gokitLog "github.com/go-kit/kit/log"
	gokitHttp "github.com/go-kit/kit/transport/http"
	mux "github.com/gorilla/mux"

	endpoints "wechat-miniprogram/services/endpoints"
	database "wechat-miniprogram/utils/database"
	healthcheck "wechat-miniprogram/utils/healthcheck"
	responses "wechat-miniprogram/utils/responses"
	server "wechat-miniprogram/utils/server"

	storeErr "wechat-miniprogram/datastore/error"
	serviceErr "wechat-miniprogram/services/errors"

	recordStore "wechat-miniprogram/datastore/record"
	detailInfoService "wechat-miniprogram/services/detailInfo"
	detailInfoHttp "wechat-miniprogram/services/detailInfo/transports/http"
)

const (
	// LogTimestampTag is the tag for timestamp log
	LogTimestampTag = "timestamp"
	// LogCallerTag is the tag for caller log
	LogCallerTag = "caller"
	// LogLayerTag is the tag for layer log
	LogLayerTag = "layer"
	// LogRouteTag is the tag for route log
	LogRouteTag = "route"
	// LogMessageTag is the tag for message log
	LogMessageTag = "message"
	// LogErrorTag is the tag for error log
	LogErrorTag = "error"

	// LayerApplication represents application layer
	LayerApplication = "application"
	// LayerEndpoint represents endpoint layer
	LayerEndpoint = "endpoint"
	// LayerTransport represents transport layer
	LayerTransport = "transport"

	// HTTPHeaderConetent is for http content type header
	HTTPHeaderConetent = "Content-Type"
	// HTTPContentJSON represents json conetent type
	HTTPContentJSON = "application/json"
	// HTTPContentUTF8 represents utf-8 conetent type
	HTTPContentUTF8 = "charset=utf-8"
	// HTTPHeaderBreak breaks in http header
	HTTPHeaderBreak = ";"

	// MessageListenAddress represents message for logging listen address
	MessageListenAddress = "http listening on "
	// MessageHalting represents message for halting
	MessageHalting = "halting!"
)

// App is a struct containing basic app configs
type App struct {
	Router          *mux.Router
	Logger          gokitLog.Logger
	AppLogger       gokitLog.Logger
	EndpointLogger  gokitLog.Logger
	TransportLogger gokitLog.Logger
	DB              database.Database
	ServerConfig    server.ServerConfig
	ErrorEncoder    gokitHttp.ServerOption
	Errs            chan error
}

// InitApp inits app configs
func (a *App) InitApp(dbConfig database.DBConfig, serverConfig server.ServerConfig) error {
	a.Router = mux.NewRouter()
	a.ServerConfig = serverConfig
	a.Errs = make(chan error)
	a.ErrorEncoder = gokitHttp.ServerErrorEncoder(errorHandler)

	a.initLoggers()
	err := a.initDB(dbConfig)
	if err != nil {
		return err
	}
	return nil
}

func (a *App) initLoggers() {
	baseLogger := gokitLog.NewLogfmtLogger(gokitLog.NewSyncWriter(os.Stderr))
	a.Logger = gokitLog.With(baseLogger,
		LogTimestampTag, gokitLog.DefaultTimestampUTC,
		LogCallerTag, gokitLog.DefaultCaller,
	)
	a.AppLogger = gokitLog.With(a.Logger, LogLayerTag, LayerApplication)
	a.EndpointLogger = gokitLog.With(a.Logger, LogLayerTag, LayerEndpoint)
	a.TransportLogger = gokitLog.With(a.Logger, LogLayerTag, LayerTransport)
}

func (a *App) initDB(dbConfig database.DBConfig) error {
	db, err := database.New(dbConfig)
	if err != nil {
		return err
	}
	a.DB = db
	return nil
}

func (a *App) initHeartbeat() {
	a.Router.Methods("GET").Path("/ping.json").Handler(http.HandlerFunc(healthcheck.Simple))
}

func (a *App) initDetailInfoHandler() {
	recordStore := recordStore.NewRecordStore(a.DB)
	detailInfoService := detailInfoService.NewDetailInfoService(recordStore)

	a.Router.Methods("GET").Path("/records/user/{host_id}/{guest_id}").Handler(gokitHttp.NewServer(
		endpoints.MakeRetrieveEndpoint(a.EndpointLogger, detailInfoService, endpoints.ServiceDetailInfoRetrieve),
		detailInfoHttp.DecodeRetrieveRequest,
		encodeJSONResponse,
		a.ErrorEncoder,
		gokitHttp.ServerErrorLogger(gokitLog.With(a.TransportLogger, LogRouteTag, "retrieve"))))
}

// Run starts the app server
func (a *App) Run() {
	go func() {
		a.initHeartbeat()
		a.initDetailInfoHandler()
		address := a.ServerConfig.Server.ListenAddress()
		srv := &http.Server{
			Handler: a.Router,
			Addr:    address,
		}
		a.Logger.Log(LogLayerTag, LayerApplication, LogMessageTag, MessageListenAddress+address)
		a.Errs <- srv.ListenAndServe()
	}()
}

func errorHandler(_ context.Context, err error, w http.ResponseWriter) {
	w.Header().Set(HTTPHeaderConetent, HTTPContentJSON+HTTPHeaderBreak+HTTPContentUTF8)
	var response responses.ErrorResponse

	switch {
	case err == storeErr.ErrInvalidQuery || err == storeErr.ErrNotSupportedQuery:
		response = responses.Invalid(err.Error(), nil)
	case err == serviceErr.ErrIncorrectParamsFormat || err == serviceErr.ErrInsufficientParams:
		response = responses.Invalid(err.Error(), nil)
	case err == storeErr.ErrNotFound:
		response = responses.NotFound()
	default:
		response = responses.InternalError(err.Error())
	}

	w.WriteHeader(response.Status)
	json.NewEncoder(w).Encode(&response)
}

func encodeJSONResponse(_ context.Context, writer http.ResponseWriter, response interface{}) error {
	writer.Header().Set(HTTPHeaderConetent, HTTPContentJSON+HTTPHeaderBreak+HTTPContentUTF8)
	return json.NewEncoder(writer).Encode(response)
}
