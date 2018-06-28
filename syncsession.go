package v3io

import (
	"encoding/base64"
	"fmt"

	"github.com/nuclio/logger"
	"github.com/valyala/fasthttp"
)

type SyncSession struct {
	logger             logger.Logger
	context            *SyncContext
	authenticatioToken string
}

func newSyncSession(parentLogger logger.Logger,
	context *SyncContext,
	username string,
	password string,
	label string) (*SyncSession, error) {

	// generate token for basic authentication
	usernameAndPassword := fmt.Sprintf("%s:%s", username, password)
	encodedUsernameAndPassword := base64.StdEncoding.EncodeToString([]byte(usernameAndPassword))

	return &SyncSession{
		logger:             parentLogger.GetChild("session"),
		context:            context,
		authenticatioToken: "Basic " + encodedUsernameAndPassword,
	}, nil
}

func (ss *SyncSession) sendRequest(request *fasthttp.Request, response *fasthttp.Response) error {

	// add authorization token
	request.Header.Set("Authorization", ss.authenticatioToken)

	// delegate to context
	return ss.context.sendRequest(request, response)
}
