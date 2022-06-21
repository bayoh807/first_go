package v1

import (
	"github.com/gogf/gf/v2/frame/g"
)

type UserReq struct {
	g.Meta `mime:"application/json" example:"string"`
}
type UserRes struct {
	Id string
	// Name      string
	// Email     string
	// CreatedAt *gtime.Time `json:"createdAt"   dc:"创建时间"`
	// UpdatedAt *gtime.Time `json:"updatedAt"   dc:"最后修改时间"`
}
