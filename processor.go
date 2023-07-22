package gateway

import "github.com/aliworkshop/loggerlib/logger"

type Processor interface {
	Process(handler Handler, request Requester, respond bool) (ok bool)
}

func (c *controller) Process(handler Handler, request Requester, respond bool) (ok bool) {
	result, err := handler.Handle(request)
	if err != nil {
		c.log.With(logger.Field{
			"message": err.Message(),
			"detail":  err.Detail(),
		}).ErrorF("error on handler func")
		c.RespondError(request, err)
		return
	}
	if respond && !request.IsResponded() {
		c.Respond(request, StatusUnknown, result)
	}
	return true
}
