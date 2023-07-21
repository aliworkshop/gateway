package middleware

import (
	"github.com/aliworkshop/configlib"
	"github.com/aliworkshop/errorslib"
	"github.com/aliworkshop/gateway"
)

var headerMissMatchError = errorslib.New(nil).
	WithType(errorslib.TypeForbidden).
	WithId("HeaderMissMatchError").
	WithMessage("System didn't meet sufficient requirements to process your request.").
	WithDetail("insufficient header")

type header struct {
	gateway.HandlerEngine
	config struct {
		Headers []struct {
			Key   string
			Value string
		}
	}
}

func NewHeaderHandler(handlerModel gateway.HandlerEngine,
	configRegistry configlib.Registry) gateway.HandlerEngine {
	handler := &header{
		HandlerEngine: handlerModel,
	}
	handler.SetHandlerFunc(handler.handle)
	if err := configRegistry.Unmarshal(&handler.config); err != nil {
		panic(err)
	}
	return handler
}

func (h *header) handle(request gateway.Requester) (interface{}, errorslib.ErrorModel) {
	for _, header := range h.config.Headers {
		v := request.GetHeader(header.Key)
		if v != header.Value {
			return nil, headerMissMatchError
		}
	}
	return nil, nil
}
