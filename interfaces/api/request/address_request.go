package request

type AddressCreateRequest struct {
	Title      string `json:"title" validate:"required" binding:"required"`
	Street     string `json:"street" validate:"required" binding:"required"`
	Note       string `json:"note"`
	City       string `json:"city" validate:"required" binding:"required"`
	CustomerID uint   `json:"customerId"`
}

type AddressUpdateRequest struct {
	Title  string `json:"title" validate:"required" binding:"required"`
	Street string `json:"street" validate:"required" binding:"required"`
	Note   string `json:"note"`
	City   string `json:"city" validate:"required" binding:"required"`
}

type AddressFindById struct {
	ID int `uri:"id" binding:"required"`
}
