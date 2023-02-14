package handlerlib

import (
	"context"
	"github.com/aliworkshop/dfilterlib"
	"mime/multipart"
	"net/http"

	"github.com/aliworkshop/errorslib"
	"github.com/aliworkshop/handlerlib/authorization"
	"github.com/aliworkshop/loggerlib/logger"
	"github.com/nicksnyder/go-i18n/v2/i18n"
)

type Language struct {
	AcceptLanguage string
	Localizer      *i18n.Localizer
}

type Pagination interface {
	PerPage() int
	Page() int
	SetPage(int)
	SetPerPage(int)
	SortBy() string
}

type RequestModel interface {
	SetUid(uid string)
	GetUid() string

	SetConnectionContext(ctx context.Context)
	GetConnectionContext() context.Context
	WithLogger(l logger.Logger) RequestModel
	Logger() logger.Logger
	SetContext(interface{})
	GetContext() interface{}
	GetClientIp() string
	GetMethod() string
	GetPath() string
	GetHeader(key string) string
	Paging() Pagination
	Cookie(name string) (string, error)
	SetCookie(cookie interface{})

	SetAuth(auth authorization.Model)
	GetAuth() authorization.Model
	Token() (token string)
	IsAuthenticated() bool
	GetCurrentAccountId() interface{}
	GetScopes() []string
	HasScope(scopes ...string) bool

	GetBody() interface{}
	SetBody(interface{})
	BaseRequest() *http.Request
	BaseWriter() http.ResponseWriter
	HandleRequestBody(body interface{}) (err errorslib.ErrorModel)
	GetLanguage() Language
	MustLocalize(lc *i18n.LocalizeConfig) string
	ShouldLocalize(lc *i18n.LocalizeConfig) string
	Localize(msgId string, message string, params ...map[string]interface{}) string
	SetModel(interface{})
	GetQuery(key string) string
	GetParam(key string) string
	GetFile(key string) (*multipart.FileHeader, error)
	Filters() map[string][]string
	IsResponded() bool
	SetResponded(bool)
	SetDynamicFilters(filter []dfilterlib.Filter)
	GetDynamicFilters() []dfilterlib.Filter
	// SetTemp sets temp into current request to handle during processing
	// request
	SetTemp(key string, value interface{})
	// GetTemp gets the existing temp value from request, returns nil if
	// nothing found for given key
	GetTemp(key string) interface{}
	GetStatusCode() int
}
