package middleware

import (
	"github.com/aliworkshop/configlib"
	"github.com/aliworkshop/handlerlib"
)

func Get(handlerModel handlerlib.HandlerModel,
	configRegistry configlib.Registry,
	middlewareType string) handlerlib.HandlerModel {
	switch middlewareType {
	case "header":
		return NewHeaderHandler(handlerModel, configRegistry)
	}
	return nil
}
