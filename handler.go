package gateway

import (
	"github.com/aliworkshop/errors"
)

type Action string

const (
	Create Action = "CREATE"
	Read   Action = "READ"
	Update Action = "UPDATE"
	Delete Action = "DELETE"
)

type Handler interface {
	Handle(request HttpRequester) (any, errors.ErrorModel)
}
