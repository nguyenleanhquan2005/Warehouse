package payload

import (
	"app/internal/domain/entity"
)

const defaultPageSize uint32 = 10

// PagingRequest holds paging query request
type PagingRequest struct {
	Size *int `form:"size" json:"size" validate:"omitempty,gte=1,lte=500"` // number of items per page
	Page *int `form:"page" json:"page" validate:"omitempty,gte=1"`         // page number
}

// StructName returns payload name
func (p PagingRequest) StructName() string {
	return "PagingRequest"
}

// Form converts paging data from request to entity
func (p PagingRequest) Form() entity.PagingRequest {
	pg := entity.PagingRequest{
		Size: defaultPageSize,
		Page: 1,
	}
	if p.Size != nil {
		pg.Size = uint32(*p.Size)
	}

	if p.Page != nil {
		pg.Page = uint32(*p.Page)
	}

	return pg
}
