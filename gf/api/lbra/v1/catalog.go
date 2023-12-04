// Package v1 图书分类
// author : https://github.com/Wyatex
// date : 2023-11-04 11:10
// desc : 图书分类 api定义
package v1

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
	"pdmaner-test/internal/model"
	"pdmaner-test/internal/model/entity"
)

type LbraCatalogAddReq struct {
	g.Meta      `path:"/lbra_catalog/v1/add" tags:"LbraCatalog" method:"post" summary:"创建一个图书分类"`
	CatalogName string      `json:"catalogName" dc:"分类名" v:"required#分类名不能为空"`
	CatalogDesc string      `json:"catalogDesc" dc:"分类描述"`
	SortNo      *int        `json:"sortNo" dc:"排序号"`
	DemoTime    *gtime.Time `json:"demoTime" dc:"演示时间选择"`
}

type LbraCatalogAddRes bool

type LbraCatalogEditReq struct {
	g.Meta      `path:"/lbra_catalog/v1/edit" tags:"LbraCatalog" method:"post" summary:"编辑图书分类"`
	Id          *int        `json:"id" dc:"分类ID" v:"required#分类ID不能为空"`
	CatalogName string      `json:"catalogName" dc:"分类名" v:"required#分类名不能为空"`
	CatalogDesc string      `json:"catalogDesc" dc:"分类描述"`
	SortNo      *int        `json:"sortNo" dc:"排序号"`
	DemoTime    *gtime.Time `json:"demoTime" dc:"演示时间选择"`
	CreatedBy   string      `json:"createdBy" dc:"创建人"`
	CreatedAt   *gtime.Time `json:"createdAt" dc:"创建时间"`
	UpdatedBy   string      `json:"updatedBy" dc:"更新人"`
	UpdatedAt   *gtime.Time `json:"updatedAt" dc:"更新时间"`
	DeletedBy   string      `json:"deletedBy" dc:"删除人"`
	DeletedAt   *gtime.Time `json:"deletedAt" dc:"删除时间"`
}

type LbraCatalogEditRes struct{}

type LbraCatalogDelReq struct {
	g.Meta `path:"/lbra_catalog/v1/del" tags:"LbraCatalog" method:"post" summary:"删除图书分类"`
	Id     *int  `json:"id" dc:"删除单个id"`
	Ids    []int `json:"ids" dc:"批量删除"`
}

type LbraCatalogDelRes struct{}

type LbraCatalogFindOneReq struct {
	g.Meta `path:"/lbra_catalog/v1/fine_one" tags:"LbraCatalog" method:"post" summary:"根据id查询图书分类"`
	Id     int `json:"id" dc:"删除单个id" v:"required#id不能为空"`
}

type LbraCatalogFindOneRes entity.LbraCatalog

type LbraCatalogPageReq struct {
	g.Meta       `path:"/lbra_catalog/v1/page" tags:"LbraCatalog" method:"post" summary:"分页查询图书分类"`
	CatalogName  string      `json:"catalogName" dc:"分类名"`
	CatalogDesc  string      `json:"catalogDesc" dc:"分类描述"`
	SortNo       *int        `json:"sortNo" dc:"排序号"`
	DemoTime     *gtime.Time `json:"demoTime" dc:"演示时间选择"`
	CreatedStart *gtime.Time `json:"createdStart" dc:"搜索创建时间范围"`
	CreatedEnd   *gtime.Time `json:"createdEnd" dc:"搜索创建时间范围"`
	UpdatedStart *gtime.Time `json:"updatedStart" dc:"搜索更新时间范围"`
	UpdatedEnd   *gtime.Time `json:"updatedEnd" dc:"搜索更新时间范围"`
	PageNo       int         `json:"pageNo" dc:"查找第几页" v:"required#查询页数不能为空"`
	PageItem     int         `json:"pageItem" dc:"每页个数" v:"required#每页个数不能为空"`
}

type LbraCatalogPageRes model.LbraCatalogFindPageOutput

type LbraCatalogListReq struct {
	g.Meta       `path:"/lbra_catalog/v1/list" tags:"LbraCatalog" method:"post" summary:"查询所有图书分类"`
	CatalogName  string      `json:"catalogName" dc:"分类名"`
	CatalogDesc  string      `json:"catalogDesc" dc:"分类描述"`
	SortNo       *int        `json:"sortNo" dc:"排序号"`
	DemoTime     *gtime.Time `json:"demoTime" dc:"演示时间选择"`
	CreatedStart *gtime.Time `json:"createdStart" dc:"搜索创建时间范围"`
	CreatedEnd   *gtime.Time `json:"createdEnd" dc:"搜索创建时间范围"`
	UpdatedStart *gtime.Time `json:"updatedStart" dc:"搜索更新时间范围"`
	UpdatedEnd   *gtime.Time `json:"updatedEnd" dc:"搜索更新时间范围"`
}
type LbraCatalogListRes []entity.LbraCatalog
