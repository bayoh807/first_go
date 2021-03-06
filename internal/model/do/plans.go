// =================================================================================
// Code generated by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// Plans is the golang structure of table plans for DAO operations like Where/Data.
type Plans struct {
	g.Meta    `orm:"table:plans, do:true"`
	Id        interface{} //
	Name      interface{} //
	Price     interface{} //
	Unit      interface{} //
	CreatedAt *gtime.Time //
	UpdatedAt *gtime.Time //
	DeletedAt *gtime.Time //
}
