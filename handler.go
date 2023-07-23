package gateway

import (
	"github.com/aliworkshop/error"
)

type HandlerFunc func(request Requester) (any, error.ErrorModel)

type ActionType string

const (
	ActionCreate ActionType = "CREATE"
	ActionRead   ActionType = "READ"
	ActionUpdate ActionType = "UPDATE"
	ActionDelete ActionType = "DELETE"
)

type Handler interface {
	Handle(request Requester) (any, error.ErrorModel)
}
