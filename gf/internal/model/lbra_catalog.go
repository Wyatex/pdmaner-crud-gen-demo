// Package model 图书分类
// author : https://github.com/Wyatex
// date : 2023-11-04 11:09
// desc : 图书分类 controller - service 传输数据类型定义
package model

import (
	"github.com/gogf/gf/v2/os/gtime"
	"pdmaner-test/internal/model/entity"
)

// LbraCatalogAddInput 添加图书分类
type LbraCatalogAddInput struct {
	CatalogName string      // 分类名
	CatalogDesc string      // 分类描述
	SortNo      *int        // 排序号
	DemoTime    *gtime.Time // 演示时间选择
	CreatedBy   string      // 创建者
}

// LbraCatalogFindPageInput 分页查找图书分类
type LbraCatalogFindPageInput struct {
	CatalogName  string      // 分类名
	CatalogDesc  string      // 分类描述
	SortNo       *int        // 排序号
	DemoTime     *gtime.Time // 演示时间选择
	CreatedStart *gtime.Time // 搜索创建时间范围
	CreatedEnd   *gtime.Time // 搜索创建时间范围
	UpdatedStart *gtime.Time // 搜索更新时间范围
	UpdatedEnd   *gtime.Time // 搜索更新时间范围
	PageNo       int         // 查找第几页
	PageItem     int         // 每页个数
}

// LbraCatalogFindPageOutput 分页查找图书分类
type LbraCatalogFindPageOutput struct {
	List      []entity.LbraCatalog
	ItemCount int // 搜索结果有多少个
	PageCount int // 搜索结果有多少页
	PageNo    int // 查找第几页
	PageItem  int // 每页个数
}

// LbraCatalogFindListInput 全部查找图书分类
type LbraCatalogFindListInput struct {
	CatalogName  string      // 分类名
	CatalogDesc  string      // 分类描述
	SortNo       *int        // 排序号
	DemoTime     *gtime.Time // 演示时间选择
	CreatedStart *gtime.Time // 搜索创建时间范围
	CreatedEnd   *gtime.Time // 搜索创建时间范围
	UpdatedStart *gtime.Time // 搜索更新时间范围
	UpdatedEnd   *gtime.Time // 搜索更新时间范围
}
