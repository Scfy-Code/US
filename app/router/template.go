package router

import (
	"net/http"

	"github.com/Scfy-Code/US/sys"
)

type loginTemplate struct {
}

// NewloginTemplate 创建登录页面路由
func NewloginTemplate() http.Handler {
	return loginTemplate{}
}
func (lr loginTemplate) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		var data map[string]string = map[string]string{"action": "/login.action", "msg": ""}
		sys.ReturnTemplate("login.scfy").Execute(w, data)
	}
}

type registTemplate struct {
}

// NewRegistTemplate 创建注册页面路由
func NewRegistTemplate() http.Handler {
	return registTemplate{}
}
func (rt registTemplate) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		var data map[string]interface{} = map[string]interface{}{"action": "regist.action"}
		sys.ReturnTemplate("regist.scfy").Execute(w, data)
	}

}
