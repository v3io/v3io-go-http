package v3io

import (
	"github.com/valyala/fasthttp"
	"github.com/nuclio/logger"
)

type SyncSession struct {
	logger     logger.Logger
	context    *SyncContext
	sessionKey string
}

func newSyncSession(parentLogger logger.Logger,
	context *SyncContext,
	username string,
	password string,
	label string) (*SyncSession, error) {
	return &SyncSession{
		logger:  parentLogger.GetChild("session"),
		context: context,
	}, nil
}

func (ss *SyncSession) sendRequest(request *fasthttp.Request, response *fasthttp.Response) error {

	// add session key
	// TODO

	// delegate to context
	return ss.context.sendRequest(request, response)
}
