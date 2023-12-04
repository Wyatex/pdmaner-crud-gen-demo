package lbra

import (
	"context"
	"github.com/gogf/gf/v2/util/gconv"
	"pdmaner-test/internal/model"
	"pdmaner-test/internal/service"

	"pdmaner-test/api/lbra/v1"
)

func (c *ControllerV1) LbraCatalogList(ctx context.Context, req *v1.LbraCatalogListReq) (res *v1.LbraCatalogListRes, err error) {
	in := new(model.LbraCatalogFindListInput)
	err = gconv.Struct(req, in)
	if err != nil {
		return nil, err
	}
	list, err := service.LbraCatalog().FindList(ctx, in)
	return (*v1.LbraCatalogListRes)(&list), err
}
