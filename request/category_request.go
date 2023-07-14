package request

type CategoryCreateInput struct {
	Title string `json:"title" validate:"required" binding:"required"`
}

type CategoryUpdateInput struct {
	Title string `json:"title" validate:"required" binding:"required"`
}

type CategoryFindById struct {
	ID int `uri:"id" binding:"required"`
}
