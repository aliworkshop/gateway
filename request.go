package gateway

import (
	"context"
	"github.com/aliworkshop/dfilter"
	"io"
	"io/fs"
	"mime/multipart"
	"net/http"

	errors "github.com/aliworkshop/error"
	"github.com/aliworkshop/gateway/v2/authorization"
	"github.com/nicksnyder/go-i18n/v2/i18n"
)

type Paginator interface {
	PerPage() int
	Page() int
	SetPage(int)
	SetLimit(int)
	SetSort(string)
	SetTotal(uint64)
	SortBy() string
	Total() uint64
}

type Requester interface {
	SetUid(uid string)
	GetUid() string

	SetConnectionContext(ctx context.Context)
	GetConnectionContext() context.Context
	SetContext(any)
	GetContext() any
	GetClientIp() string
	GetMethod() string
	GetPath() string
	GetHeader(key string) string
	Paginator() Paginator
	Cookie(name string) (string, error)
	SetCookie(cookie any)

	SetAuth(auth authorization.Authorizer)
	GetAuth() authorization.Authorizer
	Token() (token string)
	IsAuthenticated() bool
	GetCurrentAccountId() any
	GetScopes() []string
	HasScope(scopes ...string) bool

	GetBody() any
	SetBody(any)
	Request() *http.Request
	Writer() http.ResponseWriter
	BindRequest(body any) (err errors.ErrorModel)
	MustLocalize(lc *i18n.LocalizeConfig) string
	ShouldLocalize(lc *i18n.LocalizeConfig) string
	Localize(msgId string, message string, params ...map[string]any) string
	GetQuery(key string) string
	GetParam(key string) string
	GetFile(key string) (*multipart.FileHeader, error)
	Filters() map[string][]string
	IsResponded() bool
	SetResponded(bool)
	SetDynamicFilters(filter []dfilter.Filter)
	GetDynamicFilters() []dfilter.Filter
	// SetTemp sets temp into current request to handle during processing
	// request
	SetTemp(key string, value any)
	// GetTemp gets the existing temp value from request, returns nil if
	// nothing found for given key
	GetTemp(key string) any
	GetStatusCode() int

	RespondBlob(status Status, contentType string, body []byte) errors.ErrorModel
	RespondStream(status Status, contentType string, reader io.Reader) errors.ErrorModel
	RespondFile(file string) errors.ErrorModel
	RespondFsFile(file string, filesystem fs.FS) errors.ErrorModel
}
