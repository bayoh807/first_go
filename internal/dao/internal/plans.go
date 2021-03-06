// ==========================================================================
// Code generated by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"
	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// PlansDao is the data access object for table plans.
type PlansDao struct {
	table   string       // table is the underlying table name of the DAO.
	group   string       // group is the database configuration group name of current DAO.
	columns PlansColumns // columns contains all the column names of Table for convenient usage.
}

// PlansColumns defines and stores column names for table plans.
type PlansColumns struct {
	Id        string //
	Name      string //
	Price     string //
	Unit      string //
	CreatedAt string //
	UpdatedAt string //
	DeletedAt string //
}

//  plansColumns holds the columns for table plans.
var plansColumns = PlansColumns{
	Id:        "id",
	Name:      "name",
	Price:     "price",
	Unit:      "unit",
	CreatedAt: "created_at",
	UpdatedAt: "updated_at",
	DeletedAt: "deleted_at",
}

// NewPlansDao creates and returns a new DAO object for table data access.
func NewPlansDao() *PlansDao {
	return &PlansDao{
		group:   "default",
		table:   "plans",
		columns: plansColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *PlansDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *PlansDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *PlansDao) Columns() PlansColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *PlansDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *PlansDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *PlansDao) Transaction(ctx context.Context, f func(ctx context.Context, tx *gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
