package handlerlib

type Responder interface {
	Respond(RequestModel, Status, interface{})
	RespondWithError(model RequestModel, err errors.ErrorModel)
	LanguageHandler
}
