package gateway

import (
	"github.com/aliworkshop/error"
)

type HandlerFunc func(request Requester) (any, error.ErrorModel)

type Action string

const (
	Create Action = "CREATE"
	Read   Action = "READ"
	Update Action = "UPDATE"
	Delete Action = "DELETE"
)

type Handler interface {
	Handle(request Requester) (any, error.ErrorModel)
}
