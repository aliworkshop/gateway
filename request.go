package gateway

import (
	"context"
	"github.com/aliworkshop/dfilterlib"
	"mime/multipart"
	"net/http"

	"github.com/aliworkshop/errorslib"
	"github.com/aliworkshop/gateway/authorization"
	"github.com/aliworkshop/loggerlib/logger"
	"github.com/nicksnyder/go-i18n/v2/i18n"
)

type Language struct {
	AcceptLanguage string
	Localizer      *i18n.Localizer
}

type Paginator interface {
	PerPage() int
	Page() int
	SetPage(int)
	SetPerPage(int)
	SortBy() string
}

type Requester interface {
	SetUid(uid string)
	GetUid() string

	SetConnectionContext(ctx context.Context)
	GetConnectionContext() context.Context
	WithLogger(l logger.Logger) Requester
	Logger() logger.Logger
	SetContext(any)
	GetContext() any
	GetClientIp() string
	GetMethod() string
	GetPath() string
	GetHeader(key string) string
	Pager() Paginator
	Cookie(name string) (string, error)
	SetCookie(cookie any)

	SetAuth(auth authorization.Model)
	GetAuth() authorization.Model
	Token() (token string)
	IsAuthenticated() bool
	GetCurrentAccountId() any
	GetScopes() []string
	HasScope(scopes ...string) bool

	GetBody() any
	SetBody(any)
	Request() *http.Request
	Writer() http.ResponseWriter
	BindRequest(body any) (err errorslib.ErrorModel)
	GetLanguage() Language
	MustLocalize(lc *i18n.LocalizeConfig) string
	ShouldLocalize(lc *i18n.LocalizeConfig) string
	Localize(msgId string, message string, params ...map[string]any) string
	SetModel(any)
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
	SetTemp(key string, value any)
	// GetTemp gets the existing temp value from request, returns nil if
	// nothing found for given key
	GetTemp(key string) any
	GetStatusCode() int
}
