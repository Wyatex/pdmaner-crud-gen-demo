// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package lbra

import (
	"context"

	"pdmaner-test/api/lbra/v1"
)

type ILbraV1 interface {
	LbraCatalogAdd(ctx context.Context, req *v1.LbraCatalogAddReq) (res *v1.LbraCatalogAddRes, err error)
	LbraCatalogEdit(ctx context.Context, req *v1.LbraCatalogEditReq) (res *v1.LbraCatalogEditRes, err error)
	LbraCatalogDel(ctx context.Context, req *v1.LbraCatalogDelReq) (res *v1.LbraCatalogDelRes, err error)
	LbraCatalogFindOne(ctx context.Context, req *v1.LbraCatalogFindOneReq) (res *v1.LbraCatalogFindOneRes, err error)
	LbraCatalogPage(ctx context.Context, req *v1.LbraCatalogPageReq) (res *v1.LbraCatalogPageRes, err error)
	LbraCatalogList(ctx context.Context, req *v1.LbraCatalogListReq) (res *v1.LbraCatalogListRes, err error)
}
