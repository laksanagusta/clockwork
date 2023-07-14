package request

type RegisterUserRequest struct {
	Username    string `json:"username" binding:"required"`
	Name        string `json:"name" binding:"required"`
	Occupation  string `json:"occupation" binding:"required"`
	Address     string `json:"address" binding:"required"`
	Email       string `json:"email" binding:"required,email"`
	Password    string `json:"password" binding:"required"`
	PhoneNumber string `json:"phoneNumber" binding:"required,numeric,min=12,max=12"`
	Role        string `json:"role"`
}

type UpdateUserRequest struct {
	Username    string `json:"username" binding:"required"`
	Name        string `json:"name" binding:"required"`
	Occupation  string `json:"occupation" binding:"required"`
	Email       string `json:"email" binding:"required,email"`
	PhoneNumber string `json:"phoneNumber" binding:"required"`
}

type UpdateUserRoleRequest struct {
	Name       string `json:"name" binding:"required"`
	Occupation string `json:"occupation" binding:"required"`
	Email      string `json:"email" binding:"required,email"`
	Role       string `json:"role"`
}

type LoginRequest struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}

type GetUserDetailRequest struct {
	ID uint64 `uri:"id" binding:"required"`
}
