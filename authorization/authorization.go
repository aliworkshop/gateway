package authorization

import "time"

type Authorizer interface {
	Token() (token string)
	SetToken(token string)
	SetClaim(claim Claim)
	GetClaim() Claim
	IsAuthenticated() bool
	SetIsAuthenticated(isAuthenticated bool)
	GetCurrentAccountId() interface{}
	HasScope(scopes ...string) bool
	GetScopes() []string
}

type Claim interface {
	HasScope(scopes ...string) bool
	GetScopes() []string
	GetName() string
	GetEmail() string
	GetMobile() string
	GetExpireAt() time.Time
	GetUserId() uint64
}
