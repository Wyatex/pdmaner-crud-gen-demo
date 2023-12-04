package lbra

import (
	"context"
	"github.com/gogf/gf/v2/util/gconv"
	"pdmaner-test/api/lbra/v1"
	"pdmaner-test/internal/model"
	"pdmaner-test/internal/service"
)

func (c *ControllerV1) LbraCatalogAdd(ctx context.Context, req *v1.LbraCatalogAddReq) (res *v1.LbraCatalogAddRes, err error) {
	in := new(model.LbraCatalogAddInput)
	err = gconv.Struct(req, in)
	if err != nil {
		return nil, err
	}
	err = service.LbraCatalog().Add(ctx, in)
	return nil, err
}
