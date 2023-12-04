package cmd

import (
	"context"
	"pdmaner-test/internal/controller/lbra"
	"pdmaner-test/internal/middleware"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/os/gcmd"

	"pdmaner-test/internal/controller/hello"
)

var (
	Main = gcmd.Command{
		Name:  "main",
		Usage: "main",
		Brief: "start http server",
		Func: func(ctx context.Context, parser *gcmd.Parser) (err error) {
			s := g.Server()
			s.Group("/", func(group *ghttp.RouterGroup) {
				group.Middleware(ghttp.MiddlewareHandlerResponse, middleware.DemoSetCtx)
				group.Bind(
					hello.NewV1(),
					lbra.NewV1(),
				)
			})
			s.Run()
			return nil
		},
	}
)
