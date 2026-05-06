package gateway

import (
	"github.com/aliworkshop/errors"
	"github.com/nicksnyder/go-i18n/v2/i18n"
)

type Responder interface {
	Respond(HttpRequester, Status, any)
	RespondError(model HttpRequester, err errors.ErrorModel)
	LanguageBundle() *i18n.Bundle
}

type Response struct {
	Page    int    `json:"page"`
	PerPage int    `json:"per_page"`
	Items   any    `json:"items"`
	Total   uint64 `json:"total"`
}
