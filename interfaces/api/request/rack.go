package request

type RackCreateInput struct {
	Title  string `json:"title" validate:"required" binding:"required"`
	Code   string `json:"code" validate:"required" binding:"required"`
	MaxQty string `json:"maxQty" validate:"required" binding:"required"`
}

type RackUpdateInput struct {
	Title  string `json:"title" validate:"required" binding:"required"`
	Code   string `json:"code" validate:"required" binding:"required"`
	MaxQty string `json:"maxQty" validate:"required" binding:"required"`
}

type RackFindById struct {
	ID int `uri:"id" binding:"required"`
}
