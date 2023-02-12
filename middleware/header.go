package middleware

import (
	"github.com/aliworkshop/configlib"
	"github.com/aliworkshop/errorslib"
	"github.com/aliworkshop/handlerlib"
)

var headerMissMatchError = errorslib.New(nil).
	WithType(errorslib.TypeForbidden).
	WithId("HeaderMissMatchError").
	WithMessage("System didn't meet sufficient requirements to process your request.").
	WithDetail("insufficient header")

type header struct {
	handlerlib.HandlerModel
	config struct {
		Headers []struct {
			Key   string
			Value string
		}
	}
}

func NewHeaderHandler(handlerModel handlerlib.HandlerModel,
	configRegistry configlib.Registry) handlerlib.HandlerModel {
	handler := &header{
		HandlerModel: handlerModel,
	}
	handler.SkipSuccessResponses()
	handler.SetHandlerFunc(handler.handle)
	if err := configRegistry.Unmarshal(&handler.config); err != nil {
		panic(err)
	}
	return handler
}

func (h *header) handle(request handlerlib.RequestModel,
	args ...interface{}) (interface{}, errorslib.ErrorModel) {
	for _, header := range h.config.Headers {
		v := request.GetHeader(header.Key)
		if v != header.Value {
			return nil, headerMissMatchError
		}
	}
	return nil, nil
}
