package cmd

import (
	"context"
	"gf_0620/internal/controller"

	"github.com/gogf/gf/v2/frame/g"
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
			group.GET("/hello", controller.Hello)

			s.Run()
			return nil
		},
	}
)
