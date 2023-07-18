package handlerlib

import "github.com/aliworkshop/loggerlib/logger"

type HandlerModel interface {
	Handler
	Responder
	// SkipSuccessResponses disables functionality for success responses it wont
	// disable error response handler
	SkipSuccessResponses()
	// IsSkippingSuccessResponses returns state if responder is enabled for
	// success responses
	IsSkippingSuccessResponses() bool
}

type handlerModelStruct struct {
	Handler
	Responder
	skipSuccessResponses bool
}

func (h *handlerModelStruct) SkipSuccessResponses() {
	h.skipSuccessResponses = true
}

func (h *handlerModelStruct) IsSkippingSuccessResponses() bool {
	return h.skipSuccessResponses
}

func NewModel(h Handler, res Responder) HandlerModel {
	return &handlerModelStruct{
		Handler:   h,
		Responder: res,
	}
}

func Handle(model HandlerModel, request RequestModel, respond bool) bool {
	result, err := model.HandlerFunc()(request)
	if err != nil {
		model.Logger().With(logger.Field{
			"message": err.Message(),
			"detail":  err.Detail(),
		}).ErrorF("error on handler func")
		model.RespondWithError(request, err)
		return false
	}
	if respond && !model.IsSkippingSuccessResponses() && !request.IsResponded() {
		model.Respond(request, "", result)
		return true
	}
	return true
}
