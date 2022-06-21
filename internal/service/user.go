package service

import (
	"context"
	v1 "gf_0620/api/v1"
	"gf_0620/internal/model/entity"

	"github.com/gogf/gf/v2/frame/g"
)

var (
	User = sUser{}
)

type sUser struct{}

func (s *sUser) GetUsers() (users *entity.Users) {

	g.Model("users").Where("id", 1).Scan(&users)

	return
}

func (s *sUser) CreateUser(ctx context.Context) (users *v1.UserRes) {

	// fmt.Println(users)
	return users
}
