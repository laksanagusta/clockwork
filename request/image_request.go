package request

type ImageCreateRequest struct {
	Url       string `json:"url" validate:"required" binding:"required"`
	ProductID string `json:"productId" validate:"required" binding:"required"`
}

type ImageRemoveRequest struct {
	ID int8 `uri:"id" binding:"required"`
}
