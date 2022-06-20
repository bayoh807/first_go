// =================================================================================
// Code generated by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// PersonalAccessTokens is the golang structure of table personal_access_tokens for DAO operations like Where/Data.
type PersonalAccessTokens struct {
	g.Meta        `orm:"table:personal_access_tokens, do:true"`
	Id            interface{} //
	TokenableType interface{} //
	TokenableId   interface{} //
	Name          interface{} //
	Token         interface{} //
	Abilities     interface{} //
	LastUsedAt    *gtime.Time //
	CreatedAt     *gtime.Time //
	UpdatedAt     *gtime.Time //
}
