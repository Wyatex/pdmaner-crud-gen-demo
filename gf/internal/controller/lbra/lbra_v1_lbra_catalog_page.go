package lbra

import (
	"context"
	"github.com/gogf/gf/v2/util/gconv"
	"pdmaner-test/internal/model"
	"pdmaner-test/internal/service"

	"pdmaner-test/api/lbra/v1"
)

func (c *ControllerV1) LbraCatalogPage(ctx context.Context, req *v1.LbraCatalogPageReq) (res *v1.LbraCatalogPageRes, err error) {
	in := new(model.LbraCatalogFindPageInput)
	err = gconv.Struct(req, in)
	if err != nil {
		return nil, err
	}
	list, err := service.LbraCatalog().FindPage(ctx, in)
	return (*v1.LbraCatalogPageRes)(list), err
}
