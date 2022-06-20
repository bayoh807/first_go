// ==========================================================================
// Code generated by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"
	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// PersonalAccessTokensDao is the data access object for table personal_access_tokens.
type PersonalAccessTokensDao struct {
	table   string                      // table is the underlying table name of the DAO.
	group   string                      // group is the database configuration group name of current DAO.
	columns PersonalAccessTokensColumns // columns contains all the column names of Table for convenient usage.
}

// PersonalAccessTokensColumns defines and stores column names for table personal_access_tokens.
type PersonalAccessTokensColumns struct {
	Id            string //
	TokenableType string //
	TokenableId   string //
	Name          string //
	Token         string //
	Abilities     string //
	LastUsedAt    string //
	CreatedAt     string //
	UpdatedAt     string //
}

//  personalAccessTokensColumns holds the columns for table personal_access_tokens.
var personalAccessTokensColumns = PersonalAccessTokensColumns{
	Id:            "id",
	TokenableType: "tokenable_type",
	TokenableId:   "tokenable_id",
	Name:          "name",
	Token:         "token",
	Abilities:     "abilities",
	LastUsedAt:    "last_used_at",
	CreatedAt:     "created_at",
	UpdatedAt:     "updated_at",
}

// NewPersonalAccessTokensDao creates and returns a new DAO object for table data access.
func NewPersonalAccessTokensDao() *PersonalAccessTokensDao {
	return &PersonalAccessTokensDao{
		group:   "default",
		table:   "personal_access_tokens",
		columns: personalAccessTokensColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *PersonalAccessTokensDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *PersonalAccessTokensDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *PersonalAccessTokensDao) Columns() PersonalAccessTokensColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *PersonalAccessTokensDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *PersonalAccessTokensDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *PersonalAccessTokensDao) Transaction(ctx context.Context, f func(ctx context.Context, tx *gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
