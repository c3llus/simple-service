package http

import (
	"github.com/c3llus/cdk/app/http"
)

type Server struct {
}

func NewServer() Server {
	return Server{}
}

func (s Server) RegisterHandler(r *http.Router) {
	r.HandleFunc("/ping", ping, "GET")
}

func ping(httpCtx http.CtxHTTP) error {
	httpCtx.Writer().Write([]byte("pong"))
	return nil
}
