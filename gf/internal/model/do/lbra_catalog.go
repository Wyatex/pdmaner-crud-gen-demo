// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// LbraCatalog is the golang structure of table lbra_catalog for DAO operations like Where/Data.
type LbraCatalog struct {
	g.Meta      `orm:"table:lbra_catalog, do:true"`
	Id          interface{} // 分类ID
	CatalogName interface{} // 分类名
	CatalogDesc interface{} // 分类描述
	SortNo      interface{} // 排序号
	DemoTime    *gtime.Time // 演示时间选择
	CreatedBy   interface{} // 创建人
	CreatedAt   *gtime.Time // 创建时间
	UpdatedBy   interface{} // 更新人
	UpdatedAt   *gtime.Time // 更新时间
	DeletedBy   interface{} // 删除人
	DeletedAt   *gtime.Time // 删除时间
}
