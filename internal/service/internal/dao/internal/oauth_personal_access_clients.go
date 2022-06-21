// ==========================================================================
// Code generated by GoFrame CLI tool. DO NOT EDIT. Created at 2022-06-21 11:39:19
// ==========================================================================

package internal

import (
	"context"
	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// OauthPersonalAccessClientsDao is the data access object for table oauth_personal_access_clients.
type OauthPersonalAccessClientsDao struct {
	table   string          // table is the underlying table name of the DAO.
	group   string          // group is the database configuration group name of current DAO.
	columns OauthPersonalAccessClientsColumns // columns contains all the column names of Table for convenient usage.
}

// OauthPersonalAccessClientsColumns defines and stores column names for table oauth_personal_access_clients.
type OauthPersonalAccessClientsColumns struct {
	Id         string //   
    ClientId   string //   
    CreatedAt  string //   
    UpdatedAt  string //
}

//  oauthPersonalAccessClientsColumns holds the columns for table oauth_personal_access_clients.
var oauthPersonalAccessClientsColumns = OauthPersonalAccessClientsColumns{
	Id:        "id",          
            ClientId:  "client_id",   
            CreatedAt: "created_at",  
            UpdatedAt: "updated_at",
}

// NewOauthPersonalAccessClientsDao creates and returns a new DAO object for table data access.
func NewOauthPersonalAccessClientsDao() *OauthPersonalAccessClientsDao {
	return &OauthPersonalAccessClientsDao{
		group:   "default",
		table:   "oauth_personal_access_clients",
		columns: oauthPersonalAccessClientsColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *OauthPersonalAccessClientsDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *OauthPersonalAccessClientsDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *OauthPersonalAccessClientsDao) Columns() OauthPersonalAccessClientsColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *OauthPersonalAccessClientsDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *OauthPersonalAccessClientsDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *OauthPersonalAccessClientsDao) Transaction(ctx context.Context, f func(ctx context.Context, tx *gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}