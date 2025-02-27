package web

type CustomerCreateRequest struct {
	Name    string `validate:"required,min=1,max=100" json:"name"`
	Email   string `validate:"required,email" json:"customer_email"`
	Phone   string `validate:"required" json:"customer_phone"`
	Address string `validate:"required,min=1,max=100" json:"customer_address"`
}

type CustomerUpdateRequest struct {
	Id         uint64 `validate:"required"`
	Name       string `validate:"required,max=200,min=1" json:"customer_name"`
	Email      string `validate:"required,email" json:"customer_email"`
	Phone      string `validate:"required" json:"customer_phone"`
	Address    string `validate:"required,min=1,max=100" json:"customer_address"`
	LoyaltyPts int    `json:"loyalty_pts"`
}

type CustomerResponse struct {
	Id         uint64 `json:"id"`
	Name       string `json:"customer_name"`
	Email      string `json:"customer_email"`
	Phone      string `json:"customer_phone"`
	Address    string `json:"customer_address"`
	LoyaltyPts int    `json:"loyalty_pts"`
}
