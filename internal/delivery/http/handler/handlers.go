package handler

type Handlers struct {
	AuthHandler AuthHandler
}

func NewHandlers(authHandler AuthHandler) *Handlers {
	return &Handlers{
		AuthHandler: authHandler,
	}
}
