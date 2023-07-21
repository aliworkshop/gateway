package gateway

import "github.com/aliworkshop/loggerlib/logger"

type HandlerEngine interface {
	Handler
	Responder
}

type handlerEngine struct {
	Handler
	Responder
}

func NewEngine(h Handler, res Responder) HandlerEngine {
	return &handlerEngine{
		Handler:   h,
		Responder: res,
	}
}

func Handle(model HandlerEngine, request Requester, respond bool) bool {
	result, err := model.HandlerFunc()(request)
	if err != nil {
		model.Logger().With(logger.Field{
			"message": err.Message(),
			"detail":  err.Detail(),
		}).ErrorF("error on handler func")
		model.RespondError(request, err)
		return false
	}
	if respond && !request.IsResponded() {
		model.Respond(request, "", result)
		return true
	}
	return true
}
