package request

import "clockwork-server/model"

type ProductCreateInput struct {
	Title        string `json:"title" validate:"required" binding:"required"`
	UnitPrice    int    `json:"unitPrice" validate:"required" binding:"required,numeric,min=1"`
	Description  string `json:"description" validate:"required" binding:"required"`
	SerialNumber string `json:"serialNumber" validate:"required" binding:"required,alphanum"`
	Attributes   []int  `json:"attributes" validate:"required" binding:"required"`
	CategoryID   uint   `json:"categoryId" validate:"required" binding:"required,numeric"`
	User         model.User
}
type ProductUpdateInput struct {
	Title        string `json:"title" validate:"required" binding:"required"`
	UnitPrice    int    `json:"unitPrice" validate:"required" binding:"required"`
	Description  string `json:"description" validate:"required" binding:"required"`
	SerialNumber string `json:"serialNumber" validate:"required" binding:"required"`
	Attributes   []int  `json:"attributes" validate:"required" binding:"required"`
	CategoryID   uint   `json:"categoryId" validate:"required" binding:"required,numeric"`
}

type ProductFindById struct {
	ID int `uri:"id" binding:"required"`
}

type ProductFindBySerialNumber struct {
	SerialNumber string `uri:"serialNumber" binding:"required"`
}
