package middleware

import (
	"github.com/aliworkshop/configlib"
	"github.com/aliworkshop/gateway/v2"
)

func Get(configRegistry configlib.Registry,
	middlewareType string) gateway.Handler {
	switch middlewareType {
	case "header":
		return NewHeaderHandler(configRegistry)
	}
	return nil
}
