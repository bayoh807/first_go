package controller

import (
	"context"
	"fmt"

	// "gf_0620/internal/dao"

	v1 "gf_0620/api/v1"
)

var (
	User = cUser{}
)

type cUser struct{}

func (c *cUser) Index(ctx context.Context, req *v1.UserReq) (users *v1.UserRes, err error) {

	// users, err = service.User().GetUsers()

	// g.RequestFromCtx(ctx).Response.Writeln(users)

	// // users = dao.
	// // service.Users()
	// // users = service.getUsers()
	// // users := service.User
	// // service.getUsers()
	// // result, err = service.User.GetUsers(ctx)

	// if err != nil {
	// 	return nil, err
	// }

	// // available, err := service.User().IsNicknameAvailable(ctx, req.Nickname)
	fmt.Println(123)
	return
}

func (c *cUser) Store(ctx context.Context, req *v1.UserReq) (res *v1.UserRes, err error) {
	// g.RequestFromCtx(ctx).Response.Writeln("Store User!")

	// service.User.GetUsers()
	return
}

// func Index(ctx context.Context, req *v1.HelloReq) (res *v1.HelloRes, err error) {
// 	g.RequestFromCtx(ctx).Response.Writeln("Hello World!")
// 	return
// }
