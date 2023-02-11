package authorization

import (
	"time"
)

type SessionModel interface {
	GetType() string
	GetUserId() uint64
	GetUserRefId() string
	SetAccessToken(token string)
	GetAccessToken() string
	SetRefreshToken(token string)
	GetRefreshToken() string
	SetCSRFToken(token string)
	GetCSRFToken() string
	GetCreateAt() time.Time
	GetExpireAt() time.Time
	SetScopes(scopes []string)
	GetScopes() []string
	SetRoles(roles []string)
	GetRoles() []string
	GetAgent() UserAgent
	GetIp() string
	GetClient() string
	GetMetadata() map[string]interface{}
	HasAccessToSubAccount(refId string, scopes ...string) bool
	GetUser() UserModel
	IsKYCVerified() bool
}

type Model interface {
	Token() (token string)
	SetToken(token string)
	IsAuthenticated() bool
	SetIsAuthenticated(isAuthenticated bool)
	GetCurrentAccountId() interface{}
	GetScopes() []string
	HasScope(scopes ...string) bool
	GetRoles() []string
	HasRole(roles ...string) bool
	SetSession(session SessionModel)
	GetSession() SessionModel
}

type UserAgent interface {
	Name() string
	Raw() string
	Version() string
	Device() string
	Os() string
	OsVersion() string
	IsDesktop() bool
	IsBot() bool
	IsTablet() bool
	IsMobile() bool
}
