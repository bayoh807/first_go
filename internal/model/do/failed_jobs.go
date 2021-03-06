// =================================================================================
// Code generated by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// FailedJobs is the golang structure of table failed_jobs for DAO operations like Where/Data.
type FailedJobs struct {
	g.Meta     `orm:"table:failed_jobs, do:true"`
	Id         interface{} //
	Uuid       interface{} //
	Connection interface{} //
	Queue      interface{} //
	Payload    interface{} //
	Exception  interface{} //
	FailedAt   *gtime.Time //
}
