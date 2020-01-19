package sys

import (
	"net/http"
)

var (
	// UniversalHandler 通用处理器
	UniversalHandler *universalHandler
	urls             map[string]string
)

func init() {
	UniversalHandler = &universalHandler{
		http.DefaultServeMux,
	}
}

type universalHandler struct {
	serverMux *http.ServeMux
}

func (ush *universalHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	handler, pattern := ush.serverMux.Handler(r)
	if _, ok := urls[pattern]; ok {

		return
	}
	handler.ServeHTTP(w, r)
}
func (ush *universalHandler) Handle(pattern string, handler http.Handler) {
	http.Handle(pattern, handler)
	urls[pattern] = ""
}
