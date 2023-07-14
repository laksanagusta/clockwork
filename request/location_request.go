package request

type LocationCreateInput struct {
	Name    string `json:"name" validate:"required" binding:"required"`
	Code    string `json:"code" validate:"required" binding:"required"`
	Address string `json:"address" validate:"required" binding:"required"`
}

type LocationUpdateInput struct {
	Name    string `json:"name" validate:"required" binding:"required"`
	Code    string `json:"code" validate:"required" binding:"required"`
	Address string `json:"address" validate:"required" binding:"required"`
}

type LocationFindById struct {
	ID int `uri:"id" binding:"required"`
}
