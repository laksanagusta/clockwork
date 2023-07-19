package request

type AttributeCreateRequest struct {
	Title      string `json:"title" validate:"required" binding:"required"`
	IsMultiple *bool  `json:"isMultiple" validate:"required" binding:"required"`
	IsRequired *bool  `json:"isRequired" validate:"required" binding:"required"`
}

type AttributeUpdateRequest struct {
	Title      string `json:"title" validate:"required" binding:"required"`
	IsMultiple *bool  `json:"isMultiple" validate:"required" binding:"required"`
	IsRequired *bool  `json:"isRequired" validate:"required" binding:"required"`
}

type AttributeFindById struct {
	ID int `uri:"id" binding:"required"`
}
