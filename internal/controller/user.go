package controller

import (
	"context"

	// "gf_0620/internal/dao"

	v1 "gf_0620/api/v1"
	"gf_0620/internal/service"
)

var (
	User = cUser{}
)

type cUser struct{}

func (c *cUser) Index(ctx context.Context, req *v1.UserReq) (users *v1.UserRes, err error) {

	users, err = service.User.GetUsers()

	return
}

func (c *cUser) Store(ctx context.Context, req *v1.UserReq) (users *v1.UserRes, err error) {
	// g.RequestFromCtx(ctx).Response.Writeln("Store User!")

	// service.User.GetUsers()
	users, err = service.User.CreateUser(ctx)
	return
}
