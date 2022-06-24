package service

import (
	"context"
	"fmt"
	v1 "gf_0620/api/v1"

	"github.com/gogf/gf/v2/frame/g"
	"golang.org/x/crypto/bcrypt"
)

var (
	User  = sUser{}
	query = g.Model("users")
)

type sUser struct{}

func (s *sUser) GetUsers() (users *v1.UserRes, err error) {

	fmt.Print(query.One())

	return
}

func (s *sUser) CreateUser(ctx context.Context) (users *v1.UserRes, err error) {

	passwordOK := "123456"

	hash, err1 := bcrypt.GenerateFromPassword([]byte(passwordOK), bcrypt.DefaultCost)

	if err1 != nil {
		fmt.Print("err1 : ", err)
	}

	encodePW := string(hash)

	// test, err2 := query.Data(g.Map{"account": "john123", "password": encodePW, "name": "john"}).InsertIgnore()
	test2, err3 := query.Data(g.Map{"account": "john123", "password": encodePW, "name": "john"}).Insert(g.Map{"account": "john123", "password": encodePW, "name": "john"})

	// if err2 != nil {
	// 	fmt.Println("err2 :", err2)
	// }
	if err3 != nil {
		fmt.Println("err3 :", err3)
	}

	// str, err := test.e
	fmt.Println("re3: ", test2)
	// fmt.Println("s")
	return
}
