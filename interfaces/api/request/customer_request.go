package request

type RegisterCustomerRequest struct {
	Name        string `json:"name" binding:"required"`
	Email       string `json:"email" binding:"required,email"`
	Password    string `json:"password" binding:"required"`
	PhoneNumber string `json:"phoneNumber" binding:"required,numeric,min=12,max=12"`
}

type UpdateCustomerRequest struct {
	Name        string `json:"name" binding:"required"`
	Email       string `json:"email" binding:"required,email"`
	Password    string `json:"password" binding:"required"`
	PhoneNumber string `json:"phoneNumber" binding:"required,numeric,min=12,max=12"`
}

type UserLoginRequest struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}

type CustomerFindById struct {
	ID uint64 `uri:"id" binding:"required"`
}
