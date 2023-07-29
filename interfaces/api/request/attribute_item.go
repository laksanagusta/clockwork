package request

type AttributeItemCreateRequest struct {
	Title            string `json:"title" validate:"required" binding:"required"`
	AdditionalCharge int    `json:"additionalCharge" validate:"required" binding:"required,number"`
	AttributeID      uint   `json:"attributeId" validate:"required" binding:"required,number"`
}

type AttributeItemUpdateRequest struct {
	Title            string `json:"title" validate:"required" binding:"required"`
	AdditionalCharge int    `json:"additionalCharge" validate:"required" binding:"required,number"`
	AttributeID      uint   `json:"attributeId" validate:"required" binding:"required,number"`
}

type AttributeItemFindById struct {
	ID int `uri:"id" binding:"required"`
}
