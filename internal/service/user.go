package service

import (
	"fmt"
	v1 "gf_0620/api/v1"

	"github.com/gogf/gf/v2/frame/g"
)

var (
	User  = sUser{}
	query = g.Model("users")
)

type sUser struct{}

func (s *sUser) GetUsers() (users *v1.UserRes, err error) {

	user := query

	user.Where("id", 1)
	fmt.Print(user.All())

	return
}

func (s *sUser) CreateUser() (users *v1.UserRes, err error) {

	// passwordOK := "123456"

	// hash, err := bcrypt.GenerateFromPassword([]byte(passwordOK), bcrypt.DefaultCost)

	// if err != nil {
	// 	fmt.Println("err: %v", err)
	// }

	// encodePW := string(hash)

	// test, err := query.Data(g.Map{"account": "john123", "password": encodePW, "name": "john"}).Save()

	// if err != nil {
	// 	fmt.Println(err)
	// }

	// fmt.Println(test)
	fmt.Println("s")
	return
}
