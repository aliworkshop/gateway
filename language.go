package gateway

import "github.com/nicksnyder/go-i18n/v2/i18n"

type Language interface {
	Localize(lc *i18n.LocalizeConfig) (string, error)
}

type language struct {
	bundle *i18n.Bundle
	*i18n.Localizer
}

func NewLanguage(bundle *i18n.Bundle, langs ...string) Language {
	l := &language{
		bundle:    bundle,
		Localizer: i18n.NewLocalizer(bundle, langs...),
	}
	return l
}
