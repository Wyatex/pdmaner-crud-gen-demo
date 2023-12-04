// Package lbra 图书分类
// author : https://github.com/Wyatex
// date : 2023-11-04 11:09
// desc : 图书分类 logic代码
package lbra

import (
	"context"
	"github.com/gogf/gf/v2/frame/g"
	"math"
	"pdmaner-test/internal/dao"
	"pdmaner-test/internal/model"
	"pdmaner-test/internal/model/entity"
	"pdmaner-test/internal/service"
)

type sLbraCatalog struct{}

func init() {
	service.RegisterLbraCatalog(New())
}

func New() *sLbraCatalog {
	return &sLbraCatalog{}
}

// Add 添加一个图书分类
func (s *sLbraCatalog) Add(ctx context.Context, in *model.LbraCatalogAddInput) error {
	uid := service.BizCtx().Get(ctx).User.Id
	in.CreatedBy = uid
	_, err := dao.LbraCatalog.Ctx(ctx).Insert(in)
	return err
}

// Edit 编辑图书分类
func (s *sLbraCatalog) Edit(ctx context.Context, in *entity.LbraCatalog) error {
	uid := service.BizCtx().Get(ctx).User.Id
	in.UpdatedBy = uid
	_, err := dao.LbraCatalog.Ctx(ctx).Update(in, "id", in.Id)
	return err
}

// Delete 删除图书分类
func (s *sLbraCatalog) Delete(ctx context.Context, id *int, ids []int) (err error) {
	uid := service.BizCtx().Get(ctx).User.Id
	if id != nil {
		_, err = dao.LbraCatalog.Ctx(ctx).Data(g.Map{"deleted_by": uid}).Where("id", id).Update()
	} else {
		_, err = dao.LbraCatalog.Ctx(ctx).Data(g.Map{"deleted_by": uid}).WhereIn("id", ids).Update()
	}
	if err != nil {
		return err
	}
	if id != nil {
		_, err = dao.LbraCatalog.Ctx(ctx).Where("id", id).Delete()
	} else {
		_, err = dao.LbraCatalog.Ctx(ctx).WhereIn("id", ids).Delete()
	}
	return err
}

// FindOne 根据id查找图书分类
func (s *sLbraCatalog) FindOne(ctx context.Context, id int) (*entity.LbraCatalog, error) {
	var m *entity.LbraCatalog
	err := dao.LbraCatalog.Ctx(ctx).Where("id", id).Scan(&m)
	return m, err
}

// FindPage 分页查询图书分类
func (s *sLbraCatalog) FindPage(ctx context.Context, in *model.LbraCatalogFindPageInput) (*model.LbraCatalogFindPageOutput, error) {
	orm := dao.LbraCatalog.Ctx(ctx).Safe(false)
	items := make([]entity.LbraCatalog, 0)
	total := 0
	if in.CatalogName != "" {
		orm.WhereLike(dao.LbraCatalog.Columns().CatalogName, in.CatalogName)
	}
	if in.CatalogDesc != "" {
		orm.WhereLike(dao.LbraCatalog.Columns().CatalogDesc, in.CatalogDesc)
	}
	if in.DemoTime != nil {
		orm.Where(dao.LbraCatalog.Columns().DemoTime, in.DemoTime)
	}
	if in.CreatedStart != nil {
		orm.WhereGTE(dao.LbraCatalog.Columns().CreatedAt, in.CreatedStart)
	}
	if in.CreatedEnd != nil {
		orm.WhereLT(dao.LbraCatalog.Columns().CreatedAt, in.CreatedEnd)
	}
	if in.UpdatedStart != nil {
		orm.WhereGTE(dao.LbraCatalog.Columns().UpdatedAt, in.UpdatedStart)
	}
	if in.UpdatedEnd != nil {
		orm.WhereLT(dao.LbraCatalog.Columns().UpdatedAt, in.UpdatedEnd)
	}
	err := orm.Limit((in.PageNo-1)*in.PageItem, in.PageItem).
		ScanAndCount(&items, &total, true)
	if err != nil {
		return nil, err
	}
	return &model.LbraCatalogFindPageOutput{
		List:      items,
		ItemCount: total,
		PageCount: int(math.Ceil(float64(total) / float64(in.PageItem))),
		PageItem:  in.PageItem,
		PageNo:    in.PageNo,
	}, nil
}

// FindList 查找所有图书分类
func (s *sLbraCatalog) FindList(ctx context.Context, in *model.LbraCatalogFindListInput) ([]entity.LbraCatalog, error) {
	orm := dao.LbraCatalog.Ctx(ctx).Safe(false)
	items := make([]entity.LbraCatalog, 0)
	if in.CatalogName != "" {
		orm.WhereLike(dao.LbraCatalog.Columns().CatalogName, "%"+in.CatalogName+"%")
	}
	if in.CatalogDesc != "" {
		orm.WhereLike(dao.LbraCatalog.Columns().CatalogDesc, "%"+in.CatalogDesc+"%")
	}
	if in.DemoTime != nil {
		orm.Where(dao.LbraCatalog.Columns().DemoTime, in.DemoTime)
	}
	if in.CreatedStart != nil {
		orm.WhereGTE(dao.LbraCatalog.Columns().CreatedAt, in.CreatedStart)
	}
	if in.CreatedEnd != nil {
		orm.WhereLT(dao.LbraCatalog.Columns().CreatedAt, in.CreatedEnd)
	}
	if in.UpdatedStart != nil {
		orm.WhereGTE(dao.LbraCatalog.Columns().UpdatedAt, in.UpdatedStart)
	}
	if in.UpdatedEnd != nil {
		orm.WhereLT(dao.LbraCatalog.Columns().UpdatedAt, in.UpdatedEnd)
	}
	err := orm.Scan(&items)
	return items, err
}
