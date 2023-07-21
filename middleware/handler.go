package middleware

import (
	"github.com/aliworkshop/configlib"
	"github.com/aliworkshop/gateway"
)

func Get(handlerModel gateway.HandlerEngine,
	configRegistry configlib.Registry,
	middlewareType string) gateway.HandlerEngine {
	switch middlewareType {
	case "header":
		return NewHeaderHandler(handlerModel, configRegistry)
	}
	return nil
}
