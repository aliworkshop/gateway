package authorization

import "time"

type Authorizer interface {
	Token() (token string)
	SetToken(token string)
	SetClaim(claim Claim)
	GetClaim() Claim
	IsAuthenticated() bool
	SetIsAuthenticated(isAuthenticated bool)
	GetCurrentAccountId() uint64
	HasScope(scopes ...string) bool
	GetScopes() map[string]uint16
	HasRole(roles ...string) bool
	GetRoles() map[string]uint16
}

type Claim interface {
	HasScope(scopes ...string) bool
	GetScopes() map[string]uint16
	HasRole(roles ...string) bool
	GetRoles() map[string]uint16
	GetEmail() string
	GetExpireAt() time.Time
	GetUserId() uint64
	GetIssuer() string
	GetUuid() string
	GetMobile() string
}
