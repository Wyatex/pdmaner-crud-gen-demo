package lbra

import (
	"context"
	"github.com/gogf/gf/v2/util/gconv"
	"pdmaner-test/internal/model/entity"
	"pdmaner-test/internal/service"

	"pdmaner-test/api/lbra/v1"
)

func (c *ControllerV1) LbraCatalogEdit(ctx context.Context, req *v1.LbraCatalogEditReq) (res *v1.LbraCatalogEditRes, err error) {
	in := new(entity.LbraCatalog)
	err = gconv.Struct(req, in)
	if err != nil {
		return nil, err
	}
	err = service.LbraCatalog().Edit(ctx, in)
	return nil, err
}
