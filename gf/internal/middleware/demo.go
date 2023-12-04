package middleware

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"pdmaner-test/internal/model"
	"pdmaner-test/internal/service"
)

func DemoSetCtx(r *ghttp.Request) {
	// 初始化，务必最开始执行
	customCtx := &model.Context{
		Session: r.Session,
		Data:    make(g.Map),
		User:    &model.ContextUser{Id: "123123123"},
	}
	service.BizCtx().Init(r, customCtx)
	r.Middleware.Next()
}
