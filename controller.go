package gateway

import (
	"github.com/aliworkshop/loggerlib/logger"
	"github.com/nicksnyder/go-i18n/v2/i18n"
)

type Controller interface {
	Responder
	Processor
}

type controller struct {
	Responder
	bundle *i18n.Bundle
	log    logger.Logger
}

func NewController(responder Responder, log logger.Logger) Controller {
	c := &controller{
		Responder: responder,
		log:       log,
	}
	return c
}
