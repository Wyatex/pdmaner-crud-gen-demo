// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// LbraCatalogDao is the data access object for table lbra_catalog.
type LbraCatalogDao struct {
	table   string             // table is the underlying table name of the DAO.
	group   string             // group is the database configuration group name of current DAO.
	columns LbraCatalogColumns // columns contains all the column names of Table for convenient usage.
}

// LbraCatalogColumns defines and stores column names for table lbra_catalog.
type LbraCatalogColumns struct {
	Id          string // 分类ID
	CatalogName string // 分类名
	CatalogDesc string // 分类描述
	SortNo      string // 排序号
	DemoTime    string // 演示时间选择
	CreatedBy   string // 创建人
	CreatedAt   string // 创建时间
	UpdatedBy   string // 更新人
	UpdatedAt   string // 更新时间
	DeletedBy   string // 删除人
	DeletedAt   string // 删除时间
}

// lbraCatalogColumns holds the columns for table lbra_catalog.
var lbraCatalogColumns = LbraCatalogColumns{
	Id:          "id",
	CatalogName: "catalog_name",
	CatalogDesc: "catalog_desc",
	SortNo:      "sort_no",
	DemoTime:    "demo_time",
	CreatedBy:   "created_by",
	CreatedAt:   "created_at",
	UpdatedBy:   "updated_by",
	UpdatedAt:   "updated_at",
	DeletedBy:   "deleted_by",
	DeletedAt:   "deleted_at",
}

// NewLbraCatalogDao creates and returns a new DAO object for table data access.
func NewLbraCatalogDao() *LbraCatalogDao {
	return &LbraCatalogDao{
		group:   "default",
		table:   "lbra_catalog",
		columns: lbraCatalogColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *LbraCatalogDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *LbraCatalogDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *LbraCatalogDao) Columns() LbraCatalogColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *LbraCatalogDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *LbraCatalogDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *LbraCatalogDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
