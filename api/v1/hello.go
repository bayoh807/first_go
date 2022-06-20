package v1

import (
	"github.com/gogf/gf/v2/frame/g"
)

type HelloReq struct {
	g.Meta `mime:"text/html" example:"string"`
}
type HelloRes struct {
	g.Meta `mime:"text/html" example:"string"`
}
