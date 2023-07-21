package gateway

import "github.com/aliworkshop/errorslib"

type Responder interface {
	SetTotal(uint) Responder
	Respond(Requester, Status, any)
	RespondError(model Requester, err errorslib.ErrorModel)
	LanguageHandler
}

type Response struct {
	Page    int  `json:"page"`
	PerPage int  `json:"per_page"`
	Items   any  `json:"items"`
	Total   uint `json:"total"`
}
