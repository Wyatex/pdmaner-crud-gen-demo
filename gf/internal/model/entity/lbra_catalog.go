// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// LbraCatalog is the golang structure for table lbra_catalog.
type LbraCatalog struct {
	Id          int         `json:"id"          ` // 分类ID
	CatalogName string      `json:"catalogName" ` // 分类名
	CatalogDesc string      `json:"catalogDesc" ` // 分类描述
	SortNo      int         `json:"sortNo"      ` // 排序号
	DemoTime    *gtime.Time `json:"demoTime"    ` // 演示时间选择
	CreatedBy   string      `json:"createdBy"   ` // 创建人
	CreatedAt   *gtime.Time `json:"createdAt"   ` // 创建时间
	UpdatedBy   string      `json:"updatedBy"   ` // 更新人
	UpdatedAt   *gtime.Time `json:"updatedAt"   ` // 更新时间
	DeletedBy   string      `json:"deletedBy"   ` // 删除人
	DeletedAt   *gtime.Time `json:"deletedAt"   ` // 删除时间
}
