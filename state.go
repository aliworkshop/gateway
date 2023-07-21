package gateway

type Status string

const (
	StatusMovedPermanently  Status = "MOVED"              // like http status 301
	StatusFound             Status = "FOUND"              // like http status 302
	StatusPermanentRedirect Status = "PERMANENT_REDIRECT" // like http status 307
	StatusTemporaryRedirect Status = "TEMPORARY_REDIRECT" // like http status 308
	StatusOK                Status = "OK"
	StatusCreated           Status = "CREATED"
	StatusNoContent         Status = "NO_CONTENT"
	StatusBadInput          Status = "BAD REQUEST"
	StatusConflict          Status = "CONFLICT"
)
