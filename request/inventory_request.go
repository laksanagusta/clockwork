package request

type InventoryCreateInput struct {
	StockQty  int  `json:"stockQty" validate:"required" binding:"required"`
	IsInStock bool `json:"isInStock" validate:"required" binding:"required"`
	ProductID int  `json:"productId" binding:"required"`
}

type InventoryUpdateInput struct {
	StockQty  int  `json:"stockQty" validate:"required" binding:"required"`
	IsInStock bool `json:"isInStock" validate:"required" binding:"required"`
	ProductID int  `json:"productId" binding:"required"`
}

type InventoryFindById struct {
	ID int `uri:"id" binding:"required"`
}

type InventoryFindByProductId struct {
	ProductID int `uri:"productId" binding:"required"`
}
