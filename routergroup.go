package gateway

import "net/http"

type RouterGroupModel interface {
	Group(relativePath string) RouterGroupModel
	READ(path string, handlers ...Handler)
	CREATE(path string, handlers ...Handler)
	UPDATE(path string, handlers ...Handler)
	DELETE(path string, handlers ...Handler)
	STATIC(path string)
	ServeHttp(w http.ResponseWriter, req *http.Request)
	// middleware
	Middleware(handlers ...Handler)
}

func RegisterRouters(model RouterGroupModel, path string, action Action, handlers ...Handler) {
	switch action {
	case Read:
		model.READ(path, handlers...)
	case Create:
		model.CREATE(path, handlers...)
	case Update:
		model.UPDATE(path, handlers...)
	case Delete:
		model.DELETE(path, handlers...)
	}
}
