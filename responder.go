package gateway

import (
	"github.com/aliworkshop/errorslib"
	"github.com/nicksnyder/go-i18n/v2/i18n"
)

type Responder interface {
	SetTotal(uint) Responder
	Respond(Requester, Status, any)
	RespondError(model Requester, err errorslib.ErrorModel)
	LanguageBundle() *i18n.Bundle
	SetLanguageBundle(bundle *i18n.Bundle)
}

type Response struct {
	Page    int    `json:"page"`
	PerPage int    `json:"per_page"`
	Items   any    `json:"items"`
	Total   uint64 `json:"total"`
}
