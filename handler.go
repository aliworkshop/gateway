package gateway

import (
	"github.com/aliworkshop/errorslib"
	"github.com/aliworkshop/loggerlib/logger"
)

type HandlerFunc func(request Requester) (any, errorslib.ErrorModel)

type ActionType string

const (
	ActionCreate ActionType = "CREATE"
	ActionRead   ActionType = "READ"
	ActionUpdate ActionType = "UPDATE"
	ActionDelete ActionType = "DELETE"
)

type Handler interface {
	Logger() logger.Logger
	Upgrade(model Requester) (WebSocketHandler, error)
	HandlerFunc() HandlerFunc
	SetHandlerFunc(hand HandlerFunc)
}
