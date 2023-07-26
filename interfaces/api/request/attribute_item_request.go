package request

type AttributeItemCreateRequest struct {
	Title            string `json:"title" validate:"required" binding:"required"`
	AdditionalCharge int    `json:"additionalCharge" validate:"required" binding:"required,numeric"`
	AttributeID      uint   `json:"attributeId" validate:"required" binding:"required,numeric"`
}

type AttributeItemUpdateRequest struct {
	Title            string `json:"title" validate:"required" binding:"required"`
	AdditionalCharge int    `json:"additionalCharge" validate:"required" binding:"required,numeric"`
	AttributeID      uint   `json:"attributeId" validate:"required" binding:"required,numeric"`
}

type AttributeItemFindById struct {
	ID int `uri:"id" binding:"required"`
}
