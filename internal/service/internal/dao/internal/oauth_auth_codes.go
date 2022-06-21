// ==========================================================================
// Code generated by GoFrame CLI tool. DO NOT EDIT. Created at 2022-06-21 11:39:19
// ==========================================================================

package internal

import (
	"context"
	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// OauthAuthCodesDao is the data access object for table oauth_auth_codes.
type OauthAuthCodesDao struct {
	table   string          // table is the underlying table name of the DAO.
	group   string          // group is the database configuration group name of current DAO.
	columns OauthAuthCodesColumns // columns contains all the column names of Table for convenient usage.
}

// OauthAuthCodesColumns defines and stores column names for table oauth_auth_codes.
type OauthAuthCodesColumns struct {
	Id         string //   
    UserId     string //   
    ClientId   string //   
    Scopes     string //   
    Revoked    string //   
    ExpiresAt  string //
}

//  oauthAuthCodesColumns holds the columns for table oauth_auth_codes.
var oauthAuthCodesColumns = OauthAuthCodesColumns{
	Id:        "id",          
            UserId:    "user_id",     
            ClientId:  "client_id",   
            Scopes:    "scopes",      
            Revoked:   "revoked",     
            ExpiresAt: "expires_at",
}

// NewOauthAuthCodesDao creates and returns a new DAO object for table data access.
func NewOauthAuthCodesDao() *OauthAuthCodesDao {
	return &OauthAuthCodesDao{
		group:   "default",
		table:   "oauth_auth_codes",
		columns: oauthAuthCodesColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *OauthAuthCodesDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *OauthAuthCodesDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *OauthAuthCodesDao) Columns() OauthAuthCodesColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *OauthAuthCodesDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *OauthAuthCodesDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *OauthAuthCodesDao) Transaction(ctx context.Context, f func(ctx context.Context, tx *gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}