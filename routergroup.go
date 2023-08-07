package gateway

type RouterGroupModel interface {
	Group(relativePath string) RouterGroupModel
	READ(path string, handlers ...Handler)
	CREATE(path string, handlers ...Handler)
	UPDATE(path string, handlers ...Handler)
	DELETE(path string, handlers ...Handler)
	STATIC(path string)
	// middleware
	Middleware(handlers ...Handler)
}

func RegisterRouters(model RouterGroupModel, path string, action ActionType, handlers ...Handler) {
	switch action {
	case ActionRead:
		model.READ(path, handlers...)
	case ActionCreate:
		model.CREATE(path, handlers...)
	case ActionUpdate:
		model.UPDATE(path, handlers...)
	case ActionDelete:
		model.DELETE(path, handlers...)
	}
}
