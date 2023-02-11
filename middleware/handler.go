package middleware

import (
	"github.com/aliworkshop/handlerlib"
)

func Get(handlerModel handlerlib.HandlerModel,
	configRegistry configcore.Registry,
	middlewareType string) handlerlib.HandlerModel {
	switch middlewareType {
	case "header":
		return NewHeaderHandler(handlerModel, configRegistry)
	}
	return nil
}
