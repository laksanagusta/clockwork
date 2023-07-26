package request

type OrganizationCreateInput struct {
	Name    string `json:"name" validate:"required" binding:"required"`
	Address string `json:"address" validate:"required" binding:"required"`
}

type OrganizationUpdateInput struct {
	Name    string `json:"name" validate:"required" binding:"required"`
	Address string `json:"address" validate:"required" binding:"required"`
}

type OrganizationFindById struct {
	ID int `uri:"id" binding:"required"`
}
