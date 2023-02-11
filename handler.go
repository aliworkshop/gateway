package handlerlib

import (
	"github.com/aliworkshop/errorslib"
	"github.com/aliworkshop/loggerlib/logger"
)

type HandlerFunc func(request RequestModel, args ...interface{}) (interface{}, errorslib.ErrorModel)

type ActionType string

const (
	ActionCreate ActionType = "CREATE"
	ActionRead   ActionType = "READ"
	ActionUpdate ActionType = "UPDATE"
	ActionDelete ActionType = "DELETE"
)

type Handler interface {
	Logger() logger.Logger
	Upgrade(model RequestModel) (WebSocketModel, error)
	HandlerFunc() HandlerFunc
	SetHandlerFunc(hand HandlerFunc)
}
