package gateway

import "github.com/nicksnyder/go-i18n/v2/i18n"

type LanguageHandler interface {
	LanguageBundle() *i18n.Bundle
	SetLanguageBundle(bundle *i18n.Bundle)
}
