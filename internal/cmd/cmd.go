package cmd

import (
	"context"
	"gf_0620/internal/controller"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/os/gcmd"
)

var (
	Main = gcmd.Command{
		Name:  "main",
		Usage: "main",
		Brief: "start http server",
		Func: func(ctx context.Context, parser *gcmd.Parser) (err error) {
			s := g.Server()

			group := s.Group("/api")

			group.Group("/user", func(g *ghttp.RouterGroup) {
				g.GET("/", controller.User.Index)
				g.POST("/", controller.User.Store)
				// group.GET("/{uuid}", controller.User.Get)
				// group.PATCH("/{uuid}", controller.User.Update)
				// group.DELETE("/{uuid}", controller.User.Delete)
			})

			s.Run()
			return nil
		},
	}
)
