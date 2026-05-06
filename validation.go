package gateway

import (
	"github.com/aliworkshop/errors"
	"github.com/go-playground/validator/v10"
)

type Validatable interface {
	Validate(gValidator *validator.Validate, lang Language) errors.ErrorModel
}
