package handlerlib

import "github.com/aliworkshop/errorslib"

type Responder interface {
	Respond(RequestModel, Status, interface{})
	RespondWithError(model RequestModel, err errorslib.ErrorModel)
	LanguageHandler
}
