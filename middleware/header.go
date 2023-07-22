package middleware

import (
	"github.com/aliworkshop/configlib"
	"github.com/aliworkshop/errorslib"
	"github.com/aliworkshop/gateway/v2"
)

var headerMissMatchError = errorslib.New(nil).
	WithType(errorslib.TypeForbidden).
	WithId("HeaderMissMatchError").
	WithMessage("System didn't meet sufficient requirements to process your request.").
	WithDetail("insufficient header")

type header struct {
	config struct {
		Headers []struct {
			Key   string
			Value string
		}
	}
}

func NewHeaderHandler(configRegistry configlib.Registry) gateway.Handler {
	handler := new(header)
	if err := configRegistry.Unmarshal(&handler.config); err != nil {
		panic(err)
	}
	return handler
}

func (h *header) Handle(request gateway.Requester) (any, errorslib.ErrorModel) {
	for _, header := range h.config.Headers {
		v := request.GetHeader(header.Key)
		if v != header.Value {
			return nil, headerMissMatchError
		}
	}
	return nil, nil
}
