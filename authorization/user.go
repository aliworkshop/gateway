package authorization

import "time"

type UserModel interface {
	GetCacheID() uint64
	GetId() uint64
	GetCreationDate() time.Time
	GetUpdateDate() *time.Time
	GetPasswordChangedAt() *time.Time
	GetMobile() *string
	GetNormalizedMobile() *string
	GetEmail() *string
	GetPlan() string
	GetReferralLevel() int
	GetStatus64() uint64
	GetCrmId() *string
	GetRefId() string
}
