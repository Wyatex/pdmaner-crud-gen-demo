package lbra

import (
	"context"
	"pdmaner-test/internal/service"

	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"

	"pdmaner-test/api/lbra/v1"
)

func (c *ControllerV1) LbraCatalogDel(ctx context.Context, req *v1.LbraCatalogDelReq) (res *v1.LbraCatalogDelRes, err error) {
	IdOrIdsNil := gcode.New(-1, "id或者ids不能为空", nil)
	if req.Id == nil {
		if req.Ids == nil {
			return nil, gerror.NewCode(IdOrIdsNil)
		} else if len(req.Ids) == 0 {
			return nil, gerror.NewCode(IdOrIdsNil)
		}
	}
	err = service.LbraCatalog().Delete(ctx, req.Id, req.Ids)
	return nil, err
}
