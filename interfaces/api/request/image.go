package request

type ImageCreateRequest struct {
	ProductID int  `form:"productId" validate:"required" binding:"required"`
	IsPrimary bool `form:"isPrimary" validate:"required" binding:"required"`
}

type ImageUpdateRequest struct {
	IsPrimary bool `form:"isPrimary" validate:"required" binding:"required"`
}

type ImageRemoveRequest struct {
	ID int8 `uri:"id" binding:"required"`
}

type ImageFindByIdRequest struct {
	ID int8 `uri:"id" binding:"required"`
}
