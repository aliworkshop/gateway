package middleware

import (
	"github.com/aliworkshop/configer"
	"github.com/aliworkshop/gateway/v2"
)

func Get(configRegistry configer.Registry,
	middlewareType string) gateway.Handler {
	switch middlewareType {
	case "header":
		return NewHeaderHandler(configRegistry)
	}
	return nil
}
