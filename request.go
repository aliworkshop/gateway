package gateway

import (
	"context"
	"io"
	"io/fs"
	"mime/multipart"
	"net/http"

	ad "github.com/aliworkshop/authorizer/port"
	"github.com/aliworkshop/dfilter"
	"github.com/aliworkshop/errors"
	"github.com/nicksnyder/go-i18n/v2/i18n"
)

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
	Paginator() IPaginator
	SetPaginator(paging IPaginator)
	Cookie(name string) (string, error)
	SetCookie(cookie any)
	Sorter() Sorter
	SetSorter(sorting Sorter)

	SetAuth(auth ad.Authorizer)
	GetAuth() ad.Authorizer
	Token() (token string)
	IsAuthenticated() bool
	GetCurrentAccountId() uint64
	GetCurrentAccountEmail() string
	GetCurrentAccountUuid() string
	GetScopes() map[string]uint16
	HasScope(scopes ...string) bool
	GetRequestScopes() []string
	SetRequestScopes(scopes ...string)
	GetRoles() map[string]uint16
	GetRoleId(role string) uint16
	HasRole(roles ...string) bool
	GetIssuer() string
	RequestUUID() string
	SetRequestUUID(str string)
	SetDynamicFilters(filter []dfilter.Filter)
	GetDynamicFilters() []dfilter.Filter
	Filters() map[string][]string
	SetKey(key string, value any)
	GetKey(key string) (value any, exists bool)
	Translate(msgId, message string, params ...any) string
	SetLanguage(Language)
	GetLanguage() Language
}

type HttpRequester interface {
	Requester
	SetIsResponded(bool)
	IsResponded() bool
	GetBody() any
	SetBody(any)
	Request() *http.Request
	Writer() http.ResponseWriter
	BindRequest(req Validatable) (err errors.ErrorModel)
	ShouldLocalize(lc *i18n.LocalizeConfig) string
	MustLocalize(lc *i18n.LocalizeConfig) string
	Localize(msgId string, message string, params ...map[string]any) string
	GetHttpContext() any
	SetCookie(cookie any)
	FormValue(key string) string
	GetQuery(key string) string
	GetParam(key string) string
	GetFile(key string) (*multipart.FileHeader, error)
	GetFiles(key string) ([]*multipart.FileHeader, error)
	GetAllFiles() (map[string][]*multipart.FileHeader, error)
	GetStatusCode() int
	Websocket() (WebSocketHandler, errors.ErrorModel)
	RespondBlob(status Status, contentType string, body []byte) errors.ErrorModel
	RespondStream(status Status, contentType string, reader io.Reader) errors.ErrorModel
	RespondFile(file string) errors.ErrorModel
	RespondFsFile(file string, filesystem fs.FS) errors.ErrorModel
	RespondHtml(status int, name string, body any) errors.ErrorModel
}
