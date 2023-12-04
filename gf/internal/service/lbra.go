// ================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// You can delete these comments if you wish manually maintain this interface file.
// ================================================================================

package service

import (
	"context"
	"pdmaner-test/internal/model"
	"pdmaner-test/internal/model/entity"
)

type (
	ILbraCatalog interface {
		// Add 添加一个图书分类
		Add(ctx context.Context, in *model.LbraCatalogAddInput) error
		// Edit 编辑图书分类
		Edit(ctx context.Context, in *entity.LbraCatalog) error
		// Delete 删除图书分类
		Delete(ctx context.Context, id *int, ids []int) (err error)
		// FindOne 根据id查找图书分类
		FindOne(ctx context.Context, id int) (*entity.LbraCatalog, error)
		// FindPage 分页查询图书分类
		FindPage(ctx context.Context, in *model.LbraCatalogFindPageInput) (*model.LbraCatalogFindPageOutput, error)
		// FindList 查找所有图书分类
		FindList(ctx context.Context, in *model.LbraCatalogFindListInput) ([]entity.LbraCatalog, error)
	}
)

var (
	localLbraCatalog ILbraCatalog
)

func LbraCatalog() ILbraCatalog {
	if localLbraCatalog == nil {
		panic("implement not found for interface ILbraCatalog, forgot register?")
	}
	return localLbraCatalog
}

func RegisterLbraCatalog(i ILbraCatalog) {
	localLbraCatalog = i
}
