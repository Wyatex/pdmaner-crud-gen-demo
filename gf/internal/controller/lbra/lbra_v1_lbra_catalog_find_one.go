package lbra

import (
	"context"
	"pdmaner-test/internal/service"

	"pdmaner-test/api/lbra/v1"
)

func (c *ControllerV1) LbraCatalogFindOne(ctx context.Context, req *v1.LbraCatalogFindOneReq) (res *v1.LbraCatalogFindOneRes, err error) {
	m, err := service.LbraCatalog().FindOne(ctx, req.Id)
	return (*v1.LbraCatalogFindOneRes)(m), err
}
