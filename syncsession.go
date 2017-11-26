package v3io

import (
	"github.com/valyala/fasthttp"
	"github.com/nuclio/nuclio-sdk"
)

type SyncSession struct {
	logger     nuclio.Logger
	context    *SyncContext
	sessionKey string
}

func newSyncSession(parentLogger nuclio.Logger,
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
